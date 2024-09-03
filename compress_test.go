package compress

import (
	"os"
	"testing"
)

func TestCompressImage(t *testing.T) {
	b, err := os.ReadFile("img.jpg")
	if err != nil {
		t.Fatal(err)
	}
	b, err = Do(b, 0, 100, 200, 10)
	if err != nil {
		t.Fatal(err)
	}
	os.WriteFile("img2.jpg", b, 0644)
}
