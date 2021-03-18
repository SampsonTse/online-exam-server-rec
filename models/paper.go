package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// 试卷表
type Paper struct {
	Id         int
	IsDelete   bool
	CreateTime time.Time `orm:"auto_now_add"`

	Name        string `orm:"description(试卷名)"`
	Status      string `orm:"description(考试状态：开始、暂停、结束...)"`
	Tips        string `orm:"description(试卷说明)"`
	Range       string `orm:"description(试卷范围)"`
	Type        string `orm:"description(考试还是作业)"`
	CreateModel string `orm:"description(试卷创建模式)"`
	ExamMode    string `orm:"description(考试模式)"`
	TotalScore  int    `orm:"description(总分)"`

	Duration  int       `orm:"description(考试时长)"`
	StartTime time.Time `orm:"description(开始时间)"`
	EndTime   time.Time `orm:"description(结束时间)"`

	PaperCreator *Teacher           `orm:"description(试题创建人);rel(fk)"`
	Course       *Course            `orm:"description(课程);rel(fk)"`
	Questions    []*HistoryQuestion `orm:"description(试题);reverse(many)"`
}

func (t *Paper) TableName() string {
	return "paper"
}

func init() {
	orm.RegisterModel(new(Paper))
}