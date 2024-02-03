package main

import "testing"

// go test -bench=ReadOnly_100_100M_ -benchtime=10s -run=NONE -v ./
// result: 30798691 ns/op = 30.79ms
func BenchmarkReadOnly_100_100M_(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := ReadAllFilesDiscard("tmp_100_100M", 100)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// go test -bench=ReadOnly_100_1G_ -benchtime=10s -run=NONE -v ./
// result: 268767161 ns/op = 268.76ms
func BenchmarkReadOnly_100_1G_(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := ReadAllFilesDiscard("tmp_100_1G", 100)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// go test -bench=ReadOnly_100_5G_ -benchtime=10s -run=NONE -v ./
// result: 1455964925 ns/op = 1455.96ms
func BenchmarkReadOnly_100_5G_(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := ReadAllFilesDiscard("tmp_100_5G", 100)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// go test -bench=ReadOnly_100_10G_ -benchtime=10s -run=NONE -v ./
// result: 4093329008 ns/op = 4093.32ms
func BenchmarkReadOnly_100_10G_(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := ReadAllFilesDiscard("tmp_100_10G", 100)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// go test -bench=ReadOnly_10000_5G_ -benchtime=10s -run=NONE -v ./
// result: 13858609664 ns/op = 13.85s
// speed: 369.67 M/s
func BenchmarkReadOnly_10000_5G_(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := ReadAllFilesDiscard("tmp_10000_5G", 10000)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// go test -bench=DiscardConcurrent_10000_5G_1000Ch_20Go_ -benchtime=10s -run=NONE -v ./
// result: 1877701836 ns/op = 1877.70ms = 1.87s
// read speed: 5G/1.87s = 2737 M/s
func BenchmarkDiscardConcurrent_10000_5G_1000Ch_20Go_(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := ReadAllFilesDiscardConcurrent("tmp_10000_5G", 10000, 1000, 20)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// go test -bench=DiscardConcurrent_10000_5G_1000Ch_40Go_ -benchtime=10s -run=NONE -v ./
// result: 1632719568 ns/op = 1632.72ms = 1.63s
// read speed: 5G/1.63s = 3141 M/s
func BenchmarkDiscardConcurrent_10000_5G_1000Ch_40Go_(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := ReadAllFilesDiscardConcurrent("tmp_10000_5G", 10000, 1000, 40)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// go test -bench=DiscardConcurrent_10000_5G_1000Ch_50Go_ -benchtime=10s -run=NONE -v ./
// result: 1718289015 ns/op = 1.71s
// read speed: 5G/1.71s = 2994 M/s
func BenchmarkDiscardConcurrent_10000_5G_1000Ch_50Go_(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := ReadAllFilesDiscardConcurrent("tmp_10000_5G", 10000, 1000, 50)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// go test -bench=DiscardConcurrent_10000_5G_1000Ch_100Go_ -benchtime=10s -run=NONE -v ./
// result: 1696218382 ns/op = 1.69s
// read speed: 5G/16.13s = 3029 M/s
func BenchmarkDiscardConcurrent_10000_5G_1000Ch_100Go_(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := ReadAllFilesDiscardConcurrent("tmp_10000_5G", 10000, 1000, 100)
		if err != nil {
			b.Fatal(err)
		}
	}
}
