# go utils


## 字符串操作
1. CheckCharDoSpecial 检查字符串,去掉特殊字符


## 数据库
> 实现golang的数据库增删改查操作,无需传入表结构,快捷操作数据库

### 初使数据库连接
```
var config = DBConfig{
	Host:    "localhost",
	Port:    3308,
	Name:    "root",
	Pass:    "123456",
	DBName:  "sgfoot",
	Charset: "utf8mb4",
}
var (
	MasterDB Modeler //定义主DB
)

func init() {
	db := InitDB(config)
	model := NewDB(db)
	MasterDB = model
}
```

### 查询所有数据

```
ls, err := MasterDB.Find("select * from book")
```

### 查询单行数据

```
ls, err := MasterDB.First("select * from book where id = ?", 1)
```

### 查询单列数据

```
ls, err := MasterDB.Pluck("select * from book where id = ?", "book_name", 3)
```

### 增加数据
```
ls, err := MasterDB.Insert("insert into book set book_name=?, book_author=?, book_province=?", "论语", "孔子", "山东")
```


### 修改数据
```
ls, err := MasterDB.Update("update book set book_name=? where id=?", "国学-论语", 3)
```


### 删除数据
```
ls, err := MasterDB.Delete("delete from book where id = ?", 1)
```
