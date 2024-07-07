package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"regexp"

	"github.com/golang/freetype"
	"golang.org/x/image/draw"
	"golang.org/x/image/font/gofont/goregular"
)

type (
	//Label is a struct
	Label struct {
		Text  string
		Size  float64
		Color color.Color
		Rect  image.Rectangle
	}
)

func addLabel(icon *image.RGBA, label Label) {
	font, _ := freetype.ParseFont(goregular.TTF)

	c := freetype.NewContext()
	c.SetDPI(120)
	c.SetFont(font)
	c.SetFontSize(label.Size)
	c.SetClip(icon.Bounds())
	c.SetSrc(&image.Uniform{label.Color})

	// To measure total width, creating temp image
	measureBoard := image.NewRGBA(image.Rect(0, 0, 1, 1))
	c.SetDst(measureBoard)

	//positioning the label to center
	centerX := label.Rect.Min.X + label.Rect.Dx()/2
	centerY := label.Rect.Min.Y + label.Rect.Dy()/2 - int(label.Size)
	pt := freetype.Pt(centerX, centerY+int(c.PointToFixed(label.Size)>>6))
	newPt, _ := c.DrawString(label.Text, pt)
	fmt.Println(pt)
	fmt.Println(newPt)
	// Aligning text to center
	pt.X = pt.X - (newPt.X-pt.X)/2
	pt.Y = pt.Y - (newPt.Y-pt.Y)/2

	// Setting orignal icon
	c.SetDst(icon)
	c.DrawString(label.Text, pt)
}

func openIcon(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	return img, err
}

func saveIcon(icon image.Image, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	err = png.Encode(f, icon)
	if err != nil {
		return err
	}
	return nil
}

func resize(icon image.Image, size image.Rectangle) image.Image {
	resizedIcon := image.NewRGBA(size)
	draw.NearestNeighbor.Scale(resizedIcon, resizedIcon.Rect, icon, icon.Bounds(), draw.Over, nil)
	return resizedIcon
}

func createOverlay(rect image.Rectangle, primaryColor, secondaryColor color.Color) image.Image {
	// Overlay to be added on base icon
	iconOverlay := image.NewRGBA(rect)

	zp := image.Point{}
	sectionHeight := rect.Dy() / 100 * 12
	sideGap := rect.Dx() / 100 * 20 / 2
	center := iconOverlay.Bounds().Dx() / 2

	versionRect := image.Rect(sideGap, 0, center, sectionHeight)
	buildRect := image.Rect(center, 0, iconOverlay.Rect.Dx()-sideGap, sectionHeight)
	titleRect := image.Rect(0, iconOverlay.Rect.Dy()-sectionHeight, iconOverlay.Rect.Dx(), iconOverlay.Rect.Dy())

	draw.Draw(iconOverlay, versionRect, &image.Uniform{primaryColor}, zp, draw.Over)
	draw.Draw(iconOverlay, buildRect, &image.Uniform{secondaryColor}, zp, draw.Over)
	draw.Draw(iconOverlay, titleRect, &image.Uniform{primaryColor}, zp, draw.Over)

	addLabel(iconOverlay, Label{Text: "3.4.0", Color: color.White, Size: 60, Rect: versionRect})
	addLabel(iconOverlay, Label{Text: "3890", Color: color.White, Size: 60, Rect: buildRect})
	addLabel(iconOverlay, Label{Text: "Production", Color: color.White, Size: 60, Rect: titleRect})
	return iconOverlay
}

func getFiles(dir string, regex *string) []os.DirEntry {
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return []os.DirEntry{}
	}

	if regex != nil {
		matcher, _ := regexp.Compile(*regex)
		filteredFiles := []os.DirEntry{}
		for _, file := range files {
			if len(matcher.FindAllString(file.Name(), 1)) > 0 {
				filteredFiles = append(filteredFiles, file)
			}
		}
		return filteredFiles
	}
	return files
}

func main() {

	primaryColor := color.NRGBA{R: 255, G: 0, B: 0, A: 255}
	secondaryColor := color.NRGBA{R: 0, G: 0, B: 255, A: 255}

	// Base icon
	iconBase, _ := openIcon("./icon.png")

	// Overlay to be added on base icon
	iconOverlay := createOverlay(image.Rect(0, 0, 1024, 1024), primaryColor, secondaryColor)

	// Final icon combined both base + overlay
	icon := image.NewRGBA(iconBase.Bounds())
	draw.Draw(icon, icon.Bounds(), &image.Uniform{color.Black}, image.Point{0, 0}, draw.Over)
	draw.Draw(icon, icon.Bounds(), iconBase, image.Point{0, 0}, draw.Over)

	resizedOverlay := resize(iconOverlay, icon.Rect)
	draw.Draw(icon, icon.Bounds(), resizedOverlay, image.Point{0, 0}, draw.Over)

	// Saving final icon
	saveIcon(icon, "outimage.png")

	// env.LoadEnvironment()

	// fmt.Println("Applying Function....")
	// println("-----------------------------------------")
	// switch env.FUNCTION {
	// case function.SKIP_VERIFICATION:
	// 	function.PerformSkipVerification()
	// case function.CHECK_PULL_REQUESTS_DEADLINE:
	// 	function.PerformCheckPullRequestsDeadline()
	// case function.CALL_API:
	// 	function.PerformApiCall()
	// default:
	// 	fmt.Println("Invalid function selected !!")
	// }
	// env.DumpOutputs()
	// println("-----------------------------------------")

	// fmt.Println()
	// fmt.Println("Function is completed !")
	// fmt.Println("Tools signing off :)")
}
