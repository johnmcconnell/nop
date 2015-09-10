package nop

// Writer ...
type Writer struct {
}

// Reader ...
type Reader struct {
}

// NewWriter ...
func NewWriter() *Writer {
	return &Writer{}
}

// NewReader ...
func NewReader() *Reader {
	return &Reader{}
}

// Read ...
func (w *Reader) Read(b []byte) (int, error) {
	return len(b), nil
}

// Writer ...
func (w *Writer) Write(b []byte) (int, error) {
	return len(b), nil
}
