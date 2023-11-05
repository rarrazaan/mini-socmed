package usecase

import (
	"activity-reporter/repository"
	"activity-reporter/shared/dto"
	"activity-reporter/shared/helper"
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

type resetPWUsecase struct {
	rdb      *redis.Client
	userRepo repository.UserRepo
}

type ResetPWUsecase interface {
	ForgetPW(ctx context.Context, email string) (dto.ForgetPWRes, error)
}

func NewResetPWUsecase(rdb *redis.Client, userRepo repository.UserRepo) *resetPWUsecase {
	return &resetPWUsecase{
		rdb:      rdb,
		userRepo: userRepo,
	}
}

func (ru *resetPWUsecase) ForgetPW(ctx context.Context, email string) (dto.ForgetPWRes, error) {
	_, err := ru.userRepo.FindUserByIdentifier(ctx, email)
	if err != nil {
		if errors.Is(err, helper.ErrUserNotFound) {
			return dto.ForgetPWRes{}, helper.ErrCredential
		}
		return dto.ForgetPWRes{}, helper.ErrInternalServer
	}
	key := "reset_pw_token"
	token := "test"
	ttl := time.Duration(24) * time.Hour

	op1 := ru.rdb.Set(ctx, key, token, ttl)
	if err := op1.Err(); err != nil {
		return dto.ForgetPWRes{}, helper.ErrInternalServer
	}

	op2 := ru.rdb.Get(context.Background(), key)
	if err := op2.Err(); err != nil {
		return dto.ForgetPWRes{}, helper.ErrInternalServer
	}
	res, err := op2.Result()
	if err != nil {
		return dto.ForgetPWRes{}, helper.ErrInternalServer
	}
	return dto.ForgetPWRes{
		Token: res,
	}, nil
}
