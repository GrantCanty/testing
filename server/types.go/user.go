package types

import (
	"crypto/rand"
	"image"
	"log"
	"os"
)

type User struct {
	Name         string
	About        string
	Occupation   string
	ProfilePhoto image.Image
}

func (u *User) LoadImage(filePath string) {
	imgFile, err := os.Open(filePath)
	defer imgFile.Close()
	if err != nil {
		log.Println(err)
	}

	img, _, err := image.Decode(imgFile)
	if err != nil {
		log.Println(err)
	}
	u.ProfilePhoto = img.(*image.NRGBA)
}

func (u *User) CreateImage() {
	rect := image.Rect(0, 0, 150, 150)
	pixels := make([]uint8, rect.Dx()*rect.Dy()*4)
	rand.Read(pixels)
	pic := &image.NRGBA{
		Pix:    pixels,
		Stride: rect.Dx() * 4,
		Rect:   rect,
	}
	u.ProfilePhoto = pic
}
