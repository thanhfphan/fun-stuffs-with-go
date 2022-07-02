package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Create("thebest.ppm")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	width := 200
	height := 200
	// http://people.uncw.edu/tompkinsj/112/texnh/assignments/imageFormat.html
	n, err := file.Write([]byte(fmt.Sprintf("P3 %d %d 255\n", width, height)))
	if err != nil {
		log.Fatal(err)
	}

	// TODO: replace with beauty picture instead the red square :D
	content := ""
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			content = content + "255 0 0 "
		}
	}

	n, err = file.WriteAt([]byte(content), int64(n))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(n)
}
