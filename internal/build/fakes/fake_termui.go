package fakes

import (
	"io"

	"github.com/buildpacks/pack/internal/build"
	"github.com/buildpacks/pack/internal/container"
)

type FakeTermui struct {
	handler container.Handler
}

func NewFakeTermui(handler container.Handler) *FakeTermui {
	return &FakeTermui{
		handler: handler,
	}
}

func (f *FakeTermui) Run(funk func()) error {
	return nil
}

func (f *FakeTermui) Handler() container.Handler {
	return f.handler
}

func WithTermui(screen build.Termui) func(*build.LifecycleOptions) {
	return func(opts *build.LifecycleOptions) {
		opts.Interactive = true
		opts.Termui = screen
	}
}

func (f *FakeTermui) Debug(msg string) {
	// not implemented
}

func (f *FakeTermui) Debugf(fmt string, v ...interface{}) {
	// not implemented
}

func (f *FakeTermui) Info(msg string) {
	// not implemented
}

func (f *FakeTermui) Infof(fmt string, v ...interface{}) {
	// not implemented
}

func (f *FakeTermui) Warn(msg string) {
	// not implemented
}

func (f *FakeTermui) Warnf(fmt string, v ...interface{}) {
	// not implemented
}

func (f *FakeTermui) Error(msg string) {
	// not implemented
}

func (f *FakeTermui) Errorf(fmt string, v ...interface{}) {
	// not implemented
}

func (f *FakeTermui) Writer() io.Writer {
	// not implemented
	return nil
}

func (f *FakeTermui) IsVerbose() bool {
	// not implemented
	return false
}
