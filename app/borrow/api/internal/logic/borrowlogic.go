package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"library/app/search/rpc/search"
	"time"

	"library/app/login/rpc/login"
	"library/common/errorx"
	"log"
	"strconv"

	"library/app/borrow/api/internal/svc"
	"library/app/borrow/api/internal/types"

	"github.com/jinzhu/gorm"
	"github.com/zeromicro/go-zero/core/logx"
	"library/app/borrow/model"
)

type BorrowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBorrowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BorrowLogic {
	return &BorrowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BorrowLogic) Borrow(req *types.BorrowReq) (*types.BorrowResq, error) {
	// todo: 借书
	db, err := gorm.Open("mysql", l.svcCtx.Config.Mysql.DataSource)

	if err != nil {
		log.Println("grom打开数据库错误，原因为：", err)
	}
	defer db.Close()
	db.AutoMigrate(&model.Borrow{})

	//定义数据库传出的类型数据
	userIdNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value("userId")))
	logx.Infof("userId: %s", userIdNumber)
	userId, err := userIdNumber.Int64()
	if err != nil {
		return nil, errorx.NewCodeError(10003, "请重新登录!")
	}

	userinfo, err := l.svcCtx.UserRpc.Login(l.ctx, &login.IdReq{Id: strconv.Itoa(int(userId))})
	logx.Info("user info:", userinfo.Name)
	bookinfo, err := l.svcCtx.SearchRpc.Search(l.ctx, &search.SearchReq{Name: req.Name})

	//bookinfo, err := l.svcCtx.BooksModel.FindOneByName(l.ctx, req.Name)
	if err == model.ErrNotFound {
		return nil, errorx.NewCodeError(10005, "查无此书，无法借阅!")
	} else if err != nil {
		return nil, err
	}
	if bookinfo.CountNow > 0 {
		//可以借书
		var tmpinfo model.Borrow
		fmt.Println("userinfo.Number, bookinfo.Id == ", userinfo.Number, bookinfo.Id)
		db.Debug().Table("borrow").Where("user_id = ? AND book_id = ?", userinfo.Number, bookinfo.Id).First(&tmpinfo)

		fmt.Println("tmpinfo == ", tmpinfo.Id)
		if tmpinfo.Id != 0 && tmpinfo.Isreturn == 0 {
			//已经借了一本
			return &types.BorrowResq{Status: false}, errorx.NewCodeError(10007, "已经借了一本")
		} else {
			//可以借
			tmpinfo.UserId = userinfo.Number
			tmpinfo.BookId = bookinfo.Id
			tmpinfo.BorrowTime = time.Now()
			tmpinfo.ReturnTime = time.Now()
			tmpinfo.Isreturn = 0
			fmt.Println("tmpinfo  == ", tmpinfo)
			_, err = l.svcCtx.BorrowModel.Insert(l.ctx, &tmpinfo)
			if err != nil {
				fmt.Println("err = ", err)
				return nil, errorx.NewCodeError(10008, "借阅失败")
			}
			//通过rpc告诉search服务把书给借走
			_, err := l.svcCtx.SearchRpc.Borrow(l.ctx, &search.BorrowReq{Id: bookinfo.Id})
			if err != nil {
				return &types.BorrowResq{Status: false}, nil
			}
			//
		}

	} else {
		return &types.BorrowResq{Status: false}, nil
	}

	return &types.BorrowResq{Status: true}, nil
}
