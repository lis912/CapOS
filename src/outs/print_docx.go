package outs

import (
	"baliance.com/gooxml/color"
	"baliance.com/gooxml/document"
	"baliance.com/gooxml/measurement"

	"baliance.com/gooxml/schema/soo/wml"
	"fmt"
)

func Outdocx(filename string, Docdate []Singctr) {
	filename = fmt.Sprint(filename, ".docx")
	// 记录每一个大测评要求中有几个测评项目
	var num int

	// 新建一个文本文件
	doc := document.New()

	{
		// 测试表：

		// 新建一个表
		table := doc.AddTable()

		// 表的总宽度 100%
		table.Properties().SetWidthPercent(100)

		// 边框样式
		borders := table.Properties().Borders()
		// s个参数 wml.ST_BorderSingle 颜色  1*measurement.Point
		borders.SetAll(wml.ST_BorderSingle, color.Auto, 1*measurement.Point)

		for i, singdoc := range Docdate {

			if singdoc.Nam != "" {
				// 如果是最后一个元素，那么放弃使用单元格模式
				if i == (len(Docdate) - 1) {
					// 新建一行
					row := table.AddRow()
					// 新建单元格
					cell := row.AddCell()
					// 设置合并列单元格 Restart
					cell.Properties().SetWidthPercent(12)
					// cell.Properties().SetVerticalMerge(wml.ST_MergeRestart)
					cell.AddParagraph().AddRun().AddText(singdoc.Nam)
					row.AddCell().AddParagraph().AddRun().AddText(singdoc.Ctr)
					row.AddCell().AddParagraph().AddRun().AddText(singdoc.Ret)
					// 第一行第4个单元设置宽度，首行设置后以后，下面就可以自动对齐了
					cell = row.AddCell()
					cell.Properties().SetWidthPercent(12)
					cell.AddParagraph().AddRun().AddText(singdoc.Deg)
					break
				}
				// 新建一行
				row := table.AddRow()
				// 新建单元格
				cell := row.AddCell()
				// 设置合并列单元格 Restart
				cell.Properties().SetWidthPercent(12)
				cell.Properties().SetVerticalMerge(wml.ST_MergeRestart)
				cell.AddParagraph().AddRun().AddText(singdoc.Nam)
				row.AddCell().AddParagraph().AddRun().AddText(singdoc.Ctr)
				row.AddCell().AddParagraph().AddRun().AddText(singdoc.Ret)
				// 第一行第4个单元设置宽度，首行设置后以后，下面就可以自动对齐了
				cell = row.AddCell()
				cell.Properties().SetWidthPercent(12)
				cell.AddParagraph().AddRun().AddText(singdoc.Deg)
				num = 2

			} else {
				row := table.AddRow()
				cell := row.AddCell()
				cell.Properties().SetVerticalMerge(wml.ST_MergeContinue)
				// 和上一单元格合并 Vertical Merge 2 大概是表示 第二行第一个单元格 和 "身份鉴别" 这个单元格合并
				// num记录多少有多少行，在最后++
				mergenum := "Vertical Merge " + string(num)
				cell.AddParagraph().AddRun().AddText(mergenum)
				row.AddCell().AddParagraph().AddRun().AddText(singdoc.Ctr)
				row.AddCell().AddParagraph().AddRun().AddText(singdoc.Ret)
				row.AddCell().AddParagraph().AddRun().AddText(singdoc.Deg)
				num = num + 1

			}

		}

	}

	// 保存文档
	doc.SaveToFile(filename)
}
