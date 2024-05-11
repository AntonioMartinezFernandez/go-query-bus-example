package main

import (
	"fmt"

	user "github.com/AntonioMartinezFernandez/go-query-bus-example/internal/user"
	query_bus "github.com/AntonioMartinezFernandez/go-query-bus-example/pkg/query-bus"
)

func main() {
	// Create a new bus
	bus := query_bus.NewQueryBus()

	// Register handler
	getUserHandler := user.NewGetUserHandler(user.NewInMemoryUserRepository())
	bus.Register(&user.GetUserQuery{}, getUserHandler)

	// Send a Query to get user 1
	qRes, qErr := bus.Dispatch(user.NewGetUserQuery("1"))
	if qErr != nil {
		fmt.Println("Error: ", qErr)
	}
	// Cast the query result
	res, ok := qRes.(*user.User)
	if !ok {
		panic("error casting query result")
	}

	// Print result
	fmt.Println("User -> ", res.Id(), res.Name(), res.Birthdate().Local())

	// Send a Query to get user 2
	qRes2, qErr2 := bus.Dispatch(user.NewGetUserQuery("2"))
	if qErr2 != nil {
		fmt.Println("Error: ", qErr2)
	}

	// Cast the query result
	res2, ok2 := qRes2.(*user.User)
	if !ok2 {
		panic("error casting query result")
	}

	// Print result
	fmt.Println("User 2 -> ", res2.Id(), res2.Name(), res2.Birthdate().Local())
}
