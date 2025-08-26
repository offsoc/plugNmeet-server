package authservice

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"encoding/xml"
	"fmt"
	"io"
	"net/url"
	"strings"
	"time"

	"github.com/mynaparrot/plugnmeet-server/pkg/config"
	"gopkg.in/square/go-jose.v2/jwt"

	"github.com/go-ldap/ldap/v3"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// AuthMethod 身份验证方法类型
type AuthMethod string

const (
	APIKeyAuth   AuthMethod = "api_key"
	OAuth2Auth   AuthMethod = "oauth2"
	SAMLAuth     AuthMethod = "saml"
	LDAPAuth     AuthMethod = "ldap"
	JWTTokenAuth AuthMethod = "jwt_token"
)

// MultiAuthService 多因素身份验证服务
type MultiAuthService struct {
	app *config.AppConfig
}

// OAuth2Config OAuth2配置
type OAuth2Config struct {
	ClientID     string   `yaml:"client_id"`
	ClientSecret string   `yaml:"client_secret"`
	RedirectURL  string   `yaml:"redirect_url"`
	Scopes       []string `yaml:"scopes"`
	AuthURL      string   `yaml:"auth_url"`
	TokenURL     string   `yaml:"token_url"`
}

// SAMLConfig SAML配置
type SAMLConfig struct {
	IDPIssuerURL                string `yaml:"idp_issuer_url"`
	IDPCert                     string `yaml:"idp_cert"`
	SPCert                      string `yaml:"sp_cert"`
	SPKey                       string `yaml:"sp_key"`
	AssertionConsumerServiceURL string `yaml:"acs_url"`
}

// LDAPConfig LDAP配置
type LDAPConfig struct {
	Server   string `yaml:"server"`
	Port     int    `yaml:"port"`
	BaseDN   string `yaml:"base_dn"`
	BindDN   string `yaml:"bind_dn"`
	BindPass string `yaml:"bind_pass"`
}

// JWTConfig JWT配置
type JWTConfig struct {
	Issuer   string `yaml:"issuer"`
	Audience string `yaml:"audience"`
	Key      string `yaml:"key"`
}

// AuthConfig 身份验证配置
type AuthConfig struct {
	Method AuthMethod   `yaml:"method"`
	OAuth2 OAuth2Config `yaml:"oauth2,omitempty"`
	SAML   SAMLConfig   `yaml:"saml,omitempty"`
	LDAP   LDAPConfig   `yaml:"ldap,omitempty"`
	JWT    JWTConfig    `yaml:"jwt,omitempty"`
}

// AuthResult 身份验证结果
type AuthResult struct {
	UserID      string            `json:"user_id"`
	UserName    string            `json:"user_name"`
	Email       string            `json:"email"`
	Attributes  map[string]string `json:"attributes"`
	Valid       bool              `json:"valid"`
	AccessToken string            `json:"access_token,omitempty"`
}

// NewMultiAuthService 创建新的多因素身份验证服务
func NewMultiAuthService(app *config.AppConfig) *MultiAuthService {
	return &MultiAuthService{
		app: app,
	}
}

// Authenticate 执行身份验证
func (s *MultiAuthService) Authenticate(ctx context.Context, config AuthConfig, credentials map[string]string) (*AuthResult, error) {
	switch config.Method {
	case OAuth2Auth:
		return s.authenticateOAuth2(ctx, config.OAuth2, credentials)
	case SAMLAuth:
		return s.authenticateSAML(ctx, config.SAML, credentials)
	case LDAPAuth:
		return s.authenticateLDAP(ctx, config.LDAP, credentials)
	case JWTTokenAuth:
		return s.authenticateJWT(ctx, config.JWT, credentials)
	default:
		return nil, fmt.Errorf("unsupported authentication method: %s", config.Method)
	}
}

// authenticateOAuth2 OAuth2身份验证
func (s *MultiAuthService) authenticateOAuth2(ctx context.Context, config OAuth2Config, credentials map[string]string) (*AuthResult, error) {
	// 创建OAuth2配置
	oauthConfig := &oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		RedirectURL:  config.RedirectURL,
		Scopes:       config.Scopes,
	}

	// 根据提供商设置端点
	if strings.Contains(config.AuthURL, "google") {
		oauthConfig.Endpoint = google.Endpoint
	} else {
		oauthConfig.Endpoint = oauth2.Endpoint{
			AuthURL:  config.AuthURL,
			TokenURL: config.TokenURL,
		}
	}

	// 获取授权码
	code, ok := credentials["code"]
	if !ok {
		return nil, fmt.Errorf("missing authorization code")
	}

	// 交换授权码获取令牌
	token, err := oauthConfig.Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange token: %w", err)
	}

	// 获取用户信息
	client := oauthConfig.Client(ctx, token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %w", err)
	}
	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read user info: %w", err)
	}

	// 解析用户信息（这里简化处理，实际应解析JSON）
	// 在实际实现中，应解析JSON并提取用户信息

	result := &AuthResult{
		UserID:      "", // 从用户信息中提取
		UserName:    "", // 从用户信息中提取
		Email:       "", // 从用户信息中提取
		Valid:       true,
		AccessToken: token.AccessToken,
	}

	return result, nil
}

