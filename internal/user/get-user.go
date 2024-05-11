package user

import (
	"errors"

	query_bus "github.com/AntonioMartinezFernandez/go-query-bus-example/pkg/query-bus"
)

type GetUserQuery struct {
	id string
}

func NewGetUserQuery(id string) *GetUserQuery {
	return &GetUserQuery{
		id: id,
	}
}

func (q *GetUserQuery) Data() map[string]interface{} {
	return map[string]interface{}{"id": q.id}
}

type GetUserHandler struct {
	userRepo UserRepository
}

func NewGetUserHandler(userRepo UserRepository) *GetUserHandler {
	return &GetUserHandler{
		userRepo: userRepo,
	}
}

func (h *GetUserHandler) Handle(query query_bus.Query) (interface{}, error) {
	q, ok := query.(*GetUserQuery)
	if !ok {
		return nil, errors.New("invalid command")
	}

	userId := q.Data()["id"].(string)
	return h.userRepo.GetUser(userId)
}
