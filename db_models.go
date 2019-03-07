package utils

import (
	"database/sql"
	"fmt"
)

// model 接口
type Modeler interface {
	Find(sql string, args ...interface{}) ([]map[string]interface{}, error)    //查询所有数据,即2维数据
	First(sql string, args ...interface{}) (map[string]interface{}, error)     //查询一行数据,即1维数据
	Pluck(sql string, name string, args ...interface{}) ([]interface{}, error) //查询一列数据,即1维数据 map
	Update(sql string, args ...interface{}) (int64, error)                     //更新
	Delete(sql string, args ...interface{}) (int64, error)                     //删除
	Insert(sql string, args ...interface{}) (int64, error)                     //增加
}

// model结构
type ModelS struct {
	DB *sql.DB
}

//数据库配置结构
type DBConfig struct {
	Host    string //地址
	Port    int    //端口
	Name    string //名称
	Pass    string //密码
	DBName  string //库名
	Charset string //编码
}

//连接数据库
func InitDB(cfg DBConfig) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", cfg.Name, cfg.Pass, cfg.Host, cfg.Port, cfg.DBName, cfg.Charset)
	connection, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return connection
}

//实例一个数据库对象
func NewDB(db *sql.DB) Modeler {
	return &ModelS{
		DB: db,
	}
}

//查询数据库
func (m *ModelS) Find(sql string, args ...interface{}) ([]map[string]interface{}, error) {
	stmt, err := m.DB.Prepare(sql)
	defer stmt.Close()
	rows, err := stmt.Query(args...)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	//获取列名称
	var columns []string
	if columns, err = rows.Columns(); err != nil {
		return nil, err
	}
	//新建存储结果变量
	result := make([]map[string]interface{}, 0)
	count := len(columns)
	//存储数据变量
	values := make([]interface{}, count)
	//扫描变量
	scanValues := make([]interface{}, count)
	for rows.Next() {
		//将带&赋值给scanValues切片,然后全部给Scan方法
		for i := 0; i < count; i++ {
			scanValues[i] = &values[i]
		}
		err = rows.Scan(scanValues...) //赋值给scan
		if err != nil {
			continue
		}
		//新建一个实体,用于存储数据
		entity := make(map[string]interface{})
		for i, field := range columns {
			var v interface{}              //定义一个通用变量
			val := values[i]               //获取一个结果数据
			if b, ok := val.([]byte); ok { //判定是否是切片字节
				v = string(b)
			} else {
				v = val
			}
			entity[field] = v //以字段名为键,存储数据
		}
		//追加到结果集里
		result = append(result, entity)
	}
	return result, nil
}

//查询一行数据,即1维数据
func (m *ModelS) First(sql string, args ...interface{}) (map[string]interface{}, error) {
	stmt, err := m.DB.Prepare(sql)
	defer stmt.Close()
	rows, err := stmt.Query(args...)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var columns []string
	columns, err = rows.Columns()
	if err != nil {
		return nil, err
	}
	values := make([]interface{}, len(columns))
	scanValues := make([]interface{}, len(columns))
	result := make(map[string]interface{})
	for rows.Next() {
		for i := range columns {
			scanValues[i] = &values[i]
		}
		err = rows.Scan(scanValues...)
		if err != nil {
			continue
		}
		entity := make(map[string]interface{})
		for i, field := range columns {
			var v interface{}
			val := values[i]
			if b, ok := val.([]byte); ok {
				v = string(b)
			} else {
				v = val
			}
			entity[field] = v
		}
		result = entity
		break
	}
	return result, nil
}

//查询一列数据,即1维数据 map
func (m *ModelS) Pluck(sql string, name string, args ...interface{}) ([]interface{}, error) {
	stmt, err := m.DB.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(args...)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var columns []string
	columns, err = rows.Columns()
	if err != nil {
		return nil, err
	}
	values := make([]interface{}, len(columns))
	scanValues := make([]interface{}, len(columns))
	result := make([]interface{}, 0)
	for rows.Next() {
		for i := range columns {
			scanValues[i] = &values[i]
		}
		err = rows.Scan(scanValues...)
		if err != nil {
			continue
		}
		var one interface{}
		for i, field := range columns {
			var v interface{}
			val := values[i]
			if b, ok := val.([]byte); ok {
				v = string(b)
			} else {
				v = val
			}
			if field == name {
				one = v
				break
			}
		}
		result = append(result, one)
	}
	return result, nil
}

//更新
func (m *ModelS) Update(sql string, args ...interface{}) (int64, error) {
	stmt, err := m.DB.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(args...)
	if err != nil {
		return 0, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, nil
}

//删除
func (m *ModelS) Delete(sql string, args ...interface{}) (int64, error) {
	stmt, err := m.DB.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(args...)
	if err != nil {
		return 0, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, nil
}

//增加
func (m *ModelS) Insert(sql string, args ...interface{}) (int64, error) {
	stmt, err := m.DB.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(args...)
	if err != nil {
		return 0, err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastId, nil
}
