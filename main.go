package main

import (
	"fmt"
	"gexcel/routers"
)

func main() {
	fmt.Println("广播系统生成excel表格")
	// aa := "33;45;34;35;36;37;38;39;46;41;42;56;43;58;44;15;60;62;61;59;18;52;29;30;4;13;54;55;53;21;22;23;14;19;32;8;24;28;10;57;9;25;27;3;12;31;26;17;6;2;11;20;7;5;63;69;68;70;47;67;64;"
	// arr := strings.Split(aa, ";")
	// ar := arr[:len(arr)-1]
	// rs := Sort(ar)
	// fmt.Println(rs)
	route := routers.Newrouter()
	route.Run(":8080")
}

// func Sort(arr []string) []int {
// 	a := []int{}
// 	for _, val := range arr {
// 		bi, _ := strconv.Atoi(val)
// 		a = append(a, bi)
// 	}
// 	newarr := []int{}
// 	newarr = append(newarr, a[0])
// 	for i := 1; i < len(a); i++ {
// 		k := 0
// 		//newarr = append(newarr, a[i])
// 		for j := len(newarr) - 1; j >= 0; j-- {
// 			if a[i] < newarr[j] {
// 				newarr[j+1] = newarr[j]
// 				k = j
// 			}

// 		}
// 		newarr[k] = a[i]

// 	}
// 	return newarr
// }
