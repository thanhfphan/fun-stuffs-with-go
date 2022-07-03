package main

import (
	"fmt"

	"github.com/thanhfphan/global-id/gid"
)

func main() {
	gid := gid.New(1)
	fmt.Printf("ID %d\n", gid.GenarateID())
}
