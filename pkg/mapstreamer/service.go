package mapstreamer

import (
	"context"
	"io"
	"log"
	"sync"

	"google.golang.org/protobuf/types/known/emptypb"

	mapstreampb "github.com/numaproj/numaflow-go/pkg/apis/proto/mapstream/v1"
)

const (
	uds                   = "unix"
	defaultMaxMessageSize = 1024 * 1024 * 64
	address               = "/var/run/numaflow/mapstream.sock"
	serverInfoFilePath    = "/var/run/numaflow/mapstreamer-server-info"
)

// Service implements the proto gen server interface and contains the map
// streaming function.
type Service struct {
	mapstreampb.UnimplementedMapStreamServer

	MapperStream MapStreamer
}

// IsReady returns true to indicate the gRPC connection is ready.
func (fs *Service) IsReady(context.Context, *emptypb.Empty) (*mapstreampb.ReadyResponse, error) {
	return &mapstreampb.ReadyResponse{Ready: true}, nil
}

// MapStreamFn applies a function to each request element and streams the results back.
func (fs *Service) MapStreamFn(d *mapstreampb.MapStreamRequest, stream mapstreampb.MapStream_MapStreamFnServer) error {
	var hd = NewHandlerDatum(d.GetValue(), d.GetEventTime().AsTime(), d.GetWatermark().AsTime(), d.GetHeaders())
	ctx := stream.Context()
	messageCh := make(chan Message)

	log.Printf("MDW: I'm in MapStreamFn")

	done := make(chan bool)
	go func() {
		fs.MapperStream.MapStream(ctx, d.GetKeys(), hd, messageCh)
		done <- true
	}()
	finished := false
	for {
		select {
		case <-done:
			finished = true
		case message, ok := <-messageCh:
			if !ok {
				// Channel already closed, not closing again.
				return nil
			}
			element := &mapstreampb.MapStreamResponse{
				Result: &mapstreampb.MapStreamResponse_Result{
					Keys:  message.Keys(),
					Value: message.Value(),
					Tags:  message.Tags(),
				},
			}
			err := stream.Send(element)
			// the error here is returned by stream.Send() which is already a gRPC error
			if err != nil {
				// Channel may or may not be closed, as we are not sure leave it to GC.
				return err
			}
		default:
			if finished {
				close(messageCh)
				return nil
			}
		}
	}
}

func (fs *Service) MapStreamBatchFn(stream mapstreampb.MapStream_MapStreamBatchFnServer) error {
	log.Printf("MDW: Enter MapStreamBatchFn...")

	var (
		// resultList    []*sinkpb.SinkResponse_Result  // MDW: Want the output STREAM
		wg            sync.WaitGroup
		datumStreamCh = make(chan Datum)
		ctx           = stream.Context()
	)

	// Make call to kick off stream handling
	messageCh := make(chan Message)
	done := make(chan bool)
	go func() {
		fs.MapperStream.MapStreamBatch(ctx, datumStreamCh, messageCh)
		done <- true
	}()

	// Read messages and push to read channel
	go func() {
		for {
			d, err := stream.Recv()
			if err == io.EOF {
				close(datumStreamCh)
				log.Printf("MDW: Stream closed")
				return
			}
			if err != nil {
				close(datumStreamCh)
				log.Printf("MDW: Error maybe -- %s", err)
				// TODO: research on gRPC errors and revisit the error handler
				return
			}
			var hd = &handlerDatum{
				value:     d.GetValue(),
				eventTime: d.GetEventTime().AsTime(),
				watermark: d.GetWatermark().AsTime(),
				headers:   d.GetHeaders(),
			}
			log.Printf("MDW: Send Datum")
			datumStreamCh <- hd
		}
	}()

	wg.Wait()

	// // var hd = NewHandlerDatum(d.GetValue(), d.GetEventTime().AsTime(), d.GetWatermark().AsTime(), d.GetHeaders())
	// // ctx := stream.Context()
	// // messageCh := make(chan Message)

	finished := false
	for {
		select {
		case <-done:
			finished = true
		case message, ok := <-messageCh:
			if !ok {
				// Channel already closed, not closing again.
				return nil
			}
			log.Printf("MDWR: Got message")
			element := &mapstreampb.MapStreamResponse{
				Result: &mapstreampb.MapStreamResponse_Result{
					Keys:  message.Keys(),
					Value: message.Value(),
					Tags:  message.Tags(),
				},
			}
			err := stream.Send(element)
			// the error here is returned by stream.Send() which is already a gRPC error
			if err != nil {
				log.Printf("MDWR: Got error %s", err)
				// Channel may or may not be closed, as we are not sure leave it to GC.
				return err
			}
		default:
			if finished {
				close(messageCh)
				return nil
			}
		}
	}

	log.Printf("MDW: Leaving MapStreamBatchFn...")
	return nil
}
