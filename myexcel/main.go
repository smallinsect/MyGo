package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tealeg/xlsx"
)

//获取指定目录下的所有文件和目录
func main() {

	//设置路径，文件夹放在main的同级目录下
	PthSep := string(os.PathSeparator)
	Pthdir := "./excel"

	dir, err := ioutil.ReadDir(Pthdir)
	if err != nil {
		fmt.Printf("open dir failed: %s\n", err.Error())
	}

	//申明合并后的文件
	var new_file *xlsx.File
	var new_sheet *xlsx.Sheet
	new_file = xlsx.NewFile()
	var new_err error
	new_sheet, new_err = new_file.AddSheet("Sheet1")

	for _, fi := range dir {

		fmt.Printf("open success: %s\n", Pthdir+PthSep+fi.Name())
		if new_err != nil {
			fmt.Printf(new_err.Error())
		}

		//读取文件
		xlFile, err := xlsx.OpenFile(Pthdir + PthSep + fi.Name())
		if err != nil {
			fmt.Printf("open failed: %s\n", err)
		}
		for _, sheet := range xlFile.Sheets {
			fmt.Printf("Sheet Name: %s\n", sheet.Name)

			num := 0
			for _, row := range sheet.Rows {
				num++
				//跳过前5行，将后面的行写入新的文件
				if num > 5 {
					new_row := new_sheet.AddRow()
					new_row.SetHeightCM(1)

					for _, cell := range row.Cells {
						text := cell.String()
						fmt.Printf("%s\n", text)

						new_cell := new_row.AddCell()
						new_cell.Value = text
					}
				}
			}
		}
	}

	//写入文件
	new_err = new_file.Save("merge.xlsx")
	if new_err != nil {
		fmt.Printf(new_err.Error())
	}
}
