//request yang akan di tampilkan ketika user/admin login

package payload

type Logins struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
