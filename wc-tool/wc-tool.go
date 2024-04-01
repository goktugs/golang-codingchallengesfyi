package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Kullanım: ccwc -c dosya_adı")
		return
	}

	filename := os.Args[1]

	fmt.Println("Dosya adı:", filename)

	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Dosya açılamadı", err)
		return
	}

	defer file.Close()

	var byteCount int

	buf := make([]byte, 1024)

	for {
		n, err := file.Read(buf)

		if err != nil {
			fmt.Println("Dosya okunamadı", err)
			return
		}

		if n == 0 {
			break
		}

		byteCount += n
	}

	fmt.Printf("%d %s\n", byteCount, filename)

}
