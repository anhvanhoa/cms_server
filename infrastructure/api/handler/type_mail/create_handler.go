package typeMail

import (
	pkgres "cms-server/infrastructure/service/response"

	"github.com/gofiber/fiber/v2"
)

func (h *TypeMailHandlerImpl) Create(c *fiber.Ctx) error {
	var body createReq
	if err := c.BodyParser(&body); err != nil {
		err = pkgres.NewErr("Dữ liệu không hợp lệ").BadReq()
		return h.log.Log(c, err)
	}

	if err := h.validate.ValidateStruct(&body); err != nil {
		err := pkgres.NewErr(err.Message).SetData(err.Data).BadReq()
		return h.log.Log(c, err)
	}

	CreatedBy := "1" // fix
	if err := h.createUseCase.Create(body.Name, CreatedBy); err != nil {
		err = pkgres.Err(err).BadReq()
		return h.log.Log(c, err)
	}
	return c.JSON(pkgres.NewRes("Tạo loại mail thành công"))
}
