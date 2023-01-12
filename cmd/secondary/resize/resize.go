package resize

import (
	"image"

	"github.com/disintegration/imaging"
)

func Resize(img image.Image, size int, imgChan chan<- image.Image) {
	dstImage := imaging.Resize(img, size, size, imaging.Lanczos)
	imgChan <- dstImage

}
