package download

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"errors"
)

type Download interface{
	DownloadPage () (content map[string]*goquery.Document,err error)
}

type DownLoader struct {
	Urls []string
}

func NewLoad(Urls []string) *DownLoader {
	return &DownLoader{
		Urls : Urls,
	}
}

func (d *DownLoader)DownloadPage() (content map[string]*goquery.Document , err error) {
	content = make(map[string]*goquery.Document)
	if le := len(d.Urls); le > 0  {
		for _,url := range d.Urls {
			document , err := goquery.NewDocument(url)
			if err != nil {
				log.Fatal("request error : ", err.Error())
			}
			content[url] = document
			err = nil
		}
		println(content)
	}else {
		err = errors.New(" Urls is empty! ")
		log.Fatal("Urls is empty!")
	}
	return
}

