package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func FileRead() (content string, err error) {
	var baseSlice []string
	file, err := os.OpenFile("../fileStorage/student.json", os.O_CREATE|os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("文件打开失败...", err)
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')

		if len(line) != 0 {
			// fmt.Println(line)
			baseSlice = append(baseSlice, line)

		}

		if err != nil {
			switch err {
			case io.EOF:
				fmt.Println("文件读取完成！")
			default:
				fmt.Println("读取文件失败：", err)
			}
			break
		}
	}
	content = strings.Join(baseSlice, "")
	// fmt.Println(content)
	return
}

func FileWrite(content string) (err error) {
	file, err := os.OpenFile("../fileStorage/student.json", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("打开文件失败:", err)
	}

	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString(content)
	writer.Flush()
	return

}
