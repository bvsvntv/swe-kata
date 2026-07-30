package actors

import (
	repo "dvdrental-api/internal/adapters/postgresql/sqlc"
	"dvdrental-api/internal/types"
)

type ActorsResponse struct {
	types.MessageResponse
	Actors []repo.Actor `json:"actors"`
	types.PaginatedResponse
}

type ActorResponse struct {
	types.MessageResponse
	Actor repo.Actor `json:"actor"`
}

type CreateActorRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UpdateActorRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UpdateActorPartialRequest struct {
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
}

type ActorFilmsResponse struct {
	types.MessageResponse
	Films []repo.Film `json:"films"`
}
