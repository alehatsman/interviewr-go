package images

import (
	"bytes"
	"errors"
	"image"
	"image/jpeg"
	"io"
	"net/http"
	"strconv"

	"gopkg.in/mgo.v2/bson"

	"github.com/atsman/interviewr-go/handlers/utils"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("handlers.images")

type CropData struct {
	x, y, w, h float32
}

func convertToRectangle(w, h, x, y int) image.Rectangle {
	y0 := y + h
	y1 := y
	x0 := x
	x1 := x + w
	return image.Rect(x0, y0, x1, y1)
}

func parseVal(key string, c *gin.Context) (error, int) {
	val, err := strconv.ParseFloat(c.Request.FormValue(key), 32)
	if err != nil {
		return err, 0
	}
	return err, int(val)
}

func Create(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.Error(err)
		return
	}

	_, h := parseVal("h", c)
	_, w := parseVal("w", c)
	_, x := parseVal("x", c)
	_, y := parseVal("y", c)

	img, _, err := image.Decode(file)
	if err != nil {
		c.Error(err)
		return
	}

	log.Debug("Crop w:%v h:%v x:%v y:%v", w, h, x, y)

	rect := convertToRectangle(int(w), int(h), int(x), int(y))

	log.Debug("Rect", rect)

	croppedImg := imaging.Crop(img, rect)

	db := utils.GetDb(c)
	gridFile, err := db.GridFS("fs").Create(fileHeader.Filename)
	if err != nil {
		c.Error(err)
		return
	}

	buffer := new(bytes.Buffer)
	jpeg.Encode(buffer, croppedImg, nil)

	gridFile.SetName(fileHeader.Filename)
	gridFile.SetContentType(fileHeader.Header.Get("Content-Type"))
	_, err = io.Copy(gridFile, buffer)
	if err != nil {
		c.Error(err)
		return
	}

	file.Close()
	gridFile.Close()

	c.JSON(http.StatusOK, gin.H{"_id": gridFile.Id()})
}

func GetOne(c *gin.Context) {
	db := utils.GetDb(c)
	imageId := c.Params.ByName("id")

	if !bson.IsObjectIdHex(imageId) {
		c.Error(errors.New("imageId is not a ObjectIdHex"))
		return
	}

	gfile, err := db.GridFS("fs").OpenId(bson.ObjectIdHex(imageId))
	if err != nil {
		c.Error(err)
		return
	}

	_, err = io.Copy(c.Writer, gfile)
	if err != nil {
		c.Error(err)
		return
	}
	gfile.Close()
	c.Header("Content-Disposition", "attachment; filename="+gfile.Name())
	c.Header("Content-Type", "application/x-download")
}
