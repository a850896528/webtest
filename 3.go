package main

import (
	"fmt"
	"os"
)

type Student struct {
	id int
	password string
	name string
	money int
}
type studentMgr struct {
	allStudents []*Student
}
//使用newStudent构造Student
func newStudent(id int,name string,password string)*Student{
	return &Student{
		id: id,
		name: name,
		password:password,
	}
}
//使用newStudentMgr构造studentMgr
func newStudentMgr()*studentMgr{
	return &studentMgr{
		allStudents:make([]*Student,0,100),
	}
}
//注册
func (u *studentMgr)addUser(newuser *Student){
	for _,v:=range u.allStudents{
		if newuser.id==v.id{
			return
		}
	}
	u.allStudents=append(u.allStudents,newuser)
}
//登录
func (u *studentMgr)signin(newuer *Student){
	for _,v:=range u.allStudents{
		if v.id==newuer.id{
			if v.password==newuer.password{
				fmt.Println("登录成功")
			}
		}
	}
	return
}
func (s *studentMgr)addStudent(newStu *Student){
	s.allStudents=append(s.allStudents,newStu)
}
func (s *studentMgr)showStudent(){
	for _,v:=range s.allStudents{
		fmt.Printf("学号：%d,姓名：%s,金额：%d\n",v.id,v.name,v.money)
	}
}
func showmap(){
	fmt.Println("欢迎来到学生管理系统")
	fmt.Println("1.登录")
	fmt.Println("2.注册")
	fmt.Println("3.")
	fmt.Println("4.退出")
}
func getInput()*Student{
	var(
		id int
		name string
		money int
	  password string
	)
	fmt.Println("输入你的id:")
	fmt.Scanf("%d\n",&id)
	fmt.Println("输入你的名字:")
	fmt.Scanf("%s\n",&name)
	fmt.Println("输入你的password:")
	fmt.Scanf("%d\n",&password)
	stu:=newStudent(id,name,password)
	return stu
}
func main() {
		showmap()
		message:=newStudentMgr()
		for{
			fmt.Println("请输入你的数字：")
			var input int
			fmt.Scanf("%d\n",&input)
			fmt.Println("你所输入的数字是：",input)
			switch input {
			case 1:
				a:=getInput()
				message.addStudent(a)
			case 2:
				a:=getInput()
			case 3:
			case 4:
				os.Exit(0)
			}
		}
	}


