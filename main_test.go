package main

import (
	"context"
	"testing"
	"time"
)

func BenchmarkConcorrentFetch(b *testing.B) {
	for b.Loop() {
		cuncurr(context.Background(), time.Millisecond*15)
	}
}
