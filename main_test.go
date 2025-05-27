package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

// Test para verificar que la ruta principal sirva archivos correctamente
func TestRutaPrincipalSirveArchivo(t *testing.T) {
    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatalf("❌ No se pudo crear la solicitud: %v", err)
    }

    rr := httptest.NewRecorder()
    handler := http.FileServer(http.Dir("./"))
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("❌ Esperado código de estado %d, pero se obtuvo %d", http.StatusOK, status)
    } else {
        t.Logf("✅ Código de estado correcto: %d", status)
    }

    if rr.Body.String() == "" {
        t.Error("❌ La respuesta está vacía, se esperaba contenido HTML u otro archivo servido")
    } else {
        t.Log("✅ La respuesta contiene contenido, la petición fue exitosa")
    }
}
