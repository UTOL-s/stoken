package stoken

import (
	"log"
	
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

func Init() {
	
	conf := config.Config{
		SuperTokenUrl:    config.SuperTokenURL(),
		SuperTokenApiKey: config.SuperTokensKey(),
		EmailHost:        config.EmailHost(),
		FromEmail:        config.EmailFrom(),
		Password:         config.EmailPassword(),
	}
	
	tService := token.NewTokenService(conf)
	
	err := tService.TokenInit()
	if err != nil {
		log.Fatal(err.Error())
	}
	
	//return err
}
