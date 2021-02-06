package cliargument

import (
	"fmt"
	"github.com/sadmansakib/gractical/common"
	"os"
	"sync"
)

//CliRunner struct is the definition of CLI Runner process
type CliRunner struct {
	Run         int    `csv:"Run"`
	Title       string `csv:"Title"`
	Message1    string `csv:"Message 1"`
	Message2    string `csv:"Message 2"`
	StreamDelay int    `csv:"Stream Delay"`
	RunTimes    int    `csv:"Run Times"`
}

//GenerateStreamer Creates an instance of CLI Streamer Process
func (cliRunner *CliRunner) GenerateStreamer() CliStreamer {
	return CliStreamer{
		Title:       cliRunner.Title,
		Message1:    cliRunner.Message1,
		Message2:    cliRunner.Message2,
		StreamDelay: cliRunner.StreamDelay,
		RunTimes:    cliRunner.RunTimes,
	}
}

//Execute function takes all runners and run all necessary operations
//creates a file named result.log where we are going to dump our outputs
//creates a wait group and channel to synchronize instances of CLI streamer
//traverses through all cli runners and creates instances of cli streamer on goroutine
//and attaches value to channel.
//Creates a mutex instance that locks goroutine execution for printing output
//and writing log file
func Execute(cliRunners []CliRunner) {
	f, err := os.Create("./log/result.log")
	common.Check(err)

	defer f.Close()

	var wg sync.WaitGroup

	var mutex = &sync.Mutex{}

	var ch = make(chan string)

	for _, runner := range cliRunners {
		wg.Add(1)
		go runner.GenerateStreamer().Stream(&wg, ch)
	}

	go log(f, ch, &wg, mutex)

	wg.Wait()
	close(ch)
}

//log function prints and writes output to logfile
func log(file *os.File, ch chan string, wg *sync.WaitGroup, mutex *sync.Mutex) {
	mutex.Lock()
	for msg := range ch {
		fmt.Print(msg)
		_, err := file.WriteString(msg)
		common.Check(err)
		err = file.Sync()
		common.Check(err)
	}
	mutex.Unlock()
	wg.Done()
}
