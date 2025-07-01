package typeMail

import (
	pkgres "cms-server/infrastructure/service/response"

	"github.com/gofiber/fiber/v2"
)

func (h *TypeMailHandlerImpl) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var body updateReq
	if err := c.BodyParser(&body); err != nil {
		err = pkgres.Err(err).BadReq()
		return h.log.Log(c, err)
	}

	if err := h.validate.ValidateStruct(&body); err != nil {
		err := pkgres.NewErr(err.Message).SetData(err.Data).BadReq()
		return h.log.Log(c, err)
	}

	if err := h.updateUseCase.Update(id, body.Name); err != nil {
		err = pkgres.Err(err).BadReq()
		return h.log.Log(c, err)
	}
	return c.JSON(pkgres.NewRes("Cập nhật loại mail thành công"))
}
