package stoken

import (
	"github.com/UTOL-s/stoken/internal/token"
	"github.com/ankorstore/yokai/config"
	"github.com/supertokens/supertokens-golang/ingredients/emaildelivery"
	"github.com/supertokens/supertokens-golang/recipe/passwordless"
	"github.com/supertokens/supertokens-golang/recipe/passwordless/plessmodels"
	"github.com/supertokens/supertokens-golang/supertokens"
	"go.uber.org/fx"
)

const ModuleName = "stoken"

var FxSTokenModule = fx.Module(
	ModuleName,
	fx.Provide(
		TokenInit,
		//fx.Annotate(token.NewDefaultTokenClientFactory, fx.As(new(token.TokenClientFactory))),
		//NewTokenClientInit,
	
	),
)

type FxTokenClientParam struct {
	fx.In
	Lifecycle fx.Lifecycle
	Config    *config.Config
	Factory   token.TokenClientFactory
}

func TokenInit(f FxTokenClientParam) error {
	
	email := f.Config.GetString("modules.stoken.email.username")
	
	apiBasePath := "/api/auth"
	webBasePath := "/api/auth"
	
	err := supertokens.Init(
		
		supertokens.TypeInput{
			
			Supertokens: &supertokens.ConnectionInfo{
				ConnectionURI: f.Config.GetString("modules.stoken.apiUrl"),
				APIKey:        f.Config.GetString("modules.stoken.apiKey"),
			},
			
			AppInfo: supertokens.AppInfo{
				AppName:       "UTOL",
				WebsiteDomain: "http://localhost:3000",
				APIDomain:     "http://localhost:8080",
				//APIDomain:       config.AppUrl(),
				APIBasePath:     &apiBasePath,
				WebsiteBasePath: &webBasePath,
				APIGatewayPath:  nil,
			},
			
			RecipeList: []supertokens.Recipe{
				passwordless.Init(
					plessmodels.TypeInput{
						
						FlowType: "USER_INPUT_CODE_AND_MAGIC_LINK",
						ContactMethodEmailOrPhone: plessmodels.ContactMethodEmailOrPhoneConfig{
							Enabled: true,
						},
						
						//Override: &plessmodels.OverrideStruct{
						//	Functions: OverRideSignIn,
						//	APIs:      nil,
						//},
						
						EmailDelivery: &emaildelivery.TypeInput{
							
							Service: passwordless.MakeSMTPService(emaildelivery.SMTPServiceConfig{
								
								Settings: emaildelivery.SMTPSettings{
									Host: f.Config.GetString("modules.stoken.email.host"),
									From: emaildelivery.SMTPFrom{
										Name:  "OTP",
										Email: email,
									},
									Port: 465,
									//Username: &smtpUsername, // this is optional. In case not given, from.email will be used
									Username: &email,
									Password: f.Config.GetString("modules.stoken.email.password"),
									Secure:   false,
								},
							}),
						},
					},
				),
			},
		},
	)
	
	if err != nil {
		return err
	}
	
	return nil
}
