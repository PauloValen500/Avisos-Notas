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

    id, _ := result.LastInsertId()
    json.NewEncoder(w).Encode(map[string]interface{}{ "id": id, "fecha": aviso.Fecha, "mensaje": aviso.Mensaje })
}

func deleteAviso(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    _, err := db.Exec("DELETE FROM avisos WHERE id = ?", id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

func updateAviso(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    var aviso struct {
        Fecha   string `json:"fecha"`
        Mensaje string `json:"mensaje"`
    }

    if err := json.NewDecoder(r.Body).Decode(&aviso); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    _, err := db.Exec("UPDATE avisos SET fecha = ?, mensaje = ? WHERE id = ?", aviso.Fecha, aviso.Mensaje, id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

func getNotas(w http.ResponseWriter, r *http.Request) {
    rows, err := db.Query("SELECT id, contenido FROM notas")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var notas []map[string]string
    for rows.Next() {
        var id, contenido string
        if err := rows.Scan(&id, &contenido); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        notas = append(notas, map[string]string{"id": id, "contenido": contenido})
    }

    json.NewEncoder(w).Encode(notas)
}

func createNota(w http.ResponseWriter, r *http.Request) {
    var nota struct {
        Contenido string `json:"contenido"`
    }
    if err := json.NewDecoder(r.Body).Decode(&nota); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    result, err := db.Exec("INSERT INTO notas (contenido) VALUES (?)", nota.Contenido)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    id, _ := result.LastInsertId()
    json.NewEncoder(w).Encode(map[string]interface{}{ "id": id, "contenido": nota.Contenido })
}

func deleteNota(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    _, err := db.Exec("DELETE FROM notas WHERE id = ?", id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

func updateNota(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    var nota struct {
        Contenido string `json:"contenido"`
    }

    if err := json.NewDecoder(r.Body).Decode(&nota); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    _, err := db.Exec("UPDATE notas SET contenido = ? WHERE id = ?", nota.Contenido, id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
