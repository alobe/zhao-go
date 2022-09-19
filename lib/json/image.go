package json

type Image struct{ images []string }

var ImageIns Json[Image]

func InitImage() {
	Init("./config/image.json", Image{images: []string{}}, ImageIns)
}
