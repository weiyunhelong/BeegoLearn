package models 

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// 用户
type User struct{
	Id              int64    `orm:"auto"`
	Name            string   `orm:"size(100)"`
	Nickname        string   `orm:"size(100)"`
    Pwd             string   `orm:"size(100)"`
    Email           string   `orm:"size(100)"`
    Sex             string   `orm:"size(2)"`
	Roleid          string   `orm:"size(100)"`
	Status          int64    
	Phone           string   `orm:"size(16)"`
}

//根据用户数据总个数
func GetRecordNum(search string) int64 {
	
	o := orm.NewOrm()
	qs := o.QueryTable("user")
	if search !=""{
		qs=qs.Filter("Name",search)
	}
	var us []User
	num, err :=  qs.All(&us)
	if err == nil {
		return num
	}else{
		return 0
	}	
}
func SearchDataList(pagesize,pageno int,search string) (users []User) {
	o := orm.NewOrm()
	qs := o.QueryTable("user")
	if search !=""{
		qs=qs.Filter("Name",search)
	}
	var us []User
	cnt, err :=  qs.Limit(pagesize, (pageno-1)*pagesize).All(&us)
	if err == nil {
		fmt.Println("count", cnt)
	}
	return us
}
//初始化模型
func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User))
}