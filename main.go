package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/yi-jiayu/datamall/v3"
)

func createClient(ctx context.Context) *firestore.Client {
	// Sets your Google Cloud Platform project ID.
	projectID := "BUS_TIMINGS_DATA_TEST"

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	// Close client when done with
	// defer client.Close()
	return client
}

func main() {
	API_ACCOUNT_KEY := os.Getenv("API_ACCOUNT_KEY")

	busStopCode := flag.String("stop", "", "5 Digit Bus Stop Code")

	service := flag.String("service", "", "Service")
	flag.Parse()

	fmt.Printf("Querying arrivals for %s at %s\n", *busStopCode, *service)

	datamallClient := datamall.NewDefaultClient(API_ACCOUNT_KEY)

	arrival, err := datamallClient.GetBusArrival(*busStopCode, *service)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	firestoreClient := createClient(ctx)

	for _, service := range arrival.Services {
		buses := []datamall.ArrivingBus{
			service.NextBus,
			service.NextBus2,
			service.NextBus3,
		}
		fmt.Printf("Service %s\n", service.ServiceNo)

		for _, bus := range buses {
			fmt.Printf("%s\n", bus.EstimatedArrival)

			_, _, err = firestoreClient.Collection("observations").Add(ctx, map[string]interface{}{
				"service": service.ServiceNo,
				"arrival": bus.EstimatedArrival,
				"load":    bus.Load,
				"type":    bus.Type,
			})
			if err != nil {
				log.Fatalf("%v", err)
			}
		}
	}

}
