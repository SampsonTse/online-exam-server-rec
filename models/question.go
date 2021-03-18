package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// Question 题目表
type Question struct {
	Id         int
	IsDelete   bool
	CreateTime time.Time `orm:"auto_now_add"`

	Type       string `orm:"description(题型)"`
	Range      string `orm:"description(范围)"`
	Difficulty string `orm:"description(难度)"`

	Question       string `orm:"description(题目)"`
	EncodedChoices string `orm:"description(编码后的选项);type(text)"`
	EncodedAnswer  string `orm:"description(答案，根据类型 0 1... 代表 A B... 或者 错误 正确)"`

	QuestionCreator *Teacher `orm:"description(出题人外键);rel(fk)"`
	Course          *Course  `orm:"description(课程外键);rel(fk)"`
}

func (q *Question) TableName() string {
	return "question"
}

func init() {
	orm.RegisterModel(new(Question))
}