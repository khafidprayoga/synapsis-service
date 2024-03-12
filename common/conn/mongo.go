package conn

import (
	"context"
	"fmt"
	codecs "github.com/amsokol/mongo-go-driver-protobuf"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"golang.org/x/net/proxy"
	"net"
	"os"
)

type myDialer struct {
	dialer proxy.Dialer
}

func (d myDialer) DialContext(ctx context.Context, network, address string) (net.Conn, error) {
	if os.Getenv("SERVE_MODE") == "localhost" {
		auth := proxy.Auth{
			User:     os.Getenv("MONGODB_USERPROXY"),
			Password: os.Getenv("MONGODB_PASSWORDPROXY"),
		}
		dialer, err := proxy.SOCKS5("tcp", os.Getenv("MONGODB_HOSTPROXY"), &auth, proxy.Direct)
		if err != nil {
			return nil, err
		}
		return dialer.Dial(network, address)
	}

	return d.dialer.Dial(network, address)
}

func MongoConnect(ctx context.Context) (*mongo.Client, error) {
	connectionString := fmt.Sprintf(
		"mongodb+srv://%v:%v@%v/?retryWrites=true&w=majority",
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
		SetDialer(myDialer{dialer: proxy.Direct}).
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

	if err := client.Database(os.Getenv("MONGODB_DATABASE")).
		RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		fmt.Println("Failed to connect to MongoDB:", err)
	}
	if err != nil {
		return nil, err
	}

	return client, nil
}
