package queries

import (
	"context"
	"testing"
	"time"

	"github.com/freedom-sketch/sub2go/internal/database"
	"github.com/freedom-sketch/sub2go/internal/database/models"
)

func TestCreateAndGetSubscription(t *testing.T) {
	db, err := database.ConnectInMemory()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}

	ctx := context.Background()

	user := models.User{
		UserID: 123,
	}
	err = db.WithContext(ctx).Create(&user).Error
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	now := time.Now().Truncate(time.Second)
	sub := &models.Subscription{
		UserID:    user.UserID,
		StartDate: now,
		EndDate:   now.Add(30 * 24 * time.Hour),
	}

	if err := CreateSubscription(ctx, db, sub); err != nil {
		t.Fatalf("Error creating subscription: %v", err)
	}

	got, err := GetSubscriptionByUserID(ctx, db, user.UserID)
	if err != nil {
		t.Fatalf("Failed to get subscription: %v", err)
	}
	t.Logf("Subscription created and received:: %+v", got)

	if !got.StartDate.Equal(now) {
		t.Errorf("StartDate does not match: expected %v, received %v", now, got.StartDate)
	}
}

func TestExtendSubscription(t *testing.T) {
	db, err := database.ConnectInMemory()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}

	ctx := context.Background()

	user := models.User{
		UserID: 123,
	}
	err = db.WithContext(ctx).Create(&user).Error
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	now := time.Now().Truncate(time.Second)
	endDate := now.Add(30 * 24 * time.Hour)
	sub := &models.Subscription{
		UserID:    user.UserID,
		StartDate: now,
		EndDate:   endDate,
	}

	if err := CreateSubscription(ctx, db, sub); err != nil {
		t.Fatalf("Error creating subscription: %v", err)
	}
	t.Logf("The subscription is valid until %v", endDate)

	err = ExtendSubscription(ctx, db, user.UserID, 7)
	if err != nil {
		t.Fatalf("Error renewing subscription: %v", err)
	}

	got, err := GetSubscriptionByUserID(ctx, db, user.UserID)
	if err != nil {
		t.Fatalf("Failed to get subscription: %v", err)
	}

	t.Logf("After renewal, the subscription is valid until %v", got.EndDate)
}
