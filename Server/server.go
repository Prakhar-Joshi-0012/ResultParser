package server

import (
	Parser "ResultParser/Parser"
	"fmt"
	"html/template"
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("Server/templates/index.html"))
	t.Execute(w, "Here are your file entries......")

}

var uploadpath string

const MaxUploadSize int = 2 * 1024 * 1024

func DownloadHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseMultipartForm(int64(MaxUploadSize)); err != nil {
		fmt.Printf("Could not parse Multipart form: %v\n", err)
		renderError(w, "Cant Parse Form", http.StatusInternalServerError)
		return
	}
	file, fileHeader, err := r.FormFile("ufile")
	std, _ := strconv.Atoi(r.FormValue("Standard"))
	if err != nil {
		renderError(w, "Invalid File", http.StatusBadRequest)
		return
	}
	defer file.Close()

	filesize := fileHeader.Size
	fmt.Printf("File size (bytes): %v\n", filesize)
	if int64(filesize) > int64(MaxUploadSize) {
		renderError(w, "FileTooBig", http.StatusBadRequest)
		return
	}
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		renderError(w, "Invalid File", http.StatusBadRequest)
		return
	}
	fileType := http.DetectContentType(fileBytes)
	fileName := "UploadedFile"
	fileEndings, err := mime.ExtensionsByType(fileType)
	if err != nil {
		renderError(w, "Cannot Read File Type", http.StatusInternalServerError)
		return
	}
	newFileName := fileName + fileEndings[0]
	newPath := filepath.Join(uploadpath, newFileName)
	fmt.Printf("FileType : %s, File: %s\n", fileType, newPath)

	// write file
	newFile, err := os.Create(newPath)
	if err != nil {
		renderError(w, "Cannot Write File", http.StatusInternalServerError)
		return
	}
	defer file.Close()
	if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
		renderError(w, "Cannot Write File", http.StatusInternalServerError)
		return
	}
	Parser.Parse(std, newPath)
	xlfile, _ := os.Open("test.xlsx")
	xlfilebytes, _ := io.ReadAll(xlfile)
	xlfileType := http.DetectContentType(xlfilebytes)

	w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote(xlfile.Name()))
	w.Header().Set("Content-Type", xlfileType+";"+xlfile.Name())
	// w.Header().Set("Content-Length", strconv.Itoa(int(filesize)))

	xlfile.Seek(0, 0)
	io.Copy(w, xlfile)
}

func renderError(w http.ResponseWriter, msg string, statusCode int) {
	w.WriteHeader(statusCode)
	w.Write([]byte(msg))
}

func CreateServer() {
	server := http.Server{
		Addr: "127.0.0.1:8000",
	}
	http.HandleFunc("/", Handler)
	var err error
	uploadpath, err = os.Getwd()
	if err != nil {
		fmt.Printf("Failed to get uploadpath")
	}
	uploadpath += "/Uploaded"
	http.HandleFunc("/download", DownloadHandler)
	fs := http.FileServer(http.Dir(uploadpath))
	http.Handle("/files/", http.StripPrefix("/files", fs))
	err = server.ListenAndServe()
	if err != nil {
		fmt.Printf(err.Error())
	}

}
