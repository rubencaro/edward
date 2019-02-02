package main

// Taken partly from github.com/Nr90/imgsim
// which was taken from github.com/jteeuwen/imghash

import (
	"bytes"
	"image"

	"github.com/disintegration/imaging"
)

// AverageHash calculates the average hash of an image. First the image is resized to 8x8.
// Then it is converted to grayscale. Lastly the average hash is computed.
func AverageHash(data []byte) (uint64, error) {
	buff := bytes.NewBuffer(data)
	img, _, err := image.Decode(buff)
	if err != nil {
		return 0, err
	}
	img = imaging.Resize(img, 8, 8, imaging.Box)
	img = imaging.Grayscale(img)
	mean := mean(img)
	return calcAvgHash(img, mean), nil
}

// mean computes the mean of all pixels.
func mean(img image.Image) uint32 {
	rect := img.Bounds()
	w := rect.Max.X - rect.Min.X
	h := rect.Max.Y - rect.Min.Y
	t := uint32(w * h)

	if t == 0 {
		return 0
	}

	var x, y int
	var r, sum uint32
	for x = rect.Min.X; x < rect.Max.X; x++ {
		for y = rect.Min.Y; y < rect.Max.Y; y++ {
			r, _, _, _ = img.At(x, y).RGBA()
			sum += r
		}
	}

	return sum / t
}

// calcAvgHash computes the average hash for the given image and mean.
func calcAvgHash(img image.Image, mean uint32) uint64 {
	var x, y int
	var hash, p uint64
	p = 1
	var r uint32

	rect := img.Bounds()

	for y = rect.Min.Y; y < rect.Max.Y; y++ {
		for x = rect.Min.X; x < rect.Max.X; x++ {
			r, _, _, _ = img.At(x, y).RGBA()

			if r > mean {
				hash |= p
			}

			p = p << 1
		}
	}

	return hash
}

// Distance calculates the Hamming Distance between the two input hashes.
func Distance(a, b uint64) uint64 {
	var dist uint64
	for val := a ^ b; val != 0; val &= val - 1 {
		dist++
	}
	return dist
}
