package controllers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/mynaparrot/plugnmeet-protocol/utils"
	"github.com/mynaparrot/plugnmeet-server/pkg/config"
	authservice "github.com/mynaparrot/plugnmeet-server/pkg/services/auth"
)

// MultiAuthController 多因素身份验证控制器
type MultiAuthController struct {
	app         *config.AppConfig
	authService *authservice.MultiAuthService
}

// NewMultiAuthController 创建新的多因素身份验证控制器
func NewMultiAuthController(app *config.AppConfig) *MultiAuthController {
	return &MultiAuthController{
		app:         app,
		authService: authservice.NewMultiAuthService(app),
	}
}

// HandleOAuth2Callback 处理OAuth2回调
func (c *MultiAuthController) HandleOAuth2Callback(provider string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// 检查OAuth2是否已启用
		if c.app.Client.AuthMethods == nil || c.app.Client.AuthMethods.OAuth2 == nil {
			return utils.SendCommonProtoJsonResponse(ctx, false, "OAuth2 authentication is not configured")
		}

		var oauthConfig *config.OAuth2Config
		switch provider {
		case "google":
			if c.app.Client.AuthMethods.OAuth2.Google == nil || !c.app.Client.AuthMethods.OAuth2.Google.Enabled {
				return utils.SendCommonProtoJsonResponse(ctx, false, "Google OAuth2 authentication is not enabled")
			}
			oauthConfig = c.app.Client.AuthMethods.OAuth2.Google
		case "microsoft":
			if c.app.Client.AuthMethods.OAuth2.Microsoft == nil || !c.app.Client.AuthMethods.OAuth2.Microsoft.Enabled {
				return utils.SendCommonProtoJsonResponse(ctx, false, "Microsoft OAuth2 authentication is not enabled")
			}
			oauthConfig = c.app.Client.AuthMethods.OAuth2.Microsoft
		case "github":
			if c.app.Client.AuthMethods.OAuth2.GitHub == nil || !c.app.Client.AuthMethods.OAuth2.GitHub.Enabled {
				return utils.SendCommonProtoJsonResponse(ctx, false, "GitHub OAuth2 authentication is not enabled")
			}
			oauthConfig = c.app.Client.AuthMethods.OAuth2.GitHub
		default:
			return utils.SendCommonProtoJsonResponse(ctx, false, "Unsupported OAuth2 provider")
		}

		// 构造认证配置
		authConfig := authservice.AuthConfig{
			Method: authservice.OAuth2Auth,
			OAuth2: authservice.OAuth2Config{
				ClientID:     oauthConfig.ClientID,
				ClientSecret: oauthConfig.ClientSecret,
				RedirectURL:  oauthConfig.RedirectURL,
				Scopes:       oauthConfig.Scopes,
				AuthURL:      oauthConfig.AuthURL,
				TokenURL:     oauthConfig.TokenURL,
			},
		}

		// 获取授权码
		code := ctx.Query("code")
		if code == "" {
			return utils.SendCommonProtoJsonResponse(ctx, false, "Missing authorization code")
		}

		// 执行身份验证
		credentials := map[string]string{
			"code": code,
		}

		result, err := c.authService.Authenticate(context.Background(), authConfig, credentials)
		if err != nil {
			return utils.SendCommonProtoJsonResponse(ctx, false, err.Error())
		}

		if !result.Valid {
			return utils.SendCommonProtoJsonResponse(ctx, false, "Authentication failed")
		}

		// 生成plugNmeet访问令牌
		// 这里应该调用现有的令牌生成逻辑
		// 为简化起见，返回用户信息

		return ctx.JSON(fiber.Map{
			"status": true,
			"msg":    "Authentication successful",
			"data": fiber.Map{
				"user_id":      result.UserID,
				"user_name":    result.UserName,
				"email":        result.Email,
				"access_token": result.AccessToken,
			},
		})
	}
}

// HandleSAMLCallback 处理SAML回调
func (c *MultiAuthController) HandleSAMLCallback(ctx *fiber.Ctx) error {
	// 检查SAML是否已启用
	if c.app.Client.AuthMethods == nil || c.app.Client.AuthMethods.SAML == nil || !c.app.Client.AuthMethods.SAML.Enabled {
		return utils.SendCommonProtoJsonResponse(ctx, false, "SAML authentication is not enabled")
	}

	// 构造认证配置
	authConfig := authservice.AuthConfig{
		Method: authservice.SAMLAuth,
		SAML: authservice.SAMLConfig{
			IDPIssuerURL:                c.app.Client.AuthMethods.SAML.IDPIssuerURL,
			IDPCert:                     c.app.Client.AuthMethods.SAML.IDPCert,
			SPCert:                      c.app.Client.AuthMethods.SAML.SPCert,
			SPKey:                       c.app.Client.AuthMethods.SAML.SPKey,
			AssertionConsumerServiceURL: c.app.Client.AuthMethods.SAML.AssertionConsumerServiceURL,
		},
	}

	// 获取SAML响应
	samlResponse := ctx.FormValue("SAMLResponse")
	if samlResponse == "" {
		return utils.SendCommonProtoJsonResponse(ctx, false, "Missing SAML response")
	}

	// 执行身份验证
	credentials := map[string]string{
		"SAMLResponse": samlResponse,
	}

	result, err := c.authService.Authenticate(context.Background(), authConfig, credentials)
	if err != nil {
		return utils.SendCommonProtoJsonResponse(ctx, false, err.Error())
	}

	if !result.Valid {
		return utils.SendCommonProtoJsonResponse(ctx, false, "Authentication failed")
	}

	// 生成plugNmeet访问令牌
	// 这里应该调用现有的令牌生成逻辑
	// 为简化起见，返回用户信息

	return ctx.JSON(fiber.Map{
		"status": true,
		"msg":    "Authentication successful",
		"data": fiber.Map{
			"user_id":   result.UserID,
			"user_name": result.UserName,
			"email":     result.Email,
		},
	})
}

