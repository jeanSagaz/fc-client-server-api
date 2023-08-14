package services

import (
	"errors"
	"fmt"
	"os"

	"github.com/jeanSagaz/client/internal/dto"
)

func createPath(path string) error {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}

func CreateFile(response *dto.Response) error {
	createPath("txt")

	f, err := os.OpenFile("txt/cotacao.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	size, err := f.Write([]byte(fmt.Sprintf("Dólar: {%s}\n", response.Bid)))
	if err != nil {
		return err
	}
	fmt.Printf("File created successfully! Size: %d bytes\n", size)

	return nil
}
