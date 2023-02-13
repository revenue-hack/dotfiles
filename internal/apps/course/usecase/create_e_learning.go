package usecase

import (
	"context"

	"gitlab.kaonavi.jp/ae/sardine/internal/apps/course"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/database"
	"gitlab.kaonavi.jp/ae/sardine/internal/ctxt"
	"gitlab.kaonavi.jp/ae/sardine/internal/errs"
)

func NewCreateELearning(
	connFactory database.ConnFactory,
	repos course.CreateRepository,
) course.CreateUseCase {
	return &createELearning{connFactory: connFactory, repos: repos}
}

type createELearning struct {
	connFactory database.ConnFactory
	repos       course.CreateRepository
}

func (uc *createELearning) Exec(ctx context.Context) (*course.CreateOutput, error) {
	conn, err := uc.connFactory.Create(ctx)
	if err != nil {
		return nil, errs.Wrap("[createElearning.Exec]connFactory.Createのエラー", err)
	}

	authedUser, err := ctxt.AuthenticatedUser(ctx)
	if err != nil {
		return nil, errs.Wrap("[createElearning.Exec]ctxt.AuthenticatedUserのエラー", err)
	}

	newRecord, err := uc.repos.CreateELearning(ctx, conn, authedUser)
	if err != nil {
		return nil, errs.Wrap("[createElearning.Exec]repos.CreateELearningのエラー", err)
	}
	return &course.CreateOutput{
		Course: *newRecord,
	}, nil
}
