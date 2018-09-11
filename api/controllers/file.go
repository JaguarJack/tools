package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"mime/multipart"
	"errors"
	"os"
	"io"
	"math/rand"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"image/gif"
	"image"
	"image/draw"
	"image/color/palette"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"sync"
	"log"
	"github.com/nfnt/resize"
	"image/jpeg"
	"image/png"
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
			"msg": "success",
			"data": msg,
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
	newFileName := randomNString(5) + "." + extension
	img, err := os.Create("F:/image/" + newFileName)
	defer img.Close()
	if err != nil {
		return "", err
	}
	_, e := io.Copy(img, file)
	if e != nil {
		return "", e
	}

	// 图片压缩
	_i, _ := os.Open("F:/image/" + newFileName)
	defer _i.Close()
	_img, _, err := image.Decode(_i)
	if err != nil {
		log.Fatalln(err)
	}
	bounds := _img.Bounds()
	if bounds.Dx() < 150 && bounds.Dx() < 150 {
		return newFileName, nil
	}
	canvas := resize.Resize(300, 300, _img, resize.NearestNeighbor)
	file_out, err := os.Create("F:/image/" + newFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file_out.Close()
	if extension == "jpg" || extension == "jpeg" {
		jpeg.Encode(file_out, canvas, &jpeg.Options{80})
	} else {
		png.Encode(file_out, canvas)
	}

	return newFileName, nil

}

func uploadGif(file multipart.File, header *multipart.FileHeader) (msg string, err error) {
	if header.Size > 10240 * 1000 {
		return "", errors.New("GIF文件不能超过 10M")
	}
	newFileName := randomNString(10) + ".gif"
	img, err := os.Create("F:/image/" + newFileName)
	defer img.Close()
	if err != nil {
		return "", err
	}
	_, e := io.Copy(img, file)

	if e != nil {
		return "", e
	}

	return newFileName, nil
}

func makeGif(c *gin.Context)  {
	c.Request.ParseForm()
	if len(c.Request.PostForm) < 1 {
		c.JSON(http.StatusForbidden, gin.H{
			"msg": "Please Upload Pictures",
			"data": "",
		})
		c.Abort()
		return
	}
	newGif := &gif.GIF{}
	_palette := append(palette.WebSafe, color.Transparent)
	for _, value := range c.Request.PostForm {
		fmt.Println(value[0])
		f, err := os.Open("F:/image/" + value[0])
		if err != nil {
			log.Fatalln(err)
		}
		defer f.Close()
		img, _, err := image.Decode(f)
		if err != nil {
			log.Fatalln(err)
		}
		// 图片格式转换
		bounds := img.Bounds()
		palettedImage := image.NewPaletted(bounds, _palette)
		draw.Draw(palettedImage, bounds, img, image.ZP, draw.Src)
		newGif.Image = append(newGif.Image, palettedImage)
		newGif.Delay = append(newGif.Delay, 20)
	}
	_gif := randomNString(6) + ".gif"
	f, _ := os.Create(_gif)
	defer f.Close()
	gif.EncodeAll(f, newGif)

	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
		"data": _gif,
	})
}
func MakeGif(c *gin.Context) {
	c.Request.ParseForm()
	if len(c.Request.PostForm) < 1 {
		c.JSON(http.StatusForbidden, gin.H{
			"msg": "Please Upload Pictures",
			"data": "",
		})
		c.Abort()
		return
	}
	newGif := &gif.GIF{}
	var wg sync.WaitGroup
	_palette := append(palette.WebSafe, color.Transparent)
	_p := make(chan *image.Paletted)
	for _, value := range c.Request.PostForm {
		wg.Add(1)
			go func(name string, s []color.Color) {
				defer wg.Done()
				f, err := os.Open("F:/image/" + name)
				if err != nil {
					log.Fatalln(err)
				}
				defer f.Close()
				img, _, err := image.Decode(f)
				if err != nil {
					log.Fatalln(err)
				}
				// 图片格式转换
				bounds := img.Bounds()
				palettedImage := image.NewPaletted(bounds, s)
				draw.Draw(palettedImage, bounds, img, image.ZP, draw.Src)
				_p <- palettedImage
			}(value[0], _palette)
	}
	go func() {
		wg.Wait()
		close(_p)
	}()
	for v := range _p {
		newGif.Image = append(newGif.Image, v)
		newGif.Delay = append(newGif.Delay, 20)
	}
	fmt.Println(newGif)
	_gif := randomNString(6) + ".gif"
	f, _ := os.Create(_gif)
	defer f.Close()
	gif.EncodeAll(f, newGif)

	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
		"data": _gif,
	})
}
/**
产生一个随机字符串
 */
func randomNString(n int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJSKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	res := []byte{}
	for i :=0; i < n; i++ {
		res = append(res, bytes[rand.Intn(len(bytes))])
	}
	ctx := md5.New()
	ctx.Write(res)
	return hex.EncodeToString(ctx.Sum(nil))
}