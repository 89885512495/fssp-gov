// Download csv file, pasring it and add info in DataBase
package main

import (
	"fmt"

	"github.com/dennis/download"
	"github.com/dennis/parsing"
)

func main() {

	// Get date from now
	date := time.Now()

	// Get the required format for link (yymmdd)
	todayIs := date.Format("20060102")
	fileUrl := fmt.Sprintf("http://opendata.fssprus.ru/7709576929-iplegallist/data-" + todayIs + "-structure-20200401.csv")

	// Give the name for download document
	documentName := fmt.Sprintf("data-" + todayIs + "-structure-20200401.csv")

	// Start to download
	fmt.Println("Downloading..")
	download.DownloadFile(documentName, fileUrl)

	// Read csv file and add info in DataBase
	parsing.AddDB(documentName)
}
