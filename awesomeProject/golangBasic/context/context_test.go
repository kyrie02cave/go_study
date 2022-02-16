package context

import "testing"

func TestBasic(t *testing.T) {
	Basic()
}

func TestGlobalVariable(t *testing.T) {
	GlobalVariable()
}

func TestChan(t *testing.T) {
	Chan()
}

func TestStd(t *testing.T) {
	Std()
}

func TestWithCancelDemo(t *testing.T) {
	WithCancelDemo()
}

func TestWithDeadlineDemo(t *testing.T) {
	WithDeadlineDemo()
}

func TestWithTimeoutDemo(t *testing.T) {
	WithTimeoutDemo()
}

func TestWithValueDemo(t *testing.T) {
	WithValueDemo()
}

func TestServerContext(t *testing.T) {
	ServerContext()
}

func TestClientContext(t *testing.T) {
	ClientContext()
}