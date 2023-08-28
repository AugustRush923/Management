package core

import (
	"Management/models"
	"Management/utils"
	"encoding/json"
	"fmt"
)

type StudentManagement struct {
	StudentList []*models.Student
}

func (sm *StudentManagement) add_student() {
	var name, gender, tel string
	fmt.Println("请输入姓名：")
	fmt.Scanln(&name)
	fmt.Println("请输入性别：")
	fmt.Scanln(&gender)
	fmt.Println("请输入电话：")
	fmt.Scanln(&tel)
	student := models.InitStudent(name, gender, tel)
	sm.StudentList = append(sm.StudentList, student)
	fmt.Println("新增成功！")
	fmt.Println("****************************")
}

func (sm StudentManagement) get_student_info() {
	var name string
	fmt.Println("请输入要查找的学生的姓名：")
	fmt.Scanln(&name)
	for _, student := range sm.StudentList {
		if name == student.Name {
			fmt.Println(student)
			fmt.Println("****************************")
			return
		}
	}
	fmt.Println("查无此人！")
	fmt.Println("****************************")
}

func (sm StudentManagement) get_students_info() {
	if len(sm.StudentList) == 0 {
		fmt.Println("暂无数据...")
		fmt.Println("****************************")
	}
	for _, v := range sm.StudentList {
		fmt.Println(v)
	}
	fmt.Println("****************************")
}

func (sm *StudentManagement) del_student() {
	var name string
	fmt.Println("请输入要查找的学生的姓名：")
	fmt.Scanln(&name)
	for index, student := range sm.StudentList {
		if name == student.Name {
			sm.StudentList = append(sm.StudentList[:index], sm.StudentList[index+1:]...)
			fmt.Println("删除成功！")
			fmt.Println("****************************")
			return
		}
	}
	fmt.Println("查无此人！")
	fmt.Println("****************************")
}

func (sm *StudentManagement) update_student_info() {
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
			fmt.Println("****************************")
			return
		}
	}
	fmt.Println("查无此人！")
	fmt.Println("****************************")
}

func (sm *StudentManagement) load_student_info() (err error) {
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

func (sm *StudentManagement) save_student_info() (err error) {
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

func showMenu() {
	fmt.Println("****************************")
	fmt.Println("欢迎进入xxx学员管理系统 v1.0")
	fmt.Println("1. 增加学员信息")
	fmt.Println("2. 删除学员信息")
	fmt.Println("3. 更新学员信息")
	fmt.Println("4. 获取指定学员信息")
	fmt.Println("5. 展示所有学员信息")
	fmt.Println("6. 保存学员信息")
	fmt.Println("7. 打印帮助信息")
	fmt.Println("8. 退出系统")
	fmt.Println("****************************")
}

func Run() {
	sm := StudentManagement{}
	load_err := sm.load_student_info()
	if load_err != nil {
		panic(load_err)
	}
	showMenu()
	for {
		var choice string
		fmt.Println("请输入您需要的功能序号：")
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			sm.add_student()
		case "2":
			sm.del_student()
		case "3":
			sm.update_student_info()
		case "4":
			sm.get_student_info()
		case "5":
			sm.get_students_info()
		case "6":
			save_err := sm.save_student_info()
			if save_err != nil {
				panic(save_err)
			}
			fmt.Println("保存成功！")
		case "7":
			showMenu()
		case "8", "q":
			return
		default:
			fmt.Println("请输入正确的编号！")
			continue
		}
	}

}
