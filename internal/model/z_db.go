package model

import (
	"database/sql/driver"
	"github.com/go-courier/sqlx/v2"
	"github.com/go-courier/sqlx/v2/builder"
	"time"

	"github.com/rookie-luochao/gin-demo/internal/utils"
)

var DB = sqlx.NewDatabase("")

type BasicField struct {
	PrimaryID
	OperationTime
}

type PrimaryID struct {
	ID uint `db:"F_id,autoincrement" json:"key"`
}

type OperationTime struct {
	CreatedAt utils.Timestamp `db:"f_created_at" json:"createdAt" orderByLabel:"创建时间"`
	UpdatedAt utils.Timestamp `db:"f_updated_at" json:"updatedAt" orderByLabel:"更新时间"`
	DeletedAt utils.Timestamp `db:"f_deleted_at,default=0" json:"-"`
}

func (o *OperationTime) SetNow() {
	o.UpdatedAt = utils.Timestamp(time.Now())
}

func (o *OperationTime) SetNowForCreate() {
	o.SetNow()
	o.CreatedAt = o.UpdatedAt
}

type Tag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Tags []Tag

func (tags *Tags) Scan(v interface{}) error {
	return utils.JSONScan(v, tags)
}

func (tags Tags) Value() (driver.Value, error) {
	return utils.JSONValue(tags)
}

func (Tags) DataType(driverName string) string {
	return "text"
}

type ConditionRules struct {
}

func NewCondRules() *ConditionRules {
	return &ConditionRules{}
}

func (rule *ConditionRules) When(epr bool, condition builder.SqlCondition) builder.SqlCondition {
	if epr {
		return condition
	}
	return nil
}
