package ui

import (
	"bytes"
	"image"
	"strings"

	"github.com/charmbracelet/x/mosaic"
	"github.com/mattn/go-sixel"
)

type ImageRenderer interface {
	Render(img image.Image) string
	Resize(width int, height int)
}
type MosaicRenderer struct {
	_mosaic mosaic.Mosaic
}
func NewMosaicRenderer(w int, h int) ImageRenderer {
	return &MosaicRenderer{
		_mosaic: mosaic.New().Height(h).Width(w),
	}
}
func (b *MosaicRenderer) Resize(w int, h int) {
	b._mosaic.Width(w)
	b._mosaic.Height(h)
}
func (b *MosaicRenderer) Render(img image.Image) string {
	return b._mosaic.Render(img)
}

type SixelRenderer struct {
	_width int
	_height int
}
func NewSixelRenderer(w int, h int) ImageRenderer {
	return &SixelRenderer{
		_width: w,
		_height: h,
	}
}
func (b *SixelRenderer) Resize(w int, h int) {
	b._width = w
	b._height = h
}
func (b *SixelRenderer) Render(img image.Image) string {
	var buf bytes.Buffer
	encoder := sixel.NewEncoder(&buf)
	encoder.Width = b._width * 5
	encoder.Height = b._height * 10
	encoder.Encode(img)
	var result strings.Builder
	result.WriteString(buf.String())
	placeholder := strings.Repeat(" ", b._width) + "\n"
	for range b._height {
		result.WriteString(placeholder)
	}
	return result.String()
}

