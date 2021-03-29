/***
** Dependency Inversion Principle :
** Entities must depend on abstractions, not on concretions
**/
package main

import (
	"fmt"
)

type MySQL struct{}

func (db MySQL) Query() interface{} {
	return []string{"alex", "john", "mike"}
}

type PostgreSQL struct{}

func (db PostgreSQL) Query() interface{} {
	return map[string]string{
		"a3f69c2b-d153-48fd-b10c-5b641657477b": "alex",
		"a4f69c2b-d153-48fd-b10c-5b641657477a": "john",
		"a5f69c2b-d153-48fd-b10c-5b641657477c": "mike",
	}
}

type DBConn interface {
	Query() interface{}
}

type UsersRepository struct {
	db DBConn
}

func (r UsersRepository) GetUsers() []string {
	var users []string
	res := r.db.Query()

	switch res.(type) {
	case map[string]string:
		for _, u := range res.(map[string]string) {
			users = append(users, u)
		}
		return users
	case []string:
		return res.([]string)
	}

	return []string{}
}

func main() {
	fmt.Println("Dependency Inversion Principle After")

	mysqlDB := MySQL{}
	postreSQLDB := PostgreSQL{}
	repoMySQL := UsersRepository{db: mysqlDB}
	repoPostgreSQL := UsersRepository{db: postreSQLDB}
	fmt.Println(repoMySQL.GetUsers())
	fmt.Println(repoPostgreSQL.GetUsers())
}
