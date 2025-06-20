package stoken

import (
	"github.com/UTOL-s/stoken/internal/token"
	"github.com/UTOL-s/stoken/pkg/config"
)

func New(cfg config.Config) error {
	
	tService := token.NewTokenService(cfg)
	
	err := tService.TokenInit()
	
	return err
}
