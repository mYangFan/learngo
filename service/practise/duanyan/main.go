package main

import "fmt"

type Basic interface {
	GetName() string
	SetName(name string) error
}

type User struct {
	Name string
}

func (u *User) GetName() string {
	return u.Name
}

func SetName(name interface{}) error {
	if sname,ok := name.(map[string]string); ok {
		fmt.Println(sname["name"])
	}
	return nil
}

func main() {
	name := map[string]string{
		"name":"yangfan",
	}
	SetName(name)
}
