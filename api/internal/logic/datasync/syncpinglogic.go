package logic

import (
	"context"

	"github.com/tal-tech/go-zero/core/logx"
	"mwp3000/api/internal/svc"
)

type SyncPingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSyncPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) SyncPingLogic {
	return SyncPingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SyncPingLogic) SyncPing() error {
	// todo: add your logic here and delete this line

	return nil
}
