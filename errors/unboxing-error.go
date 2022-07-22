package main

import (
	"errors"
	"fmt"
	"os"
)

// たとえば os.Open() 関数の返り値のエラー型は以下の内部状態を持っている。
// PathError records an error and the operation and file path that caused it.
// type PathError struct {
//     Op   string
//     Path string
//     Err  error
// }

func main() {
	file, err := os.Open("no-file.go")

	// switch e := err.(type) {
	// case *os.PathError:
	// 	if errno, ok := e.Err.(syscall.Errno); ok {

	// 		switch errno {
	// 		case syscall.ENOENT:
	// 			fmt.Fprintln(os.Stderr, "file does not exist.")
	// 		case syscall.ENOTDIR:
	// 			fmt.Fprintln(os.Stderr, "directory does not exist.")
	// 		default:
	// 			fmt.Fprintln(os.Stderr, "Errno =", errno)
	// 		}

	// 	} else {
	// 		fmt.Fprintln(os.Stderr, "other path error")
	// 	}

	// default:
	// 	fmt.Fprintln(os.Stderr, "other error except PathError")
	// }

	var perr *os.PathError
	if errors.As(err, &perr) {
		fmt.Fprintf(os.Stderr, "file is \"%v\"\n", perr.Path)
	}

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}()
}
