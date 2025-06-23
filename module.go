package stoken

import (
	"github.com/UTOL-s/stoken/internal/token"
	"go.uber.org/fx"
)

const ModuleName = "stoken"

var FxSTokenModule = fx.Module(
	ModuleName,
	fx.Provide(
		fx.Annotate(token.NewDefaultTokenClientFactory, fx.As(new(token.TokeClientFactory))),
	),
)
