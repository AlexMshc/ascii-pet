package main

import (
	"net"
	"net/http"
	"strconv"

	"ascii-pet/internal/gen/restapi"
	ops "ascii-pet/internal/gen/restapi/operations"
	petops "ascii-pet/internal/gen/restapi/operations/pet"
	"ascii-pet/internal/server"
	"ascii-pet/internal/store"

	"github.com/go-openapi/loads"
	"github.com/rs/cors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	doc, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load swagger spec")
	}

	api := ops.NewPetAPI(doc)
	fs := store.NewFileStore("data/pet.json")
	impl := server.NewAPI(fs)
	api.PetGetPetHandler = petops.GetPetHandlerFunc(impl.HandleGetPet)
	api.PetUploadPetHandler = petops.UploadPetHandlerFunc(impl.HandleUploadPet)
	api.PetDeletePetHandler = petops.DeletePetHandlerFunc(impl.HandleDeletePet)

	srv := restapi.NewServer(api)
	defer srv.Shutdown()
	srv.ConfigureAPI()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPut, http.MethodOptions, http.MethodDelete},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	orig := srv.GetHandler()
	srv.SetHandler(c.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info().
			Str("method", r.Method).
			Str("path", r.URL.Path).
			Msg("incoming request")
		orig.ServeHTTP(w, r)
	})))

	specHost := doc.Spec().Host
	host, portStr, err := net.SplitHostPort(specHost)
	if err != nil {
		host = specHost
	}

	if host == "localhost" {
		host = "0.0.0.0"
	}
	srv.Host = host

	if portStr != "" {
		if port, err := strconv.Atoi(portStr); err == nil {
			srv.Port = port
		}
	}

	log.Info().Msgf("listening on %s:%d", srv.Host, srv.Port)
	if err := srv.Serve(); err != nil {
		log.Fatal().Err(err).Msg("server stopped")
	}
}
