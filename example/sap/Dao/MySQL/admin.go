package MySQL

import (
	"go.uber.org/zap"
	"sap/Model"
)

//SelectMySQL() 查询数据库所有数据
//func SelectMySQL()(msgs []Model.Godown){
//	var msg Model.Godown
//	//issue_date,inventory,synopsis,price
//	sqlstr:="select title,author,category,issue_date,inventory,synopsis,price from godown "
//	rows,err:=db.Query(sqlstr)
//	if err!=nil{
//		zap.L().Error("Query MySQL failed",zap.Error(err))
//		return
//	}
//	for rows.Next() {
//		errn := rows.Scan(&msg.Title,&msg.Author,&msg.Category,&msg.IssueDate,&msg.Inventory,&msg.Synopsis,&msg.Price)
//		if errn != nil {
//			zap.L().Error("Query MySQL rows failed",zap.Error(err))
//		}
//		msgs=append(msgs,msg)
//	}
//	return msgs
//}

//SelectMySQL 按分页查询
func SelectMySQL()(msgs []Model.Godown,count int){
	var msg Model.Godown
	tx,err:=db.Beginx() //申明事务
	if err!=nil{
		if tx!=nil{
			tx.Rollback()//回滚
		}
		zap.L().Error("CountMySQL failed",zap.Error(err))
		return
	}
	sqlstr1:="select count(*) from godown"
	tx.MustExec(sqlstr1) //事务1
	sqlstr2:="select title,author,category,issue_date,inventory,synopsis,price from godown limit ?"
	tx.MustExec(sqlstr2,2) //事务2
	err=tx.Commit() //事务提交 提交之后执行
	if err!=nil{
		tx.Rollback()
		zap.L().Error("CountMySQL failed",zap.Error(err))
	}
	err =db.Get(&count,sqlstr1)
	if err!=nil{
		zap.L().Error("Query MySQL count failed",zap.Error(err))
		return
	}
	rows,err:=db.Query(sqlstr2,3)
	if err!=nil{
		zap.L().Error("Query MySQL list failed",zap.Error(err))
		return
	}
	for rows.Next() {
		errn := rows.Scan(&msg.Title,&msg.Author,&msg.Category,&msg.IssueDate,&msg.Inventory,&msg.Synopsis,&msg.Price)
		if errn != nil {
			zap.L().Error("Query MySQL rows failed",zap.Error(err))
		}
		msgs=append(msgs,msg)
	}
	return msgs,count
}

//InsertIntoMySQL 往数据库中增加数据
func InsertIntoMySQL(godown *Model.Godown)(err error){
	sqlstr:="insert into godown (title, author,price,inventory,issue_date,synopsis,category) values (?,?,?,?,?,?,?)"
	_, err =db.Exec(sqlstr,
		godown.Title,
		godown.Author,
		godown.Price,
		godown.Inventory,
		godown.IssueDate,
		godown.Synopsis,
		godown.Category)
	if err!=nil{
		zap.L().Error("insert into godown failed",zap.Error(err))
		return
	}
	return

}

//DeleteMySQL 删除数据
func DeleteMySQL(title string)(err error){
	var inventory int
	//sqlstr1:="delete from godown where title=?"
	sqlstr1:="select inventory from godown where title=?"
	err =db.Get(&inventory,sqlstr1,title)
	if err!=nil {
		zap.L().Error("db.get",zap.Error(err))
		return
	}else if inventory==0{
		zap.L().Error("inventor is zero",zap.Error(err))
		return
	}
	sqlstr2:="delete from godown where title=?"
	_,err =db.Exec(sqlstr2,title)
	if err!=nil {
		zap.L().Error("delete msg failed",zap.Error(err))
		return
	}
	return
}