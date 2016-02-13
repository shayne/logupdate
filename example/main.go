package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/shayne/logupdate"
)

var lu logupdate.LogUpdate

func render(tick <-chan time.Time) {
	i, frames := 0, []string{"-", "\\", "|", "/"}
	for {
		select {
		case <-tick:
			i++ // increment frame
			frame := frames[i%len(frames)]
			lu.Render(fmt.Sprintf(
				`
               ♥♥
          %s unicorns %s
               ♥♥
				`, frame, frame))
		}
	}
}

func main() {
	lu = logupdate.LogUpdate{}
	tick := time.Tick(80 * time.Millisecond)
	go render(tick)

	brk := make(chan os.Signal)
	signal.Notify(brk, os.Interrupt)
	<-brk // wait for Ctrl-c

	lu.Clear()
	lu.Done()
}
