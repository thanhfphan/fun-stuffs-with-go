package main

import (
	"fmt"

	"github.com/thanhfphan/fun-stuff-with-go/buffer/src"
)

type Demo struct {
	Data int
}

func (d *Demo) String() string {
	return fmt.Sprintf("%d", d.Data)
}

func main() {
	initSize := 4
	testSize := 10
	d := src.NewDeque[Demo](initSize)

	for i := 0; i < testSize; i++ {
		tmp := Demo{
			Data: i,
		}
		d.PushRight(tmp)
	}

	for {
		val, ok := d.PopLeft()
		if !ok {
			fmt.Println(val)
			break
		}

		fmt.Println(val.String())
	}
}
