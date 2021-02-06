package cli_argument

import (
	"fmt"
	"sync"
	"time"
)

//CliStreamer struct is the definition of CLI Streamer process
type CliStreamer struct {
	Title       string `csv:"Title"`
	Message1    string `csv:"Message 1"`
	Message2    string `csv:"Message 2"`
	StreamDelay int    `csv:"Stream Delay"`
	RunTimes    int    `csv:"Run Times"`
}

//Stream Streams Creates messages to pass them to channel created by CLI runner
func (streamer CliStreamer) Stream(wg *sync.WaitGroup, ch chan string) {
	for i := 0; i < streamer.RunTimes; i++ {
		ch <- message(i+1,streamer.Message1)
		time.Sleep(time.Duration(streamer.StreamDelay) * time.Second)
		ch <- message(i+1,streamer.Message2)
	}
	wg.Done()
	return
}

func message(invoke int, message string) string {
	return fmt.Sprintf("CLI Invoke %v -> %s \n", invoke, message)
}
