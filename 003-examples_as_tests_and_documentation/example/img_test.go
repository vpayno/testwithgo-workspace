package example

import (
	"fmt"
	"os"
	"testing"
)

// Only the contents of an overview example is shown when it's not in its own file.
// This comment for Example_crop2() gets used as the description text for the code only example block.
func Example_crop2() {
	f, err := os.Open("img.png")
	if err != nil {
		panic(err)
	}

	img, err := Decode(f)
	if err != nil {
		panic(err)
	}

	err = Crop(img, 0, 0, 20, 20)
	if err != nil {
		panic(err)
	}

	out, err := os.Create("out.png")
	if err != nil {
		panic(err)
	}

	err = Encode(img, out)
	if err != nil {
		panic(err)
	}

	fmt.Println("See out.png for the cropped image.")
	// Output:
	// See out.png for the cropped image.
}

// This is how the Decode() function is used.
func TestDecode(t *testing.T) {}

// This is how the Crop() function is used.
func TestCrop(t *testing.T) {}

// This is how the Encode() function is used.
func TestEncode(t *testing.T) {}
