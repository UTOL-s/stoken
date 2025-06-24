package token

import (
	"github.com/ankorstore/yokai/config"
	"github.com/supertokens/supertokens-golang/ingredients/emaildelivery"
	"github.com/supertokens/supertokens-golang/recipe/passwordless"
	"github.com/supertokens/supertokens-golang/recipe/passwordless/plessmodels"
	"github.com/supertokens/supertokens-golang/supertokens"
)

var _ TokenClientFactory = (*DefaultTokenClientFactory)(nil)

type TokenClientFactory interface {
	TokenInit() error
}

// DefaultTokenClientFactory handles token operations
type DefaultTokenClientFactory struct {
	config *config.Config
}

// NewDefaultTokenClientFactory creates a new TokenService with the provided configuration
func NewDefaultTokenClientFactory(cfg *config.Config) *DefaultTokenClientFactory {
	return &DefaultTokenClientFactory{
		config: cfg,
	}
}

func (f *DefaultTokenClientFactory) TokenInit() error {
	
	email := f.config.GetString("modules.stoken.email.username")
	
	apiBasePath := "/api/auth"
	webBasePath := "/api/auth"
	
	err := supertokens.Init(
		
		supertokens.TypeInput{
			
			Supertokens: &supertokens.ConnectionInfo{
				ConnectionURI: f.config.GetString("modules.stoken.apiUrl"),
				APIKey:        f.config.GetString("modules.stoken.apiKey"),
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
									Host: f.config.GetString("modules.stoken.email.host"),
									From: emaildelivery.SMTPFrom{
										Name:  "OTP",
										Email: email,
									},
									Port: 465,
									//Username: &smtpUsername, // this is optional. In case not given, from.email will be used
									Username: &email,
									Password: f.config.GetString("modules.stoken.email.password"),
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
