package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

/**
序列化反序列化
*/

const file = "user.gob"

type User struct {
	id   int
	Name string
}

func (u *User) String() string {
	return fmt.Sprintf("{id:%d,Name:%s}", u.id, u.Name)
}

func main() {
	user := &User{1, "superman"}
	fmt.Println("source obj:", user)
	obj := new(User)
	Save(file, user)
	Load(file, obj)
	fmt.Println("decode obj:", obj)

}

//Encode object to file
// 序列对象到文件 类似 python pickle.dumps(file)
func Save(path string, object interface{}) error {
	file, err := os.Create(path)
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(object)
	}
	file.Close()
	return err
}

// Decode gob file
//从文件反序列化到对象  类似 python pickle.loads(file)
func Load(path string, object interface{}) error {
	file, err := os.Open(path)
	if err == nil {
		decoder := gob.NewDecoder(file)
		decoder.Decode(object)
	}
	file.Close()
	return err
}
