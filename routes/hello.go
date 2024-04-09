package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"routes/util"
)

func main() {
	http.HandleFunc("/somar/", handleSum)
	http.HandleFunc("/subtrair/", handleSubtract)
	http.HandleFunc("/multiplicar/", handleMultiply)
	http.HandleFunc("/dividir/", handleDivide)

	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}

func handleSum(w http.ResponseWriter, r *http.Request) {
	a, b, err := parseURLValues(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := util.Sum(a, b)
	fmt.Fprintf(w, "Resultado da soma: %d", result)
}

func handleSubtract(w http.ResponseWriter, r *http.Request) {
	a, b, err := parseURLValues(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := util.Subtract(a, b)
	fmt.Fprintf(w, "Resultado da subtração: %d", result)
}

func handleMultiply(w http.ResponseWriter, r *http.Request) {
	a, b, err := parseURLValues(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := util.Multiply(a, b)
	fmt.Fprintf(w, "Resultado da multiplicação: %d", result)
}

func handleDivide(w http.ResponseWriter, r *http.Request) {
	a, b, err := parseURLValues(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := util.Divide(a, b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Resultado da divisão: %d", result)
}

func parseURLValues(path string) (int, int, error) {
	parts := strings.Split(path, "/")
	if len(parts) != 4 {
		return 0, 0, fmt.Errorf("URL inválida")
	}

	a, err := strconv.Atoi(parts[2])
	if err != nil {
		return 0, 0, fmt.Errorf("valor de a inválido")
	}

	b, err := strconv.Atoi(parts[3])
	if err != nil {
		return 0, 0, fmt.Errorf("valor de b inválido")
	}

	return a, b, nil
}
