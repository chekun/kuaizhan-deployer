package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/chekun/kuaizhan"
)

func main() {
	filePath := flag.String("file", "", "path to source file")
	appKey := flag.String("appkey", "", "kuaizhan app key")
	appSecret := flag.String("appsecret", "", "kuaizhan app secret")
	siteID := flag.String("site", "", "kuaizhan site id")
	pageID := flag.String("page", "", "kuaizhan page id, optional")

	flag.Parse()

	if *filePath == "" || *appKey == "" || *appSecret == "" || *siteID == "" {
		log.Println("insufficient arguments")
		os.Exit(1)
	}

	client := kuaizhan.NewClient(*appKey, *appSecret, nil)

	if *pageID == "" {
		pages, err := client.TbkGetPageName(*siteID)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		if len(pages) == 0 {
			log.Println("no pages in site")
			os.Exit(1)
		}
		*pageID = fmt.Sprintf("%d", pages[0].PageID)
	}

	code, err := ioutil.ReadFile(*filePath)
	if err != nil {
		log.Println("failed to read file, ", err)
		os.Exit(1)
	}

	err = client.TbkModifyPageJs(*siteID, *pageID, string(code), false)
	if err != nil {
		log.Println("failed to modify page, ", err)
		os.Exit(1)
	}

	pageURL, err := client.TbkPublishPage(*siteID, *pageID)
	if err != nil {
		log.Println("failed to publish page, ", err)
		os.Exit(1)
	}

	log.Println("page modified and published: ", pageURL)
}
