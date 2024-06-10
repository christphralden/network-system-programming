package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Received %s request for %s\n", r.Method, r.URL.Path)
		host, port, _ := net.SplitHostPort(r.RemoteAddr)
		fmt.Printf("HOST: %v, PORT: %v\n", host, port)

		if r.Method == http.MethodGet || r.Method == http.MethodPost {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
}

func fileProcessingHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10 MB max file size
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Unable to get file from form", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File size: %d bytes\n", len(fileBytes))
}

func handleRoutes() {
	mux := http.NewServeMux()

	mux.Handle("/upload", loggingMiddleware(http.HandlerFunc(fileProcessingHandler)))

	go func() {
		if err := http.ListenAndServe(":8080", mux); err != nil {
			log.Fatalf("Server failed: %v", err)
		}
	}()
}

func writeFile(filename string) {
	fmt.Print("Enter the content to write: ")

    in := bufio.NewReader(os.Stdin)
    content, err := in.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	err = ioutil.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
}

func readFile(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	fmt.Println("Read from file:", string(data))
}

func simulateClient(filename string, endpoint string) {
	imageData, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file file: %v", err)
	}

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		log.Fatalf("Error creating form file: %v", err)
	}
	part.Write(imageData)

	writer.Close()

	req, err := http.NewRequest(http.MethodPost, endpoint, body)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	fmt.Println("Server response:", string(respData))
}

func main() {
    var filename string
    fmt.Print("Enter filename: ")
    fmt.Scanln(&filename)
    splitted := strings.Split(filename, ".")

    if len(splitted) > 1 && splitted[len(splitted)-1] == "txt"{
        writeFile(filename)
        readFile(filename)
    }

	handleRoutes()

	time.Sleep(1 * time.Second)
	// clientWithTimeout()
    simulateClient(filename, "http://localhost:8080/upload")
}