package handler

import (
	"net/http"

	logic "mwp3000/api/internal/logic/datasync"
	"mwp3000/api/internal/svc"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func SyncPingHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewSyncPingLogic(r.Context(), ctx)
		err := l.SyncPing()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
