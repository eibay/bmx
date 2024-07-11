package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os/exec"
    "strconv"
    "github.com/gorilla/mux"
    "github.com/gorilla/handlers"
)

type Book struct {
    ID     int    `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
}

var books []Book
var nextID int = 1

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/api/books", getBooks).Methods("GET")
    r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
    r.HandleFunc("/api/books", createBook).Methods("POST")
    r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
    r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
    r.HandleFunc("/api/run-script", runScript).Methods("POST")
    r.HandleFunc("/api/run-command", runCommand).Methods("POST")

    // CORS settings
    corsObj := handlers.AllowedOrigins([]string{"*"}) // Adjust this to be more restrictive if needed
    headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
    methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})

    log.Println("Server started on :8080")
     log.Fatal(http.ListenAndServe(":8080", handlers.CORS(corsObj, headersOk, methodsOk)(r)))
}

func runCommand(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    // Define the shell command to run
    cmd := exec.Command("ls", "-l") // Example: List directory contents

    // Run the command and capture its output
    output, err := cmd.CombinedOutput()
    if err != nil {
        http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
        return
    }

    // Return the output as a JSON response
    json.NewEncoder(w).Encode(map[string]string{
        "output": string(output),
    })
}

func runScript(w http.ResponseWriter, r *http.Request) {
    cmd := exec.Command("bash", "scripts/az_login.sh")
    output, err := cmd.CombinedOutput()
    if err != nil {
        http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
        return
    }
    fmt.Fprintf(w, "VNet creation output: %s", output)
}

func getBooks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    for _, book := range books {
        if book.ID == id {
            json.NewEncoder(w).Encode(book)
            return
        }
    }

    http.Error(w, "Book not found", http.StatusNotFound)
}

func createBook(w http.ResponseWriter, r *http.Request) {
    fmt.Println("CREATE BOOK ---> ")
    w.Header().Set("Content-Type", "application/json")
    var book Book
    if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    book.ID = nextID
    nextID++
    books = append(books, book)
    json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    for i, book := range books {
        if book.ID == id {
            var updatedBook Book
            if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
                http.Error(w, err.Error(), http.StatusBadRequest)
                return
            }
            updatedBook.ID = id
            books[i] = updatedBook
            json.NewEncoder(w).Encode(updatedBook)
            return
        }
    }

    http.Error(w, "Book not found", http.StatusNotFound)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    for i, book := range books {
        if book.ID == id {
            books = append(books[:i], books[i+1:]...)
            w.WriteHeader(http.StatusNoContent)
            return
        }
    }

    http.Error(w, "Book not found", http.StatusNotFound)
}


// package main

// import (
//     "fmt"
//     "net/http"
//     "os/exec"
//     "log"
// )

// func helloHandler(w http.ResponseWriter, r *http.Request) {
//     fmt.Fprintf(w, "Hello from Go backend!")
// }

// func goodbyeHandler(w http.ResponseWriter, r *http.Request) {
//     fmt.Fprintf(w, "Goodbye from Go backend!")
// }

// func createVNetHandler(w http.ResponseWriter, r *http.Request) {
//     cmd := exec.Command("/bin/bash", "create_vnet.sh")
//     output, err := cmd.CombinedOutput()
//     if err != nil {
//         http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
//         return
//     }
//     fmt.Fprintf(w, "VNet creation output: %s", output)
// }

// func main() {
//     http.HandleFunc("/api/hello", helloHandler)
//     http.HandleFunc("/api/goodbye", goodbyeHandler)
//     http.HandleFunc("/api/create-vnet", createVNetHandler)
//     fmt.Println("Server is running on port 8080")
//     log.Fatal(http.ListenAndServe(":8080", nil))
// }
