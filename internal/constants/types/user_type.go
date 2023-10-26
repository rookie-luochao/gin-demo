package types

// swag:enum
//
//go:generate klctl gen enum UserType
type UserType int64

const (
	USER_TYPE_UNKNOWN UserType = iota
	USER_TYPE__ADMIN           // 管理员
	USER_TYPE__USER            // 用户
)
