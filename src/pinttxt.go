package outs

import (
	"capos"
	"fmt"
	"github.com/bndr/gotabulate"
)

func ReturnRow(os OsDate) string {
	// 空接口切片 一维切片 []interface{}{初始化}
	row_1 := []interface{}{1989, 923423421342345511, 12, "hello"}
	row_2 := []interface{}{2010, 4}
	row_3 := []interface{}{4435345, "dfas"}

	// Create an object from 2D interface array
	t := gotabulate.Create([][]interface{}{row_1, row_2, row_3})

	t.Title = "mytest"

	// 是否要隐藏表格顶行 "top"
	t.HideLines = append(t.HideLines, "belowheader")

	// Set the Headers (optional)
	t.SetHeaders([]string{"year", "month"})

	// Set the Empty String (optional)
	// t.SetEmptyString("lishichnag")

	// Set Align (Optional)
	t.SetAlign("center")

	// Print the result: grid, or simple
	// fmt.Println(t.Render("grid"))
	// fmt.Println("\n\n")

	return t.Render("grid")
}
