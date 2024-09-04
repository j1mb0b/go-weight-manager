package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"

	"github.com/gocql/gocql"
	wmServer "github.com/j1mb0b/go-weight-manager/cmd/server"
	pb "github.com/j1mb0b/go-weight-manager/proto"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	goroutineCount = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "goroutine_count",
		Help: "Number of Go routines",
	})
)

func init() {
	prometheus.MustRegister(goroutineCount)
}

func updateMetrics() {
	goroutineCount.Set(float64(runtime.NumGoroutine()))
}

func metricsHandler() {
	http.Handle("/metrics", promhttp.Handler())
	server := &http.Server{Addr: ":2112"}
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Metrics server failed: %v", err)
	}
}

func startGRPCServer(ctx context.Context) (*grpc.Server, net.Listener) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	server := wmServer.NewServer()

	pb.RegisterWMServiceServer(grpcServer, server.(pb.WMServiceServer))
	reflection.Register(grpcServer)

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Println("Stopping metric updates")
				return
			default:
				updateMetrics()
				time.Sleep(10 * time.Second)
			}
		}
	}()

	return grpcServer, lis
}

func setupCluster() *gocql.Session {
	cluster := gocql.NewCluster("127.0.0.1") // replace with your Cassandra IP address
	cluster.Keyspace = "weight_manager"      // replace with your keyspace
	cluster.Consistency = gocql.Quorum

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal("Unable to connect to Cassandra: ", err)
	}

	// Create table to match protobuf schema
	createTableQuery := `CREATE TABLE IF NOT EXISTS weight_entries (
        id TEXT PRIMARY KEY,
        uid INT,
        date TEXT,
        weight FLOAT
    )`
	if err := session.Query(createTableQuery).Exec(); err != nil {
		log.Fatal("Failed to create table: ", err)
	}

	return session
}

func main() {
	// Setup signal handling for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start gRPC server
	grpcServer, lis := startGRPCServer(ctx)

	// Start metrics HTTP server
	go metricsHandler()

	// Setup Cassandra cluster
	session := setupCluster()
	defer session.Close()

	// Handle shutdown signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, os.Kill)
	go func() {
		sig := <-sigChan
		log.Printf("Received signal: %v", sig)
		cancel()
		grpcServer.GracefulStop()
	}()

	// Start gRPC server
	log.Println("Starting gRPC server...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}

	log.Println("Server stopped gracefully")
}
