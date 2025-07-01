package handler

import (
	authModel "cms-server/infrastructure/model/auth"
	pkgres "cms-server/infrastructure/service/response"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
)

func (rth *authHandlerImpl) ResetCode(c *fiber.Ctx) error {
	var req authModel.ResetPasswordByCodeRequest
	if err := c.BodyParser(&req); err != nil {
		err := pkgres.NewErr("Dữ liệu không hợp lệ").BadReq()
		return rth.log.Log(c, err)
	}

	if _, err := govalidator.ValidateStruct(req); err != nil {
		err := pkgres.Err(err).UnprocessableEntity()
		return rth.log.Log(c, err)
	}

	if req.Password != req.ConfirmPassword {
		err := pkgres.NewErr("Mật khẩu không khớp").BadReq()
		return rth.log.Log(c, err)
	}

	userID, err := rth.resetCodeUc.VerifySession(req.Code, req.Email)
	if err != nil {
		err = pkgres.Err(err).BadReq()
		return rth.log.Log(c, err)
	}

	if err := rth.resetCodeUc.ResetPass(userID, req.Password, req.ConfirmPassword); err != nil {
		err = pkgres.Err(err).BadReq()
		return rth.log.Log(c, err)
	}

	return c.JSON(pkgres.NewRes("Cập nhật mật khẩu thành công"))
}
