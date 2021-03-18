package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type HistoryRecord struct {
	Id         	int
	UserId		int
	ExamId		int
	IP			string
	CreateTime 	time.Time 	`orm:"auto_now_add"`
	Status		string 		`orm:"description(考试状态：开始、暂停、结束...)"`
	Duration  	int       	`orm:"description(考试时长)"`
	StartTime 	time.Time 	`orm:"description(开始时间)"`
	EndTime   	time.Time 	`orm:"description(结束时间)"`
}


func (t *HistoryRecord) TableName() string {
	return "history_record"
}

func init() {
	orm.RegisterModel(new(HistoryRecord))
}