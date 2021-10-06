package business

import (
	"bytes"
	"context"
	"fmt"
	"image/jpeg"
	"io"
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/caohieu04/real-estate/common"
	"github.com/caohieu04/real-estate/component/upcloud"
	"github.com/caohieu04/real-estate/module/upload/model"
)

type CreateImageStore interface {
	CreateImage(context context.Context, data *common.Image) error
}

type uploadBusiness struct {
	provider upcloud.UploadProvider
	imgStore CreateImageStore
}

func NewUploadBusiness(provider upcloud.UploadProvider, imgStore CreateImageStore) *uploadBusiness {
	return &uploadBusiness{provider: provider, imgStore: imgStore}
}

func (biz *uploadBusiness) Upload(ctx context.Context, data []byte, folder, fileName string) (*common.Image, error) {
	fileBytes := bytes.NewBuffer(data)

	w, h, err := getImageDimension(fileBytes)

	if err != nil {
		return nil, model.ErrFileIsNotImage(err)
	}

	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}

	fileExt := filepath.Ext(fileName)
	fileName = fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt)

	img, err := biz.provider.SaveFileUploaded(ctx, data, fmt.Sprintf("%s/%s", folder, fileName))

	if err != nil {
		return nil, model.ErrCannotSaveFile(err)
	}

	img.Width = w
	img.Height = h
	img.Extension = fileExt

	return img, nil
}

// func convertToPNG(w io.Writer, r io.Reader) error {
// 	img, _, err := image.Decode(r)
// 	if err != nil {
// 		return err
// 	}
// 	return png.Encode(w, img.)
// }

func getImageDimension(reader io.Reader) (int, int, error) {
	img, err := jpeg.DecodeConfig(reader)
	if err != nil {
		log.Println(err)
		return 0, 0, err
	}
	return img.Width, img.Height, nil
}
