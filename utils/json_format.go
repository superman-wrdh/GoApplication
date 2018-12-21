package main

import (
	"fmt"
	"strings"
)

func main() {
	s := `{'age':23,'aihao':['pashan','movies','中国'],'name':{'firstName':'zhang','lastName':'san','aihao':['pashan','movies','name':{'firstName':'zhang','lastName':'san','aihao':['pashan','movies']}]}}`
	fmt.Println(s)
	var n_str = ""
	Tab_Num := 0
	for _, i := range s {
		index := string(i)
		if index == "{" {
			Tab_Num++
			n_str = strings.Join([]string{n_str, index, "\r\n", indent(Tab_Num)}, "")
		}
		if index == "}" {
			Tab_Num--
			n_str = strings.Join([]string{n_str, index, "\r\n", indent(Tab_Num)}, "")
		}

		if index == "[" {
			Tab_Num++
			n_str = strings.Join([]string{n_str, index, "\r\n", indent(Tab_Num)}, "")
		}
		if index == "]" {
			Tab_Num--
			n_str = strings.Join([]string{n_str, index, "\r\n", indent(Tab_Num)}, "")
		}

		n_str = strings.Join([]string{n_str, index}, "")

	}
	fmt.Println("********************************")
	fmt.Println(n_str)

}

func indent(number int) string {
	TAB := "   "
	for i := 0; i < number; i++ {
		TAB += TAB
	}
	return TAB
}
