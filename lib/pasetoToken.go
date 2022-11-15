package lib

import (
	"errors"
	"os"
	"time"

	"github.com/api-auth/config"
	"github.com/api-auth/model"
	"github.com/o1egl/paseto"
)

type PasetoMaker struct {
	paseto *paseto.V2
}

func Init() *PasetoMaker {
	return &PasetoMaker{
		paseto: paseto.NewV2(),
	}
}

func (pm *PasetoMaker) GetToken(id int, username string) (string, *model.Payload, string, *model.Payload, error) {
	token, tokenPayload, err := pm.CreateToken(id, username)
	if err != nil {
		return "", nil, "", nil, err
	}
	tokenRefresh, tokenRefreshPayload, err := pm.CreateRefreshToken(id, username)
	if err != nil {
		return "", nil, "", nil, err
	}
	return token, tokenPayload, tokenRefresh, tokenRefreshPayload, err

}
func (pm *PasetoMaker) CreateToken(id int, username string) (string, *model.Payload, error) {
	TOKEN_SYMMETRIC_KEY := os.Getenv("TOKEN_SYMMETRIC_KEY")
	tokenSymmetricKey := []byte(TOKEN_SYMMETRIC_KEY)
	duration := ParseDuration(config.TOKEN_DURATION)
	payload, err := model.NewPayload(id, username, duration)
	if err != nil {
		return "", payload, err
	}

	token, err := pm.paseto.Encrypt(tokenSymmetricKey, payload, nil)
	return token, payload, err
}

func (pm *PasetoMaker) CreateRefreshToken(id int, username string) (string, *model.Payload, error) {
	var duration time.Duration
	TOKEN_REFRESH_SYMMETRIC_KEY := os.Getenv("TOKEN_SYMMETRIC_REFRESH_KEY")
	tokenRefreshSymmetricKey := []byte(TOKEN_REFRESH_SYMMETRIC_KEY)
	today := time.Now()
	if IsWeekend(today) {
		duration = ParseDuration(config.TOKEN_REFRESH_WEEKEND_DURATION)
	} else {
		duration = ParseDuration(config.TOKEN_REFRESH_WEEKDAYS_DURATION)
	}
	payload, err := model.NewPayload(id, username, duration)
	if err != nil {
		return "", payload, err
	}

	token, err := pm.paseto.Encrypt(tokenRefreshSymmetricKey, payload, nil)
	return token, payload, err
}

func CheckToken(token string) (*model.Payload, error) {
	var maker PasetoMaker
	TOKEN_SYMMETRIC_KEY := os.Getenv("TOKEN_SYMMETRIC_KEY")
	symmetricKey := []byte(TOKEN_SYMMETRIC_KEY)
	payload := &model.Payload{}
	err := maker.paseto.Decrypt(token, symmetricKey, payload, nil)

	if err != nil {
		return nil, errors.New("token is invalid")
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}
	return payload, nil
}

func CheckRefreshToken(token string) (*model.Payload, error) {
	var maker PasetoMaker
	TOKEN_REFRESH_SYMMETRIC_KEY := os.Getenv("TOKEN_SYMMETRIC_REFRESH_KEY")
	symmetricKey := []byte(TOKEN_REFRESH_SYMMETRIC_KEY)
	payload := &model.Payload{}
	err := maker.paseto.Decrypt(token, symmetricKey, payload, nil)

	if err != nil {
		return nil, errors.New("token is invalid")
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
