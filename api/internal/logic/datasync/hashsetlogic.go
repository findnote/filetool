package logic

import (
	"context"
	"encoding/json"
	"time"

	"mwp3000/api/internal/svc"
	"mwp3000/api/internal/types"
	"mwp3000/pkg/p3000"
	"mwp3000/pkg/redis/redisapi"
	"mwp3000/pkg/utils"

	log "github.com/pion/ion-log"
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
	strTmp := utils.GetJsonStrByStruct(req)
	log.Infof("HashSet >>req: %v", strTmp)

	var syncInfo p3000.SetHashReq
	str, err := json.Marshal(req)
	if err != nil {
		log.Errorf("HashSet >> Marshal Error: %v", err)
		return &types.SetHashResp{
			Success: "false",
			Errcode: "600",
			ErrMsg:  "HashSet >> Marshal Error",
		}, nil
	}

	json.Unmarshal(str, &syncInfo)
	if err != nil {
		log.Errorf("HashSet >> Unmarshal Error: %v", err)
		return &types.SetHashResp{
			Success: "false",
			Errcode: "601",
			ErrMsg:  "HashSet >> Unmarshal Error",
		}, nil
	}

	rcli := redisapi.GetRedisClient()
	for _, index := range req.Tdrvs {
		// log.Infof("XXXXUUUUUUXXX %v", index.Hash)
		for _, info := range index.Values {
			// log.Infof("XXXXXXX>> %v XXXXX >> %v", info.Key, info.Value)
			rcli.HSet(index.Hash, info.Key, info.Value)
		}
	}
	time.Sleep(30 * time.Millisecond)
	_, errCode, err := l.svcCtx.P3000Sync.PostSync(syncInfo)
	if errCode != 0 {
		log.Errorf("HashSet >> PostSync Error: %v", err)
		return &types.SetHashResp{
			Success: "false",
			Errcode: "602",
			ErrMsg:  "HashSet >> PostSync Error",
		}, nil
	}

	return &types.SetHashResp{
		Success: "true",
		Errcode: "200",
	}, nil
}
