package main

import (
	"math/rand"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Employee struct {
	Ename    string `gorm:"column:ENAME;not null"`
	Essn     string `gorm:"column:ESSN;primaryKey;autoIncrement:false;not null;unique"`
	Address  string `gorm:"column:ADDRESS;not null"`
	Salary   int    `gorm:"column:SALARY;not null"`
	Superssn string `gorm:"column:SUPERSSN;not null"`
	Dno      int    `gorm:"column:DNO;not null"`
}

type Department struct {
	Dname        string    `gorm:"column:DNAME;not null"`
	Mgrssn       string    `gorm:"column:MGRSSN;not null;unique"`
	Mgrstartdate time.Time `gorm:"column:MGRSTARTDATE;not null"`
	Dno          int       `gorm:"column:DNO;primaryKey;autoIncrement:false;not null;unique"`
}

type Project struct {
	Pname     string `gorm:"column:PNAME;not null"`
	Plocation string `gorm:"column:PLOCATION;not null"`
	Pno       int    `gorm:"column:PNO;primaryKey;autoIncrement:false;not null;unique"`
	Dno       int    `gorm:"column:DNO;not null;unique"`
}

type WorksOn struct {
	Essn  string `gorm:"column:ESSN;primaryKey;autoIncrement:false;not null"`
	Pno   int    `gorm:"column:PNO;primaryKey;autoIncrement:false;not null"`
	Hours int    `gorm:"column:HOURS;not null"`
}

func Connect() *gorm.DB {
	if db, err := gorm.Open(mysql.Open("root:root@/company?charset=utf8mb4&parseTime=True&loc=Local")); err != nil {
		println("Incorrect!")
		return nil
	} else {
		println("Connect OK!")
		return db
	}
}

func InsertData() {
	db := Connect()

	fdepartment := &Department{
		Dname:        "",
		Mgrssn:       "",
		Mgrstartdate: time.Time{},
		Dno:          0,
	}

	for k, v := range []rune("赵钱孙李") {
		println(k, string(v))
		fdepartment.Dname = string(v) + "部"
		fdepartment.Mgrssn = "1232" + strconv.Itoa(k)
		fdepartment.Mgrstartdate = time.Now()
		fdepartment.Dno = k + 1
		db.Table("DEPARTMENT").Create(fdepartment)
	}
	fdepartment.Dno = 5
	fdepartment.Dname = "研发" + "部"
	fdepartment.Mgrssn = "12324"
	db.Table("DEPARTMENT").Create(fdepartment)

	fproject := &Project{
		Pname:     "",
		Plocation: "",
		Pno:       0,
		Dno:       0,
	}

	for k, v := range []rune("赵钱孙李周吴郑王冯陈") {
		println(k, string(v))
		fproject.Pname = string(v) + "工程"
		fproject.Plocation = "北京"
		fproject.Pno = k + 1
		fproject.Dno = k%5 + 1
		db.Table("PROJECT").Create(fproject)
	}

	fwork := &WorksOn{
		Essn:  "",
		Pno:   0,
		Hours: 0,
	}

	for k, v := range []rune("赵钱孙李周吴郑王冯陈褚卫蒋沈韩杨朱秦尤许何吕施张孔曹严华金魏陶姜戚谢") {
		println(k, string(v))
		fwork.Essn = "1232" + strconv.Itoa(k)
		fwork.Pno = k%10 + 1
		fwork.Hours = rand.Intn(10) + 1
		db.Table("WORKS_ON").Create(fwork)
	}

	femployee := &Employee{
		Ename:    "",
		Essn:     "",
		Address:  "",
		Salary:   0,
		Superssn: "",
		Dno:      0,
	}

	for k, v := range []rune("赵钱孙李周吴郑王冯陈褚卫蒋沈韩杨朱秦尤许何吕施张孔曹严华金魏陶姜戚谢邹喻柏水窦章云苏潘葛奚范彭郎鲁韦昌马苗凤花方") {
		println(k, string(v))
		femployee.Ename = string(v) + "红"
		femployee.Salary = 2000 * (rand.Intn(5) + 1)
		femployee.Essn = "1232" + strconv.Itoa(k)
		femployee.Address = "快乐老家"
		femployee.Dno = k%5 + 1
		femployee.Superssn = "1232" + strconv.Itoa(k%5)
		db.Table("EMPLOYEE").Create(femployee)
	}
}

func main() {
	InsertData()
}
