package model

import "github.com/rookie-luochao/gin-demo/internal/constants/types"

//	User 用户
//
// @def primary ID
// @def unique_index idx_user UserID
// @def unique_index idx_mobile Mobile
//
//go:generate klctl gen model2 User --database DB
type User struct {
	BasicField
	RefUser
	UserInfo
}

type RefUser struct {
	UserID int64 `db:"f_user_id" json:"userID"`
}

type UserInfo struct {
	// 昵称
	NickName string `db:"F_nick_name" json:"nickName"`
	// 电话
	Mobile string `db:"f_mobile" json:"mobile"`
	// 用户角色
	UserType types.UserType `db:"f_user_type,default=2" json:"userType"`
	// 电话运营商
	MobileType types.MobileType `db:"f_mobile_type" json:"mobileType"`
}
