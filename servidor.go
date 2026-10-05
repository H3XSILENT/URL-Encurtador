package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Headers map[string]string

type url struct {
	Id      string
	Criacao time.Time
	Destino string
}

var (
	port    int
	urlBase string
)

func init() {
	port = 8888
	urlBase = fmt.Sprintf("http://localhost:%d", port)
}

func main() {
	http.HandleFunc("/api/encurtar", Encurtador)
	http.HandleFunc("/r/", Redirecionador)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%d", port), nil))
}

func Encurtador(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		responderCom(w, http.StatusMethodNotAllowed, Headers{
			"Allow": "POST",
		})
		return
	}
	url, nova, err := url.BuscarOuCriarNovaUrl(extrairUrl(r))

	if err != nil {
		responderCom(w, http.StatusBadRequest, nil)
		return
	}

	var status int
	if nova {
		status = http.StatusCreated
	} else {
		status = http.StatusOK
	}

	urlCurta := fmt.Sprintf("%s/r/%s", urlBase, url.Id)
	responderCom(w, status, Headers{"Location": urlCurta})
}

func responderCom(
	w http.ResponseWriter,
	status int,
	headers Headers,
) {
	for k, v := range headers {
		w.Header().Set(k, v)
	}
	w.WriteHeader(status)
}

func ExtrairUrl(r *http.Request) string {
	url := make([]byte, r.ContentLength, r.ContentLength)
	r.Body.Read(url)
	return string(url)
}

func Redirecionador(w http.ResponseWriter, r *http.Request) {
	caminho := string.Split(r.URL.Path, "/")
	id := caminho[len(caminho)-1]

	if url := url.Buscar(id); url != nil {
		http.Redirect(w, r, url.Destino, http.StatusMovedPermanently)
	} else {
		http.NotFound(w, r)
	}
}
