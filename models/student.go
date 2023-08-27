package models

import (
	"Management/utils"
	"encoding/json"
	"fmt"
)

type Student struct {
	Name, Gender, Tel string
}

func (s Student) String() string {
	return fmt.Sprintf("姓名：%v\t性别：%v\t电话：%v\n", s.Name, s.Gender, s.Tel)
}

func InitStudent(name, gender, tel string) *Student {
	// 工厂函数
	return &Student{Name: name, Gender: gender, Tel: tel}
}

type StudentManagerment struct {
	Student
	StudentList []*Student
}

func (sm *StudentManagerment) Add() {
	var name, gender, tel string
	fmt.Println("请输入姓名：")
	fmt.Scanln(&name)
	fmt.Println("请输入性别：")
	fmt.Scanln(&gender)
	fmt.Println("请输入电话：")
	fmt.Scanln(&tel)
	student := InitStudent(name, gender, tel)
	sm.StudentList = append(sm.StudentList, student)
}

func (sm StudentManagerment) Get() {
	var name string
	fmt.Println("请输入要查找的学生的姓名：")
	fmt.Scanln(&name)
	for _, student := range sm.StudentList {
		if name == student.Name {
			fmt.Println(student)
			return
		}
	}
	fmt.Println("查无此人！")
}

func (sm StudentManagerment) GetAll() {
	for _, v := range sm.StudentList {
		fmt.Println(v)
	}
}

func (sm *StudentManagerment) Delete() {
	var name string
	fmt.Println("请输入要查找的学生的姓名：")
	fmt.Scanln(&name)
	for index, student := range sm.StudentList {
		if name == student.Name {
			sm.StudentList = append(sm.StudentList[:index], sm.StudentList[index+1:]...)
			return
		}
	}
	fmt.Println("查无此人！")
}

func (sm *StudentManagerment) Edit() {
	var name string
	fmt.Println("请输入要查找的学生的姓名：")
	fmt.Scanln(&name)
	for _, student := range sm.StudentList {
		var choice, value string
		if name == student.Name {
			fmt.Println("需要更新什么信息？1.姓名,2.性别,3电话：")
			fmt.Scanln(&choice)

			switch choice {
			case "1":
				fmt.Println("请输入姓名：")
				fmt.Scanln(&value)
				student.Name = value
			case "2":
				fmt.Println("请输入性别：")
				fmt.Scanln(&value)
				student.Gender = value
			case "3":
				fmt.Println("请输入电话：")
				fmt.Scanln(&value)
				student.Tel = value
			}
			fmt.Println("更新成功！")
			return
		}
	}
	fmt.Println("查无此人！")
}

func (sm *StudentManagerment) LoadInfo() (err error) {
	file_data, fileRead_err := utils.FileRead()
	if fileRead_err != nil {
		fmt.Println(fileRead_err)
		return fileRead_err
	}

	unmarshal_err := json.Unmarshal([]byte(file_data), &sm.StudentList)
	if unmarshal_err != nil {
		fmt.Println(unmarshal_err)
		return unmarshal_err
	}
	return
}

func (sm *StudentManagerment) SaveInfo() (err error) {
	json_data, err := json.Marshal(sm.StudentList)
	if err != nil {
		fmt.Println("序列化错误：", err)
		return
	}

	err = utils.FileWrite(string(json_data))
	if err != nil {
		fmt.Println("数据持久化失败：", err)
		return
	}
	return
}
