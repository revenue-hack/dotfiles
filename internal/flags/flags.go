package flags

var (
	buildVersion    = "v0.0.0"
	environmentName = "local"
)

// BuildVersion はビルド時のバイナリのバージョンを取得します
func BuildVersion() string {
	return buildVersion
}

// EnvironmentName は環境名を取得します
func EnvironmentName() string {
	return environmentName
}

// IsProduction は実行環境が本番の場合にtrueを返します
func IsProduction() bool {
	return EnvironmentName() == "prod"
}

// IsDevelopment は実行環境が開発環境の場合にtrueを返します
func IsDevelopment() bool {
	return EnvironmentName() == "local"
}
