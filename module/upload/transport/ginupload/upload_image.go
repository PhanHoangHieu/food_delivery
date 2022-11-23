package ginupload

import (
	"fmt"
	"g05-food-delivery/common"
	"g05-food-delivery/component/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func UploadImage(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")

		if err != nil {
			panic(err)
		}

		fileName := fileHeader.Filename
		extension := strings.Split(fileName, ".")[len(strings.Split(fileName, "."))-1]

		if err := c.SaveUploadedFile(fileHeader, fmt.Sprintf("static/%s", fileName)); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(common.Image{
			Id:        0,
			Url:       fmt.Sprintf("http://localhost:8080/static/%s", fileName),
			Width:     0,
			Height:    0,
			CloudName: "local",
			Extension: extension,
		}, nil, nil))
	}
}
