package authModel

type ResetPasswordByTokenRequest struct {
	Token           string `valid:"required~Vui lòng nhập mã xác thực"`
	Password        string `valid:"required~Vui lòng nhập mật khẩu, minstringlength(6)~Mật khẩu phải có ít nhất 6 ký tự"`
	ConfirmPassword string `valid:"required~Vui lòng nhập xác nhận khẩu mật mới, minstringlength(6)~Xác nhận khẩu mật khẩu mới phải có ít nhất 6 ký tự"`
}

type ResetPasswordByCodeRequest struct {
	Code            string `valid:"required~Vui lòng nhập mã xác thực"`
	Email           string `valid:"required~Vui lòng nhập email, email~Email không đúng định dạng"`
	Password        string `valid:"required~Vui lòng nhập mật khẩu, minstringlength(6)~Mật khẩu phải có ít nhất 6 ký tự"`
	ConfirmPassword string `valid:"required~Vui lòng nhập xác nhận khẩu mật khẩu mới, minstringlength(6)~Xác nhận khẩu mật khẩu mới phải có ít nhất 6 ký tự"`
}

type CheckCodeReq struct {
	Code  string `valid:"required~Vui lòng nhập mã xác thực"`
	Email string `valid:"required~Vui lòng nhập email, email~Email không đúng định dạng"`
}
