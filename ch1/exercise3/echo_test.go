package main

import "testing"

func BenchmarkStdJoin(b *testing.B) {
	//10000  173328 ns/op  96 B/op  2 allocs/op
	for i := 0; i < b.N; i++ {
		stdJoin()
	}
}

func BenchmarkInEfficient(b *testing.B) {
	//30000 89208 ns/op 128 B/op 3 allocs/op
	for i := 0; i < b.N; i++ {
		inEfficient()
	}
}
