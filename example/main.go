package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/shayne/logupdate"
)

var lu logupdate.LogUpdate

func startRenderer() {
	i := 0
	tick := time.Tick(80 * time.Millisecond)
	for {
		select {
		case <-tick:
			i++
			render(i)
		}
	}
}

func render(i int) {
	frames := []string{"-", "\\", "|", "/"}
	frame := frames[i%len(frames)]
	lu.Render(fmt.Sprintf(
		`
               ♥♥
          %s unicorns %s
               ♥♥
				`, frame, frame))
}

func main() {
	lu = logupdate.LogUpdate{}
	defer func() {
		lu.Clear()
		lu.Done()
	}()

	go startRenderer()

	brk := make(chan os.Signal)
	signal.Notify(brk, os.Interrupt)
	<-brk // wait for Ctrl-c
}
