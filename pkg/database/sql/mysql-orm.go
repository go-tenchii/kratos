package sql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"math/rand"
)

const (
	Read  = "Read"
	Write = "Write"
)

var dbs map[string]map[string]*[]*gorm.DB

type ConnInfo struct {
	DBName  string
	Dsn     string
	MaxIdle int
	MaxOpen int
	Type    string
}

// 初始化mysql
func StartMysql(ConnInfos []ConnInfo) (err error) {
	dbs = map[string]map[string]*[]*gorm.DB{}
	for _, v := range ConnInfos {
		var mysqlDB *gorm.DB
		mysqlDB, err = gorm.Open("mysql", v.Dsn)
		if err == nil {
			mysqlDB.DB().SetMaxIdleConns(v.MaxIdle)
			mysqlDB.DB().SetMaxOpenConns(v.MaxOpen)
			mysqlDB.DB().SetConnMaxLifetime(0)
		} else {
			return
		}

		var type2dbs map[string]*[]*gorm.DB
		var ok bool
		if type2dbs, ok = dbs[v.DBName]; !ok {
			type2dbs = map[string]*[]*gorm.DB{}
			dbs[v.DBName] = type2dbs
		}

		var dbConns *[]*gorm.DB
		if dbConns, ok = type2dbs[v.Type]; !ok {
			dbConns = &[]*gorm.DB{}
			type2dbs[v.Type] = dbConns
		}

		*dbConns = append(*dbConns, mysqlDB)
	}
	return
}

// 获取mysql连接
func GetMysql(dbName string, t string) *gorm.DB {
	if type2dbs, ok := dbs[dbName]; ok {
		if dbConns, ok := type2dbs[t]; ok {
			//random 后续可选roundrobin，hash等
			lens := len(*dbConns)
			index := rand.Intn(lens)
			return (*dbConns)[index]
		}
	}
	return nil
}

// 关闭mysql
func CloseMysql() {
	for _, v := range dbs {
		for _, dbConns := range v {
			for _, conn := range *dbConns {
				conn.Close()
			}
		}
	}
}
