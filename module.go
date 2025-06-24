package stoken

import (
	"github.com/UTOL-s/stoken/internal/token"
	"github.com/ankorstore/yokai/config"
	"go.uber.org/fx"
)

const ModuleName = "stoken"

var FxSTokenModule = fx.Module(
	ModuleName,
	fx.Provide(
		fx.Annotate(
			token.NewDefaultTokenClientFactory,
			fx.As(new(token.TokenClientFactory)),
		),
		NewTokenClientInit,
	),
)

type FxTokenClientParam struct {
	fx.In
	Lifecycle fx.Lifecycle
	Config    *config.Config
	Factory   token.TokenClientFactory
}

func NewTokenClientInit(p FxTokenClientParam) error {
	
	err := p.Factory.TokenInit()
	if err != nil {
		return err
	}
	
	return nil
}
