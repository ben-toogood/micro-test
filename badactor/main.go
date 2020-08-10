package main

import (
	"badactor/handler"
	"fmt"

	gostore "github.com/micro/go-micro/v3/store"
	"github.com/micro/go-micro/v3/store/cockroach"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"github.com/micro/micro/v3/service/store"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("badactor"),
		service.Version("latest"),
	)

	store.DefaultStore = cockroach.NewStore(
		gostore.Nodes("postgresql://root@cockroachdb-cluster-public.default.svc.cluster.local:26257?ssl=false"),
	)

	res, err := store.DefaultStore.List()
	fmt.Println(res, err)

	// Register handler
	srv.Handle(new(handler.Badactor))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
