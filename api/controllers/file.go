package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"mime/multipart"
	"errors"
	"os"
	"io"
	"time"
)

func Upload(c *gin.Context)  {
	file, header, err := c.Request.FormFile("file")

	filename := header.Filename

	extension := strings.Split(filename, ".")[1]
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"msg": err.Error(),
			"data": "",
		})
		c.Abort()
		return
	}
	if extension == "gif" {
		msg, err := uploadGif(file, header)

		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"msg": err.Error(),
				"data": "",
			})
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"msg": msg,
			"data": "",
		})
	} else {
		msg, err := uploadNotGif(file, header)

		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"msg": err.Error(),
				"data": "",
			})
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"msg": msg,
			"data": "",
		})
	}

}

/**
上传非gif图片
 */
func uploadNotGif(file multipart.File, header *multipart.FileHeader) (msg string, err error) {
	filename := header.Filename

	if header.Size > 2048 * 1000 {
		return "", errors.New(filename + "文件过大，最大不能超过2M")
	}
	// 文件后缀是否正确
	extension := strings.Split(filename, ".")[1]
	/*extentions := make(map[string]int);
	extentions["png"] = 1
	extentions["jpeg"] = 1
	extentions["jpg"] = 1
	if _, in := extentions[extension]; !in {
		return "", errors.New(filename + " 文件格式错误");
	}*/
	extensions := []string{"png", "jpeg", "jpg"}
	status := 0
	for _, value := range extensions {
		if extension == value {
			status++
		}
	}
	if status == 0 {
		return "", errors.New(filename + " 文件格式错误");
	}

	img, err := os.Create("F:/" + filename)
	defer img.Close()
	if err != nil {
		return "", err
	}
	
	_, e := io.Copy(img, file)
	if e != nil {
		return "", e
	}

	return "success", nil
}

func uploadGif(file multipart.File, header *multipart.FileHeader) (msg string, err error) {
	if header.Size > 10240 * 1000 {
		return "", errors.New("GIF文件不能超过 10M")
	}
	img, err := os.Create("F:/" + string(time.Now().Unix())+ ".gif")
	defer img.Close()
	if err != nil {
		return "", err
	}
	_, e := io.Copy(img, file)

	if e != nil {
		return "", e
	}

	return "upload Success", nil
}