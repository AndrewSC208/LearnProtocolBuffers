package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/AndrewSC208/GoByExample/ProtocolBuffers/example2/workout"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type workoutServer struct {
	store map[int32]*pb.Item
}

// create a workout
func (s *workoutServer) CreateWorkout(ctx context.Context, workout *pb.WorkoutRequest) (*pb.Item, error) {
	newWorkout := &pb.Item{Id: int32(len(s.store) + 1), Workout: workout}

	s.store[newWorkout.Id] = newWorkout

	return newWorkout, nil
}

// get a workout by id
func (s *workoutServer) ReadWorkout(ctx context.Context, req *pb.ItemRequest) (*pb.Item, error) {
	return s.store[req.Id], nil
}

var (
	port = flag.Int("port", 10000, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	store := make(map[int32]*pb.Item)

	pb.RegisterWorkoutServer(s, &workoutServer{store})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
