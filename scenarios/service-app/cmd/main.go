package main

import (
	"html"
	"io"
	"net/http"
	"time"

	"github.com/talbs1986/simplego/app/pkg/app"
	"github.com/talbs1986/simplego/app/pkg/logger"
	simplego_config "github.com/talbs1986/simplego/configs/pkg/app"
	"github.com/talbs1986/simplego/scenarios/service-app/pkg/scenarios"
	simplego_server "github.com/talbs1986/simplego/server/pkg/app"
	"github.com/talbs1986/simplego/server/pkg/server"
)

type SomeConfig struct {
}

func main() {
	scenarios.StartService[SomeConfig](
		&scenarios.ServiceConfig[SomeConfig]{
			AppConfig: app.AppConfig{
				Name:                "service-exmaple",
				Version:             "1.0.0",
				ServiceCloseTimeout: time.Second * 30,
			},
		},
		proc)
}

func proc(appObj *app.App) error {
	cfg, err := simplego_config.GetConfig[scenarios.ServiceConfig[SomeConfig]](appObj)
	if err != nil {
		return err
	}
	srvr := simplego_server.GetServerService(appObj)
	middlewares := srvr.GetMiddlewares()
	if err := srvr.RegisterRoute(server.ServerRoute{
		Method: http.MethodPost,
		Route:  "/echo",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			bs, err := io.ReadAll(r.Body)
			if err != nil {
				appObj.Logger.Log().Error(err, "service: failed read request")
				_, _ = w.Write([]byte("failed to read request"))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			if _, err := w.Write([]byte(html.EscapeString(string(bs)))); err != nil {
				appObj.Logger.Log().With(&logger.LogFields{"req": string(bs)}).Error(err, "service: failed to write response")
				_, _ = w.Write([]byte("failed to write response"))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		},
	}); err != nil {
		appObj.Logger.Log().Fatal(err, "simplego service: failed to register route")
	}
	appObj.Logger.Log().With(&logger.LogFields{"config": cfg, "middlewares": middlewares}).Info("service: finnished registering echo route")
	return nil
}
