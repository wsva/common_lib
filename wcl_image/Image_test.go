package wcl_image_test

import (
	"fmt"
	"testing"

	"github.com/wsva/common_lib/wcl_image"
)

func TestImage(T *testing.T) {
	img, err := wcl_image.GetImage("testimage.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(wcl_image.GetImageSize(img))
	img, err = wcl_image.CropImageHeight(img, 100)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(wcl_image.GetImageSize(img))
	img, err = wcl_image.CropImageWidth(img, 100)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(wcl_image.GetImageSize(img))
}
