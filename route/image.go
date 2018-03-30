package route

import (
	"dienlanhphongvan-cdn/app/entity"
	"dienlanhphongvan-cdn/errors"
	"dienlanhphongvan-cdn/util"
	"io"

	"github.com/gin-gonic/gin"
)

const (
	ImageKey = "file"
)

type Image struct {
	Image *entity.Image
}

func NewImage(image *entity.Image) *Image {
	return &Image{
		Image: image,
	}
}

func (r Image) Compress(c *gin.Context) {
	file, _, err := c.Request.FormFile(ImageKey)
	if util.CaseError(c, errors.ErrorBadParams(err)) {
		return
	}

	img, err := r.Image.Compress(file)
	if util.CaseError(c, errors.ErrorInternalServer(err)) {
		return
	}

	_, err = io.Copy(c.Writer, img)
	if util.CaseError(c, errors.ErrorInternalServer(err)) {
		return
	}
}

func (r Image) Crop(c *gin.Context) {
	var params struct {
		Width int `form:"width" json:"width" validate:"min=1"`
	}
	err := util.BindForm(c, &params)
	if util.CaseError(c, errors.ErrorBadParams(err)) {
		return
	}
	file, _, err := c.Request.FormFile(ImageKey)
	if util.CaseError(c, errors.ErrorBadParams(err)) {
		return
	}

	img, err := r.Image.Crop(file, params.Width)
	if util.CaseError(c, errors.ErrorInternalServer(err)) {
		return
	}

	_, err = io.Copy(c.Writer, img)
	if util.CaseError(c, errors.ErrorInternalServer(err)) {
		return
	}
}

func (r Image) Resize(c *gin.Context) {
	var params struct {
		Width int `form:"width" json:"width" validate:"min=1"`
	}
	err := util.BindForm(c, &params)
	if util.CaseError(c, errors.ErrorBadParams(err)) {
		return
	}
	file, _, err := c.Request.FormFile(ImageKey)
	if util.CaseError(c, errors.ErrorBadParams(err)) {
		return
	}

	img, err := r.Image.Resize(file, params.Width)
	if util.CaseError(c, errors.ErrorInternalServer(err)) {
		return
	}

	_, err = io.Copy(c.Writer, img)
	if util.CaseError(c, errors.ErrorInternalServer(err)) {
		return
	}
}
