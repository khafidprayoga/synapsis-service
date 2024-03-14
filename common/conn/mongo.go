package conn

import (
	"context"
	"fmt"
	codecs "github.com/amsokol/mongo-go-driver-protobuf"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"os"
)

func MongoConnect(ctx context.Context) (*mongo.Client, error) {
	connectionString := fmt.Sprintf(
		"mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority",
		os.Getenv("MONGODB_USERDB"),
		os.Getenv("MONGODB_PASSWORDDB"),
		os.Getenv("MONGODB_HOST"),
	)
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	reg := codecs.Register(bson.NewRegistryBuilder()).Build()

	monitor := &event.CommandMonitor{
		Started: func(_ context.Context, e *event.CommandStartedEvent) {
			fmt.Println(e.Command)
		},
		Succeeded: func(_ context.Context, e *event.CommandSucceededEvent) {
			fmt.Println(e.Reply)
		},
		Failed: func(_ context.Context, e *event.CommandFailedEvent) {
			fmt.Println(e.Failure)
		},
	}

	clientOptions := options.
		Client().
		ApplyURI(connectionString).
		SetServerAPIOptions(serverAPI).
		SetMonitor(monitor).
		SetRegistry(reg)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// Ping the MongoDB server to verify the connection
	if err := client.Database(os.Getenv("MONGODB_DATABASE")).
		RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		fmt.Println("Failed to connect to MongoDB:", err)
	}

	if err != nil {
		return nil, err
	}

	return client, nil
}
