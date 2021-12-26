package logic

import (
	"context"

	"mwp3000/api/internal/svc"
	"mwp3000/api/internal/types"
	"mwp3000/pkg/redis/redisapi"

	"github.com/tal-tech/go-zero/core/logx"
)

type HashSetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHashSetLogic(ctx context.Context, svcCtx *svc.ServiceContext) HashSetLogic {
	return HashSetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HashSetLogic) HashSet(req types.SetHashReq) (*types.SetHashResp, error) {
	// todo: add your logic here and delete this line
	// log.Infof("XXXXXXXXXXXXXXX %v", req.Tdrvs)
	rcli := redisapi.GetRedisClient()
	for _, index := range req.Tdrvs {
		// log.Infof("XXXXUUUUUUXXX %v", index.Hash)
		for _, info := range index.Values {
			// log.Infof("XXXXXXX>> %v XXXXX >> %v", info.Key, info.Value)
			rcli.HSet(index.Hash, info.Key, info.Value)
		}
	}

	return &types.SetHashResp{
		Success: "true",
		Errcode: "200",
	}, nil
}
