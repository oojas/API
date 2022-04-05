package imageProcessing

import (
	"io/ioutil"
	"log"
)

func ImageToBase64(path string) []byte { // Converting image into bytes
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return bytes

}
