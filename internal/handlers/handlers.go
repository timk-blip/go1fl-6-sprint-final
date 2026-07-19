package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/Users/kvako/GolandProjects/go1fl-6-sprint-final/index.html")
}

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resultStr := service.IsParseable(string(fileBytes))

	ext := filepath.Ext(header.Filename)
	nameTime := time.Now().UTC().Format("20060102-150405")
	fileLocal, err := os.Create(nameTime + ext)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer fileLocal.Close()

	_, err = fileLocal.WriteString(resultStr)
	if err != nil {
		log.Fatalf("Ошибка при записи в файл: %v", err)
	}
	fmt.Println(resultStr)
	return
}
