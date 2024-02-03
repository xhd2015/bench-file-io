# Benchmark file io operations
Benchmark file io operations including: read, write and copy.

# Conclusion
Read > Write > Copy

Increase goroutine num to for example 50 would increase io throughput.

# Read file benchmark
Read file benchmark does only read files, and discard file's content, to measure only read speed.

Code: [benchmark_read_only_test.go](benchmark_read_only_test.go)

|GoNum|Buffer|Files|Total Size|Cost|Speed|
|-|-|-|-|-|-|
|1|8K|10000|5G|15.54s|369 M/s|
|20|8K|10000|5G|1.87s|2737 M/s|
|40|8K|10000|5G|1.63s|3141 M/s|
|50|8K|10000|5G|1.71s|2994 M/s|
|100|8K|10000|5G|1.69s|3029 M/s|

Note that io.Discard always uses buffer with size 8192.

# Write file benchmark
Write file benchmark does only write the file, to measure only write speed

Code: [benchmark_write_only_test.go](benchmark_write_only_test.go)

|GoNum|Files|Total Size|Cost|Speed|
|-|-|-|-|-|
|1|10000|5G|3.08s|1662 M/s| 
|20|10000|5G|2.84s|1802 M/s|
|50|10000|5G|3.25s|1575 M/s|

Note that for write, there is no buffer invovled, all data are continueously committed to syscall write at it's max speed.

# Copy file benchmark
When doing copy, the main bottleneck should be write.However since read and write are both IO operations,so there should be a compromise in between

Code: [benchmark_copy_test.go](benchmark_copy_test.go)

|GoNum|Buffer|Ch|Files|Total Size|Cost|Speed|
|-|-|-|-|-|-|-|
|1|4M|1000|100|5G|5.98s|856 M/s|
|1|4M|1000|1000|5G|11.19s|457 M/s|
|1|4M|1000|10000|5G|25.43s|201 M/s|
|20|4M|1000|10000|5G|7.23s|708 M/s|
|50|4M|1000|10000|5G|6.65|769 M/s|
|50|4M|1000|1000|5G|5.39|949 M/s|

The classic `cp` command:
```
time cp -R tmp_10000_5G tmp_10000_5G_copied3
real    0m12.154s
user    0m0.169s
sys     0m6.051s
```

# How to reproduce?

## Step 1: Prepare the test dirs
```sh
# for specific test:
#    10000 files, total 5G
go run ./ genFiles tmp_10000_5G 10000 5G


```
## Step 2: Run the benchmark
```sh
go test -bench=ReadOnly_10000_5G_ -benchtime=10s -run=NONE -v ./
```
# `du` shows unmatched size?
It's possible that `du` gives a different size (`du` is available on MacOS and Linux).

```sh
$ du -ch tmp_10000_5G/file_9999.txt
528K    tmp_10000_5G/file_9999.txt
528K    total

$ du -ch tmp_10000_5G_copied/file_9999.txt
576K    tmp_10000_5G_copied/file_9999.txt
576K    total
```

In the above cases du shows disk usage, not accurate file size.

To compare byte-by-byte, use file-by-file comparison:
```sh
diff tmp_10000_5G tmp_10000_5G_copied
```

Or the go impl:
```sh
go run ./ diffDir tmp_10000_5G tmp_10000_5G_copied
```