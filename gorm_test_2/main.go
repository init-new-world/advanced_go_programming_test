package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	if db, err := gorm.Open(mysql.Open("root:root@/company?charset=utf8mb4&parseTime=True&loc=Local")); err != nil {
		println("Incorrect!")
		return nil
	} else {
		println("Connect OK!")
		return db
	}
}

func Question1(db *gorm.DB, args ...string) {
	type Res struct {
		Essn string `gorm:"column:ESSN"`
	}

	res := []Res{}

	db.Raw("SELECT EMPLOYEE.ESSN FROM EMPLOYEE,WORKS_ON WHERE EMPLOYEE.ESSN=WORKS_ON.ESSN AND WORKS_ON.PNO=?;", args[0]).Scan(&res)

	for k, v := range res {
		fmt.Printf("%d %s\n", k, v.Essn)
	}
}

func Question2(db *gorm.DB, args ...string) {
	type Res struct {
		Ename string `gorm:"column:ENAME"`
	}

	res := []Res{}

	db.Raw("SELECT EMPLOYEE.ENAME FROM EMPLOYEE,WORKS_ON,PROJECT WHERE EMPLOYEE.ESSN=WORKS_ON.ESSN AND PROJECT.PNO=WORKS_ON.PNO AND PROJECT.PNAME=?;", args[0]).Scan(&res)

	for k, v := range res {
		fmt.Printf("%d %s\n", k, v.Ename)
	}
}

func Question3(db *gorm.DB, args ...string) {
	type Res struct {
		Ename   string `gorm:"column:ENAME"`
		Address string `gorm:"column:ADDRESS"`
	}

	res := []Res{}

	db.Raw("SELECT EMPLOYEE.ENAME,EMPLOYEE.ADDRESS FROM EMPLOYEE,DEPARTMENT WHERE EMPLOYEE.SUPERSSN=DEPARTMENT.MGRSSN AND DNAME=?;", args[0]).Scan(&res)

	for k, v := range res {
		fmt.Printf("%d %s %s\n", k, v.Ename, v.Address)
	}
}

func Question4(db *gorm.DB, args ...interface{}) {
	type Res struct {
		Ename   string `gorm:"column:ENAME"`
		Address string `gorm:"column:ADDRESS"`
	}

	res := []Res{}

	db.Raw("SELECT EMPLOYEE.ENAME,EMPLOYEE.ADDRESS FROM EMPLOYEE,DEPARTMENT WHERE EMPLOYEE.SUPERSSN=DEPARTMENT.MGRSSN AND DNAME=? AND SALARY<?;", args[0], args[1]).Scan(&res)

	for k, v := range res {
		fmt.Printf("%d %s %s\n", k, v.Ename, v.Address)
	}
}

func Question5(db *gorm.DB, args ...string) {
	type Res struct {
		Ename string `gorm:"column:ENAME"`
	}

	res := []Res{}

	db.Raw("SELECT DISTINCT EMPLOYEE.ENAME FROM EMPLOYEE,WORKS_ON WHERE EMPLOYEE.ESSN=WORKS_ON.ESSN AND PNO!=?;", args[0]).Scan(&res)

	for k, v := range res {
		fmt.Printf("%d %s\n", k, v.Ename)
	}
}

func Question6(db *gorm.DB, args ...string) {
	type Res struct {
		Ename string `gorm:"column:ENAME"`
		Dname string `gorm:"column:DNAME"`
	}

	res := []Res{}

	db.Raw("SELECT ENAME,DNAME FROM DEPARTMENT,EMPLOYEE WHERE DEPARTMENT.DNO=(SELECT DEPARTMENT.DNO FROM EMPLOYEE,DEPARTMENT WHERE EMPLOYEE.ENAME=? AND EMPLOYEE.ESSN=DEPARTMENT.MGRSSN) AND EMPLOYEE.DNO=DEPARTMENT.DNO;", args[0]).Scan(&res)

	for k, v := range res {
		fmt.Printf("%d %s %s\n", k, v.Ename, v.Dname)
	}
}

func Question7(db *gorm.DB, args ...string) {
	type Res struct {
		Essn string `gorm:"column:ESSN"`
	}

	res := []Res{}

	db.Raw("SELECT A.ESSN FROM (SELECT EMPLOYEE.ESSN FROM EMPLOYEE JOIN WORKS_ON ON WORKS_ON.ESSN=EMPLOYEE.ESSN AND "+
		"PNO=?) A,(SELECT EMPLOYEE.ESSN FROM EMPLOYEE JOIN WORKS_ON ON WORKS_ON.ESSN=EMPLOYEE.ESSN AND PNO=?) B WHERE A.ESSN=B.ESSN;", args[0], args[1]).Scan(&res)

	for k, v := range res {
		fmt.Printf("%d %s\n", k, v.Essn)
	}
}

func Question8(db *gorm.DB, args ...int) {
	type Res struct {
		Dname string `gorm:"column:DNAME"`
	}

	res := []Res{}

	db.Raw("SELECT A.DNAME FROM (SELECT EMPLOYEE.DNO,SALARY,DEPARTMENT.DNAME FROM EMPLOYEE JOIN DEPARTMENT ON EMPLOYEE.DNO=DEPARTMENT.DNO GROUP BY EMPLOYEE.DNO HAVING AVG(EMPLOYEE.SALARY)<?) A;;", args[0]).Scan(&res)

	for k, v := range res {
		fmt.Printf("%d %s\n", k, v.Dname)
	}
}

func Question9(db *gorm.DB, args ...interface{}) {
	type Res struct {
		Ename string `gorm:"column:ENAME"`
	}

	res := []Res{}

	db.Raw("SELECT ENAME FROM EMPLOYEE,WORKS_ON WHERE EMPLOYEE.ESSN=WORKS_ON.ESSN GROUP BY EMPLOYEE.ESSN HAVING COUNT(*)>=? AND SUM(HOURS)<=?;", args[0], args[1]).Scan(&res)

	for k, v := range res {
		fmt.Printf("%d %s\n", k, v.Ename)
	}
}

func WorkLine() {
	db := Connect()
	//Question1(db, "1")
	//Question2(db, "赵工程")
	//Question3(db, "Research Department")
	//Question4(db, "周部1", 2000)
	//Question5(db, "1")
	//Question6(db, "张红")
	//Question7(db, "1", "2")
	//Question8(db, 5000)
	Question9(db, 3, 8)
}

func main() {
	WorkLine()
}
