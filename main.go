package main

import (
	"flag"
	"fmt"
	"image"
	"log"
	"net/http"
	"strings"

	"image/gif"
	"image/jpeg"
	"image/png"

	"github.com/disintegration/gift"
)

var finalImage image.Image
var imageFormat string
var imageLocation string
var filters string

func init() {
	flag.StringVar(&imageLocation, "image_location", "https://i.imgur.com/Ed4LdEW.jpg", "an image url to transform")
	flag.StringVar(&filters, "filter_list", "grayscale", "what filter(s) you want to apply to your image")
	// Add a new command line option, perhaps listing available image filters
}

func main() {
	flag.Parse()
	src, err := retrieveImage(imageLocation)
	if err != nil {
		log.Fatalf("Unable to retrieve image: %v", err)
	}

	g := gift.New()
	dst := image.NewRGBA(g.Bounds(src.Bounds()))
	filterObjects := getFilters()
	g.Add(filterObjects...)
	g.Draw(dst, src)
	finalImage = dst.SubImage(src.Bounds())

	serve()
}

// Rename serve() to save() and try saving the image to a file
func serve() {
	//serve up image on localhost:8080/image
	fmt.Println("Please visit localhost:8080/image")
	http.HandleFunc("/image", respHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}
}

func respHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "image")
	switch imageFormat {
	case "jpg", "jpeg":
		jpeg.Encode(res, finalImage, nil)
	case "png":
		png.Encode(res, finalImage)
	case "gif":
		gif.Encode(res, finalImage, nil)
	default:
		log.Fatal("unrecognized image format")
	}
}

// Try retrieving the image from a file
func retrieveImage(imageLocation string) (image.Image, error) {
	resp, err := http.Get(imageLocation)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var src image.Image
	src, imageFormat, err = image.Decode(resp.Body)
	return src, err
}

func getFilters() []gift.Filter {
	var filterList []gift.Filter
	filterMap := make(map[string]gift.Filter)
	filterMap["grayscale"] = gift.Grayscale()
	filterMap["invert"] = gift.Invert()
	filterMap["pixelate"] = gift.Pixelate(3)
	// Add more filters here!

	filterTitles := strings.Split(filters, ",")
	for _, filter := range filterTitles {
		imageFilterObject := filterMap[filter]
		if imageFilterObject != nil {
			filterList = append(filterList, imageFilterObject)
		} else {
			log.Fatal("Sorry that image filter is not in the dictionary, please try a valid image filter")
		}
	}
	return filterList
}
