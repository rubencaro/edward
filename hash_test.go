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

	type args struct {
		img image.Image
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{"gopher image", args{img}, 9331034559709847552},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AverageHash(tt.args.img)
			tst.Eq(t, tt.want, got)
		})
	}
}

func Test_mean(t *testing.T) {
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
		{"gopher image", args{img}, 9872},
		{"0x0 image", args{emptyimg}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mean(tt.args.img)
			tst.Eq(t, tt.want, got)
		})
	}
}
