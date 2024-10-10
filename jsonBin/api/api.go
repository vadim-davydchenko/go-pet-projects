package api

import (
	"fmt"

	"jsonBin/config"
)

type API struct {
	Config *config.Config
}

func NewAPI(cfg *config.Config) *API {
	api := &API{
		Config: cfg,
	}
	fmt.Printf("API initialized with key: %s\n", api.Config.Key)

	return api
}
