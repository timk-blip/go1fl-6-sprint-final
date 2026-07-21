package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

type LoggerInfo struct {
	*log.Logger
}

func (l *LoggerInfo) Info(v ...interface{}) {
	args := append([]interface{}{"[INFO]"}, v...)
	l.Println(args...)
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
	return
}

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "только POST.", http.StatusMethodNotAllowed)
		return
	}
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
	nameTime := time.Now().UTC().String()
	fileLocal, err := os.Create(nameTime + ext)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer fileLocal.Close()

	_, err = fileLocal.WriteString(resultStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	baseLogger := log.New(os.Stdout, "", log.LstdFlags)
	logM := LoggerInfo{Logger: baseLogger}
	logM.Info("Строку ", "\">", string(fileBytes), "\"<", "конвертировали в ", resultStr)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(resultStr))
	if err != nil {
		logM.Info("Ошибка записи в поток вывода http")
		return
	}
	return
}
