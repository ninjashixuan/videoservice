package dbops

import (
	"database/sql"
	"encoding/json"
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
    mysql := conf.getConf()
	data, err := json.Marshal(conf)
	if err != nil {
		fmt.Printf(err.Error())
	}
	user := mysql.User
	pwd := mysql.Pwd
	host := mysql.Host
	port := mysql.Port
	dbname := mysql.Ddname

	connDB, err = openMysql(port, user, pwd, host, dbname)
	if err != nil {
		fmt.Printf("open fail %v", err)
		panic(err.Error())
	}
}

func openMysql(port int, host, user, pwd, dbname string) (db *sql.DB, err error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		pwd,
		host,
		port,
		dbname)

	db, err = sql.Open("mysql", dsn)

	if err != nil {
		fmt.Printf("open mysql fail %v", err)
		return
	}

	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(100)

	return
}

func (c *Conf) getConf() *Conf {
	yamlfile, err := ioutil.ReadFile("E:/goweb/src/videoservice/api/conf/conf.yaml")
	if err != nil {
		log.Printf("yaml get err: %v", err)
	}

	err = yaml.Unmarshal(yamlfile, c)
	if err != nil {
		fmt.Printf(err.Error())
	}

	return c
}
	
