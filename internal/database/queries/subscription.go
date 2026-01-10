package queries

import (
	"context"
	"fmt"
	"time"

	"errors"

	"github.com/freedom-sketch/sub2go/internal/database/models"
	"gorm.io/gorm"
)

// CreateSubscription creates a new subscription
func CreateSubscription(ctx context.Context, db *gorm.DB, sub *models.Subscription) error {
	if sub == nil {
		return fmt.Errorf("subscription is nil")
	}
	return db.WithContext(ctx).Create(sub).Error
}

// GetSubscriptionByUserUUID returns the subscription for the given user UUID
func GetSubscriptionByUserUUID(ctx context.Context, db *gorm.DB, userUUID string) (*models.Subscription, error) {
	var sub models.Subscription
	err := db.WithContext(ctx).
		Where("user_uuid = ?", userUUID).
		First(&sub).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &sub, nil
}

// IsSubscriptionActive checks if the subscription is currently active for the user
func IsSubscriptionActive(ctx context.Context, db *gorm.DB, userUUID string) (bool, error) {
	var sub models.Subscription
	err := db.WithContext(ctx).
		Where("user_uuid = ?", userUUID).
		First(&sub).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return time.Now().Before(sub.EndDate), nil
}

// ExtendSubscription extends the subscription by the specified number of days
// Returns error if subscription not found or update failed
func ExtendSubscription(ctx context.Context, db *gorm.DB, userUUID string, days int) error {
	if days <= 0 {
		return fmt.Errorf("the number of days must be positive")
	}

	var sub models.Subscription
	err := db.WithContext(ctx).
		Where("user_uuid = ?", userUUID).
		First(&sub).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("subscription for user %s not found", userUUID)
		}
		return err
	}

	newEndDate := sub.EndDate.AddDate(0, 0, days)

	result := db.WithContext(ctx).Model(&sub).Update("end_date", newEndDate)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("subscription was not updated")
	}

	return nil
}
