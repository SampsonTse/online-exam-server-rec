package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)



type HistoryQuestion struct {
	Id         int
	CreateTime time.Time `orm:"auto_now_add"`

	Score      int    `orm:"description(分值)"`
	Type       string `orm:"description(题型)"`
	Range      string `orm:"description(范围)"`
	Difficulty string `orm:"description(难度)"`

	Question       string `orm:"description(题目)"`
	EncodedChoices string `orm:"description(编码后的选项);type(text)"`
	EncodedAnswer  string `orm:"description(答案，根据类型 0 1... 代表 A B... 或者 错误 正确)"`

	Paper *Paper `orm:"description(试卷外键);rel(fk)"`
	Exam  *Exam  `orm:"description(考试外键);rel(fk)"`
}

func (t *HistoryQuestion) TableName() string {
	return "history_question"
}


func init() {
	orm.RegisterModel(new(HistoryQuestion))
}