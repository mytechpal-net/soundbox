package directories

import (
	"errors"
	"log"
	"os"
)

func CreateDirectory(path string) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}
}
