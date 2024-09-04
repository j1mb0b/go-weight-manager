package server

import (
	"context"
	"testing"

	pb "github.com/j1mb0b/go-weight-manager/proto"
	"github.com/stretchr/testify/assert"
)

func TestAddGetEntries(t *testing.T) {
	client := NewServer()

	ctx := context.Background()

	testCases := []struct {
		name  string
		entry *pb.WeightEntry
	}{
		{
			name: "First entry",
			entry: &pb.WeightEntry{
				Uid:    1,
				Date:   "2024-09-01",
				Weight: 75,
			},
		},
		{
			name: "Second entry",
			entry: &pb.WeightEntry{
				Uid:    1,
				Date:   "2024-09-02",
				Weight: 74,
			},
		},
	}

	// Table-Driven Tests
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := client.AddEntry(ctx, tc.entry)
			if err != nil {
				t.Fatalf("AddEntry failed: %v", err)
			}
			assert.Equal(t, "Success", resp.Message)
		})
	}

	resp, err := client.GetEntries(context.Background(), &pb.UserID{Uid: 1})
	if err != nil {
		t.Fatalf("GetEntries failed: %v", err)
	}

	t.Logf("Entries: %v\n", resp.Entries)

	assert.Len(t, resp.Entries, 2)
}

func TestDeleteEntries(t *testing.T) {
	client := NewServer()
	ctx := context.Background()

	// Add entries
	entries := &pb.WeightEntries{
		Entries: []*pb.WeightEntry{
			{
				Uid:    1,
				Date:   "2024-09-01",
				Weight: 75,
			},
			{
				Uid:    1,
				Date:   "2024-09-02",
				Weight: 74,
			},
		},
	}
	for _, entry := range entries.Entries {
		_, err := client.AddEntry(ctx, entry)
		if err != nil {
			t.Fatalf("AddEntry failed: %v", err)
		}
	}

	// Get entries
	resp, err := client.GetEntries(context.Background(), &pb.UserID{Uid: 1})
	if err != nil {
		t.Fatalf("GetEntries failed: %v", err)
	}
	assert.Len(t, resp.Entries, 2)
	t.Logf("Entries before: %v\n", resp.Entries)

	for _, entry := range resp.Entries {

		// Only delete first entry
		if entry.Date != "2024-09-01" {
			continue
		}

		_, err := client.DeleteEntry(ctx, &pb.EntryID{Id: entry.Id})
		if err != nil {
			t.Fatalf("DeleteEntry failed: %v", err)
		}
	}

	// Verify entry was deleted
	resp, err = client.GetEntries(context.Background(), &pb.UserID{Uid: 1})
	if err != nil {
		t.Fatalf("GetEntries failed: %v", err)
	}

	assert.Len(t, resp.Entries, 1)
	t.Logf("Entries after: %v\n", resp.Entries)
}
