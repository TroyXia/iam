package apiserver

import "github.com/TroyXia/iam/internal/apiserver/config"

func Run(cfg *config.Config) error {
	server, err := createAPIServer(cfg)
	if err != nil {
		return err
	}

	return server.PrepareRun().Run()
}
