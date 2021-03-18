package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)


// Exam 考试表
type Exam struct {
	Id         int
	IsDelete   bool
	CreateTime time.Time `orm:"auto_now_add"`

	Name            string `orm:"description(试卷说明)"`
	Status          string `orm:"description(考试状态：开始、暂停、结束...)"`
	Tips            string `orm:"description(试卷说明)"`
	Range           string `orm:"description(试卷范围)"`
	Type            string `orm:"description(考试还是作业)"`
	CreateModel     string `orm:"description(试卷创建模式, 种类参考:CreateModel变量)"`
	ExamMode        string `orm:"description(考试模式)"`
	TotalScore      int    `orm:"description(总分)"`
	TotalResitTimes int    `orm:"description(最多重考的次数, 最少为1)"`

	Duration  int       `orm:"description(考试时长)"`
	StartTime time.Time `orm:"description(开始时间)"`
	EndTime   time.Time `orm:"description(结束时间)"`

	ExamCreator *Teacher           `orm:"description(考试发起人外键);rel(fk)"`
	Course      *Course            `orm:"description(课程外键);rel(fk)"`
	Questions   []*HistoryQuestion `orm:"description(试题);reverse(many)"`
	Scores      []*Score           `orm:"description(分数外键);reverse(many)"`
	Classes     []*Class           `orm:"description(班级m2m外键);reverse(many)"`
	GenRules    []*RandomGenRule   `orm:"description(随机组卷规则);reverse(many)"`
}

func (t *Exam) TableName() string {
	return "exam"
}

func init() {
	orm.RegisterModel(new(Exam))
}