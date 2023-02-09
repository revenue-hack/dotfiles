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
		return nil, errs.NewInternalError("failed to connFactory.Create from createELearning: %v", err)
	}

	authedUser, err := ctxt.AuthenticatedUser(ctx)
	if err != nil {
		return nil, errs.NewInternalError("failed to ctxt.AuthenticatedUser from createELearning: %v", err)
	}

	newRecord, err := uc.repos.CreateELearning(ctx, conn, authedUser)
	if err != nil {
		return nil, errs.NewInternalError("failed to CreateELearning from createELearning: %v", err)
	}
	return &course.CreateOutput{
		Course: *newRecord,
	}, nil
}
