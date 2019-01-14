package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/quii/try-grpc"
	"github.com/quii/try-grpc/fridge"
	"google.golang.org/grpc"
	"log"
	"os"
	"strings"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "missing subcommand: list or add")
		os.Exit(1)
	}

	conn, err := grpc.Dial(try_grpc.FridgeAddress, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect to %s, %v", try_grpc.FridgeAddress, err)
	}

	defer conn.Close()

	client := fridge.NewFridgeServiceClient(conn)

	switch cmd := flag.Arg(0); cmd {
	case "list":
		listItems(client)
	case "add":
		addItem(client, strings.Join(flag.Args()[1:], " "))
	default:
		log.Printf("unknown command %s", cmd)
	}

}

func addItem(client fridge.FridgeServiceClient, item string) {
	response, err := client.AddItem(context.Background(), &fridge.Item{Name: item})
	if err != nil {
		log.Fatalf("problem adding item %s %v", item, err)
	}
	log.Printf("Here's the stuff in the fridge %+v", response.Items)
}

func listItems(client fridge.FridgeServiceClient) {
	request := fridge.GetItemsRequest{}
	response, err := client.GetItems(context.Background(), &request)
	if err != nil {
		log.Fatalf("problem getting items %v", err)
	}
	log.Printf("Here's the stuff in the fridge %+v", response.Items)
}
