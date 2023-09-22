package token

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/emicklei/go-restful/v3"
	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/mssola/user_agent"
)

const (
	AppName = "token"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

func NewPasswordIssueTokenRequest(username, password string) *IssueTokenRequest {
	req := NewIssueTokenRequest()
	req.GrantType = GRANT_TYPE_PASSWORD
	req.Username = username
	req.Password = password
	return req
}

func NewPrivateTokenIssueTokenRequest(accessToken, description string) *IssueTokenRequest {
	req := NewIssueTokenRequest()
	req.GrantType = GRANT_TYPE_PRIVATE_TOKEN
	req.AccessToken = accessToken
	req.Description = description
	return req
}

func NewRefreshIssueTokenRequest(accessToken, refreshToken string) *IssueTokenRequest {
	req := NewIssueTokenRequest()
	req.GrantType = GRANT_TYPE_REFRESH
	req.AccessToken = accessToken
	req.RefreshToken = refreshToken
	return req
}

func NewLDAPIssueTokenRequest(username, password string) *IssueTokenRequest {
	req := NewIssueTokenRequest()
	req.GrantType = GRANT_TYPE_LDAP
	req.Username = username
	req.Password = password
	return req
}

func NewFeishuAuthCodeIssueTokenRequest(authcode, state, domain string) *IssueTokenRequest {
	req := NewIssueTokenRequest()
	req.AuthCode = authcode
	req.State = state
	req.GrantType = GRANT_TYPE_FEISHU
	req.Username = fmt.Sprintf("@%s", domain)
	return req
}

func NewDingDingAuthCodeIssueTokenRequest(authcode, state, domain string) *IssueTokenRequest {
	req := NewIssueTokenRequest()
	req.AuthCode = authcode
	req.State = state
	req.GrantType = GRANT_TYPE_DINGDING
	req.Username = fmt.Sprintf("@%s", domain)
	return req
}

func NewWechatAuthCodeIssueTokenRequest(authcode, state, domain string) *IssueTokenRequest {
	req := NewIssueTokenRequest()
	req.AuthCode = authcode
	req.State = state
	req.GrantType = GRANT_TYPE_WECHAT_WORK
	req.Username = fmt.Sprintf("@%s", domain)
	return req
}

// NewIssueTokenRequest 默认请求
func NewIssueTokenRequest() *IssueTokenRequest {
	return &IssueTokenRequest{
		ExpiredAt: DEFAULT_ACCESS_TOKEN_EXPIRE_SECOND,
	}
}

// AbnormalUserCheckKey todo
func (req *IssueTokenRequest) AbnormalUserCheckKey() string {
	key := ""
	switch req.GrantType {
	case GRANT_TYPE_PASSWORD, GRANT_TYPE_LDAP:
		key = req.Username
	case GRANT_TYPE_PRIVATE_TOKEN:
		key = req.AccessToken
	case GRANT_TYPE_REFRESH:
		key = req.RefreshToken
	case GRANT_TYPE_AUTH_CODE:
		key = req.AuthCode
	}
	return "abnormal_" + key
}

// NewRevolkTokenRequest 撤销Token请求
func NewRevolkTokenRequest(accessToken, refreshToken string) *RevolkTokenRequest {
	return &RevolkTokenRequest{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

func NewChangeNamespaceRequest() *ChangeNamespaceRequest {
	return &ChangeNamespaceRequest{}
}

func NewDefaultToken() *Token {
	return &Token{
		Status:   NewStatus(),
		Location: NewLocation(),
	}
}

func NewToken(req *IssueTokenRequest) *Token {
	tk := &Token{
		AccessToken:      MakeBearer(24),
		RefreshToken:     MakeBearer(32),
		IssueAt:          time.Now().Unix(),
		AccessExpiredAt:  req.ExpiredAt,
		RefreshExpiredAt: req.ExpiredAt * 4,
		GrantType:        req.GrantType,
		Type:             req.Type,
		Description:      req.Description,
		Status:           NewStatus(),
		Location:         req.Location,
		Meta:             map[string]string{},
	}
	switch req.GrantType {
	case GRANT_TYPE_PRIVATE_TOKEN:
		tk.Platform = PLATFORM_API
	default:
		tk.Platform = PLATFORM_WEB
	}
	return tk
}

func NewStatus() *Status {
	return &Status{
		IsBlock: false,
	}
}

func NewNewLocationFromHttp(r *http.Request) *Location {
	l := NewLocation()

	// 解析UserAgent
	ua := r.UserAgent()
	if ua != "" {
		ua := user_agent.New(ua)
		l.UserAgent = &UserAgent{
			Os:       ua.OS(),
			Platform: ua.Platform(),
		}
		l.UserAgent.EngineName, l.UserAgent.EngineVersion = ua.Engine()
		l.UserAgent.BrowserName, l.UserAgent.BrowserVersion = ua.Browser()
	}

	// 解析地理位置
	rip := request.GetRemoteIP(r)
	l.IpLocation.RemoteIp = rip

	return l
}

func NewLocation() *Location {
	return &Location{
		IpLocation: &IPLocation{},
		UserAgent:  &UserAgent{},
	}
}

func NewPlatform(p PLATFORM) *PLATFORM {
	return &p
}

// 查询用户最近一次登陆记录
func NewQueryUserWebLastToken(uid string) *QueryTokenRequest {
	return &QueryTokenRequest{
		Page:     request.NewPageRequest(1, 1),
		Platform: NewPlatform(PLATFORM_WEB),
		UserId:   uid,
	}
}

func (s *TokenSet) Length() int {
	return len(s.Items)
}

// 优先从认证头中获取, 如果头没有从Cookie中获取
func GetAccessTokenFromHTTP(r *http.Request) string {
	auth := r.Header.Get(ACCESS_TOKEN_HEADER_KEY)
	info := strings.Split(auth, " ")
	if len(info) > 1 {
		return info[1]
	}

	ck, err := r.Cookie(ACCESS_TOKEN_COOKIE_KEY)
	if err != nil {
		zap.L().Warnf("get tk from cookie: %s error, %s", ACCESS_TOKEN_COOKIE_KEY, err)
		return r.URL.Query().Get(ACCESS_TOKEN_QUERY_KEY)
	}

	return ck.Value
}

func NewValidateTokenRequest(accessToken string) *ValidateTokenRequest {
	return &ValidateTokenRequest{
		AccessToken: accessToken,
	}
}

// MakeBearer https://tools.ietf.org/html/rfc6750#section-2.1
// b64token    = 1*( ALPHA / DIGIT /"-" / "." / "_" / "~" / "+" / "/" ) *"="
func MakeBearer(lenth int) string {
	charlist := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	t := make([]string, lenth)
	r := rand.New(rand.NewSource(time.Now().UnixNano() + int64(lenth) + rand.Int63n(10000)))
	for i := 0; i < lenth; i++ {
		rn := r.Intn(len(charlist))
		w := charlist[rn : rn+1]
		t = append(t, w)
	}

	token := strings.Join(t, "")
	return token
}

// Validate 校验参数
func (m *ValidateTokenRequest) Validate() error {
	if err := validate.Struct(m); err != nil {
		return err
	}

	if m.AccessToken == "" {
		return errors.New("access_token required one")
	}

	return nil
}

// BlockMessage todo
func (t *Status) BlockMessage() string {
	if !t.IsBlock {
		return ""
	}

	return fmt.Sprintf("token blocked at %d, reason: %s", t.BlockAt, t.BlockReason)
}

// 基于令牌创建HTTP Cookie 用于Web登陆场景
func (c *Token) SetCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     ACCESS_TOKEN_COOKIE_KEY,
		Value:    url.QueryEscape(c.AccessToken),
		MaxAge:   0,
		Path:     "/",
		Domain:   "",
		SameSite: http.SameSiteDefaultMode,
		Secure:   false,
		HttpOnly: true,
	})
}

