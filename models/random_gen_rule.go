package models

import "github.com/astaxie/beego/orm"



// RandomGenRule 随机规则？
type RandomGenRule struct {
	Id               int
	QuestionType     string `orm:"description(题目类型)"`
	QuestionCount    int    `orm:"description(题目数量)"`
	PreQuestionScore int    `orm:"description(每道题目分数)"`
	// 题目范围, 分割符为';', 如：第三章;第八章
	QuestionRange string `orm:"description(题目范围)"`
	Exam          *Exam  `orm:"description(考试);rel(fk)"`
}

func (t *RandomGenRule) TableName() string {
	return "random_gen_rule"
}

func init() {
	orm.RegisterModel(new(RandomGenRule))
}