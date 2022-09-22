package goroutine

import "testing"

func TestChannelSum(t *testing.T) {
	ChannelSum()
}

func TestBufferedChannel(t *testing.T) {
	BufferedChannel()
}

func TestFibonacciRun(t *testing.T) {
	FibonacciRun()
}

func TestFibonacciSelectRun(t *testing.T) {
	FibonacciSelectRun()
}

func TestDefaultSelection(t *testing.T) {
	DefaultSelection()
}

func TestSafeCounterRun(t *testing.T) {
	SafeCounterRun()
}
