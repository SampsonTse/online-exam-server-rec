package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// Score成绩单
type Score struct {
	Id         int
	IsDelete   bool
	CreateTime time.Time `orm:"auto_now_add"`

	BeginTime time.Time `orm:"description(开始考试时间)"`
	EndTime   time.Time `orm:"description(提交时间)"`

	Score              int    `orm:"description(分值)"`
	ResitTimes         int    `orm:"description(当前重考的次数)"`
	Finish             bool   `orm:"description(是否提交)"`
	EncodedAnswers     string `orm:"type(text);description(学生做题答案);type(text)"`
	EncodedQuestionIds string `orm:"type(text);description(随机组卷时才使用,保存试题的所有id)"`

	Exam    *Exam    `orm:"description(考试);rel(fk)"`
	Class   *Class   `orm:"description(班级);rel(fk)"`
	Student *Student `orm:"description(学生);rel(fk)"`
}

func (t *Score) TableName() string {
	return "score"
}

func init() {
	orm.RegisterModel(new(Score))
}