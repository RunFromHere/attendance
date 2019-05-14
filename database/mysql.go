package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//数据库配置
const (
	User0     = "root"
	Password0 = "root"
	Network0  = "tcp"
	Ip0       = "127.0.0.1"
	LocalIp0  = "localhost"
	Port0     = 3306
	DBName0   = "attendance"
)

//数据库配置
const (
	User11     = "root"
	Password11 = "root"
	Network11  = "tcp"
	Ip11       = ""
	LocalIp11  = "localhost"
	Port11     = 3306
	DBName11   = "attendance"
)

//数据库配置
const (
	User2     = "root"
	Password2 = "Gkl616!@#$"
	Network2  = "tcp"
	LocalIp2  = "localhost"
	Port2     = 3306
	DBName2   = "attendance"
)

var SqlDB *sql.DB

func init() {
	var err error
	SqlDB, err = sql.Open("mysql", ParseSQLUrl(User2, Password2, Network2, LocalIp2, Port2, DBName2))
	if err != nil {
		log.Println("open database fail" + err.Error())
	}

	err = SqlDB.Ping()
	if err != nil {
		log.Println("connect database fail" + err.Error())
	}

	//SqlDB.SetConnMaxLifetime(100 * time.Second) //最大连接周期，超过时间的连接就close
	SqlDB.SetMaxOpenConns(100) //设置最大连接数
	SqlDB.SetMaxIdleConns(20)  //设置闲置连接数

	log.Println("connnect success")
}

//构建连接, 格式是：”用户名:密码@tcp(IP:端口)/数据库?charset=utf8”
func ParseSQLUrl(user, psd, network, ip string, port int, dbname string) string {
	//sqlUrl := user + ":" + password + "@" + "tcp(" + ip + ":" + port + ")/" + dbName + "?charset=utf8&parseTime=True&loc=Local"
	sqlUrl := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", user, psd, network, ip, port, dbname)
	log.Println(sqlUrl)
	return sqlUrl
}

func ParseLikeMatchForString(text string) string {
	return "%" + text + "%"
}
