package controller

import (
	"context"
	"github.com/go-courier/logr"
	"github.com/go-courier/sqlx/v2"
	"github.com/go-courier/sqlx/v2/builder"
	"github.com/rookie-luochao/gin-demo/internal/common"
	"github.com/rookie-luochao/gin-demo/internal/model"
	"github.com/rookie-luochao/gin-demo/internal/utils"
	"github.com/sirupsen/logrus"
)

type UserMgr struct {
	idGen utils.IGenID
	db    sqlx.DBExecutor
}

func NewUserMgr(db sqlx.DBExecutor, idGen utils.IGenID) *UserMgr {
	return &UserMgr{
		idGen: idGen,
		db:    db,
	}
}

type CreateOrUpdateUserData struct {
	// 用户名
	Name string `json:"name"`
	// 手机号
	Mobile string `json:"mobile"`
}

func (c *UserMgr) CreateOrUpdateUser(body CreateOrUpdateUserData) error {
	m := model.User{}
	m.Mobile = body.Mobile
	err := m.FetchByMobile(c.db)
	if err != nil {
		if sqlx.DBErr(err).IsNotFound() {
			m.UserID = c.idGen.ID()
			m.NickName = body.Name
			if createErr := m.Create(c.db); createErr != nil {
				return createErr
			}
			return nil
		}
		return err
	}

	// update
	m.NickName = body.Name
	err = m.UpdateByIDWithStruct(c.db)
	if err != nil {
		return err
	}
	return nil
}

type ListUserParams struct {
	common.PageParam
	// 姓名
	Name string `in:"query" name:"name,omitempty"`
	// 手机号
	Mobile string `in:"query" name:"mobile,omitempty"`
}

type ListUserResp struct {
	Data  []model.User `json:"data"`
	Total int          `json:"total"`
}

func (c *UserMgr) ListUser(ctx context.Context, params *ListUserParams) (*ListUserResp, error) {
	// 支持两种打印方式 链路中需要打印的 从ctx中取 TODO: 为何这条日志是递归的
	logr.FromContext(ctx).Debug("4444")
	// 一般调试日志 可以直接用logrus
	logrus.Debug("333")
	resp := &ListUserResp{}
	m := model.User{}
	cs := model.NewCondRules()
	where := builder.And(
		cs.When(params.Mobile != "", m.FieldMobile().Eq(params.Mobile)),
		cs.When(params.Name != "", m.FieldNickName().Eq(params.Name)),
	)

	// 只有传入context ，DEBUG下才会打印SQL语句
	count, err := m.Count(c.db.WithContext(ctx), where)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return resp, nil
	}

	// 如果不传入context ，DEBUG下不会打印SQL语句
	data, err := m.List(c.db, where, params.PageAddition())
	if err != nil {
		return resp, nil
	}

	logr.FromContext(ctx).Debug("3333")

	resp.Data = data
	return resp, nil
}
