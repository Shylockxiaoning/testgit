package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args
	if args == nil || len(args)!=3{
		fmt.Println("格式为：源文件 目标文件")
		return
	}
	srcPath := args[1]
	brtPath := args[2]
	fmt.Printf("源文件为:%s,目标文件为:%s\n",srcPath,brtPath)
	if (srcPath==brtPath){
		fmt.Println("无法将当前目录拷贝到当前目录")
		return
	}
	srcFile,err1 :=os.Open(srcPath)
	if err1 != nil {
		fmt.Println("源文件错误",err1)
		return
	}
	brtFile,err2 :=os.Create(brtPath)
	if err1 != nil {
		fmt.Println("目标文件错误",err2)
		return
	}

	buf :=make([]byte,1024)
	for {
		n,err :=srcFile.Read(buf)
		if err != nil && err == io.EOF {
			fmt.Println("读取源文件失败",err)
			break
		}
		brtFile.Write(buf[:n])
		if n == 0 {
			fmt.Println("操作完成")
			return
		}
		srcFile.Close()
		brtFile.Close()

	}

}