package actors

import (
	repo "dvdrental-api/internal/adapters/postgresql/sqlc"
	"dvdrental-api/internal/types"
)

type ActorsResponse struct {
	Actors     []repo.Actor     `json:"actors"`
	Pagination types.Pagination `json:"pagination"`
}

type ActorResponse struct {
	Actor repo.Actor `json:"actor"`
}
