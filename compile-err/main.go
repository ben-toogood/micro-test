package main

import (
	"compile-err/handler"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("compile-err"),
		service.Version("latest"),
	)

	// Register handler
	srv.Handle(new(handler.CompileErr))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
