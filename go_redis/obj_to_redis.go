package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/go-redis/redis"
	"unsafe"
)

type User struct {
	Name string
	Age  int
}

type SliceMock struct {
	addr uintptr
	len  int
	cap  int
}

func (u *User) String() string {
	return fmt.Sprintf("{ name:%v,age:%v}", u.Name, u.Age)
}
func GetClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.199.199:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	//pong, err := client.Ping().Result()
	//fmt.Println(pong, err)
	return client
	// Output: PONG <nil>
}

func StructToByte() {
	var u = &User{"hc", 22}
	fmt.Println("原始数据为 ", u)
	Len := unsafe.Sizeof(*u)
	testBytes := &SliceMock{
		addr: uintptr(unsafe.Pointer(u)),
		cap:  int(Len),
		len:  int(Len),
	}
	data := *(*[]byte)(unsafe.Pointer(testBytes))
	fmt.Println("序列化成 byte 结果为 : ", data)
	// to redis
	client := GetClient()
	err := client.Set("objxx", data, 0).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("存储到redis成功")

}

func ByteToStruct() {

	//data := *(*[]byte)(unsafe.Pointer(testBytes))
	//fmt.Println("[]byte is : ", data)
	client := GetClient()
	data, _ := client.Get("objxx").Result()
	fmt.Println("获取redis数据成功")
	var r_user = *(**User)(unsafe.Pointer(&data))
	fmt.Printf("反序列化数据为 %v : ", r_user)

}

func ObjToRedisByGob() {
	var u = &User{"hc", 22}
	fmt.Println("原始数据为 ", u)
	var buf bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&buf) // Will write to network.

	// Encode (send) some values.
	err := enc.Encode(u)
	if err != nil {
		panic("序列化错误")
	}
	client := GetClient()
	e := client.Set("ob_gob", buf.Bytes(), 0).Err()
	if e != nil {
		panic("to redis error")
	}
}

func RedisToObjByGob() {
	client := GetClient()
	//data, _ := client.Get("ob_gob").Result()
	data, _ := client.Get("ob_gob").Bytes()

	dec := gob.NewDecoder(bytes.NewBuffer(data))
	var u User
	dec.Decode(&u)
	fmt.Println(u)
}

/**
使用redis缓存复杂结构数据
运行打印输出

方式1
原始数据为  { name:hc,age:22}
序列化成 byte 结果为 :  [148 1 98 0 0 0 0 0 2 0 0 0 0 0 0 0 22 0 0 0 0 0 0 0]
存储到redis成功
获取redis数据成功
反序列化数据为 { name:hc,age:22} :

方式2
直接使用gob将对象序列化成byte存入redis
从redis取出byte数据 反序列化 成对象

*/
func main() {
	// 方式1
	//对象序列化成byte 存储到redis
	//StructToByte()
	//redis取数据 反序列化成对戏
	//ByteToStruct()

	// 方式2
	ObjToRedisByGob()
	RedisToObjByGob()

}
