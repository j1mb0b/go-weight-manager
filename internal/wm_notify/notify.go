package wm_notify

import (
	"fmt"
	"sync"
	"time"

	pb "github.com/j1mb0b/go-weight-manager/proto"
	"github.com/prometheus/client_golang/prometheus"
)

type Notification struct {
	UserID  uint32
	Message string
}

type UserData struct {
	UserID  pb.UserID
	Entries []*pb.WeightEntry
}

var userData *UserData

var (
	notificationsProcessed = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "notifications_processed_total",
		Help: "Total number of notifications processed.",
	})

	notificationsSent = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "notifications_sent_total",
		Help: "Total number of notifications sent.",
	})

	notificationErrors = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "notification_errors_total",
		Help: "Total number of errors encountered while processing notifications.",
	})
)

func init() {
	prometheus.MustRegister(notificationsProcessed)
	prometheus.MustRegister(notificationsSent)
	prometheus.MustRegister(notificationErrors)

	// Simulate loading user data (replace with actual logic)
	time.Sleep(1 * time.Second)

	userData = &UserData{
		UserID: pb.UserID{Uid: 1},
		Entries: []*pb.WeightEntry{
			{Uid: 1, Weight: 90.0, Date: "2022-01-01"},
			{Uid: 1, Weight: 89.0, Date: "2022-01-02"},
			{Uid: 1, Weight: 86.0, Date: "2022-01-03"},
			{Uid: 1, Weight: 87.0, Date: "2022-01-04"},
			{Uid: 1, Weight: 85.0, Date: "2022-01-05"},
		},
	}
}

func (ud *UserData) analyizeWeight(notifications chan<- Notification) error {
	// Simulate checking weight (replace with actual logic)
	time.Sleep(10 * time.Second)

	if len(ud.Entries) == 0 {
		notificationErrors.Inc()
		return fmt.Errorf("no entries for user")
	}

	entries := ud.Entries

	// Get user ID from entries
	firstEntry := entries[0]
	userID := ud.UserID.Uid

	// Calculate average weight
	var avgWeight float32
	for _, entry := range ud.Entries {
		avgWeight += entry.Weight
	}
	avgWeight /= float32(len(ud.Entries))
	if avgWeight > 0 {
		notificationsProcessed.Inc()
		notifications <- Notification{UserID: userID, Message: fmt.Sprintf("Your average weight is %.2f kg", avgWeight)}
	}

	// Calculate weight vairance
	var weightVariance float32
	for _, entry := range ud.Entries {
		weightVariance += (entry.Weight - avgWeight) * (entry.Weight - avgWeight)
	}
	weightVariance /= float32(len(ud.Entries))

	gainedOrLost := "maintained"
	if weightVariance > firstEntry.Weight {
		gainedOrLost = "gained"
	} else if weightVariance < firstEntry.Weight {
		gainedOrLost = "lost"
	}
	notificationsProcessed.Inc()
	notifications <- Notification{UserID: userID, Message: fmt.Sprintf("Your have %s %.2f kg", gainedOrLost, weightVariance)}

	return nil
}

func sendNotification(notification Notification, wg *sync.WaitGroup) {
	defer wg.Done()

	// Simulate sending a notification (e.g., email or push notification)
	time.Sleep(500 * time.Millisecond)

	// Simulate sending a notification (e.g., email or push notification)
	fmt.Printf("Sending notification to user %d: %s\n", notification.UserID, notification.Message)

	notificationsSent.Inc()
}

// startNotifications is a function that starts the notification processing and sending routines.
func StartNotifications() {
	notifications := make(chan Notification)
	var wg sync.WaitGroup

	// 1. Launch Go routine for checking weights
	fmt.Println("Launching analyzeWeight routine...")
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := userData.analyizeWeight(notifications); err != nil {
			fmt.Println("Error:", err)
		}
		// Close the channel after all notifications have been sent
		close(notifications)
	}()

	// 2. Launch Go routine to process and send notifications
	fmt.Println("Launching notification sending routines...")
	wg.Add(1)
	go func() {
		defer wg.Done()
		for notification := range notifications {
			wg.Add(1)
			go sendNotification(notification, &wg)
		}
	}()

	// Wait for all Go routines to complete
	wg.Wait()

	fmt.Println("All tasks completed.")
}
