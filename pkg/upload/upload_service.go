package upload

import (
	"errors"
	"image"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/disintegration/imaging"
	"github.com/google/uuid"
)

const MAX_FILE_SIZE = 2 * 1024 * 1024 // 2MB

func UploadImage(file *multipart.FileHeader, fileDirectory string) (string, error) {
	if file.Size > MAX_FILE_SIZE {
		return "", errors.New("file size exceeds 2MB")
	}

	ext := filepath.Ext(file.Filename)
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		return "", errors.New("invalid file format")
	}

	uploadPath := os.Getenv("UPLOAD_PATH")

	filename := uuid.New().String() + ext
	path := uploadPath + "/" + fileDirectory + "/" + filename

	// Create Upload Folder if not Exists
	if err := os.MkdirAll(uploadPath+"/"+fileDirectory+"/", os.ModePerm); err != nil {
		return "", err
	}

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	img, _, err := image.Decode(src)
	if err != nil {
		return "", err
	}

	// img = imaging.Resize(img, 1300, 1300, imaging.Lanczos)

	err = imaging.Save(img, path)
	if err != nil {
		return "", err
	}

	return filename, nil
}
