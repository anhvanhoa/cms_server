package typeMail

import (
	pkgres "cms-server/infrastructure/service/response"

	"github.com/gofiber/fiber/v2"
)

func (h *TypeMailHandlerImpl) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	type req struct {
		Name      string `json:"name"`
		UpdatedBy string `json:"updated_by"`
	}
	var body req
	if err := c.BodyParser(&body); err != nil {
		err = pkgres.Err(err).BadReq()
		return h.log.Log(c, err)
	}
	if err := h.updateUseCase.Update(id, body.Name, body.UpdatedBy); err != nil {
		err = pkgres.Err(err).BadReq()
		return h.log.Log(c, err)
	}
	return c.JSON(pkgres.NewRes("Cập nhật loại mail thành công"))
}
