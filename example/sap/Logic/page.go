package Logic

//分页判断
func Page(count int)(s []int,limit int){
	page := count/2
	limit = count%2
	j:=0
	for i:=0;i<page;i++{
		j++
		s=append(s,j)
	}
	return
}

//func Getpage(page string)(number int){
//	if a,err:=strconv.Atoi(page);err!=nil{
//		switch page {
//		case "reduce":
//		case "add":
//		case "last":
//			number=
//		}
//	}else {
//
//	}
//	return
//}
