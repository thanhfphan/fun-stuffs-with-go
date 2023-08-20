package main

import (
	"fmt"

	"github.com/thanhfphan/global-id/gid"
)

const base62Characters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	links := []string{
		"https://web.archive.org/web/20080603022239/http://www.gsfc.nasa.gov/topstory/20020926landcover.html",
		"http://nwas.org/ej/pdf/2007-FTT1.pdf",
		"https://education.arm.gov/outreach/publications/08augnewsltr.pdf",
		"https://www.bbc.co.uk/weather/features/understanding/fog.shtml",
		"https://go.dev/doc/modules/gomod-ref",
		"https://books.google.com.vn/books?id=gfeCXlElJTwC&pg=PA221&redir_esc=y",
		"https://www.theguardian.com/theguardian/2010/nov/08/harry-fensom-obituary",
		"https://en.wikipedia.org/wiki/Honeywell,_Inc._v._Sperry_Rand_Corp.",
		"https://web.archive.org/web/20191217200937/https://www.uspto.gov/about-us/news-updates/remarks-director-iancu-2019-international-intellectual-property-conference",
		"https://books.google.com.vn/books?id=lyJGAQAAIAAJ&redir_esc=y",
	}

	baseURL := "https://short.link/"
	shard := gid.New(1)
	for _, link := range links {
		key := convertUint64ToBase62String(shard.GenarateID())
		shortLink := baseURL + key
		fmt.Printf("%s ----> %s\n", shortLink, link)
	}

}

func convertUint64ToBase62String(n uint64) string {

	var result []byte

	for n != 0 {
		m := n % 62
		n = n / 62
		result = append(result, base62Characters[m])
	}

	return string(result)
}
