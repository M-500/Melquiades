package wmm_orm

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-04 14:17

// (衍生类型)
type op string // 定义一个枚举
// 这个叫别名
//type ops = string

func (o op) String() string {
	return string(o)
}

const (
	opEq  = "="
	opLT  = "<"
	opGT  = ">"
	opNot = "NOT"
	opAnd = "AND"
	opOr  = "OR"
)

// Condition
// @Description: 查询条件
type Condition struct {
	left  Expression
	opt   op
	right Expression
}

func (left Condition) expr() {
	//TODO implement me
	panic("implement me")
}

type Column struct {
	name string
}

func (c Column) expr() {
	//TODO implement me
	panic("implement me")
}

// 使用的时候  C("id").Eq(12)
func (c Column) Eq(arg any) Condition {
	return Condition{
		left:  c,
		opt:   opEq,
		right: value{val: arg},
	}
}

// Not(C("name").Eq("tom"))
func Not(p Condition) Condition {
	return Condition{
		opt:   opNot,
		right: p,
	}
}

// 用起来是这个效果  C("id").Eq(12).And(C("name").Eq("tom"))
func (left Condition) And(right Condition) Condition {
	return Condition{
		left:  left,
		opt:   opAnd,
		right: right,
	}
}

// 用起来是这个效果  C("id").Eq(12).Or(C("name").Eq("tom"))
func (left Condition) Or(right Condition) Condition {
	return Condition{
		left:  left,
		opt:   opOr,
		right: right,
	}
}

type value struct {
	val any
}

func (v value) expr() {
	//TODO implement me
	panic("implement me")
}

func C(name string) Column {
	return Column{
		name: name,
	}
}
