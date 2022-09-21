package json

type Image struct{ images []string }

var ImageIns Json[Image]

func InitImage() {
	ImageIns = Init("./config/image.json", Image{images: []string{}})
}
