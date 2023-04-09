package temporal

import (
	"go.temporal.io/sdk/client"
	"os"
)

const (
	WorkerQueue = "WORKER_QUEUE_NAME"
)

var tpc client.Client

func InitClient() error {
	var err error

	temporalHost := os.Getenv("TEMPORAL_HOST")

	tpc, err = client.Dial(client.Options{
		HostPort:  temporalHost,
		Namespace: "default",
	})

	return err
}

func CloseClient() {
	tpc.Close()
}

func Client() client.Client {
	return tpc
}
