package config

type Mime struct {
	PDF  string
	Jpg  string
	Png  string
	Gif  string
	Webp string
}

var MIME_TYPE = Mime{
	"application/pdf",
	"image/jpeg",
	"image/png",
	"image/gif",
	"image/webp",
}
