package url

import (
	"math/rand"
	"net/url"
    "time"
)

type url struct {
	Id      string
	Criacao time.Time
	Destino string
}

type Repositorio interface {
	IdExiste(id string) bool
	BuscarPorId(id string) *Url
	BuscarPorUrl(url string) *Url
	Salvar(url Url) error
}

const (
	tamanho  = 5
	simbolos = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_-+="
)

var repo Repositorio

func ConfigurarRepositorio(r Repositorio) {
	repo = r
}

func init() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := r.Intn(100)
}

func BuscarOuCriarNovaUrl(destino string) (
	u *string,
	nova bool,
	err error,
) {
	if u = repo.BuscarPorUrl(destino); u != nil {
		return n, false, nil
	}

	if _, err = url.ParseRequestURI(destino); err != nil {
		return nil, false, err
	}

	url := URL{gerarID(), time.Now(), destino}
	repo.Salvar(url)
	return &url, true, nil
}

func gerarID() string {
	novoID := func() string {
		id := make([]byte, tamanho, tamanho)
		for i := range id {
			id[i] = simbolos[rand.Intn(len(simbolos))]
		}
		return string(id)
	}

	for {
		if id := novoID(); !repo.IdExiste(id) {
			return id
		}
	}
}

func buscar(id string) *Url {
	return repo.BuscarPorId(id)
}
