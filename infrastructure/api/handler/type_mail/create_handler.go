package typeMail

import (
	pkgres "cms-server/infrastructure/service/response"

	"github.com/gofiber/fiber/v2"
)

func (h *TypeMailHandlerImpl) Create(c *fiber.Ctx) error {
	type req struct {
		Name      string `json:"name"`
		CreatedBy string `json:"created_by"`
	}
	var body req
	if err := c.BodyParser(&body); err != nil {
		err = pkgres.Err(err).BadReq()
		return h.log.Log(c, err)
	}
	if err := h.createUseCase.Create(body.Name, body.CreatedBy); err != nil {
		err = pkgres.Err(err).BadReq()
		return h.log.Log(c, err)
	}
	return c.JSON(pkgres.NewRes("Tạo loại mail thành công"))
}
