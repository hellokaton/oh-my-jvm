package main

import (
	"archive/zip"
	"fmt"
)

func main() {
	r, err := zip.OpenReader("/Users/biezhi/Downloads/金融盒子支付回调APIi.txt.zip")
	if err == nil {
		defer r.Close()
		for _, f := range r.File {
			fmt.Println(f.Name)
		}
	}
}
