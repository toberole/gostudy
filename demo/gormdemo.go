package demo

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"sync"
	"time"
)

type Product struct {
	gorm.Model
	Title string
	Code  string
	Price uint
}

func GORMDemo1() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(&Product{})

	// 插入内容
	db.Create(&Product{Title: "新款手机", Code: "D42", Price: 1000})
	db.Create(&Product{Title: "新款电脑", Code: "D43", Price: 3500})

	// 读取内容
	var product Product
	db.First(&product, 1)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// 更新操作： 更新单个字段
	db.Model(&product).Update("Price", 2000)

	// 更新操作： 更新多个字段
	db.Model(&product).Updates(Product{Price: 2000, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 2000, "Code": "F42"})
}

var wg sync.WaitGroup

func fn1() {
	defer wg.Done()

	time.Sleep(time.Second * 5)
	fmt.Printf("fn1 ......\n")
}

func GORMDemo2() {
	n := 3
	wg.Add(n)

	for i := 0; i < n; i++ {
		go fn1()
	}

	wg.Wait()
	fmt.Println("GORMDemo2 ......")

}
