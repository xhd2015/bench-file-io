package main

import "testing"

// go test -bench=Copy_100_100M_ -benchtime=10s -run=NONE -v ./
// result: - ns/op
func BenchmarkCopy_100_100M_(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := ReadAllFilesToAnotherDir("tmp_100_100M", 100)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// go test -bench=Copy_100_1G_ -benchtime=10s -run=NONE -v ./
// result: - ns/op
func BenchmarkCopy_100_1G_(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := ReadAllFilesToAnotherDir("tmp_100_1G", 100)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// go test -bench=Copy_100_5G_ -benchtime=10s -run=NONE -v ./
// result: 43913857865 ns/op = 43913.85ms = 43s
//
//	with 4M buffer      6279975122 ns/op = 6279.97ms = 6.27s
func BenchmarkCopy_100_5G_(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := ReadAllFilesToAnotherDir("tmp_100_5G", 100)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// prepare files: go run ./ genFiles tmp_10000_5G 10000 5368709120
//
// go test -bench=Copy_10000_5G_ -benchtime=10s -run=NONE -v ./
//
//	with 4M buffer      22325023874 ns/op = 22325.02ms = 22.32s
func BenchmarkCopy_10000_5G_(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := ReadAllFilesToAnotherDir("tmp_10000_5G", 10000)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// go test -bench=Copy_100_10G_ -benchtime=10s -run=NONE -v ./
// result: 54434134599 ns/op = 54434.32ms = 54s
func BenchmarkCopy_100_10G_(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := ReadAllFilesToAnotherDir("tmp_100_10G", 100)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// go test -bench=CopyConcurrent_100_100M_ -benchtime=10s -run=NONE -v ./
// result: - ns/op
func BenchmarkCopyConcurrent_100_100M_(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := ReadAllFilesToAnotherDir("tmp_100_100M", 100)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// go test -bench=CopyConcurrent_100_1G_ -benchtime=10s -run=NONE -v ./
// result: - ns/op
func BenchmarkCopyConcurrent_100_1G_(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := ReadAllFilesToAnotherDir("tmp_100_1G", 100)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// go test -bench=CopyConcurrent_100_5G_ -benchtime=10s -run=NONE -v ./
// result: 43913857865 ns/op = 43913.85ms = 43s
func BenchmarkCopyConcurrent_100_5G_(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := ReadAllFilesToAnotherDir("tmp_100_5G", 100)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// go test -bench=CopyConcurrent_100_10G_ -benchtime=10s -run=NONE -v ./
// result: 54434134599 ns/op = 54434.32ms = 54s
func BenchmarkCopyConcurrent_100_10G_(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := ReadAllFilesToAnotherDir("tmp_100_10G", 100)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// go test -bench=CopyConcurrent_10000_5G_100Ch_100Go_ -benchtime=10s -run=NONE -v ./
//
//	8455526032 ns/op = 8455.52ms = 8.4s
func BenchmarkCopyConcurrent_10000_5G_100Ch_100Go_(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := ReadAllFilesToAnotherDirConcurrently("tmp_10000_5G", 10000, 100, 100)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// go test -bench=CopyConcurrent_100_5G_1000Ch_1Go_ -benchtime=10s -run=NONE -v ./
//
//	5986943060 ns/op = 5.98s
//
// 856 M/s
func BenchmarkCopyConcurrent_100_5G_1000Ch_1Go_(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := ReadAllFilesToAnotherDirConcurrently("tmp_100_5G", 100, 1000, 1)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// go test -bench=CopyConcurrent_1000_5G_1000Ch_1Go_ -benchtime=10s -run=NONE -v ./
//
//	11199865775 ns/op = 11.19s
//
// 457 M/s
func BenchmarkCopyConcurrent_1000_5G_1000Ch_1Go_(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := ReadAllFilesToAnotherDirConcurrently("tmp_1000_5G", 1000, 1000, 1)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// go test -bench=CopyConcurrent_1000_5G_1000Ch_50Go_ -benchtime=10s -run=NONE -v ./
//
//	5393256305 ns/op = 5.39s
//
// 949 M/s
func BenchmarkCopyConcurrent_1000_5G_1000Ch_50Go_(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := ReadAllFilesToAnotherDirConcurrently("tmp_1000_5G", 1000, 1000, 50)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// go test -bench=CopyConcurrent_10000_5G_1000Ch_1Go_ -benchtime=10s -run=NONE -v ./
//
//	25436430809 ns/op = 25.43s
//
// 201 M/s
func BenchmarkCopyConcurrent_10000_5G_1000Ch_1Go_(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := ReadAllFilesToAnotherDirConcurrently("tmp_10000_5G", 10000, 1000, 1)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// go test -bench=CopyConcurrent_10000_5G_1000Ch_20Go_ -benchtime=10s -run=NONE -v ./
//
//	7237071884 ns/op = 7.23s
func BenchmarkCopyConcurrent_10000_5G_1000Ch_20Go_(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := ReadAllFilesToAnotherDirConcurrently("tmp_10000_5G", 10000, 1000, 20)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// go test -bench=CopyConcurrent_10000_5G_100Ch_50Go_ -benchtime=10s -run=NONE -v ./
//
//	6479135306 ns/op = 6.47s
func BenchmarkCopyConcurrent_10000_5G_100Ch_50Go_(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := ReadAllFilesToAnotherDirConcurrently("tmp_10000_5G", 10000, 100, 50)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// go test -bench=CopyConcurrent_10000_5G_1000Ch_50Go_ -benchtime=10s -run=NONE -v ./
//
//	6655165974 ns/op = 6.65s
//
// 769 M/s
func BenchmarkCopyConcurrent_10000_5G_1000Ch_50Go_(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := ReadAllFilesToAnotherDirConcurrently("tmp_10000_5G", 10000, 1000, 50)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// go test -bench=CopyConcurrent_10000_5G_1000Ch_100Go_ -benchtime=10s -run=NONE -v ./
//
//	8141753560 ns/op = 8141.75ms = 8.14s
func BenchmarkCopyConcurrent_10000_5G_1000Ch_100Go_(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := ReadAllFilesToAnotherDirConcurrently("tmp_10000_5G", 10000, 1000, 100)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// go test -bench=CopyConcurrent_10000_5G_1000Ch_200Go_ -benchtime=10s -run=NONE -v ./
//
//	12788166332 ns/op = 12788.16ms = 12.78s
func BenchmarkCopyConcurrent_10000_5G_1000Ch_200Go_(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := ReadAllFilesToAnotherDirConcurrently("tmp_10000_5G", 10000, 1000, 200)
		if err != nil {
			b.Fatal(err)
		}
	}
}
