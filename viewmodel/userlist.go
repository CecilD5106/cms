package viewmodel
// CMS is the basic information on the page
type CMS struct {
	Title      string
	Active     string
	Users 	   []User
}

// User is the basic information in the user table
type User struct {
	ID              string `json:"user_id"`
	UserName        string `json:"user_name"`
	UserEmail       string `json:"user_email"`
	FName           string `json:"user_first_name"`
	LName           string `json:"user_last_name"`
	Password        string `json:"password"`
	PasswordChange  string `json:"password_change"`
	PasswordExpired string `json:"password_expired"`
	LastLogon       string `json:"last_logon"`
	AccountLocked   string `json:"account_locked"`
}

// NewUserList populates data for the User List page
func NewUserList() CMS {
	result := CMS{
		Title:  "AR CMS",
		Active: "config",
	}
	return result
}