// CheckAccessIsExpired 检测token是否过期
func (t *Token) CheckAccessIsExpired() bool {
	if t.AccessExpiredAt == 0 {
		return false
	}

	return t.AccessExpiredTime().Before(time.Now())
}

func (t *Token) Desense() *Token {
	t.RefreshToken = ""
	return t
}

func (t *Token) AccessExpiredTime() time.Time {
	return time.Unix(t.IssueAt, 0).Add(time.Duration(t.AccessExpiredAt) * time.Second)
}

// CheckRefreshIsExpired 检测刷新token是否过期
func (t *Token) CheckRefreshIsExpired() bool {
	// 过期时间为0时, 标识不过期
	if t.RefreshExpiredAt == 0 {
		return false
	}

	return t.RefreshExpiredTime().Before(time.Now())
}

func (t *Token) RefreshExpiredTime() time.Time {
	return time.Unix(t.IssueAt, 0).Add(time.Duration(t.RefreshExpiredAt) * time.Second)
}

func (t *Token) Json() string {
	dj, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	return string(dj)
}

func NewDescribeTokenRequest(token string) *DescribeTokenRequest {
	return &DescribeTokenRequest{
		DescribeBy:    DESCRIBY_BY_ACCESS_TOKEN,
		DescribeValue: token,
	}
}

func (req *ChangeNamespaceRequest) Validate() error {
	return validate.Struct(req)
}

// Validate 校验
func (m *DescribeTokenRequest) Validate() error {
	if err := validate.Struct(m); err != nil {
		return err
	}

	return nil
}

// NewTokenSet 实例化
func NewTokenSet() *TokenSet {
	return &TokenSet{
		Items: []*Token{},
	}
}

// Add todo
func (m *TokenSet) Add(item *Token) {
	m.Items = append(m.Items, item)
}

func (m *TokenSet) JsonFormat() string {
	dj, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	return string(dj)
}

func NewQueryTokenRequestFromHttp(r *restful.Request) *QueryTokenRequest {
	req := NewQueryTokenRequest()
	req.Page = request.NewPageRequestFromHTTP(r.Request)
	return req
}

func NewQueryTokenRequest() *QueryTokenRequest {
	return &QueryTokenRequest{
		Page: request.NewDefaultPageRequest(),
	}
}
