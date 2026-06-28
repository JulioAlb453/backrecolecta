package application_ciudadano

import (
	"context"
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/vicpoo/API_recolecta/src/Ciudadanos/domain"
)

type UpdateMyCiudadanoInput struct {
	UserID          int
	Alias           *string
	Password        *string
	CurrentPassword string
}

type UpdateMyCiudadano struct {
	repo domain.CiudadanoRepository
}

func NewUpdateMyCiudadano(repo domain.CiudadanoRepository) *UpdateMyCiudadano {
	return &UpdateMyCiudadano{repo: repo}
}

func (uc *UpdateMyCiudadano) Execute(ctx context.Context, in UpdateMyCiudadanoInput) error {
	if in.Alias == nil && in.Password == nil {
		return errors.New("debe enviar alias o password para actualizar")
	}

	ciudadano, err := uc.repo.GetByID(ctx, in.UserID)
	if err != nil {
		return err
	}
	if ciudadano == nil {
		return errors.New("ciudadano no encontrado")
	}

	if in.Password != nil {
		current := strings.TrimSpace(in.CurrentPassword)
		if current == "" {
			return errors.New("current_password es requerido para cambiar la contraseña")
		}
		if err := bcrypt.CompareHashAndPassword([]byte(ciudadano.Password), []byte(current)); err != nil {
			return errors.New("contraseña actual incorrecta")
		}
	}

	updateInput := UpdateCiudadanoInput{ID: in.UserID}
	if in.Alias != nil {
		updateInput.Alias = in.Alias
	}
	if in.Password != nil {
		updateInput.Password = in.Password
	}

	updater := NewUpdateCiudadano(uc.repo)
	return updater.Execute(ctx, updateInput)
}
