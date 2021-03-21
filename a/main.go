package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	f, err := excelize.OpenFile("/Users/ktlshya/Desktop/feeds.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	str := "INSERT INTO `feed_infos` (`code`, `name`, `moisture`, `dry_matter`, `ndf`, `adf`, `endf`, `lactic_acid`, `wsc`, `starch`, `soluble_fiber`, `total_protein`, `soluble_protein`, `rdp`, `me`, `nel`, `crude_fat`, `total_ftty_acid`, `ash`, `ca`, `p`, `mg`, `k`, `mn`, `cu`, `fe`, `zn`, `methionine`, `lysine`, `vitamina`, `vitamin_d3`, `vitamine`, `choline`, `biotin`) VALUES(%s"

	rows := f.GetRows("s1")

	ff, err := os.OpenFile("feeds.sql", os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	defer ff.Close()

	for i, row := range rows {
		if i == 0 {
			continue
		}
		line := ""
		for _, colCell := range row {
			if colCell == "" {
				colCell = "0"
			}
			line += "'" + colCell + "',"
			// fmt.Print(str+"('"+colCell+"'")
		}
		n, err := ff.WriteString(fmt.Sprintf(str, strings.TrimRight(line, ",")) + ");\n")
		fmt.Println(n, err)
		// fmt.Println(fmt.Sprintf(str, line))
		// fmt.Println()
	}
}
