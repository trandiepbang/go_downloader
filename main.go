package main

import (
	downloader "downloader/lib"
	"flag"
	"os"
)

func main() {
	token := flag.String("token", "", "a token, etc: eyJpc3MiOiJ0ZWNobWFzdGVyIiwiZX...")
	url := flag.String("url", "", "url of a file, etc : http://abc.xyz/aa.csv")
	savedPath := flag.String("saved_path", "", "a file saved path, etc /tmp/data.csv")
	method := flag.String("method", "GET", "a http method, etc POST, GET, DELETE, PUT")

	flag.Parse()

	if _, err := os.Stat(*savedPath); os.IsExist(err) {
		if err := os.Remove(*savedPath); err != nil {
			panic(err)
		}
	}

	downloader := downloader.NewDownloader(*token, *method)
	downloader.DownloadFile(*savedPath, *url)
}
