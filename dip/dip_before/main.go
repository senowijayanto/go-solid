package main

import (
	"fmt"
)

type MySQL struct{}

func (db MySQL) QueryMySQL() []string {
	return []string{"alex", "john", "mike"}
}

type PostgreSQL struct{}

func (db PostgreSQL) QueryPostgreSQL() map[string]string {
	return map[string]string{
		"a3f69c2b-d153-48fd-b10c-5b641657477b": "alex",
		"a4f69c2b-d153-48fd-b10c-5b641657477a": "john",
		"a5f69c2b-d153-48fd-b10c-5b641657477c": "mike",
	}
}

type UsersRepository struct {
	db MySQL
	// db PostgreSQL
}

func (r UsersRepository) GetUsers() []string {
	res := r.db.QueryMySQL()
	// res := r.db.QueryPostgreSQL()
	// var users []string
	// for _, u := range res {
	// 	users = append(users, u)
	// }
	// return users
	return res
}

func main() {
	fmt.Println("Dependency Inversion Principle Before")

	mysqlDB := MySQL{}
	repo := UsersRepository{db: mysqlDB}
	fmt.Println(repo.GetUsers())
}
