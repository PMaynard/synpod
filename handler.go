package main

import (
    "html"
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
)

// Hack to debug gpodder.
func YES(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
}

func UploadSubDevice(w http.ResponseWriter, r *http.Request) {
    // Reply to client.

    err := r.ParseForm()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, "")

    // Handle the data.
    vars := mux.Vars(r)
    username := html.EscapeString(vars["username"])
    deviceid := html.EscapeString(vars["deviceid"])
    
    fmt.Println("UploadSubDevice")

    for key, values := range r.Form {   // range over map
      for _, value := range values {    // range over []string
         fmt.Println(username, deviceid, key, value)
      }
    }
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
    todos := Todos{
        Todo{Name: "Write presentation"},
        Todo{Name: "Host meetup"},
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    if err := json.NewEncoder(w).Encode(todos); err != nil {
        panic(err)
    }
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    todoId, err := strconv.Atoi(html.EscapeString(vars["todoId"]))
    if err != nil {
        fmt.Println(err)
    }
    fmt.Fprintln(w, "Todo show:", FindTodo(todoId))
}

