package handlers

import (
    "fmt"
    "io"
    "net/http"
	"os"
"path/filepath"
	"time"
"github.com/Yandex-Practicum/go1fl-sprint6-final/service"
)

func MainHandle(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "index.html")
}

func Upload(w http.ResponseWriter, r *http.Request) {
   file, header, err := r.FormFile("myFile")
defer file.Close()



fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Ошибка чтения данных файла: "+err.Error(), http.StatusInternalServerError)
		return
	}

convertedString := service.Detect(string(fileBytes))
	if err != nil {
		http.Error(w, "Ошибка конвертации: "+err.Error(), http.StatusInternalServerError)
		return
	}

ext := filepath.Ext(header.Filename)
	
	timestamp := time.Now().UTC().Format("20060102_150405_000000")
	localFileName := fmt.Sprintf("converted_%s%s", timestamp, ext)

dst, err := os.Create(localFileName)
	if err != nil {
		http.Error(w, "Не удалось создать локальный файл", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

_, err = dst.WriteString(convertedString)
	if err != nil {
		http.Error(w, "Ошибка записи в файл", http.StatusInternalServerError)
		return
	}
w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(convertedString))

}

//func main() {
//    http.HandleFunc("/", mainHandle)
//    http.HandleFunc("/upload", upload)
//    err := http.ListenAndServe(":8080", nil)
//    if err != nil {
//        panic(err)
//    }
//} 
