package compress

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"

	"github.com/nfnt/resize"
)

// CompressImage
// With the minimum width minWidth as the primary constraint, the file size may exceed maxKB.
// Secondly, within n iterations, perform a binary search for files sized between minKB and maxKB. If not found, return an error;
// If found, return the file content; if found but the current image width is less than minWidth, use minWidth as the standard.
func CompressImage(content []byte, minWidth uint, minKB uint, maxKB uint, n int) ([]byte, error) {
	if n <= 0 {
		// By default, a maximum of 10 iterations is allowed to prevent infinite loops in case of unknown special conditions.		n = 10
		n = 10
	}
	img, name, err := image.Decode(bytes.NewReader(content))
	if err != nil {
		return nil, err
	}
	if name == "jpeg" && compareSize(content, minKB, maxKB) == 0 {
		return content, nil
	}
	// Files that are invalid in size or not in JPG format will be uniformly converted to JPG format here.
	b, err := resizeImage(minWidth, img)
	if err != nil {
		return nil, err
	}
	// If the minimum width is already greater than the maximum limit, scaling will no longer continue. If minWidth is 0, it is considered that minWidth is not a constraint.
	if minWidth != 0 && compareSize(b, minKB, maxKB) >= 0 {
		return b, nil
	}
	preVal, curVal := uint(img.Bounds().Dx()), uint(img.Bounds().Dx())
	for i := 0; i < n; i++ {
		ret := compareSize(b, minKB, maxKB)
		if ret > 0 {
			curVal, preVal = scaleDown(curVal, preVal), curVal
			b, err = resizeImage(curVal, img)
			if err != nil {
				return nil, err
			}
			continue
		}
		if ret < 0 {
			curVal, preVal = scaleUp(curVal, preVal), curVal
			b, err = resizeImage(curVal, img)
			if err != nil {
				return nil, err
			}
			continue
		}
		// Recheck the image width.
		if curVal < minWidth {
			return resizeImage(minWidth, img)
		}
		return b, nil
	}
	return nil, fmt.Errorf("the image size cannot be adjusted to %d~%d KB within %d iterations", minKB, maxKB, n)
}

func scaleUp(currVal uint, preVal uint) uint {
	if currVal >= preVal {
		return currVal << 1
	}
	return (currVal + preVal) >> 1
}

func scaleDown(currVal uint, preVal uint) uint {
	if currVal <= preVal {
		return currVal >> 1
	}
	return (currVal + preVal) >> 1
}

func compareSize(content []byte, minKB uint, maxKB uint) int {
	padding := uint(10)
	imgSize := uint(len(content) >> 10)

	// Be more conservative when approaching the image size.
	if maxKB-minKB <= padding*2 {
		maxKB = maxKB + padding
		minKB = minKB - padding
	}
	if imgSize >= maxKB-padding {
		return 1
	}
	if imgSize <= minKB+padding {
		return -1
	}
	return 0
}

func resizeImage(weight uint, img image.Image) ([]byte, error) {
	buf := bytes.Buffer{}
	m := resize.Resize(weight, 0, img, resize.NearestNeighbor)
	if err := jpeg.Encode(&buf, m, &jpeg.Options{Quality: 75}); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
