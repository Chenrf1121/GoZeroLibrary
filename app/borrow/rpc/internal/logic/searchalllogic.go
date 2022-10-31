package logic

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/zeromicro/go-zero/core/logx"
	"library/app/borrow/model"
	"library/app/borrow/rpc/borrow"
	"library/app/borrow/rpc/internal/svc"
	"log"
	"strconv"
)

type SearchAllLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchAllLogic {
	return &SearchAllLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchAllLogic) SearchAll(in *borrow.BorrwoIdReq) (*borrow.BorrowListResp, error) {
	// todo: add your logic here and delete this line
	//	db, err := gorm.Open("mysql", l.svcCtx.Config.Mysql.DataSource)
	fmt.Println("开启searchall远程调用！！！")
	db, err := gorm.Open("mysql", l.svcCtx.Config.Mysql.DataSource)
	if err != nil {
		log.Println("borrow连接数据库错误")
	}
	defer db.Close()
	db.AutoMigrate(&model.Borrow{})
	var borrowmodel []model.Borrow

	Idnumber, _ := strconv.Atoi(in.Id)
	db.Debug().Table("borrow").Where("user_id = ?", Idnumber).Find(&borrowmodel)

	result := &borrow.BorrowListResp{}

	for _, i := range borrowmodel {
		var tmpmodel borrow.BorrowInfoResp
		tmpmodel.BorrowTime = i.BorrowTime.String()
		tmpmodel.Isreturn = int32(i.Isreturn)
		tmpmodel.UserId = i.UserId
		tmpmodel.BookId = strconv.Itoa(int(i.BookId))
		if i.Isreturn == 1 {
			tmpmodel.ReturnTime = i.ReturnTime.String()
		}

		result.List = append(result.List, &tmpmodel)
	}
	return result, nil
}
