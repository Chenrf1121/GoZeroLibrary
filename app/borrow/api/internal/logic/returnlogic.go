package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"library/app/borrow/api/internal/svc"
	"library/app/borrow/api/internal/types"
	"library/app/borrow/model"
	"library/app/login/rpc/login"
	"library/app/search/rpc/search"
	"library/common/errorx"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReturnLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReturnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReturnLogic {
	return &ReturnLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReturnLogic) Return(req *types.ReturnReq) (resp *types.ReturnResq, err error) {
	// todo: 还书
	//获取还书用户的信息
	userIdNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value("userId")))
	logx.Infof("userIdL %s", userIdNumber)
	userId, err := userIdNumber.Int64()
	if err != nil {
		return &types.ReturnResq{Status: false}, errorx.NewCodeError(10003, "请重新登录！")
	}
	//通过rpc通信得到用户和书本的id
	userinfo, err := l.svcCtx.UserRpc.Login(l.ctx, &login.IdReq{Id: strconv.Itoa(int(userId))})
	bookinfo, err := l.svcCtx.SearchRpc.Search(l.ctx, &search.SearchReq{Name: req.Name})
	db, err := gorm.Open("mysql", l.svcCtx.Config.Mysql.DataSource)

	if err != nil {
		return nil, errorx.NewCodeError(10008, "gorm连接数据库错误！")
	}
	defer db.Close()
	db.AutoMigrate(&model.Borrow{})

	var tmpinfo model.Borrow
	//	fmt.Println("userinfo.Number, bookinfo.Id  = ", userinfo.Number, bookinfo.Id)
	db.Debug().Table("borrow").Where("user_id = ? AND book_id = ?", userinfo.Number, bookinfo.Id).Last(&tmpinfo)
	if tmpinfo.Id == 0 {
		return nil, errorx.NewCodeError(10011, "你没有借过这本书。")
	}
	tmpinfo.ReturnTime = time.Now()
	tmpinfo.Isreturn = 1
	db.Debug().Table("borrow").Save(tmpinfo)
	//更新图书数据库内容
	l.svcCtx.SearchRpc.Return(l.ctx, &search.ReturnReq{Id: bookinfo.Id})

	return &types.ReturnResq{Status: true}, nil
}
