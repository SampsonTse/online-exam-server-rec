package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)


// Teacher 老师表
type Teacher struct {
	Id         int
	IsDelete   bool
	CreateTime time.Time `orm:"auto_now_add"`

	TeacherId string `orm:"description(工号|账号);unique"`
	Password  string `orm:"description(密码)"`

	Name    string `orm:"description(姓名)"`
	Major   string `orm:"description(专业)"`
	College string `orm:"description(学院)"`
	IsAdmin int    `orm:"description(是否为管理员)"`

	Papers    []*Paper    `orm:"description(试卷);reverse(many)"`
	Exams     []*Exam     `orm:"description(考试);reverse(many)"`
	Questions []*Question `orm:"description(试题);reverse(many)"`
	Classes   []*Class    `orm:"description(班级);reverse(many)"`
	Courses   []*Course   `orm:"description(课程);rel(m2m);rel_table(m2m_teacher_course)"`
}

func (t *Teacher) TableName() string{
	return "teacher"
}

func init() {
	orm.RegisterModel(new(Teacher))
}