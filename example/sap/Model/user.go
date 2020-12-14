package Model

//User 用户模块
type User struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Position string `json:"position" db:"position"`
}

//Register 注册模块
type Register struct {
	Username 		string `json:"username"`
	Password		string `json:"password"`
	Position		string `json:"position"`
	ConfirmPassword string `json:"commit_password"`
}

//Godown 仓库模块
type Godown struct {
	Title     string `json:"title"      db:"title"`
	Author    string `json:"author"     db:"author"`
	IssueDate string `json:"issue_date" db:"issue_date"`
	Synopsis  string `json:"synopsis"   db:"synopsis"`
	Category  string `json:"category"   db:"category"`
	Price     int    `json:"price"      db:"price"`
	Inventory int    `json:"inventory"  db:"inventory"`
}
