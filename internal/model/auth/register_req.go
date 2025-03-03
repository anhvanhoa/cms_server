package auth

type RegisterReq struct {
	Email           string `valid:"required~Vui lòng nhập email,email~Email không đúng định dạng"`
	FullName        string `valid:"required~Vui lòng nhập họ và tên"`
	Password        string `valid:"required~Vui lòng nhập mật khẩu,minstringlength(6)~Mật khẩu phải có ít nhất 6 ký tự"`
	ConfirmPassword string `valid:"required~Vui lòng xác nhận mật khẩu"`
}
