package handlers

import (
	"bufio"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/log6"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func GetRoot(w http.ResponseWriter, r *http.Request) {

	fname := "index.html"
	_, err := os.Stat(fname)
	if err != nil && errors.Is(err, os.ErrNotExist) {
		log6.Err.Println("index.html не найден:", err.Error())
		http.Error(w, "index.html не найден", http.StatusInternalServerError)
		return
	}

	var buff []byte

	buff, err = os.ReadFile(fname)
	if err != nil {
		log6.Err.Println("не удалось прочитать index.html:", err.Error())
		http.Error(w, "не удалось прочитать index.html", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(buff); err != nil {
		log6.Err.Println("не удалось отправить ответ клиенту:", err.Error())
	}
}

func Upload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		log6.Err.Println("ошибка при обработке формы:", err.Error())
		http.Error(w, "ошибка при обработке формы", http.StatusBadRequest)
		return
	}

	mfile, _, err := r.FormFile("myFile")
	if err != nil {
		log6.Err.Println("ошибка при получении файла", err.Error())
		http.Error(w, "ошибка при получении файла", http.StatusBadRequest)
		return
	}
	defer mfile.Close()

	result, err := os.Create(filepath.Join(time.Now().UTC().Format("2.1.2006-15:04:05") + filepath.Ext("file.txt")))
	if err != nil {
		log6.Err.Println("ошибка при создании файла", err.Error())
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	defer result.Close()

	scanner := bufio.NewScanner(mfile)
	for scanner.Scan() {
		resultStr, err := service.MorseConvert(scanner.Text())
		if err != nil {
			log6.Err.Println("ошибка при конвертации строки:", err.Error())
			http.Error(w, "ошибка при конвертации строки", http.StatusInternalServerError)
			return
		}

		// send response with result string
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(resultStr)); err != nil {
			log6.Err.Println("не удалось отправить ответ клиенту:", err.Error())
		}

		if _, err := fmt.Fprintln(result, resultStr); err != nil {
			log6.Err.Println("ошибка при записи в итоговый файл", err.Error())
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
	}
}
