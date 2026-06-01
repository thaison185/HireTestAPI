package services

import (
	"errors"
	"time"

	"hiretest-api/configs"
	code_errors "hiretest-api/internal/common/errors"
	"hiretest-api/internal/common/utils"
	"hiretest-api/internal/models"
	"hiretest-api/internal/repositories"
)

type AuthService struct {
	Repo *repositories.AuthRepository
	Cfg  configs.AppConfig
}

type AuthTokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginResult struct {
	Tokens AuthTokens   `json:"tokens"`
	User   *models.User `json:"user"`
}

func NewAuthService(repo *repositories.AuthRepository, cfg configs.AppConfig) *AuthService {
	return &AuthService{
		Repo: repo,
		Cfg:  cfg,
	}
}

func (s *AuthService) Login(email, password string) (*LoginResult, error) {
	user, err := s.Repo.FindUserByEmail(email)
	if err != nil {
		return nil, errors.New(code_errors.CodeInvalidCredentials)
	}

	if user.IsActive == false {
		return nil, errors.New(code_errors.CodeAccountInactive)
	}

	if err := utils.CheckPassword(user.PasswordHash, password); err != nil {
		return nil, errors.New(code_errors.CodeInvalidCredentials)
	}

	// Generate JWT tokens (access and refresh)
	accessTokenExpiresIn, err := time.ParseDuration(s.Cfg.Env.JWTAccessExpiresIn)
	if err != nil {
		return nil, errors.New(code_errors.CodeInvalidJWTExp)
	}
	accessToken, err := utils.GenerateJWT(s.Cfg.Env.JWTSecret, user.ID, user.Role, accessTokenExpiresIn)
	if err != nil {
		return nil, errors.New(code_errors.CodeAccessTokenGenFailed)
	}

	refreshTokenExpiresIn, err := time.ParseDuration(s.Cfg.Env.JWTRefreshExpiresIn)
	if err != nil {
		return nil, errors.New(code_errors.CodeInvalidJWTExp)
	}
	refreshToken, err := utils.GenerateJWT(s.Cfg.Env.JWTSecret, user.ID, user.Role, refreshTokenExpiresIn)
	if err != nil {
		return nil, errors.New(code_errors.CodeRefreshTokenGenFailed)
	}

	return &LoginResult{
		Tokens: AuthTokens{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
		User: user,
	}, nil
}

func (s *AuthService) Refresh(refreshToken string) (*AuthTokens, error) {
	if refreshToken == "" {
		return nil, errors.New(code_errors.CodeRefreshTokenRequired)
	}

	claims, err := utils.ParseJWT(refreshToken, s.Cfg.Env.JWTSecret)
	if err != nil {
		return nil, errors.New(code_errors.CodeInvalidRefreshToken)
	}

	sub, ok := claims["user_id"].(string)
	if !ok || sub == "" {
		return nil, errors.New(code_errors.CodeInvalidTokenClaims)
	}

	accessTokenExpiresIn, err := time.ParseDuration(s.Cfg.Env.JWTAccessExpiresIn)
	if err != nil {
		return nil, errors.New(code_errors.CodeInvalidJWTExp)
	}
	role, ok := claims["role"].(string)
	if !ok || role == "" {
		return nil, errors.New(code_errors.CodeInvalidTokenClaims)
	}
	accessToken, err := utils.GenerateJWT(s.Cfg.Env.JWTSecret, sub, role, accessTokenExpiresIn)
	if err != nil {
		return nil, errors.New(code_errors.CodeAccessTokenGenFailed)
	}

	refreshTokenExpiresIn, err := time.ParseDuration(s.Cfg.Env.JWTRefreshExpiresIn)
	if err != nil {
		return nil, errors.New(code_errors.CodeInvalidJWTExp)
	}
	newRefreshToken, err := utils.GenerateJWT(s.Cfg.Env.JWTSecret, sub, role, refreshTokenExpiresIn)
	if err != nil {
		return nil, errors.New(code_errors.CodeRefreshTokenGenFailed)
	}

	return &AuthTokens{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
	}, nil
}
