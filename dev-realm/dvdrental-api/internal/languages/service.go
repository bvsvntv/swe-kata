package languages

import (
	"context"

	repo "dvdrental-api/internal/adapters/postgresql/sqlc"
)

type Service interface {
	FetchLanguages(ctx context.Context, arg repo.FetchLanguagesParams) ([]repo.Language, int64, error)
	GetLanguage(ctx context.Context, languageID int32) (repo.Language, error)
}

type svc struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{repo: repo}
}

func (s *svc) FetchLanguages(ctx context.Context, arg repo.FetchLanguagesParams) ([]repo.Language, int64, error) {
	languages, err := s.repo.FetchLanguages(ctx, arg)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.repo.CountLanguages(ctx)
	if err != nil {
		return nil, 0, err
	}

	return languages, total, nil
}

func (s *svc) GetLanguage(ctx context.Context, languageID int32) (repo.Language, error) {
	return s.repo.GetLanguage(ctx, languageID)
}
