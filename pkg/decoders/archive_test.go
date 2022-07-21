package decoders

import (
	"bytes"
	"io"
	"testing"
)

func TestReadToMax(t *testing.T) {
	tests := map[string]struct {
		size          int
		expectedError bool
	}{
		"base": {
			size:          10,
			expectedError: false,
		},
	}

	for testName, test := range tests {
		reader := createReader(test.size)
		data, err := ReadToMax(reader)
		switch {
		case err != nil && !test.expectedError:
			t.Errorf("%s: %s", testName, err)
		case err == nil && test.expectedError:
			t.Errorf("%s: %s", testName, "error expected, but none received")
		case len(data) != test.size:
			t.Errorf("%s: unexpected length of data. Expected: %d, got: %d", testName, test.size, len(data))
		}
	}
}

func createReader(size int) io.Reader {
	reader := &bytes.Buffer{}
	for i := 0; i < size; i++ {
		a := byte(0x0)
		reader.Write([]byte{a})
	}
	return reader
}
