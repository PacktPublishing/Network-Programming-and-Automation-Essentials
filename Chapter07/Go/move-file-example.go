package main

import (
	"fmt"
	"io"
	"os"
)

func moveFile(srcFile, dstFile string) error {
	src, err := os.Open(srcFile)
	if err != nil {
		return fmt.Errorf("os.Open: %v", err)
	}
	defer src.Close()

	dst, err := os.Create(dstFile)
	if err != nil {
		return fmt.Errorf("os.Create: %v", err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return fmt.Errorf("io.Copy: %v", err)

	}

	src.Close()
	err = os.Remove(srcFile)
	if err != nil {
		return fmt.Errorf("os.Remove: %v", err)
	}
	return nil
}

func main() {
	if err := moveFile("one.txt", "two.txt"); err != nil {
		fmt.Printf("Got and error: %v", err)
	}

}
