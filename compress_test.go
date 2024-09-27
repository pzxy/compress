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
	minKB := uint(100)
	maxKB := uint(200)
	b, err = CompressImage(b, 0, minKB, maxKB, 10)
	if err != nil {
		t.Fatal(err)
	}
	if len(b)/1024 < int(minKB) {
		t.Fatal("compress failed")
	}
	if len(b)/1024 > int(maxKB) {
		t.Fatal("compress failed")
	}
	os.WriteFile("img_compressed.jpg", b, 0644)
}
