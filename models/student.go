package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// Student 学生表
type Student struct {
	Id         int
	IsDelete   bool
	CreateTime time.Time `orm:"auto_now_add"`

	StudentId string `orm:"description(学号|账号)"`
	Password  string `orm:"description(密码)"`

	Name    string `orm:"description(姓名)"`
	Grade   string `orm:"description(年级)"`
	Major   string `orm:"description(专业)"`
	College string `orm:"description(学院)"`

	Classes []*Class `orm:"reverse(many)"`
	Scores  []*Score `orm:"reverse(many)"`
}

func (t *Student) TableName() string {
	return "student"
}

func init() {
	orm.RegisterModel(new(Student))
}


// 根据班级Id获得学生人数
func CountM2MClassStudent(class Class) (count int64,err error){
	o := orm.NewOrm()
	m2m := o.QueryM2M(&class,"Students")
	count,err = m2m.Count()
	if err!=nil{
		return 0,err
	}
	return count,nil
}