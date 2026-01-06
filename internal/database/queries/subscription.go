package queries

import (
	"context"
	"fmt"
	"time"

	"errors"

	"github.com/freedom-sketch/sub2go/internal/database/models"
	"gorm.io/gorm"
)

// Creates a subscription
func CreateSubscription(ctx context.Context, db *gorm.DB, sub *models.Subscription) error {
	if sub == nil {
		return fmt.Errorf("subscription is nil")
	}
	return db.WithContext(ctx).Create(sub).Error
}

// Returns the subscription structure by subscription UserID
func GetSubscriptionByUserID(ctx context.Context, db *gorm.DB, userID int64) (*models.Subscription, error) {
	var sub models.Subscription
	err := db.WithContext(ctx).
		Where("user_id = ?", userID).
		First(&sub).Error
	if err != nil {
		return nil, err
	}
	return &sub, nil
}

// Checks if a subscription is active for the user
func IsSubscriptionActive(ctx context.Context, db *gorm.DB, userID int64) (bool, error) {
	var sub models.Subscription
	err := db.WithContext(ctx).
		Where("user_id = ?", userID).
		First(&sub).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return time.Now().Before(sub.EndDate), nil
}

// Extends a subscription for a specified number of days.
// If there is no subscription or it is not found, it returns an error
func ExtendSubscription(ctx context.Context, db *gorm.DB, userID int64, days int) error {
	if days <= 0 {
		return fmt.Errorf("the number of days must be positive")
	}

	var sub models.Subscription
	err := db.WithContext(ctx).
		Where("user_id = ?", userID).
		First(&sub).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("subscription for user %d not found", userID)
		}
		return err
	}

	newEndDate := sub.EndDate.AddDate(0, 0, days)

	result := db.WithContext(ctx).
		Model(&sub).
		Update("end_date", newEndDate)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("the subscription was not renewed")
	}

	return nil
}