// authenticateSAML SAML身份验证
func (s *MultiAuthService) authenticateSAML(ctx context.Context, config SAMLConfig, credentials map[string]string) (*AuthResult, error) {
	// 解析SAML响应
	samlResponse, ok := credentials["SAMLResponse"]
	if !ok {
		return nil, fmt.Errorf("missing SAML response")
	}

	// 解码SAML响应
	decoded, err := base64.StdEncoding.DecodeString(samlResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to decode SAML response: %w", err)
	}

	// 解析XML
	var response struct {
		Assertion struct {
			Subject struct {
				NameID struct {
					Value string `xml:",chardata"`
				} `xml:"NameID"`
			} `xml:"Subject"`
			AttributeStatement struct {
				Attributes []struct {
					Name   string `xml:"Name,attr"`
					Values []struct {
						Value string `xml:",chardata"`
					} `xml:"AttributeValue"`
				} `xml:"Attribute"`
			} `xml:"AttributeStatement"`
		} `xml:"Assertion"`
	}

	if err := xml.Unmarshal(decoded, &response); err != nil {
		return nil, fmt.Errorf("failed to parse SAML response: %w", err)
	}

	// 提取用户信息
	result := &AuthResult{
		UserID:     response.Assertion.Subject.NameID.Value,
		UserName:   response.Assertion.Subject.NameID.Value,
		Email:      "", // 从属性中提取
		Valid:      true,
		Attributes: make(map[string]string),
	}

	// 提取属性
	for _, attr := range response.Assertion.AttributeStatement.Attributes {
		if len(attr.Values) > 0 {
			result.Attributes[attr.Name] = attr.Values[0].Value
			if attr.Name == "email" || attr.Name == "mail" {
				result.Email = attr.Values[0].Value
			}
		}
	}

	return result, nil
}

// authenticateLDAP LDAP身份验证
func (s *MultiAuthService) authenticateLDAP(ctx context.Context, config LDAPConfig, credentials map[string]string) (*AuthResult, error) {
	// 连接到LDAP服务器
	conn, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", config.Server, config.Port))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to LDAP server: %w", err)
	}
	defer conn.Close()

	// 如果配置了BindDN，则使用它进行绑定
	if config.BindDN != "" {
		err = conn.Bind(config.BindDN, config.BindPass)
		if err != nil {
			return nil, fmt.Errorf("failed to bind with BindDN: %w", err)
		}
	}

	// 构造用户DN
	username, ok := credentials["username"]
	if !ok {
		return nil, fmt.Errorf("missing username")
	}
	password, ok := credentials["password"]
	if !ok {
		return nil, fmt.Errorf("missing password")
	}

	userDN := fmt.Sprintf("uid=%s,%s", username, config.BaseDN)
	if config.BindDN == "" {
		userDN = fmt.Sprintf("%s=%s,%s", "cn", username, config.BaseDN)
	}

	// 尝试绑定用户
	err = conn.Bind(userDN, password)
	if err != nil {
		return nil, fmt.Errorf("failed to bind user: %w", err)
	}

	// 搜索用户信息
	searchRequest := ldap.NewSearchRequest(
		config.BaseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=person)(uid=%s))", username),
		[]string{"dn", "cn", "mail"},
		nil,
	)

	sr, err := conn.Search(searchRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to search user: %w", err)
	}

	if len(sr.Entries) == 0 {
		return nil, fmt.Errorf("user not found")
	}

	entry := sr.Entries[0]

	result := &AuthResult{
		UserID:     entry.DN,
		UserName:   entry.GetAttributeValue("cn"),
		Email:      entry.GetAttributeValue("mail"),
		Valid:      true,
		Attributes: make(map[string]string),
	}

	// 提取所有属性
	for _, attr := range entry.Attributes {
		result.Attributes[attr.Name] = attr.Values[0]
	}

	return result, nil
}

// authenticateJWT JWT令牌身份验证
func (s *MultiAuthService) authenticateJWT(ctx context.Context, config JWTConfig, credentials map[string]string) (*AuthResult, error) {
	// 获取JWT令牌
	tokenString, ok := credentials["token"]
	if !ok {
		return nil, fmt.Errorf("missing JWT token")
	}

	// 解析PEM格式的公钥
	block, _ := pem.Decode([]byte(config.Key))
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block containing the public key")
	}

	// 解析公钥
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}

	rsaPub, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("not an RSA public key")
	}

	// 解析JWT令牌
	token, err := jwt.ParseSigned(tokenString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JWT token: %w", err)
	}

	// 验证并解码声明
	claims := jwt.Claims{}
	customClaims := map[string]interface{}{}

	err = token.Claims(rsaPub, &claims, &customClaims)
	if err != nil {
		return nil, fmt.Errorf("failed to verify JWT token: %w", err)
	}

	// 验证声明
	if err := claims.Validate(jwt.Expected{
		Issuer:   config.Issuer,
		Audience: jwt.Audience{config.Audience},
		Time:     time.Now(),
	}); err != nil {
		return nil, fmt.Errorf("JWT token validation failed: %w", err)
	}

	// 构造结果
	result := &AuthResult{
		UserID:     claims.Subject,
		UserName:   "", // 从自定义声明中提取
		Email:      "", // 从自定义声明中提取
		Valid:      true,
		Attributes: make(map[string]string),
	}

	// 提取自定义声明
	for key, value := range customClaims {
		if str, ok := value.(string); ok {
			result.Attributes[key] = str
			switch key {
			case "name", "username":
				result.UserName = str
			case "email":
				result.Email = str
			}
		}
	}

	return result, nil
}

// GenerateSAMLAuthURL 生成SAML认证URL
func (s *MultiAuthService) GenerateSAMLAuthURL(config SAMLConfig, relayState string) string {
	// 构造SAML认证请求
	authRequest := fmt.Sprintf(
		"saml_request=%s&relay_state=%s",
		url.QueryEscape("SAMLRequest"), // 实际应生成SAML请求
		url.QueryEscape(relayState),
	)

	return fmt.Sprintf("%s?%s", config.IDPIssuerURL, authRequest)
}
