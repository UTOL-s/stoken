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
	
	),
	fx.Invoke(TokenInitialize),
)

type FxTokenClientParam struct {
	fx.In
	Lifecycle fx.Lifecycle
	Config    *config.Config
	Factory   TokenClientFactory
}

type FxTokenClientResult struct {
	fx.Out
	STokenInit bool
}

func TokenInitialize(f FxTokenClientParam) (*FxTokenClientResult, error) {
	
	f.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			err := f.Factory.TokenInit()
			
			if err != nil {
				return err
			}
			return nil
		},
	})
	
	return &FxTokenClientResult{
		STokenInit: true,
	}, nil
	
	//email := f.Config.GetString("modules.stoken.email.username")
	//
	//apiBasePath := "/api/auth"
	//webBasePath := "/api/auth"
	//
	//err := supertokens.Init(
	//
	//	supertokens.TypeInput{
	//
	//		Supertokens: &supertokens.ConnectionInfo{
	//			ConnectionURI: f.Config.GetString("modules.stoken.apiUrl"),
	//			APIKey:        f.Config.GetString("modules.stoken.apiKey"),
	//		},
	//
	//		AppInfo: supertokens.AppInfo{
	//			AppName:       "UTOL",
	//			WebsiteDomain: "http://localhost:3000",
	//			APIDomain:     "http://localhost:8080",
	//			//APIDomain:       config.AppUrl(),
	//			APIBasePath:     &apiBasePath,
	//			WebsiteBasePath: &webBasePath,
	//			APIGatewayPath:  nil,
	//		},
	//
	//		RecipeList: []supertokens.Recipe{
	//			passwordless.Init(
	//				plessmodels.TypeInput{
	//
	//					FlowType: "USER_INPUT_CODE_AND_MAGIC_LINK",
	//					ContactMethodEmailOrPhone: plessmodels.ContactMethodEmailOrPhoneConfig{
	//						Enabled: true,
	//					},
	//
	//					//Override: &plessmodels.OverrideStruct{
	//					//	Functions: OverRideSignIn,
	//					//	APIs:      nil,
	//					//},
	//
	//					EmailDelivery: &emaildelivery.TypeInput{
	//
	//						Service: passwordless.MakeSMTPService(emaildelivery.SMTPServiceConfig{
	//
	//							Settings: emaildelivery.SMTPSettings{
	//								Host: f.Config.GetString("modules.stoken.email.host"),
	//								From: emaildelivery.SMTPFrom{
	//									Name:  "OTP",
	//									Email: email,
	//								},
	//								Port: 465,
	//								//Username: &smtpUsername, // this is optional. In case not given, from.email will be used
	//								Username: &email,
	//								Password: f.Config.GetString("modules.stoken.email.password"),
	//								Secure:   false,
	//							},
	//						}),
	//					},
	//				},
	//			),
	//		},
	//	},
	//)
	//
	//if err != nil {
	//	return &FxTokenClient{STokenInit: false}, err
	//}
	//
	//return &FxTokenClient{STokenInit: true}, nil
	
}
