package main

import (
	"image"
	"image/png"
	"os"
	"testing"

	"github.com/rubencaro/edward/lib/tst"
)

func TestAverageHash(t *testing.T) {
	imgfile, err := os.Open("assets/gopher.png")
	defer imgfile.Close()
	tst.Ok(t, err)
	img, err := png.Decode(imgfile)
	tst.Ok(t, err)

	got := AverageHash(img)
	tst.Eq(t, uint64(0x817e7e7e7e666000), got)
}

func TestMean(t *testing.T) {
	imgfile, err := os.Open("assets/gopher.png")
	defer imgfile.Close()
	tst.Ok(t, err)
	img, err := png.Decode(imgfile)
	tst.Ok(t, err)

	emptyimg := image.NewGray(image.Rectangle{})

	type args struct {
		img image.Image
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{"gopher image", args{img}, 0x2590},
		{"0x0 image", args{emptyimg}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mean(tt.args.img)
			tst.Eq(t, tt.want, got)
		})
	}
}

func TestDistance(t *testing.T) {
	a := uint64(9331034559709847552)
	b := uint64(9331034559709847552)
	d := Distance(a, b)
	tst.Eq(t, uint64(0), d)

	c := uint64(9331034559709848552)
	e := Distance(a, c)
	tst.Eq(t, uint64(6), e)
}
