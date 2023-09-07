// Package resize provides a function to resize a PNG to a set of sizes
package resize

import (
	"image/png"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
)

// Size is a struct to hold the size and extension of the resized image
type Size struct {
	Ext  string
	Size int
}

// Apple macOS Icon Sizes
var macOsSizes = []Size{
	{Ext: "_512x512@2x.png", Size: 1024},
	{Ext: "_512x512.png", Size: 512},
	{Ext: "_256x256@2x.png", Size: 512},
	{Ext: "_256x256.png", Size: 256},
	{Ext: "_128x128@2x.png", Size: 256},
	{Ext: "_128x128.png", Size: 128},
	{Ext: "_32x32@2x.png", Size: 64},
	{Ext: "_32x32.png", Size: 32},
	{Ext: "_16x16@2x.png", Size: 32},
	{Ext: "_16x16.png", Size: 16},
}

// Resize takes an input PNG and resizes it to the sizes defined in the sizes array
func Resize(input string) error {
	fname := strings.TrimSuffix(input, filepath.Ext(input))

	in, err := os.Open(input)
	if err != nil {
		return err
	}
	defer in.Close()

	for _, size := range macOsSizes {
		out, err := os.Create(fname + size.Ext)
		if err != nil {
			return err
		}
		defer out.Close()

		src, err := png.Decode(in)
		if err != nil {
			return err
		}

		dst := imaging.Resize(src, size.Size, size.Size, imaging.Lanczos)

		if err := png.Encode(out, dst); err != nil {
			return err
		}

		in.Seek(0, io.SeekStart)
	}

	return nil
}
