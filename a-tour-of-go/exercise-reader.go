package main

// import "golang.org/x/tour/reader"

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	return 'A', nil
}

func main() {
	// reader.Validate(MyReader{})
}
