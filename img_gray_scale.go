/*
Package img_gray_scale provides a function to convert a color image to grayscale and save it to an output file.

Usage:
1. Import the package:
   import "github.com/jstevens8185/img_gray_scale"

2. To convert an image to grayscale and save it:

   inputPath := "input_image.png"
   outputPath := "output_image.png"
   if err := imageconverter.ConvertImageToGrayscale(inputPath, outputPath); err != nil {
       fmt.Println("Error:", err)
   }

3. Handling errors:
   The ConvertImageToGrayscale function returns an error if any issues occur during the conversion or saving process.

By following these steps, you can use the imageconverter package to convert color images to grayscale in your Go projects.
*/

package img_gray_scale

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg" // Import the JPEG package
	"image/png"
	"os"
	"strings"
)

// ConvertImageToGrayscale converts a color image to grayscale and saves it to the specified output path.
func ConvertImageToGrayscale(inputPath, outputPath string) error {
	// Open the original image
	reader, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("error opening the image: %v", err)
	}
	defer reader.Close()

	// Detect the image format based on the file extension
	var img image.Image
	if strings.HasSuffix(strings.ToLower(inputPath), ".jpg") || strings.HasSuffix(strings.ToLower(inputPath), ".jpeg") {
		img, _, err = image.Decode(reader)
		if err != nil {
			return fmt.Errorf("error decoding the image: %v", err)
		}
	} else if strings.HasSuffix(strings.ToLower(inputPath), ".png") {
		img, _, err = image.Decode(reader)
		if err != nil {
			return fmt.Errorf("error decoding the image: %v", err)
		}
	} else {
		return fmt.Errorf("unsupported image format: %s", inputPath)
	}

	// Get image bounds
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	// Create a new grayscale image
	grayImg := image.NewGray(bounds)

	// Convert the color image to grayscale
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Get the color at pixel (x, y)
			oldColor := img.At(x, y)

			// Convert to gray using the luminance formula
			grayColor := color.GrayModel.Convert(oldColor).(color.Gray)
			grayImg.Set(x, y, grayColor)
		}
	}

	// Save the grayscale image
	grayFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating the output file: %v", err)
	}
	defer grayFile.Close()

	// Determine the output format based on the file extension
	if strings.HasSuffix(strings.ToLower(outputPath), ".jpg") || strings.HasSuffix(strings.ToLower(outputPath), ".jpeg") {
		if err := jpeg.Encode(grayFile, grayImg, nil); err != nil {
			return fmt.Errorf("error encoding grayscale image: %v", err)
		}
	} else if strings.HasSuffix(strings.ToLower(outputPath), ".png") {
		if err := png.Encode(grayFile, grayImg); err != nil {
			return fmt.Errorf("error encoding grayscale image: %v", err)
		}
	} else {
		return fmt.Errorf("unsupported output image format: %s", outputPath)
	}

	fmt.Println("Grayscale image saved to", outputPath)
	return nil
}
