package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
)

func GenFiles(dir string, numFiles int, totalBytes ByteUnit) error {
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return nil
	}
	avgSize := int(totalBytes) / numFiles
	content := []byte(strings.Repeat("a", avgSize))
	for i := 0; i < numFiles; i++ {
		file := filepath.Join(dir, fmt.Sprintf("file_%d.txt", i))
		// f,err:= os.OpenFile(file,os.O_WRONLY|os.O_CREATE|os.O_TRUNC,0755)
		// if err!=nil{
		// 	return err
		// }

		// io.CopyBuffer()
		err := ioutil.WriteFile(file, content, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

// to measure the fastest speed of read io
func GenFilesConcurrently(dir string, numFiles int, totalBytes ByteUnit, chSize int, goNum int) error {
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return nil
	}

	avgSize := int(totalBytes) / numFiles
	content := []byte(strings.Repeat("a", avgSize))

	var errP atomic.Value // *error
	var actFiles int64
	ch := make(chan int, chSize)
	var wg sync.WaitGroup
	for i := 0; i < goNum; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range ch {
				if errP.Load() != nil {
					break
				}
				atomic.AddInt64(&actFiles, 1)
				file := filepath.Join(dir, fmt.Sprintf("file_%d.txt", i))
				err := ioutil.WriteFile(file, content, 0755)
				if err != nil {
					errP.Store(&err)
					break
				}
			}
		}()
	}
	for i := 0; i < numFiles; i++ {
		ch <- i
	}
	close(ch)
	wg.Wait()

	if e := errP.Load(); e != nil {
		return *(e.(*error))
	}
	if atomic.LoadInt64(&actFiles) != int64(numFiles) {
		return fmt.Errorf("mismatched files, expected: %d, atual: %d", int64(numFiles), atomic.LoadInt64(&actFiles))
	}
	return nil
}
