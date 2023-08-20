package main

import (
	"fmt"

	"github.com/sqids/sqids-go"
	"github.com/thanhfphan/global-id/gid"
)

func main() {

	s, _ := sqids.New()
	gidd := gid.New(1)

	gid := gidd.GenarateID()
	fmt.Println(gid)

	id, _ := s.Encode([]uint64{gid})
	fmt.Println(id)

	numbers := s.Decode(id)
	fmt.Println(numbers)

}
