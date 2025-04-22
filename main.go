package main

import (
    "database/sql"
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    _ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func initDB() {
    var err error
    db, err = sql.Open("sqlite3", "./notas-y-avisos.db")
    if err != nil {
        log.Fatal(err)
    }

    sqlStmt := `
    CREATE TABLE IF NOT EXISTS avisos (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        fecha TEXT,
        mensaje TEXT
    );`
    _, err = db.Exec(sqlStmt)
    if err != nil {
        log.Printf("%q: %s\n", err, sqlStmt)
        return
    }

    sqlStmt = `
    CREATE TABLE IF NOT EXISTS notas (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        contenido TEXT
    );`
    _, err = db.Exec(sqlStmt)
    if err != nil {
        log.Printf("%q: %s\n", err, sqlStmt)
        return
    }
}

func main() {
    initDB()
    defer db.Close()

    r := mux.NewRouter()

    r.HandleFunc("/avisos", getAvisos).Methods("GET")
    r.HandleFunc("/avisos", createAviso).Methods("POST")
    r.HandleFunc("/avisos/{id}", deleteAviso).Methods("DELETE")
    r.HandleFunc("/avisos/{id}", updateAviso).Methods("PUT")

    r.HandleFunc("/notas", getNotas).Methods("GET")
    r.HandleFunc("/notas", createNota).Methods("POST")
    r.HandleFunc("/notas/{id}", deleteNota).Methods("DELETE")
    r.HandleFunc("/notas/{id}", updateNota).Methods("PUT")

    r.PathPrefix("/").Handler(http.FileServer(http.Dir("./")))

    log.Println("Servidor iniciado en http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}

func getAvisos(w http.ResponseWriter, r *http.Request) {
    rows, err := db.Query("SELECT id, fecha, mensaje FROM avisos")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var avisos []map[string]string
    for rows.Next() {
        var id, fecha, mensaje string
        if err := rows.Scan(&id, &fecha, &mensaje); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        avisos = append(avisos, map[string]string{"id": id, "fecha": fecha, "mensaje": mensaje})
    }

    json.NewEncoder(w).Encode(avisos)
}

func createAviso(w http.ResponseWriter, r *http.Request) {
    var aviso struct {
        Fecha   string `json:"fecha"`
        Mensaje string `json:"mensaje"`
    }
    if err := json.NewDecoder(r.Body).Decode(&aviso); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    result, err := db.Exec("INSERT INTO avisos (fecha, mensaje) VALUES (?, ?)", aviso.Fecha, aviso.Mensaje)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }