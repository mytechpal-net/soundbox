package directories

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func CreateMainDirectory() {
	path := "./sounds"
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}
}

func CreateSbDirectory(directory string) {
	path := fmt.Sprintf("./sounds/%v", directory)
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
}
