package utils

import (
	"fmt"

	"github.com/yitter/idgenerator-go/idgen"
)

type IGenID interface {
	ID() int64
	IDToString() string
}

type IDMgr struct {
	idGen *idgen.DefaultIdGenerator
}

func NewIdMgr(node uint16) (*IDMgr, error) {
	opts := idgen.NewIdGeneratorOptions(node)
	idGenerator := idgen.NewDefaultIdGenerator(opts)
	return &IDMgr{
		idGen: idGenerator,
	}, nil
}

func (c *IDMgr) ID() int64 {
	return c.idGen.NewLong()
}

func (c *IDMgr) IDToString() string {
	return fmt.Sprintf("%d", c.idGen.NewLong())
}
