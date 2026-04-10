package services

import (
	"errors"
	"time"

	"hiretest-api/configs"
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
		return nil, errors.New("invalid credentials")
	}

	if user.IsActive == false {
		return nil, errors.New("account is inactive")
	}

	if err := utils.CheckPassword(user.PasswordHash, password); err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Generate JWT tokens (access and refresh)
	accessTokenExpiresIn, err := time.ParseDuration(s.Cfg.Env.JWTAccessExpiresIn)
	if err != nil {
		return nil, errors.New("invalid JWT expiration duration")
	}
	accessToken, err := utils.GenerateJWT(user.ID, s.Cfg.Env.JWTSecret, accessTokenExpiresIn)
	if err != nil {
		return nil, errors.New("failed to generate access token")
	}

	refreshTokenExpiresIn, err := time.ParseDuration(s.Cfg.Env.JWTRefreshExpiresIn)
	if err != nil {
		return nil, errors.New("invalid JWT expiration duration")
	}
	refreshToken, err := utils.GenerateJWT(user.ID, s.Cfg.Env.JWTSecret, refreshTokenExpiresIn)
	if err != nil {
		return nil, errors.New("failed to generate refresh token")
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
		return nil, errors.New("refresh token is required")
	}

	claims, err := utils.ParseJWT(refreshToken, s.Cfg.Env.JWTSecret)
	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	sub, ok := claims["user_id"].(string)
	if !ok || sub == "" {
		return nil, errors.New("invalid token claims")
	}

	accessTokenExpiresIn, err := time.ParseDuration(s.Cfg.Env.JWTAccessExpiresIn)
	if err != nil {
		return nil, errors.New("invalid JWT expiration duration")
	}
	accessToken, err := utils.GenerateJWT(sub, s.Cfg.Env.JWTSecret, accessTokenExpiresIn)
	if err != nil {
		return nil, errors.New("failed to generate access token")
	}

	refreshTokenExpiresIn, err := time.ParseDuration(s.Cfg.Env.JWTRefreshExpiresIn)
	if err != nil {
		return nil, errors.New("invalid JWT expiration duration")
	}
	newRefreshToken, err := utils.GenerateJWT(sub, s.Cfg.Env.JWTSecret, refreshTokenExpiresIn)
	if err != nil {
		return nil, errors.New("failed to generate refresh token")
	}

	return &AuthTokens{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
	}, nil
}
