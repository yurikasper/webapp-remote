package main

import (
	"bytes"
	"image"
	"net"
	"runtime"
	"time"

	qrcode "github.com/skip2/go-qrcode"
)

var localIp = "0.0.0.0"
var QRimage image.Image

func generateQR() {
	// For linux, it is very important to wait 2 seconds
	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}

	//get local IP
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok {
			if !ipnet.IP.IsLoopback() && ipnet.IP.DefaultMask() != nil {
				localIp = ipnet.IP.String()
			}

		}
	}

	//create IP QR code
	var QRpng []byte
	QRpng, qrErr := qrcode.Encode("http://"+localIp, qrcode.Medium, 256)
	if qrErr != nil {
		panic(qrErr)
	}
	// convert []byte to image
	QRimage, _, _ = image.Decode(bytes.NewReader(QRpng))
}
