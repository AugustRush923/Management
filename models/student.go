package models

import (
	"fmt"
)

type Student struct {
	Name, Gender, Tel string
}

func (s Student) String() string {
	return fmt.Sprintf("姓名：%v\t性别：%v\t电话：%v", s.Name, s.Gender, s.Tel)
}

func InitStudent(name, gender, tel string) *Student {
	// 工厂函数
	return &Student{Name: name, Gender: gender, Tel: tel}
}
