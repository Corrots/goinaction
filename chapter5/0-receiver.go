package main

import "fmt"

type user struct {
	name string
	email string
}

func (u user) notify()  {
	fmt.Printf("Send Email to %s<%s>\n", u.name, u.email)
}

func (u *user) changeEmail(newEmail string)  {
	u.email = newEmail
}

func (u user) emailUpdate(newEmail string)  {
	u.email = newEmail
}

func main()  {
	bill := user{
		name: "Bill",
		email: "bill@test.com",
	}

	bill.notify()
	bill.emailUpdate("bill@admin.com")
	bill.notify()
	bill.changeEmail("abc@qq.com")
	bill.notify()
}