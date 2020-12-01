package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)
func main(){
	f, err := os.Open("F:/dict.txt") //打开保存地址文件
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var trans string
	slice := []string{}
	slice1 :=[]string{}  //键值切片
	slice2 :=[]string{}  //翻译切片
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
		if err != nil || io.EOF == err {
			break
		}
		slice = append(slice,line)
	}
	for i:=0;i<len(slice);i++ {
		if i%2 == 0 || i == 0 {
			slice1 = append(slice1,slice[i])
		}else {
			slice2 = append(slice2,slice[i])
		}
	}
	fmt.Print("请输入你要翻译的单词:")
	fmt.Scanln(&trans)
	trans = "#"+ trans+"\r\n"
	for i,_ :=range slice1{
		if (trans==slice1[i]){
			fmt.Println(slice2[i])
		}
	}
}