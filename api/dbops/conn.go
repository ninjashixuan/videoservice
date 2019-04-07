package dbops

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var (
	connDB *sql.DB
	err    error
)

var conf Conf

type Conf struct {
	User   string `yaml:"user"`
	Pwd    string `yaml:"pwd"`
	Host   string `yaml:"host"`
	Ddname string `yaml:"dbname"`
	Port   int    `yaml:"port"`
}


func init() {
	yamlfile, err := ioutil.ReadFile("E:/goweb/src/videoservice/api/conf/conf.yaml")
	if err != nil {
		fmt.Printf(err.Error())
	}

	err = yaml.Unmarshal(yamlfile, &conf)
	if err != nil {
		fmt.Printf(err.Error())
	}
	user := conf.User
	pwd := conf.Pwd
	host := conf.Host
	port := conf.Port
	dbname := conf.Ddname

	connDB, err = openMysql(host, port, user, pwd,  dbname)
	if err != nil {
		fmt.Printf("open fail %v", err)
		panic(err.Error())
	}
}

func openMysql(host string, port int, user, pwd, dbname string) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		pwd,
		host,
		port,
		dbname)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Printf("open mysql fail %v", err)
	}

	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(100)

	return db, err
}

