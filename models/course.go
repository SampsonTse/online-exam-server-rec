package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// Course 课程表
type Course struct {
	Id         int
	IsDelete   bool
	CreateTime time.Time `orm:"auto_now_add"`

	Name  string `orm:"description(课程名)"`
	Grade string `orm:"description(年级)"`

	CourseCreator *Teacher    `orm:"description(课程创建人);rel(fk);null"`
	Classes       []*Class    `orm:"description(班级);rel(m2m);rel_table(m2m_course_class)"`
	Questions     []*Question `orm:"description(试题);reverse(many)"`
	Papers        []*Paper    `orm:"description(试卷);reverse(many)"`
	Teachers      []*Teacher  `orm:"description(老师);reverse(many)"`
}

func (t *Course) TableName() string {
	return "course"
}


func init() {
	orm.RegisterModel(new(Course))
}