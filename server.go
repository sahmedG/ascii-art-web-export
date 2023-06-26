package webart

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

/*
This file is what could be considered the router. it redirects all methods sent by the client, and
returns appropriate responses and response codes
*/

// * we create 2 structs. One Unmarshal JSON data sent by the client to the golang server as golang objects
type ASCII_ART struct {
	Text     string `json:"Text"`
	Banner   string `json:"Banner"`
	Newcolor string `json:"Newcolor"`
}

// * and one represents the result string
type RESULT_ASCII_ART struct {
	Result     string
	ApplyColor string
}

type RESULT_ASCII_EXPORT struct {
	AsciiArt string `json:"AsciiArt"`
}

//* this function is responsible for processing the GET request for the main page in the case it is requested.
//* is returns a Bad request error in the case something else rather than "/" is typed

func Handler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	//* if the HTTP method is GET, serve the HTML files in the template directory, otherwise, serve the
	//* custom HTML for bad requests
	case "GET":
		r.ParseForm()
		path := "../templates" + r.URL.Path
		http.ServeFile(w, r, path)
	default:
		w.WriteHeader(http.StatusBadRequest)
		http.ServeFile(w, r, "../templates/badrequest.html")
	}
}

//* this function basically appends the ascii art to the result div in HTML

func Gen_ASCII(w http.ResponseWriter, r *http.Request) {
	var ascii_art ASCII_ART
	var ascii_result RESULT_ASCII_ART

	if r.Method != http.MethodPost {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&ascii_art)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(ascii_art.Text) > 255 {
		http.Error(w, "Too much letters خوك", http.StatusBadRequest)
		return
	}

	if MapFont(ascii_art.Banner) == "Error 500" {
		http.Error(w, "Error 500: internarl server error ", http.StatusInternalServerError)
	} else {
		ascii_result.Result = PrintART(ascii_art.Text, ascii_art.Banner)
		ascii_result.ApplyColor = ascii_art.Newcolor
		jsonENC := json.NewEncoder(w)
		jsonENC.Encode(ascii_result)
	}
}

func FileDownload(w http.ResponseWriter, r *http.Request, Filename string) {
	fmt.Println("Client requests: " + Filename)

	Openfile, err := os.Open(Filename)
	if err != nil {
		fmt.Print("file doesn't excist")
	}
	defer Openfile.Close() //Close after function return

	//File is found, create and send the correct headers

	//Get the Content-Type of the file
	//Create a buffer to store the header of the file in
	FileHeader := make([]byte, 512)
	//Copy the headers into the FileHeader buffer
	Openfile.Read(FileHeader)
	//Get content type of file
	FileContentType := http.DetectContentType(FileHeader)

	//Get the file size
	FileStat, _ := Openfile.Stat()                     //Get info from file
	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

	//Send the headers
	w.Header().Set("Content-Disposition", "attachment; filename="+Filename)
	w.Header().Set("Content-Type", FileContentType)
	w.Header().Set("Content-Length", FileSize)

	//Send the file
	//We read 512 bytes from the file already, so we reset the offset back to 0
	Openfile.Seek(0, 0)
	io.Copy(w, Openfile) //'Copy' the file to the client
}

// exportHandler handle exportation system
func ExportHandler(w http.ResponseWriter, r *http.Request) {
	var ascii_art RESULT_ASCII_EXPORT
	err := json.NewDecoder(r.Body).Decode(&ascii_art)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("expothandler success")
	ExportTXT(ascii_art.AsciiArt)
	FileDownload(w, r, "../export.txt")
}

// exportTXT create a .txt file and put ascii-art inside
func ExportTXT(Text string) {
	file, err := os.Create("../export.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.WriteString(Text)
}
