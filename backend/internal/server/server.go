package server

import (
	"errors"
	"net/http"

	"ascii-pet/internal/gen/models"
	petops "ascii-pet/internal/gen/restapi/operations/pet"
	"ascii-pet/internal/store"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

type API struct {
	store *store.FileStore
}

func NewAPI(s *store.FileStore) *API {
	return &API{store: s}
}

func (a *API) HandleGetPet(_ petops.GetPetParams) middleware.Responder {
	p, err := a.store.Load()
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			return petops.NewGetPetNoContent()
		}
		return middleware.ResponderFunc(func(rw http.ResponseWriter, _ runtime.Producer) {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		})
	}
	return petops.NewGetPetOK().
		WithPayload(&models.PetProperties{
			ASCII:       p.ASCII,
			Description: p.Description,
		})
}

func (a *API) HandleUploadPet(params petops.UploadPetParams) middleware.Responder {
	body := params.PetProperties

	if body.ASCII == "" || body.Description == "" {
		return petops.NewUploadPetBadRequest()
	}

	err := a.store.Save(models.PetProperties{
		ASCII:       body.ASCII,
		Description: body.Description,
	})
	if err != nil {
		return middleware.ResponderFunc(func(rw http.ResponseWriter, _ runtime.Producer) {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		})
	}

	return petops.NewUploadPetOK()
}

func (a *API) HandleDeletePet(_ petops.DeletePetParams) middleware.Responder {
	if err := a.store.Delete(); err != nil && !errors.Is(err, store.ErrNotFound) {
		return middleware.ResponderFunc(func(rw http.ResponseWriter, _ runtime.Producer) {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		})
	}

	return petops.NewDeletePetNoContent()
}
