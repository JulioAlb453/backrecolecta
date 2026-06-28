package application

import (
	"github.com/vicpoo/API_recolecta/src/domicilio/domain"
)

type ListDomicilios struct {
	repo domain.DomicilioRepository
}

func NewListDomicilios(repo domain.DomicilioRepository) *ListDomicilios {
	return &ListDomicilios{repo}
}

func (uc *ListDomicilios) ExecuteByUsuarioID(usuarioID int) ([]domain.Domicilio, error) {
	return uc.repo.ListByUsuarioID(usuarioID)
}
