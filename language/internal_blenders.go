package language

import (
	"fmt"
	"image"
	"image/color"
	"sync"
)

type Blenders struct {
	*sync.Mutex
	blenders map[string]*Blender
}

func (bls *Blenders) add(blenders ...*Blender) error {
	bls.Lock()
	defer bls.Unlock()
	for _, b := range blenders {
		if _, ok := bls.blenders[b.name]; ok {
			return fmt.Errorf("%s is already registered", b.name)
		}
		bls.blenders[b.name] = b
	}
	return nil
}

func (bls *Blenders) get(name string) *Blender {
	bls.Lock()
	defer bls.Unlock()
	if b, ok := bls.blenders[name]; ok {
		return b
	}
	return nil
}

func (bls *Blenders) BlendColor(name string, colA, colB color.RGBA64) color.RGBA64 {
	return bls.get(name).Color(colA, colB)
}

func (bls *Blenders) BlendPixel(name string, pixel Point, imgA, imgB *image.NRGBA64) *image.NRGBA64 {
	return bls.get(name).Pixel(pixel, IClone(imgA), imgB)
}

func (bls *Blenders) BlendImages(name string, imgA, imgB *image.NRGBA64) *image.NRGBA64 {
	return bls.get(name).Images(IClone(imgA), imgB)
}

func NewBlenders(blenders ...*Blender) *Blenders {
	bls := Blenders{
		Mutex:    &sync.Mutex{},
		blenders: map[string]*Blender{},
	}
	bls.add(blenders...)
	return &bls
}
