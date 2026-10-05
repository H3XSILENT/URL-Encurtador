package url

import (
	"math/rand"
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
	rand.Seed(time.Now().UnixNano())
}

func BuscarOuCriarNovaUrl(destino string) (
	u *string,
	nova bool,
	err error,
) {
}
