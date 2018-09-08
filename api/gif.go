package main

import (
	"os"
	"fmt"
	"image/gif"
	"flag"
	"encoding/base64"
	"io/ioutil"
	"image/jpeg"
)
var (
	fontfile = flag.String("fontfile", "F:/msyh.ttf", "filename of the ttf font")
)
func main()  {
	file, err:= os.Open("F:/hhhh.gif")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}
	img, err := gif.DecodeAll(file)
	fmt.Println(img.Image[0])
	buffstore := make([]byte, 50000000)
	dist := make([]byte, 50000000)
	base64.StdEncoding.Decode(buffstore, []byte(img.Image[0])))
	/*for _, value := range img.Image {
		fontBytes,err := ioutil.ReadFile(*fontfile)
		if err != nil {
			log.Println(err)
		}
		//载入字体数据
		font,err := freetype.ParseFont(fontBytes)
		if err != nil {
			log.Println("load front fail",err)
		}
		f := freetype.NewContext()
		//设置分辨率
		f.SetDPI(72)
		//设置字体
		f.SetFont(font)
		//设置尺寸
		f.SetFontSize(26)
		f.SetClip(value.Bounds())
		//设置输出的图片
		f.SetDst(value)
		//设置字体颜色(红色)
		f.SetSrc(image.NewUniform(color.RGBA{255,0,0,255}))
		pt := freetype.Pt(40,40 + int(f.PointToFixed(26)) >> 8)
		_,err = f.DrawString("gay丽gay丽",pt)
		if err != nil {
			log.Fatal(err)
		}
	}
	f_, _ := os.Create("F:/kechi.gif")
	defer f_.Close()
	e := gif.EncodeAll(f_, img)
	fmt.Println(e)*/
	/*	fmt.Println(img.Image[0])
		for index, value := range img.Image {
			fmt.Println(index,"=>" ,value, "\n")
			break
		}*/
	//_img, err := gif.Decode(file)
	//fmt.Println(_img)

}
