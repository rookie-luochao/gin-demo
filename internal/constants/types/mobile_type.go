package types

// swag:enum
type MobileType string

const (
	USER_TYPE_Y MobileType = "移动"
	USER_TYPE_D MobileType = "联通" // 管理员
)
