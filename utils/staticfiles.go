package utils

import (
	"bus-api/config"
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/xyproto/randomstring"
)

func CreateMediaDirectories() {
	base := config.Config("MEDIA_ROOT")
	os.Mkdir(base+"/operator_logos", os.ModePerm)
	fmt.Println("Settting up 'operator_logos' media directory")
}

func SaveStaticFile(c *fiber.Ctx, file *multipart.FileHeader, media_path string) error {

	base_path := fmt.Sprintf("%s/%s", config.Config("MEDIA_ROOT"), media_path)
	path := fmt.Sprintf("%s/%s", base_path, file.Filename)
	for i := 1; i < 10; i++ {
		fmt.Println("i=", i)
		if _, err := os.Stat(path); !errors.Is(err, os.ErrNotExist) {
			str := randomstring.CookieFriendlyString(5)
			ext := filepath.Ext(path)
			path = fmt.Sprintf("%s/%s_%s%s", base_path, file.Filename, str, ext)
			fmt.Println(path)
			// break
		} else {
			break
		}
	}
	// fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
	if err := c.SaveFile(file, path); err != nil {
		return err
	}
	return nil
}
