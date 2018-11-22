package main

import (
	"fmt"

	go_redis_orm "github.com/fananchong/go-redis-orm.v2"
)

func test11() {
	dbName := "db1"

	go_redis_orm.SetNewRedisHandler(go_redis_orm.NewDefaultRedisClient)
	go_redis_orm.CreateDB(dbName, []string{"127.0.0.1:6379"}, "", 0)

	// key值为1的 TestStruct1 数据
	data1 := NewTestStruct1(dbName, 1)
	data1.SetMyb(true)
	data1.SetMyf1(1.5)
	data1.SetMys1("hello")
	data1.SetMys2([]byte("world"))

	myst1 := data1.GetMyst1(true)
	myst1.A = 100
	myst1.B = 1000.99

	myst2 := data1.GetMyst2(true)
	myst2.C = append(myst2.C, 1)
	myst2.C = append(myst2.C, 2)
	myst2.C = append(myst2.C, 3)

	err := data1.Save()
	if err != nil {
		panic(err)
	}

	data1.SetMyi8(123456)
	dirtyData, _ := data1.DirtyData()
	data1.Save2(dirtyData)

	data2 := NewTestStruct1(dbName, 1)
	err = data2.Load()

	if err == nil {
		if data2.GetMyb() != true ||
			data2.GetMyf1() != 1.5 ||
			data2.GetMys1() != "hello" ||
			string(data2.GetMys2()) != "world" {
			panic("#1")
		}
		myst1 := data2.GetMyst1(false)
		if myst1.A != 100 || myst1.B != 1000.99 {
			panic("#2")
		}
		myst2 := data2.GetMyst2(false)
		if myst2.C[0] != 1 || myst2.C[1] != 2 || myst2.C[2] != 3 {
			panic("#3")
		}
	} else {
		panic(err)
	}

	err = data2.Delete()
	if err != nil {
		panic(err)
	}

	var hasKey int
	hasKey, err = data2.HasKey()
	if hasKey != 0 {
		panic("#2")
	}

	fmt.Println(data2)
	fmt.Println("OK")
}

func test1n() {
	dbName := "db2"

	go_redis_orm.SetNewRedisHandler(go_redis_orm.NewDefaultRedisClient)
	go_redis_orm.CreateDB(dbName, []string{"127.0.0.1:6379"}, "", 0)

	data1 := NewTestStruct2(dbName, 8)
	item1 := data1.NewItem(1)
	item1.SetMyf2(99.9)

	myst1 := item1.GetMyst1(true)
	myst1.A = 100
	myst1.B = 1000.99

	myst2 := item1.GetMyst2(true)
	myst2.C = append(myst2.C, 1)
	myst2.C = append(myst2.C, 2)
	myst2.C = append(myst2.C, 3)

	item2 := data1.NewItem(2)
	item2.SetMys1("hello")
	item2.SetMys2([]byte("world"))
	item2.SetMyi3(10000)
	err := data1.Save()
	if err != nil {
		panic(err)
	}

	item2.SetMyi8(654321)
	dirtyData, _ := data1.DirtyData()
	data1.Save2(dirtyData)

	data2 := NewTestStruct2(dbName, 8)
	err = data2.Load()
	if err != nil {
		panic(err)
	}
	fmt.Printf("2: %+v\n", data2.GetItem(1))
	fmt.Printf("2: %+v\n", data2.GetItem(2))
	data2.DeleteItem(1)
	data2.Save()

	data3 := NewTestStruct2(dbName, 8)
	data3.Load()
	for _, v := range data3.GetItems() {
		fmt.Printf("3: %+v\n", v)
	}
	data3.Delete()
	data3.Save()

	data4 := NewTestStruct2(dbName, 8)
	data4.Load()
	fmt.Printf("4: item count = %d\n", len(data4.GetItems()))

	fmt.Println("OK")
}

func main() {
	test11()
	test1n()
}

