package auth

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
func (p User) AccessToken() string {
	return p.accessToken
}

// CustomerCode 顧客コードを返却します
func (p User) CustomerCode() string {
	return p.customerCode
}

// LoginAddress メールアドレスを返却します
func (p User) LoginAddress() string {
	return p.loginAddress
}

// Client ユーザー固有の職別値を返却します
func (p User) Client() string {
	return p.client
}

// UserId ユーザーIDを返却します
func (p User) UserId() uint32 {
	return p.userId
}

// RoleType ユーザーのロール種別を返却します
func (p User) RoleType() uint8 {
	return p.roleType
}
