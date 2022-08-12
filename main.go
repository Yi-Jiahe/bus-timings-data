package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/yi-jiayu/datamall/v3"
)

func main() {
	API_ACCOUNT_KEY := os.Getenv("API_ACCOUNT_KEY")

	busStopCode := flag.String("stop", "", "5 Digit Bus Stop Code")
	service := flag.String("service", "", "Service")
	flag.Parse()

	fmt.Printf("Querying arrivals for %s at %s\n", *busStopCode, *service)

	c := datamall.NewDefaultClient(API_ACCOUNT_KEY)

	arrival, err := c.GetBusArrival(*busStopCode, *service)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v+\n", arrival)
}
