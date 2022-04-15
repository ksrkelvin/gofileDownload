package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fmt.Println("iniciando...")

	fileName := "sample"
	repo := "./Downloads/"

	URL := "https://ciclovivo.com.br/wp-content/uploads/2018/10/iStock-536613027.jpg"

	err := downloadFile(URL, fileName, "jpg", repo)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("File %s downlaod in current working directory. \n", fileName)
	fmt.Println("Finalizado...")

}

func downloadFile(url string, fileName string, fileType string, downloadLocal string) (err error) {

	path := downloadLocal + fileName + "." + fileType
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("Received non 200 response code")
	}

	_, err = os.Stat(path)

	if err == nil {
		count := 1
		i := true
		for i {
			path = fmt.Sprintf("%v%v(%v).%v", downloadLocal, fileName, count, fileType)
			_, err = os.Stat(path)
			if err == nil {
				count++
			} else {
				i = false
			}

		}
	}

	file, err := os.Create(path)
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return err

}
