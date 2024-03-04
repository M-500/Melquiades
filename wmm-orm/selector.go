package wmm_orm

import (
	"context"
	"fmt"
	"reflect"
	"strings"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-04 14:20

type Selector[T any] struct {
	table string
	where []Condition

	sb   *strings.Builder
	args []any
}

func (s *Selector[T]) Build() (*Query, error) {
	/*
		strings.Builder 是 Go 语言标准库中的一个类型，用于构建字符串。
		它提供了一种高效的方式来动态地构建字符串，
		特别是当需要频繁地进行字符串拼接时，使用 strings.Builder
		可以避免因为频繁的字符串拼接导致的性能问题。
	*/
	//var sb strings.Builder
	s.sb = &strings.Builder{}
	sb := s.sb
	sb.WriteString("SELECT * FROM ")

	if s.table == "" {
		// 怎么拿到数据表的名字
		var t T
		of := reflect.TypeOf(t)
		sb.WriteByte('`')
		sb.WriteString(of.Name())
		sb.WriteByte('`')
	} else {
		sb.WriteString(s.table)
	}
	// 记录一下参数
	//args := make([]any, 0, 4)
	// 只有有where条件的时候，才会去构建Where子查询
	if len(s.where) > 0 {
		sb.WriteString(" WHERE ") // 先把Where写上去
		p := s.where[0]
		for i := 1; i < len(s.where); i++ {
			p = p.And(s.where[i])
		}
		if err := s.buildExpression(p); err != nil {
			return nil, err
		}
	}

	sb.WriteByte(';')
	return &Query{
		SQL:  sb.String(),
		Args: s.args,
	}, nil
}

// buildExpression
//
//	@Description: 递归构建表达式  这个代码写得俊啊！
//	@receiver s
//	@param expr
//	@return error
func (s *Selector[T]) buildExpression(expr Expression) error {
	switch exp := expr.(type) {
	case nil:
	case Condition:
		_, ok := exp.left.(Condition)
		if ok {
			s.sb.WriteByte('(')
		}
		if err := s.buildExpression(exp.left); err != nil {
			return err
		}
		if ok {
			s.sb.WriteByte(')')
		}
		s.sb.WriteByte(' ')
		s.sb.WriteString(exp.opt.String())
		s.sb.WriteByte(' ')
		_, ok = exp.right.(Condition)
		if ok {
			s.sb.WriteByte('(')
		}
		if err := s.buildExpression(exp.right); err != nil {
			return err
		}
		if ok {
			s.sb.WriteByte(')')
		}
	case Column:
		s.sb.WriteByte('`')
		s.sb.WriteString(exp.name)
		s.sb.WriteByte('`')
	// 剩下的暂时不考虑
	case value:
		s.sb.WriteByte('?')
		if err := s.addArg(exp.val); err != nil {
			return err
		}
	// 剩下的暂时不考虑
	default:
		return fmt.Errorf("ORM: 不支持的表达式类型 %v", expr)
	}
	return nil
}

func (s *Selector[T]) addArg(val any) error {
	if s.args == nil {
		// 分配一个预估容量  通常查询条件不会超过4个
		s.args = make([]any, 0, 6)
	}
	s.args = append(s.args, val)
	return nil
}

// From
//
//	@Description: 中间方法，要返回一个对象本身回去
//	@receiver s
//	@param table
//	@return *Selector[T]
func (s *Selector[T]) From(table string) *Selector[T] {
	s.table = table
	return s
}

// ids := []int{1,2,3}
// s.Where("id in (?,?,?)",ids) 错
// s.Where("id in (?,?,?)",ids...) 对
//func (s *Selector[T]) Where(query string, args ...any) (*T, error) {
//	panic("implement me")
//}

func (s *Selector[T]) Where(pc ...Condition) *Selector[T] {
	s.where = pc
	return s
}

func (s *Selector[T]) Get(ctx context.Context) (*T, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Selector[T]) GetMulti(ctx context.Context) ([]*T, error) {
	//TODO implement me
	panic("implement me")
}
