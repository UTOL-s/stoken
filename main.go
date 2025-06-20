package stoken

import (
	"github.com/UTOL-s/stoken/internal/token"
	"github.com/UTOL-s/stoken/pkg/config"
)

func New(cfg config.Config) error {
	
	// Get Configuration
	tService := token.NewTokenService(cfg)
	
	// Initiate Supertoken
	err := tService.TokenInit()
	
	return err
}
