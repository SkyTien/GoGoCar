package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	client_table_name          = "client"
	car_table_name             = "car"
	maintain_detail_table_name = "maintain_detail"
)

type table interface {
}

type client struct {
	Id              int `orm:"pk"`
	Name            string
	Phone           string
	Address         string
	Company_name    string
	Company_phone   string
	Company_address string
	Company_id      string
	Memo            string
}

type car struct {
	Number    string `orm:"pk"`
	Client_id int
}
type maintain_detail struct {
	Seq            int `orm:"pk"`
	Car_number     string
	Date           string
	Part_name      string
	Import_factory string
	Part_status    string
	Price          string
	Volume         string
	Miles          string
	Payment_status string
}

type DB struct {
	Username string
	Password string
	IP       string
	Name     string
}

func (d *DB) Init() {
	// register model
	//
	orm.RegisterDriver("mysql", orm.DR_MySQL)

	// set default database
	//
	cmd := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8", d.Username, d.Password, d.IP, d.Name)
	orm.RegisterDataBase("default", "mysql", cmd, 30)
	//orm.RegisterModel(new(client), new(car), new(maintain_detail))

	fmt.Println("DB Init")
}

/*
func (d *DB) QueryClientTable(tableName string) (table interface{}) {
	o := orm.NewOrm()
	o.Using("default")
	clients := []client

	num, err := o.Raw("SELECT * FROM client WHERE name = ?", tableName).QueryRows(&clients)
	//num, err := o.QueryTable("client").All(&clients)
	fmt.Printf("client: %v\n", clients)
	//cc, found := clients.(client)
	//fmt.Printf("client: %v\n", found)
	//fmt.Printf("client: %v\n", cc)
	if err == nil {

		fmt.Println("user nums: ", num)
	} else {
		fmt.Println("err: ", err)
	}
	return clients
}*/

func (d *DB) Query(queryString string) (num int64, result []orm.Params) {

	var maps []orm.Params
	o := orm.NewOrm()
	o.Using("default")
	num, err := o.Raw(queryString).Values(&maps)

	if err != nil {
		fmt.Println("err: ", err)
		return 0, nil
	}
	return num, maps
}
