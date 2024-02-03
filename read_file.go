package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
)

func ReadAllFilesDiscard(dir string, numFiles int) error {
	return readFiles(dir, numFiles, readDiscardFile)
}

func readDiscardFile(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(io.Discard, f)
	if err != nil {
		return err
	}
	return nil
}

func ReadAllFilesToAnotherDir(dir string, numFiles int) error {
	dstDir := dir + "_copied"
	err := os.MkdirAll(dstDir, 0755)
	if err != nil {
		return err
	}
	buf := make([]byte, 4*M)
	return readFiles(dir, numFiles, func(file string) error {
		return copyFile(file, filepath.Join(dstDir, filepath.Base(file)), buf)
	})
}

func ReadAllFilesToAnotherDirConcurrently(dir string, numFiles int, chSize int, goNum int) error {
	dstDir := dir + "_copied"
	err := os.MkdirAll(dstDir, 0755)
	if err != nil {
		return err
	}

	var errP atomic.Value // *error

	var actFiles int64
	ch := make(chan string, chSize)
	var wg sync.WaitGroup
	for i := 0; i < goNum; i++ {
		wg.Add(1)
		go func() {
			buf := make([]byte, 4*M)
			defer wg.Done()
			for file := range ch {
				if errP.Load() != nil {
					break
				}
				atomic.AddInt64(&actFiles, 1)
				err := copyFile(file, filepath.Join(dstDir, filepath.Base(file)), buf)
				if err != nil {
					errP.Store(&err)
					break
				}
			}
		}()
	}
	readFiles(dir, numFiles, func(file string) error {
		ch <- file
		return nil
	})
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

// to measure the fastest speed of read io
func ReadAllFilesDiscardConcurrent(dir string, numFiles int, chSize int, goNum int) error {
	var errP atomic.Value // *error
	var actFiles int64
	ch := make(chan string, chSize)
	var wg sync.WaitGroup
	for i := 0; i < goNum; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for file := range ch {
				if errP.Load() != nil {
					break
				}
				atomic.AddInt64(&actFiles, 1)
				err := readDiscardFile(file)
				if err != nil {
					errP.Store(&err)
					break
				}
			}
		}()
	}
	readFiles(dir, numFiles, func(file string) error {
		ch <- file
		return nil
	})
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

func copyFile(file string, dstFile string, buf []byte) error {
	r, err := os.Open(file)
	if err != nil {
		return err
	}
	defer r.Close()

	f, err := os.OpenFile(dstFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, r)
	return err
}

func ReadAllFilesToMemory(dir string, numFiles int) (map[string][]byte, error) {
	m := make(map[string][]byte)
	err := readFiles(dir, numFiles, func(file string) error {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}
		m[file] = content
		return nil
	})
	if err != nil {
		return nil, err
	}
	return m, nil
}

func readFiles(dir string, numFiles int, fn func(file string) error) error {
	for i := 0; i < numFiles; i++ {
		file := filepath.Join(dir, fmt.Sprintf("file_%d.txt", i))

		err := fn(file)
		if err != nil {
			return err
		}
	}
	return nil
}
