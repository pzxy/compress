package compress

import (
	"os"
	"testing"
)

func TestCompressImage(t *testing.T) {
	b, err := os.ReadFile("img.jpeg")
	if err != nil {
		t.Fatal(err)
	}
	minKB := uint(20)
	maxKB := uint(30)
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
	_ = os.WriteFile("img2.jpg", b, 0644)
}