// HandleLDAPLogin 处理LDAP登录
func (c *MultiAuthController) HandleLDAPLogin(ctx *fiber.Ctx) error {
	// 检查LDAP是否已启用
	if c.app.Client.AuthMethods == nil || c.app.Client.AuthMethods.LDAP == nil || !c.app.Client.AuthMethods.LDAP.Enabled {
		return utils.SendCommonProtoJsonResponse(ctx, false, "LDAP authentication is not enabled")
	}

	// 获取用户名和密码
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")

	if username == "" || password == "" {
		return utils.SendCommonProtoJsonResponse(ctx, false, "Username and password are required")
	}

	// 构造认证配置
	authConfig := authservice.AuthConfig{
		Method: authservice.LDAPAuth,
		LDAP: authservice.LDAPConfig{
			Server:   c.app.Client.AuthMethods.LDAP.Server,
			Port:     c.app.Client.AuthMethods.LDAP.Port,
			BaseDN:   c.app.Client.AuthMethods.LDAP.BaseDN,
			BindDN:   c.app.Client.AuthMethods.LDAP.BindDN,
			BindPass: c.app.Client.AuthMethods.LDAP.BindPass,
		},
	}

	// 执行身份验证
	credentials := map[string]string{
		"username": username,
		"password": password,
	}

	result, err := c.authService.Authenticate(context.Background(), authConfig, credentials)
	if err != nil {
		return utils.SendCommonProtoJsonResponse(ctx, false, err.Error())
	}

	if !result.Valid {
		return utils.SendCommonProtoJsonResponse(ctx, false, "Authentication failed")
	}

	// 生成plugNmeet访问令牌
	// 这里应该调用现有的令牌生成逻辑
	// 为简化起见，返回用户信息

	return ctx.JSON(fiber.Map{
		"status": true,
		"msg":    "Authentication successful",
		"data": fiber.Map{
			"user_id":   result.UserID,
			"user_name": result.UserName,
			"email":     result.Email,
		},
	})
}

// HandleJWTAuth 处理JWT令牌认证
func (c *MultiAuthController) HandleJWTAuth(ctx *fiber.Ctx) error {
	// 检查JWT是否已启用
	if c.app.Client.AuthMethods == nil || c.app.Client.AuthMethods.JWT == nil || !c.app.Client.AuthMethods.JWT.Enabled {
		return utils.SendCommonProtoJsonResponse(ctx, false, "JWT authentication is not enabled")
	}

	// 获取JWT令牌
	token := ctx.Get("Authorization")
	if token == "" {
		return utils.SendCommonProtoJsonResponse(ctx, false, "Missing JWT token")
	}

	// 移除Bearer前缀
	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}

	// 构造认证配置
	authConfig := authservice.AuthConfig{
		Method: authservice.JWTTokenAuth,
		JWT: authservice.JWTConfig{
			Issuer:   c.app.Client.AuthMethods.JWT.Issuer,
			Audience: c.app.Client.AuthMethods.JWT.Audience,
			Key:      c.app.Client.AuthMethods.JWT.Key,
		},
	}

	// 执行身份验证
	credentials := map[string]string{
		"token": token,
	}

	result, err := c.authService.Authenticate(context.Background(), authConfig, credentials)
	if err != nil {
		return utils.SendCommonProtoJsonResponse(ctx, false, err.Error())
	}

	if !result.Valid {
		return utils.SendCommonProtoJsonResponse(ctx, false, "Authentication failed")
	}

	// 生成plugNmeet访问令牌
	// 这里应该调用现有的令牌生成逻辑
	// 为简化起见，返回用户信息

	return ctx.JSON(fiber.Map{
		"status": true,
		"msg":    "Authentication successful",
		"data": fiber.Map{
			"user_id":   result.UserID,
			"user_name": result.UserName,
			"email":     result.Email,
		},
	})
}

// GenerateSAMLAuthURL 生成SAML认证URL
func (c *MultiAuthController) GenerateSAMLAuthURL(ctx *fiber.Ctx) error {
	// 检查SAML是否已启用
	if c.app.Client.AuthMethods == nil || c.app.Client.AuthMethods.SAML == nil || !c.app.Client.AuthMethods.SAML.Enabled {
		return utils.SendCommonProtoJsonResponse(ctx, false, "SAML authentication is not enabled")
	}

	// 获取relay state
	relayState := ctx.Query("relay_state", "/")

	// 生成SAML认证URL
	samlConfig := authservice.SAMLConfig{
		IDPIssuerURL: c.app.Client.AuthMethods.SAML.IDPIssuerURL,
	}

	authURL := c.authService.GenerateSAMLAuthURL(samlConfig, relayState)

	return ctx.JSON(fiber.Map{
		"status": true,
		"msg":    "SAML authentication URL generated",
		"data": fiber.Map{
			"auth_url": authURL,
		},
	})
}
