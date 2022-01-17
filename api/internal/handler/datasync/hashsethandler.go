package handler

import (
	"net/http"

	logic "mwp3000/api/internal/logic/datasync"
	"mwp3000/api/internal/svc"
	"mwp3000/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func HashSetHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SetHashReq
		// log.Infof("request:>>>", r)
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewHashSetLogic(r.Context(), ctx)
		resp, err := l.HashSet(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
