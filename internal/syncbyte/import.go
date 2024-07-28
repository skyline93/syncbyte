package syncbyte

import (
	"os"
	"path/filepath"
	"phto/internal/config"
	"phto/internal/entity"
	"phto/internal/log"

	"github.com/sirupsen/logrus"
)

var (
	logger *logrus.Logger
)

func init() {
	logger = log.NewLogger("syncbyte.log")
}

func ImportOriginals(userName string, conf *config.Config) error {
	user := entity.FindUser(userName)

	logger.Debugf("import at albums, user %s, album: %v", userName, user.Albums)
	for _, album := range user.Albums {
		photos, err := entity.FindUnimportedPhotosByAlbum(album.ID)
		if err != nil {
			continue
		}

		logger.Debugf("import at photos, %v", photos)
		for _, photo := range photos {
			filePath := filepath.Join(conf.StoragePath, "uploads", userName, album.Name, photo.FileName)

			originalsPath := filepath.Join(conf.StoragePath, "originals", userName, album.Name)
			if _, err := os.Stat(originalsPath); os.IsNotExist(err) {
				err := os.MkdirAll(originalsPath, os.ModePerm)
				if err != nil {
					continue
				}
			}

			destPath := filepath.Join(originalsPath, photo.FileName)
			logger.Debugf("rename %s to %s", filePath, destPath)
			if err := os.Rename(filePath, destPath); err != nil {
				logger.Debugf("rename failed, err: %s", err)
				continue
			}

			if err := photo.SetIsImported(true); err != nil {
				logger.Debugf("set is imported failed, err : %s", err)
				continue
			}
		}
	}

	return nil
}
