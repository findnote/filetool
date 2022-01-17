// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	datasync "mwp3000/api/internal/handler/datasync"
	"mwp3000/api/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/sync/ping",
				Handler: datasync.SyncPingHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/1.0/middle/dw/sync/setHash",
				Handler: datasync.HashSetHandler(serverCtx),
			},
		},
	)
}