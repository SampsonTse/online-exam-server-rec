package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// Class 班级表
type Class struct{
	Id         int
	IsDelete   bool
	CreateTime time.Time `orm:"auto_now_add"`

	Name string `orm:"description(班级名)"`

	ClassCreator *Teacher   `orm:"description(班级创建人);rel(fk)"`
	Courses      []*Course  `orm:"description(课程);reverse(many)"`
	Scores       []*Score   `orm:"description(分数外键);reverse(many)"`
	Exams        []*Exam    `orm:"description(课程);rel(m2m);rel_table(m2m_class_exam)"`
	Students     []*Student `orm:"description(学生);rel(m2m);rel_table(m2m_class_student)"`
}

func (t *Class) TableName() string {
	return "class"
}


func init() {
	orm.RegisterModel(new(Class))
}

// 获取查询容器
func GetClassQuery() orm.QuerySeter{
	o := orm.NewOrm()
	return o.QueryTable("class").Filter("IsDelete",false)
}

// 获取班级学生人数
func CountStudentByClassId(id int) (count int64,err error){
	o := orm.NewOrm()
	class := Class{Id:id}
	err = o.Read(&class)
	if err == nil{
		count,_ = CountM2MClassStudent(class)
	}
	return count,err

}

// 增加班级
func AddClass(class Class) (id int64,err error){
	o := orm.NewOrm()
	id,err = o.Insert(&class)
	if err != nil {
		return 0,err
	}
	return id,nil
}

// 根据教师teacherId 获得该教师所在的所有班级
func GetAllClassByTeacherId(teacherId string) (classArray []Class,err error){
	o :=  orm.NewOrm()
	qs := o.QueryTable("class").Filter("ClassCreator_teacherId",teacherId).Filter("IsDelete",0)

	var class []Class
	if _,err = qs.All(&class);err == nil{
		for _, value := range class{
			classArray = append(classArray,value)
		}
		return classArray,nil
	}
	return nil,err
}

// GetClassesByCourseId 根据课程ID获取班级 one to Many
func GetClassesByCourseId(id int)(vc []*Class,err error)  {
	o := orm.NewOrm()
	v := &Course{Id:id}
	var classes []*Class
	if err = o.Read(v); err ==nil{
		_,err = o.QueryTable("class").Filter("IsDelete",false).Filter("Courses__Course__Id",id).All(&classes)
		if err == nil{
			return classes,nil
		}
	}
	return nil,err
}

// 根据课程 courseId 获得加入的班级
func GetClassByCourseId(courseId int) (class []Class,err error){
	o := orm.NewOrm()
	qs := o.QueryTable("class").Filter("Courses__Course__Id",courseId)

	if _,err = qs.All(&class); err == nil{
		return class,err
	}
	return nil,err
}

// 获取新增班级ID(还没有增)
func GetClassNextId() (count int64,err error){
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.QueryTable("class").OrderBy("-id").Values(&maps)
	if err==nil && num>0{
		return maps[0]["Id"].(int64)+1,nil
	}
	return 0,nil
}

// 更新班级名称
func UpdateClass(id int,name string)(err error){
	o := orm.NewOrm()
	class := Class{Id:id}
	if err = o.Read(&class); err == nil{
		class.Name = name
		_,err = o.Update(&class)
		if err!=nil{
			return err
		}
	}
	return nil
}

// 删除班级，将isdelete改为true
func DeleteClass(id int)(err error){
	o := orm.NewOrm()
	class := Class{Id:id}
	if err = o.Read(&class);err == nil{
		class.IsDelete = true
		_,err = o.Update(&class)
		if err!=nil{
			return err
		}
	}
	return nil
}

// 获取班级信息 (没有学生信息)
func GetClassById(id int) (v *Class,err error){
	o := orm.NewOrm()
	v = &Class{Id:id}
	if err = o.Read(v);err == nil{
		return v,nil
	}
	return nil,err
}

// 添加学生到班级中
func AddStudentToClass(stud *Student,class *Class) (nu int64,err error){
	o := orm.NewOrm()
	m2m := o.QueryM2M(class,"Students")
	if m2m.Exist(stud){
		return 0, nil
	}
	o.Insert(stud)
	num, err := m2m.Add(stud)
	return num,err
}


//  将学生从班级中移除
func DeleteStudentFromClass(stud *Student,class *Class) (nu int64,err error){
	o := orm.NewOrm()
	m2m := o.QueryM2M(class,"students")
	num,err := m2m.Remove(stud)
	return num,err
}

// 通过班级Id获取该班级学生
func GetStudentByClassId(classId int) (stud []Student,err error){
	class, err := GetClassById(classId)
	if err!=nil{
		return nil, err
	}
	num, err := orm.NewOrm().QueryTable(new(Student)).Filter("Classes__class__Id",class.Id).Filter("IsDelete", false).All(&stud)
	if err != nil || num <= 0 {
		return nil, err
	}
	return stud, nil
}

// 搜索学生
func GetStudentByContent(content string,teacherId string) (students []Student, err error){
	classes, err := GetAllClassByTeacherId(teacherId)
	if err != nil{
		return nil, err
	}
	cond := orm.NewCondition()
	cond1 := cond.Or("StudentId__icontains",content).Or("Name__icontains", content).Or("Grade__icontains", content).Or("Major__icontains", content).Or("college__icontains", content)
	query := orm.NewOrm().QueryTable(new(Student))
	query = query.SetCond(cond1)
	var allStudents []Student
	query.Limit(0,0).All(&allStudents)

	o := orm.NewOrm()
	for _,class := range classes{
		m2m := o.QueryM2M(&class,"Students")
		for _,student := range allStudents{
			if m2m.Exist(student){
				students = append(students,student)
			}
		}
	}
	return students,err
}

// 精准搜索
func SearchOneStudent(stuId string)(v *Student,err error){
	cond := orm.NewCondition()
	cond1 := cond.And("StudentId__iexact",stuId)
	query := orm.NewOrm().QueryTable(new(Student))
	query = query.SetCond(cond1)

	var students []Student
	query.Limit(0, 0).All(&students)
	_,err =query.Count()
	for _, student := range students {
		if student.IsDelete == false {
			return &student, err
		}
	}
	return nil, err


}

