package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	badactor "badactor/proto"
)

type Badactor struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Badactor) Call(ctx context.Context, req *badactor.Request, rsp *badactor.Response) error {
	log.Info("Received Badactor.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Badactor) Stream(ctx context.Context, req *badactor.StreamingRequest, stream badactor.Badactor_StreamStream) error {
	log.Infof("Received Badactor.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&badactor.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Badactor) PingPong(ctx context.Context, stream badactor.Badactor_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&badactor.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
