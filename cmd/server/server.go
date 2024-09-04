package server

import (
	"context"
	"fmt"
	"sync"

	"github.com/google/uuid"
	notify "github.com/j1mb0b/go-weight-manager/internal/wm_notify"
	pb "github.com/j1mb0b/go-weight-manager/proto"
)

type WeightManagerServer interface {
	AnalyzeWeight(ctx context.Context, empty *pb.Empty) (*pb.EntryResponse, error)
	AddEntry(ctx context.Context, entry *pb.WeightEntry) (*pb.EntryResponse, error)
	GetEntries(ctx context.Context, u *pb.UserID) (*pb.WeightEntries, error)
	UpdateEntry(ctx context.Context, entry *pb.WeightEntry) (*pb.EntryResponse, error)
	DeleteEntry(ctx context.Context, entry *pb.EntryID) (*pb.EntryResponse, error)
}

func NewServer() WeightManagerServer {
	return &server{
		entries: make(map[string]*pb.WeightEntry),
	}
}

type server struct {
	pb.UnimplementedWMServiceServer
	entries map[string]*pb.WeightEntry
	mu      sync.RWMutex // RWMutex to protect the read and write access to the entries map
}

var _ WeightManagerServer = &server{} // Compile-time check to ensure server implements WeightManagerServer

func (s *server) AddEntry(ctx context.Context, entry *pb.WeightEntry) (*pb.EntryResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	entry.Id = uuid.New().String() // Generate a unique ID
	s.entries[entry.Id] = entry

	return &pb.EntryResponse{Message: "Success"}, nil
}

func (s *server) GetEntries(ctx context.Context, u *pb.UserID) (*pb.WeightEntries, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var entries []*pb.WeightEntry
	for _, e := range s.entries {
		if e.Uid != u.Uid {
			continue
		}
		entries = append(entries, e)
	}

	return &pb.WeightEntries{Entries: entries}, nil
}

func (s *server) UpdateEntry(ctx context.Context, entry *pb.WeightEntry) (*pb.EntryResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if entry exists.
	if _, ok := s.entries[entry.Id]; !ok {
		return &pb.EntryResponse{Message: "Entry not found"}, nil
	}
	// Update server with new entry.
	s.entries[entry.Id] = entry
	return &pb.EntryResponse{Message: "Success"}, nil
}

func (s *server) DeleteEntry(ctx context.Context, entry *pb.EntryID) (*pb.EntryResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if entry exists.
	if _, ok := s.entries[entry.Id]; !ok {
		return &pb.EntryResponse{Message: fmt.Sprintf("Entry not found %s", entry.Id)}, nil
	}
	// Delete the entry.
	delete(s.entries, entry.Id)
	return &pb.EntryResponse{Message: fmt.Sprintf("Successfully deleted entry %s", entry.Id)}, nil
}

func (s *server) AnalyzeWeight(ctx context.Context, empty *pb.Empty) (*pb.EntryResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	notify.StartNotifications()

	return nil, nil
}
