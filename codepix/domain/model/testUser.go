package main

import "fmt"

type UserTest struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

func (u *UserTest) String() string {
    return fmt.Sprintf("ID ", u.ID, "Hi! My name is %s", u.Name, "%s", u.Email)
}

func main() {
    u := &UserTest{
		ID: "1",
		Name: "Antonio",
		Email: "halisonsc5@gmail.com",
    }

    fmt.Println(u)
}