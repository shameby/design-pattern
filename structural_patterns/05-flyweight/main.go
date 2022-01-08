package main

import "fmt"

var imageFactory *ImageFlyWeightFactory

func main() {
	v1, v2 := NewImageViewer("image1.jpg"), NewImageViewer("image1.jpg")
	v1.Display()
	v2.Display()
	fmt.Println(v1.ImageFlyWeight == v2.ImageFlyWeight)
}

func GetImageFlyweightFactory() *ImageFlyWeightFactory {
	if imageFactory == nil {
		imageFactory = &ImageFlyWeightFactory{
			m: map[string]*ImageFlyWeight{},
		}
	}
	return imageFactory
}

type ImageFlyWeightFactory struct {
	m map[string]*ImageFlyWeight
}

func (this_ *ImageFlyWeightFactory) Get(fileName string) *ImageFlyWeight {
	image := this_.m[fileName]
	if image == nil {
		image = &ImageFlyWeight{data: "image data:" + fileName}
		this_.m[fileName] = image
	}
	return image
}

type ImageFlyWeight struct {
	data string
}

func (ifw *ImageFlyWeight) Data() string {
	return ifw.data
}

func NewImageViewer(fileName string) *ImageViewer {
	image := GetImageFlyweightFactory().Get(fileName)
	return &ImageViewer{
		ImageFlyWeight: image,
	}
}

type ImageViewer struct {
	*ImageFlyWeight
}

func (i *ImageViewer) Display() {
	fmt.Printf("Display %s \n", i.data)
}
