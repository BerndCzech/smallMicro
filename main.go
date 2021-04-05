package main

import (
	"smallMicro/handler"
	pb "smallMicro/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("smallmicro"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterSmallMicroHandler(srv.Server(), new(handler.SmallMicro))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
