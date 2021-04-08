package download

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/dennis/sendmsg"
)

func DownloadFile(filePath string, url string) {
	// Create the file
	out, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		sendmsg.SendMsg("Error when opening url.")
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		sendmsg.SendMsg("Error when connecting server.")
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		sendmsg.SendMsg("Error when write file.")
		log.Fatal(err)
	}
}
