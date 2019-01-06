package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"unsafe"
)

type User struct {
	name string
	age  int
}

type SliceMock struct {
	addr uintptr
	len  int
	cap  int
}

func (u *User) String() string {
	return fmt.Sprintf("{ name:%v,age:%v}", u.name, u.age)
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

/**
使用redis缓存复杂结构数据
*/
func main() {
	//对象序列化成byte 存储到redis
	StructToByte()
	//redis取数据 反序列化成对戏
	ByteToStruct()
}
