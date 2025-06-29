package authModel

type LoginReq struct {
	Identifier string `valid:"required~Vui lòng nhập email hoặc số điện thoại,email_phone(vi)~Email hoặc số điện thoại không đúng định dạng"`
	Password   string `valid:"required~Vui lòng nhập mật khẩu"`
}