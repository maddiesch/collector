package output

import (
	"fmt"
	"io"
	"strings"
	"sync"

	"github.com/fatih/color"
)

type Output struct {
	mu sync.RWMutex

	stdout io.Writer
	stderr io.Writer

	errC     printFn
	successC printFn
	warnC    printFn
}

type printFn func(io.Writer, ...any)

func New(stdout, stderr io.Writer) *Output {
	return &Output{
		stdout:   stdout,
		stderr:   stderr,
		errC:     color.New(color.FgRed).Add(color.Bold, color.Underline).FprintFunc(),
		successC: color.New(color.FgGreen).FprintFunc(),
		warnC:    color.New(color.FgYellow).Add(color.Underline).FprintFunc(),
	}
}

func (o *Output) SetErr(e io.Writer) {
	o.mu.Lock()
	defer o.mu.Unlock()

	o.stderr = e
}

func (o *Output) SetOut(e io.Writer) {
	o.mu.Lock()
	defer o.mu.Unlock()

	o.stdout = e
}

// Print writes the given formatted string to the stdout.
func (o *Output) Print(msg string, a ...any) {
	o.mu.RLock()
	defer o.mu.RUnlock()

	o.print(o.stdout, msg, a...)
}

func (o *Output) Println(a ...any) {
	o.mu.RLock()
	defer o.mu.RUnlock()

	o.print(o.stdout, fmt.Sprintln(a...))
}

// PrintE write the given formatted string to the stderr.
func (o *Output) PrintE(msg string, a ...any) {
	o.mu.RLock()
	defer o.mu.RUnlock()

	o.print(o.stderr, msg, a...)
}

func (o *Output) PrintEln(a ...any) {
	o.mu.RLock()
	defer o.mu.RUnlock()

	o.print(o.stderr, fmt.Sprintln(a...))
}

func (o *Output) Warn(msg string, a ...any) {
	o.mu.RLock()
	defer o.mu.RUnlock()

	o.printFn(o.stderr, o.warnC, msg, a...)
}

func (o *Output) Error(format string, a ...any) {
	o.mu.RLock()
	defer o.mu.RUnlock()

	o.printFn(o.stderr, o.errC, format, a...)
}

func (o *Output) Success(msg string, a ...any) {
	o.mu.RLock()
	defer o.mu.RUnlock()

	o.printFn(o.stderr, o.successC, msg, a...)
}

func (o *Output) prepare(msg string, a ...any) string {
	message := fmt.Sprintf(msg, a...)
	if !strings.HasSuffix(message, "\n") {
		message += "\n"
	}
	return message
}

func (o *Output) print(w io.Writer, msg string, args ...any) {
	fn := func(w io.Writer, a ...any) {
		fmt.Fprint(w, a...)
	}

	o.printFn(w, fn, msg, args...)
}

func (o *Output) printFn(w io.Writer, fn printFn, msg string, a ...any) {
	message := o.prepare(msg, a...)

	fn(w, message)
}
