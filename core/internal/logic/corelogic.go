package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type CoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CoreLogic {
	return &CoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CoreLogic) Core(req *types.Request) (resp *types.Response, err error) {
	// 获取用户的列表信息
	data := make([]*models.UserBasic, 0)
	err = models.Engine.Find(&data)
	if err != nil {
		log.Println("Get UserBasic Error:", err)
	}
	b, err := json.Marshal(data)
	if err != nil {
		log.Println("JSON Marshal Error:", err)
	}
	dst := new(bytes.Buffer)
	err = json.Indent(dst, b, "", " ")
	if err != nil {
		log.Println("JSON Indent Error:", err)
	}
	fmt.Println(dst.String())
	resp = new(types.Response)
	resp.Message = dst.String()
	return
}
