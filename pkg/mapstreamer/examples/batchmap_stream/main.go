package main

import (
	"context"
	"hash/fnv"
	"log"

	"github.com/numaproj/numaflow-go/pkg/mapstreamer"
)

// FlatMap is a MapStreamer that split the input message into multiple messages and stream them.
type FlatMap struct {
}

func (f *FlatMap) MapStream(ctx context.Context, keys []string, d mapstreamer.Datum, messageCh chan<- mapstreamer.Message) {

	defer close(messageCh)
	msg := d.Value()
	_ = d.EventTime() // Event time is available
	_ = d.Watermark() // Watermark is available
	// Split the msg into an array with comma.
	log.Printf("input = %s", msg)
	summer := fnv.New32a()
	summer.Write(msg)
	val := summer.Sum32()

	log.Printf("hash = %d", val)

	if val%2 == 0 {
		log.Print("even")
		messageCh <- mapstreamer.NewMessage(msg).WithKeys([]string{"even"}).WithTags([]string{"even-tag"})
	} else {
		log.Print("odd")
		messageCh <- mapstreamer.NewMessage(msg).WithKeys([]string{"odd"}).WithTags([]string{"odd-tag"})
	}
}

func (f *FlatMap) MapStreamBatch(ctx context.Context, datumCh <-chan mapstreamer.Datum, messageCh chan<- mapstreamer.Message) {
	count := 0

	for d := range datumCh {
		count += 1
		log.Printf("FN: HANDLE IT %d", count)

		msg := d.Value()
		// _ = d.EventTime() // Event time is available
		// _ = d.Watermark() // Watermark is available
		// // Split the msg into an array with comma.
		log.Printf("FN: input = %s", msg)
		summer := fnv.New32a()
		summer.Write(msg)
		val := summer.Sum32()

		// log.Printf("hash = %d", val)

		if val%2 == 0 {
			log.Print("even")
			messageCh <- mapstreamer.NewMessage(msg).WithKeys([]string{"even"}).WithTags([]string{"even-tag"})
		} else {
			log.Print("odd")
			messageCh <- mapstreamer.NewMessage(msg).WithKeys([]string{"odd"}).WithTags([]string{"odd-tag"})
		}
	}
	log.Printf("FN: Now Return from handler...")
}

// strs := strings.Split(string(msg), ",")
// for _, s := range strs {
// 	messageCh <- mapstreamer.NewMessage([]byte(s))
// }

func main() {
	err := mapstreamer.NewServer(&FlatMap{}).Start(context.Background())
	if err != nil {
		log.Panic("Failed to start map stream function server: ", err)
	}
}
