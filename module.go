package stoken

import (
	"github.com/ankorstore/yokai/config"
	"go.uber.org/fx"
)

const ModuleName = "stoken"

var FxSTokenModule = fx.Module(
	ModuleName,
	fx.Provide(
		fx.Annotate(NewDefaultTokenClientFactory, fx.As(new(TokenClientFactory))),
		TokenInit,
	),
	fx.Invoke(func(*FxTokenClient) {}),
)

type FxTokenClientParam struct {
	fx.In
	Lifecycle fx.Lifecycle
	Config    *config.Config
	Factory   TokenClientFactory
}

type FxTokenClient struct {
	STokenInit bool
}

func TokenInit(f FxTokenClientParam) (*FxTokenClient, error) {
	
	err := f.Factory.TokenInit()
	
	if err != nil {
		return &FxTokenClient{
			STokenInit: false,
		}, err
	}
	
	return &FxTokenClient{
		STokenInit: true,
	}, nil
}
