package authed

// User は認証済みのユーザー情報を表す構造体です
type User struct {
	accessToken  string
	customerCode string
	loginAddress string
	client       string
	userId       uint32
	roleType     uint8
}

func New(accessToken, customerCode, loginAddress, client string, userId uint32, roleType uint8) User {
	return User{
		accessToken:  accessToken,
		customerCode: customerCode,
		loginAddress: loginAddress,
		client:       client,
		userId:       userId,
		roleType:     roleType,
	}
}

// AccessToken 外部APIを実行するためのアクセストークンを返却します
func (u User) AccessToken() string {
	return u.accessToken
}

// CustomerCode 顧客コードを返却します
func (u User) CustomerCode() string {
	return u.customerCode
}

// LoginAddress メールアドレスを返却します
func (u User) LoginAddress() string {
	return u.loginAddress
}

// Client ユーザー固有の職別値を返却します
func (u User) Client() string {
	return u.client
}

// UserId ユーザーIDを返却します
func (u User) UserId() uint32 {
	return u.userId
}

// RoleType ユーザーのロール種別を返却します
func (u User) RoleType() uint8 {
	return u.roleType
}
