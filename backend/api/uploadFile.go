package api

import (
	"backend/analys"
	"backend/mystem"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"sync"

	"github.com/gin-gonic/gin"
)

func PostUploadFileWithLemm(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Ошибка формы: %s", err.Error())})
		return
	}

	files := form.File["files"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Файлы не найдены"})
		return
	}

	documents := make([]string, 0, len(files))
	
	for _, file := range files {
		openedFile, err := file.Open()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Не удалось открыть файл %s: %s", file.Filename, err.Error())})
			return
		}
		defer openedFile.Close()

		content, err := io.ReadAll(openedFile)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Не удалось прочитать файл %s: %s", file.Filename, err.Error())})
			return
		}
		
		documents = append(documents, string(content))
	}

	chOut := make(chan []string)
	wg := &sync.WaitGroup{}

	for _, doc := range documents {
		wg.Add(1)

		go func() {
			defer wg.Done()

			chOut <- mystem.Literalize(doc)
		}()
	}

	go func() {
		wg.Wait()
		close(chOut)
	}()

	var literalizedDocuments = make([][]string, 0, len(documents))

	for item := range(chOut) {
		literalizedDocuments = append(literalizedDocuments, item)
	}

	tf, idf := analys.GetTfIdf(literalizedDocuments...)

	c.JSON(http.StatusOK, gin.H{
		"tf": tf,
		"idf": idf,
	})
}

func PostUploadFile(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Ошибка формы: %s", err.Error())})
		return
	}

	files := form.File["files"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Файлы не найдены"})
		return
	}

	documents := make([]string, 0, len(files))
	
	for _, file := range files {
		openedFile, err := file.Open()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Не удалось открыть файл %s: %s", file.Filename, err.Error())})
			return
		}
		defer openedFile.Close()

		content, err := io.ReadAll(openedFile)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Не удалось прочитать файл %s: %s", file.Filename, err.Error())})
			return
		}
		
		documents = append(documents, string(content))
	}

	chOut := make(chan []string)
	wg := &sync.WaitGroup{}
	reg := regexp.MustCompile(`[а-яА-Яa-zA-Z]+`)

	for _, doc := range documents {
		wg.Add(1)

		go func() {
			defer wg.Done()

			chOut <- reg.FindAllString(doc, -1)
		}()
	}

	go func() {
		wg.Wait()
		close(chOut)
	}()

	var literalizedDocuments = make([][]string, 0, len(documents))

	for item := range(chOut) {
		literalizedDocuments = append(literalizedDocuments, item)
	}

	tf, idf := analys.GetTfIdf(literalizedDocuments...)

	c.JSON(http.StatusOK, gin.H{
		"tf": tf,
		"idf": idf,
	})
}