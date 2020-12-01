package main

import (
	"fmt"
)
/*40个苹果，分配给 "Mating","Sadrgf","aeidi","Evilie","Peier","Hiana","Edriano","Bgdoin","Ealizabeth"
1.名字中包含有1个i,I分1个
2.名字中包含有1个e,E分2个
3.名字中包含有1个o,O分3个
4.名字中包含有1个u,U分4个
计算每个人苹果，并剩余多少
 */
//定义苹果数目
const Apple =40
//程序主函数
func main(){
	left :=dispatchApple()
	result :=Distribution()
	fmt.Println("剩下：",left)
	for v,s:=range result{
		fmt.Printf("%v分得苹果%d个\n",v,s)
	}
}
//每个人应该得多少
func Distribution()(distribution map[string]int){
	users :=[]string{
		"Mating","Sadrgf","aeidi","Evilie","Peier","Hiana","Edriano","Bgdoin","Ealizabeth",
	}
	distribution =make(map[string]int,len(users))
	//遍历名字 v 是每个名字
	for _,v:= range users{
		//s 是每个名字里面的字母,匹配规则
		for _,s :=range v{
			a:=string(s) //类型转换
			switch a {
			case "e","E":
				distribution[v] ++
				break
			case "i","I":
				distribution[v] +=2
				break
			case "o","O":
				distribution[v] +=3
				break
			case "u","U":
				distribution[v] +=4
				break
			}
		}
	}
	return distribution
}
//dispatchApple 获取剩余苹果
func dispatchApple()(left int){
	var s int
	for _,v:=range Distribution(){
		s=s+v
	}
	left = Apple-s
	return
}