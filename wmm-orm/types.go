package wmm_orm

import (
	"context"
	"database/sql"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-04 14:13

type Querier[T any] interface {
	Get(ctx context.Context) (*T, error)
	GetMulti(ctx context.Context) ([]*T, error)
}

// Executor
// @Description: 用于增删改的接口
type Executor interface {
	Exec(ctx context.Context) (sql.Result, error)
}

type Query struct {
	SQL  string
	Args []any
}

type QueryBuilder interface {
	Build() (*Query, error)
}

// Expression
// @Description: 标记接口，代表一个表达式
type Expression interface {
	expr()
}
