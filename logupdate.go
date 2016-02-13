package logupdate

import (
	"fmt"
	"strings"

	"github.com/shayne/ansiescapes"
	"github.com/shayne/clicursor"
)

// LogUpdate blah blah
type LogUpdate struct {
	prevLineCount int
}

// Render blah blah
func (lu *LogUpdate) Render(content ...string) {
	clicursor.Hide()
	out := strings.Join(content, " ") + "\n"
	fmt.Print(ansiescapes.EraseLines(lu.prevLineCount) + out)
	lu.prevLineCount = len(strings.Split(out, "\n"))
}

// Clear blah blah
func (lu *LogUpdate) Clear() {
	fmt.Print(ansiescapes.EraseLines(lu.prevLineCount))
	lu.prevLineCount = 0
}

// Done blah blah
func (lu *LogUpdate) Done() {
	lu.prevLineCount = 0
	clicursor.Show()
}
