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
	"github.com/golang/freetype"
	"io/ioutil"
	"strconv"
	"fmt"
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

		f := resolveGif(msg)
		c.JSON(http.StatusOK, gin.H{
			"msg": msg,
			"data": f,
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
制作gif信息
 */
func MakeGifIntro(c *gin.Context) {
	_gif := strings.Split(strings.Trim(c.PostForm("gif"), ","), ",")
	info := strings.Split(strings.Trim(c.PostForm("info"), ","), ",")
	_info := make(map[int]string, 0)
	for _, v := range info {
		_f := strings.Split(v, "_")
		key, _ := strconv.Atoi(_f[1])
		_info[key] = _f[0]
	}
	newGif := &gif.GIF{}
	_palette := append(palette.WebSafe, color.Transparent)
	for index, value := range _gif {
		if _, ok := _info[index]; ok {
			drawWordsToPic(value, _info[index])
		}
	}
	_palette_ch := make(chan map[int]*image.Paletted)
	var wg sync.WaitGroup
	for index, value := range _gif {
		wg.Add(1)
		go func(name string, s []color.Color, index int) {
			defer wg.Done()
			_file, _ := os.Open("F:/image/" + value)
			img, _, _ := image.Decode(_file)
			palettedImage := image.NewPaletted(img.Bounds(), _palette)
			draw.Draw(palettedImage, img.Bounds(), img, image.ZP, draw.Src)
			data := make(map[int]*image.Paletted)
			data[index] = palettedImage
			_palette_ch <- data
		}(value, _palette, index)
	}
	go func() {
		wg.Wait()
		close(_palette_ch)
	}()

	res := make(map[int]*image.Paletted)
	for v := range _palette_ch {
		for index, value := range v {
			res[index] = value
		}
	}
	_len := len(_gif)
	for i:=0; i< _len; i++ {
		newGif.Image = append(newGif.Image, res[i])
		newGif.Delay = append(newGif.Delay, 20)
	}
	_g := randomNString(6) + ".gif"
	f, _ := os.Create("F:/" + _g)
	defer f.Close()
	gif.EncodeAll(f, newGif)

	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
		"data": _gif,
	})
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

func resolveGif(filename string) []string  {
	file, err := os.Open("F:/image/" + filename)
	defer file.Close()
	if err != nil {
		log.Fatalln(err)
	}
	img, err := gif.DecodeAll(file)
	if err != nil {
		log.Fatalln(err)
	}
	files := make([]string, 0)
	newDir := randomNString(6)
	os.Mkdir("F:/image/" + newDir, os.ModePerm)
	for _, value := range img.Image {
		fileName := randomNString(6) + ".png"
		f , _ := os.Create( "F:/image/" + newDir + "/" + fileName)
		png.Encode(f, value)
		files = append(files, newDir + "/" + fileName)
	}
	return files
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
	f, _ := os.Create("F:/image/" + _gif)
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
		newGif.Delay = append(newGif.Delay, 60)
	}

	_gif := randomNString(6) + ".gif"
	f, _ := os.Create("F:/image/" + _gif)
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

/**
给图片绘制文字
 */
func drawWordsToPic(imgName string, msg string) {
	file, err := os.Open("F:/image/" + imgName)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer file.Close()
	imageFile, _, err := image.Decode(file)
	if err != nil {
		log.Fatalln(err)
	}
	fontBytes, err := ioutil.ReadFile("F:/msyh.ttf")
	rgba := image.NewRGBA(imageFile.Bounds())
	draw.Draw(rgba, imageFile.Bounds(), imageFile, image.ZP, draw.Src)
	//载入字体数据
	_font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println("load front fail", err)
	}
	f := freetype.NewContext()
	f.SetDPI(72) //设置分辨率
	f.SetFont(_font) //设置字体
	f.SetFontSize(20) //设置尺寸
	f.SetClip(rgba.Bounds())
	f.SetDst(rgba)
	f.SetSrc(image.NewUniform(color.RGBA{255,0,0,255}))
	//设置字体的位置
	pt := freetype.Pt(40,20 * int(f.PointToFixed(38)) >> 8)
	_, err = f.DrawString(msg, pt)
	if err != nil {
		log.Fatal(err)
	}
	ext := strings.Split(imgName, ".")[1]
	_f, err := os.Create("F:/image/" + imgName)
	defer _f.Close()
	if err != nil {
		fmt.Println(err)
	}
	if (ext == "jpeg" || ext == "jpg") {
		jpeg.Encode(_f, rgba, &jpeg.Options{jpeg.DefaultQuality})
	} else {
		png.Encode(_f, rgba)
	}
}