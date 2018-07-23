package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/AndrewSC208/GoByExample/ProtocolBuffers/example2/workout"
	"google.golang.org/grpc"
)

var (
	serverAddr = flag.String("server-addr", "localhost:10000", "The server address in the formate of host:port")
)

// select a workout by Id
func createWorkout(ctx context.Context, client pb.WorkoutClient) {
	res, err := client.CreateWorkout(ctx, &pb.WorkoutRequest{Title: "First Workout", Type: "Push ups", StartDateTime: int64(23424), Duration: int64(423234)})
	if err != nil {
		log.Fatalf("could not create workout: %v", err)
	}
	log.Printf("Workout: %s", res)
}

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewWorkoutClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// createWorkout
	createWorkout(ctx, client)

}
