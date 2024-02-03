package main

import "testing"

// go test -bench=Write_10000_5G_ -benchtime=10s -run=NONE -v ./
// result: 3666011434 ns/op = 3.66s
// write speed: 5G/3.66s = 1398 M/s
func BenchmarkWrite_10000_5G_(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := GenFiles("tmp_10000_5G", 10000, 5*G)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// go test -bench=WriteConcurrent_10000_5G_1000Ch_1Go_ -benchtime=10s -run=NONE -v ./
// result: 3083167764 ns/op = 3.08s
// write speed: 5G/3.08ss = 1662 M/s
func BenchmarkWriteConcurrent_10000_5G_1000Ch_1Go_(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := GenFilesConcurrently("tmp_10000_5G", 10000, 5*G, 1000, 1)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// go test -bench=WriteConcurrent_10000_5G_1000Ch_20Go_ -benchtime=10s -run=NONE -v ./
// result: 2842253050 ns/op = 2.84s
// write speed: 5G/1.63s = 1802 M/s
func BenchmarkWriteConcurrent_10000_5G_1000Ch_20Go_(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := GenFilesConcurrently("tmp_10000_5G", 10000, 5*G, 1000, 20)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// go test -bench=WriteConcurrent_10000_5G_1000Ch_50Go_ -benchtime=10s -run=NONE -v ./
// result: 2842253050 ns/op = 2.84s
// write speed: 5G/1.63s = 1802 M/s
func BenchmarkWriteConcurrent_10000_5G_1000Ch_50Go_(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := GenFilesConcurrently("tmp_10000_5G", 10000, 5*G, 1000, 50)
		if err != nil {
			b.Fatal(err)
		}
	}
}
