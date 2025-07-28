package main

import (
	"errors"
	"testing"
)

type fakeRunner struct{ err error }

func (f fakeRunner) Run(addr ...string) error {
	return f.err
}

func TestRunServerError(t *testing.T) {
	wantErr := errors.New("listen error")
	err := runServer(fakeRunner{err: wantErr}, "1234")
	if !errors.Is(err, wantErr) {
		t.Fatalf("expected %v, got %v", wantErr, err)
	}
}
