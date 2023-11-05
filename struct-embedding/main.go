package main

import (
	"fmt"
	"time"
)

type base struct {
	created_at time.Time
	updated_at time.Time
}

type user struct {
	base
	name string
}

func (b *base) describe() {
	fmt.Println("created_at: ", b.created_at)
}
func main() {
	user := &user{
		name: "admin",
		base: base{
			created_at: time.Now(),
		},
	}

	user.describe()
}
