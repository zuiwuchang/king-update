package configure

const (
	// UserModeRoot 管理員 權限
	UserModeRoot = 1
	// UserModeUser 一般用戶 權限 (可以 瀏覽軟件包 下載 更新 軟件)
	UserModeUser = 2
)

// User ...
type User struct {
	// 用戶名
	//
	// 如果 Name 重複了 則 使用 最先出現的 配置
	Name string
	// 用戶密碼
	Pwd string
	// 用戶 權限
	Mode int
}
