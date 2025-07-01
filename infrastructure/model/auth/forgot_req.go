package authModel

import authUC "cms-server/domain/usecase/auth"

type ForgotPasswordReq struct {
	Email string                    `valid:"required~Vui lòng nhập email, email~Email không đúng định dạng"`
	Type  authUC.ForgotPasswordType `valid:"in(ForgotByCode|ForgotByToken)~Phương thức xác thực không hợp lệ"`
}
