package database

import (
	"context"
	"fmt"

	"github.com/freedom-sketch/sub2go/internal/database/models"
	"gorm.io/gorm"
)

func CreateSubscription(ctx context.Context, db *gorm.DB, sub *models.Subscription) error {
	if sub == nil {
		return fmt.Errorf("subscription is nil")
	}
	return db.WithContext(ctx).Create(sub).Error
}

func GetSubscriptionByUserID(ctx context.Context, db *gorm.DB, userID uint) (*models.Subscription, error) {
	var sub models.Subscription
	err := db.WithContext(ctx).
		Where("user_id = ?", userID).
		First(&sub).Error
	if err != nil {
		return nil, err
	}
	return &sub, nil
}
