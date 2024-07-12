csdd1008_week11
Instructor
| Instructor | Maziar Sojoudian |

Simple API Server with Go and Docker
Docker Hub Link
(https://hub.docker.com/repository/docker/keyur212/goapi/tags)

git link
https://github.com/Meet-Sutariya/assignment-week11

Overview
This project is a simple API server built using the Go programming language and Docker. The server has two endpoints:

/: Returns a welcome message.
/about: Returns information about the API.
Development
Initialize Go Module
Create a project directory and navigate into it:
mkdir assignment-week11
cd assignment-week11
Initialize a new Go module go mod init assignment-week11
create main go
package main

import ( "fmt" "log" "net/http" )

func homeHandler(w http.ResponseWriter, r *http.Request) { if r.URL.Path != "/" { http.NotFound(w, r) } fmt.Fprintln(w, "Welcome to the Go API Server!") }

func aboutHandler(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, "This is the About page.") }

func main() { mux := http.NewServeMux() mux.HandleFunc("/", homeHandler) mux.HandleFunc("/about", aboutHandler)

log.Println("Starting server on :3001")
err := http.ListenAndServe(":3001", mux)
if err != nil {
	log.Fatalf("Error starting server: %s\n", err)
}
}

