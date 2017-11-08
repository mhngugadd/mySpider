package main

import (
	"github.com/mhngugadd/mySpider/download"
	"log"
)

func main() {
	urls := []string{"http://www.jianshu.com/p/e16d23e358c0","http://www.jianshu.com/p/7031752823e7"}
	downloader := download.NewLoad(urls)

	content , err := downloader.DownloadPage()
	if err != nil {
		log.Fatal(err.Error())
	}
	println(content)
}
