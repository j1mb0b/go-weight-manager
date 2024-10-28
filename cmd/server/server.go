package server

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/gocql/gocql"
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

type server struct {
	pb.UnimplementedWMServiceServer
	session *gocql.Session
	rwmu    sync.RWMutex
}

var _ WeightManagerServer = &server{} // Compile-time check to ensure server implements WeightManagerServer

func NewServer() WeightManagerServer {
	// Initialize Cassandra session and return a new server instance
	cluster := gocql.NewCluster("cassandra") // Change to your Cassandra host
	cluster.Keyspace = "weight_manager"      // Change to your keyspace
	cluster.Consistency = gocql.Quorum

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("failed to connect to Cassandra: %v", err)
	}

	//defer session.Close()

	return &server{
		session: session,
	}
}

func (s *server) AddEntry(ctx context.Context, entry *pb.WeightEntry) (*pb.EntryResponse, error) {
	s.rwmu.Lock()
	defer s.rwmu.Unlock()

	uuidString := uuid.Must(uuid.NewRandom())
	entry.Id = uuidString.String()

	// Insert the entry into Cassandra
	query := "INSERT INTO weight_entries (id, uid, date, weight) VALUES (?, ?, ?, ?)"
	if err := s.session.Query(query, entry.Id, entry.Uid, entry.Date, entry.Weight).Exec(); err != nil {
		return nil, fmt.Errorf("my mum failed to insert entries: %v", err)
	}

	return &pb.EntryResponse{Message: "Success"}, nil
}

func (s *server) GetEntries(ctx context.Context, u *pb.UserID) (*pb.WeightEntries, error) {
	s.rwmu.RLock()
	defer s.rwmu.RUnlock()

	// Get entries from Cassandra
	var entries []*pb.WeightEntry
	var entry pb.WeightEntry
	query := "SELECT id, uid, date, weight FROM weight_entries WHERE uid = ?"
	itr := s.session.Query(query, u.Uid).Iter()
	for itr.Scan(&entry.Id, &entry.Uid, &entry.Date, &entry.Weight) {
		entries = append(entries, &entry)
	}
	if err := itr.Close(); err != nil {
		return nil, fmt.Errorf("failed to get entries: %v", err)
	}

	return &pb.WeightEntries{Entries: entries}, nil
}

func (s *server) UpdateEntry(ctx context.Context, entry *pb.WeightEntry) (*pb.EntryResponse, error) {
	s.rwmu.Lock()
	defer s.rwmu.Unlock()

	// Update existin gentry by ID
	query := "UPDATE weight_entries SET uid = ?, date = ?, weight = ? WHERE id = ?"
	if err := s.session.Query(query, entry.Uid, entry.Date, entry.Weight, entry.Id).Exec(); err != nil {
		return nil, fmt.Errorf("failed to update entry: %v", err)
	}

	return &pb.EntryResponse{Message: "Success"}, nil
}

func (s *server) DeleteEntry(ctx context.Context, entry *pb.EntryID) (*pb.EntryResponse, error) {
	s.rwmu.Lock()
	defer s.rwmu.Unlock()

	// Update existin gentry by ID
	query := "DELETE FROM weight_entries WHERE id = ?"
	if err := s.session.Query(query, entry.Id).Exec(); err != nil {
		return nil, fmt.Errorf("failed to delete entry: %v", err)
	}

	return &pb.EntryResponse{Message: fmt.Sprintf("Successfully deleted entry %s", entry.Id)}, nil
}

func (s *server) AnalyzeWeight(ctx context.Context, empty *pb.Empty) (*pb.EntryResponse, error) {
	notify.StartNotifications()
	return nil, nil
}
