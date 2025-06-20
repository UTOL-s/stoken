package token

import (
	"github.com/UTOL-s/stoken/pkg/config"
	"github.com/supertokens/supertokens-golang/ingredients/emaildelivery"
	"github.com/supertokens/supertokens-golang/recipe/passwordless"
	"github.com/supertokens/supertokens-golang/recipe/passwordless/plessmodels"
	"github.com/supertokens/supertokens-golang/supertokens"
)

// TokenService handles token operations
type TokenService struct {
	config config.Config
}

// NewTokenService creates a new TokenService with the provided configuration
func NewTokenService(cfg config.Config) *TokenService {
	return &TokenService{
		config: cfg,
	}
}

func (t *TokenService) TokenInit() error {
	
	apiBasePath := "/api/auth"
	webBasePath := "/api/auth"
	
	err := supertokens.Init(
		
		supertokens.TypeInput{
			
			Supertokens: &supertokens.ConnectionInfo{
				ConnectionURI: t.config.SuperTokenUrl,
				APIKey:        t.config.SuperTokenApiKey,
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
									Host: "smtp.mailgun.org",
									From: emaildelivery.SMTPFrom{
										Name:  "OTP",
										Email: "&t.config.FromEmail",
									},
									Port: 465,
									//Username: &smtpUsername, // this is optional. In case not given, from.email will be used
									Username: &t.config.FromEmail,
									Password: t.config.Password,
									Secure:   false,
								},
							}),
						},
					},
				),
			},
		},
	)
	
	return err
}
