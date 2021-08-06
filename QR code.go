package main

import (
	"fmt"
	"github.com/tuotoo/qrcode"
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)
func GenerateQrCode(filename ,contents string, w,h int) {
	qrBarCode,err := qr.Encode(contents,qr.M,qr.Auto)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 定义二维码的大小
	qrCode,err := barcode.Scale(qrBarCode,w,h)
	if err != nil {
		fmt.Println(err)
		return
	}

	file,err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	png.Encode(file,qrCode)
}

func DecodeQrCode(filename string)  {
	file,err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	contents,err := qrcode.Decode(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(contents.Content)

}

func main() {
	filename := "demo.png"
	GenerateQrCode(filename,"Hello world!",256,256)
	DecodeQrCode(filename)
}
