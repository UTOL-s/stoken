package stoken

import (
	"context"
	
	"github.com/ankorstore/yokai/config"
	"go.uber.org/fx"
)

const ModuleName = "stoken"

var FxSTokenModule = fx.Module(
	ModuleName,
	fx.Provide(
		fx.Annotate(NewDefaultTokenClientFactory, fx.As(new(TokenClientFactory))),
		TokenInitialize,
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

func TokenInitialize(f FxTokenClientParam) (*FxTokenClient, error) {
	
	f.Lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			err := f.Factory.TokenInit()
			
			if err != nil {
				return err
			}
			return nil
		},
	})
	
	return &FxTokenClient{
		STokenInit: true,
	}, nil
	
}
