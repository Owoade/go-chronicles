package main

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func BenchmarkSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", time.Now().UnixMilli())
	}
}

func BenchmarkItoa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strconv.Itoa(int(time.Now().UnixMilli()))
	}
}
