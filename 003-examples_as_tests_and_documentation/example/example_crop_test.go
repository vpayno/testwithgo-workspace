package example

import (
	"fmt"
	"io"

	// Needed for initialize side effect
	_ "image/png"
)

var file string = "this is not used"

// If you only have one example in this file, the whole file will be shown in the example.
// This comment won't get used in the documentation.
func Example_crop() {
	var r io.Reader
	img, err := Decode(r)
	if err != nil {
		panic(err)
	}
	err = Crop(img, 0, 0, 20, 20)
	if err != nil {
		panic(err)
	}
	var w io.Writer
	if err != nil {
		panic(err)
	}
	err = Encode(img, w)
	if err != nil {
		panic(err)
	}
	fmt.Println("See out.jpg for the cropped image.")

	// Output:
	// See out.jpg for the cropped image.
}
