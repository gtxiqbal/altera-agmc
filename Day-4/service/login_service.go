package service

import (
	"context"
	"github.com/gtxiqbal/altera-agmc/Day-4/model/dto"
)

type LoginService interface {
	DoLogin(ctx context.Context, loginDtoReq dto.LoginDtoReq) map[string]any
	DoRefreshToken(ctx context.Context, refreshDtoReq dto.RefreshDtoReq) map[string]any
}
