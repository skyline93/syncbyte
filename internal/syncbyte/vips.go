package syncbyte

import (
	"os"

	"github.com/davidbyttow/govips/v2/vips"
)

func CreateThumbnail(source string, target string) error {
	vips.Startup(nil)
	defer vips.Shutdown()

	image, err := vips.NewImageFromFile(source)
	if err != nil {
		return err
	}

	if err = image.AutoRotate(); err != nil {
		return err
	}

	ep := vips.NewDefaultJPEGExportParams()
	imageBytes, _, err := image.Export(ep)
	if err != nil {
		return err
	}

	return os.WriteFile(target, imageBytes, 0644)
}
