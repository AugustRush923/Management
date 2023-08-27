package core

import (
	"Management/models"
	"fmt"
)

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
	sm := models.StudentManagerment{}
	load_err := sm.LoadInfo()
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
			sm.Add()
		case "2":
			sm.Delete()
		case "3":
			sm.Edit()
		case "4":
			sm.Get()
		case "5":
			sm.GetAll()
		case "6":
			save_err := sm.SaveInfo()
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
