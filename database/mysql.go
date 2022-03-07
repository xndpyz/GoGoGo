package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := gin.Default()
	db := DBinit()
	defer db.Close()
	r.POST("/student", func(c *gin.Context) {
		var student Student
		_ = c.BindJSON(&student)
		db.Create(&student)
	})
	r.GET("/student/:ID", func(c *gin.Context) {
		id := c.Param("ID")
		var student Student
		_ = c.BindJSON(&student)

		db.Preload("Teachers").Preload("IDCard").Where("id = ?", id).First(&student)
		c.JSON(200, gin.H{"查询": student})
	})
	//嵌套预加载
	r.GET("/class/:ID", func(c *gin.Context) {
		id := c.Param("ID")
		var class Class
		_ = c.BindJSON(&class)
		//多层嵌套查询
		db.Preload("Students").Preload("Students.IDCard").Preload("Students.Teachers").Where("id = ?", id).First(&class)
		c.JSON(200, gin.H{"class": class})
	})

	r.Run(":6666")
}

func DBinit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/"+
		"smallbaby?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	return db
}

func manytomany() {
	db := DBinit()
	db.AutoMigrate(&Teacher{}, &Student{}, &IDCard{}, &Class{})

	s := Student{
		StudentName: "小宝",
		IDCard: IDCard{
			Num: 1,
		},
	}

	t := Teacher{
		TeacherName: "吴老师",
		Students:    []Student{s},
	}

	c := Class{
		ClassName: 1,
		Students:  []Student{s},
	}
	_ = db.Create(&c).Error
	_ = db.Create(&t).Error
}

func demo() {
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/"+
		"smallbaby?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&Product{})

	//curd操作
	//user := Product{Model: gorm.Model{}, Code: "golang", Price: 150}
	//db.Create(&user)

	var p Product
	db.First(&p, "id=?", 3)
	fmt.Println(p)
	var all []Product //查询所有数据用切片
	db.Where("Price > ?", 100).Find(&all)
	fmt.Println(all)

	db.Where("id = ?", 2).First(&Product{}).Update("Code", "php")
	db.Where("id = ?", 3).Find(&[]Product{}).Update("Code", "c#")

	db.Where("id in (?)", []int{1, 2}).Delete(&Product{})
	// 硬删除
	//db.Where("id in (?)", []int{3}).Unscoped().Delete(&Product{})
}

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

//自定义表名
func (p Product) TableName() string {
	return "my_p"
}

type Class struct {
	gorm.Model
	ClassName int
	Students  []Student
}

type Student struct {
	gorm.Model
	StudentName string
	ClassId     uint
	IDCard      IDCard
	Teachers    []Teacher `gorm:"many2many:student_teachers;"`
}

type IDCard struct {
	gorm.Model
	Num       int
	StudentId uint
}

type Teacher struct {
	gorm.Model
	TeacherName string
	Students    []Student `gorm:"many2many:student_teachers;"`
}
