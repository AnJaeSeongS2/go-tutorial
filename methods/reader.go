package methods

import (
	"fmt"
	"io"
	"strings"
)

func TourReader() {
	r := strings.NewReader("Hello, Reader!")

	for b := make([]byte, 8); ; {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])

		if err == io.EOF {
			break
		}
	}
}
