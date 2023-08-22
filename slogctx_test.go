package slogctx_test

import (
	"context"
	"log/slog"
	"testing"

	"github.com/uhthomas/slogctx"
)

func TestFrom(t *testing.T) {
	want := slog.New(slog.NewTextHandler(nil, nil))
	ctx := slogctx.With(context.Background(), want)
	if got := slogctx.From(ctx); got != want {
		t.Fatalf("got %v != want %v", got, want)
	}

	t.Run("default", func(t *testing.T) {
		if slogctx.From(context.Background()) != slog.Default() {
			t.Fatal("expected default logger")
		}
	})
}
