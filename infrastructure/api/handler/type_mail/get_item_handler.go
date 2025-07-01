package typeMail

import (
	pkgres "cms-server/infrastructure/service/response"

	"github.com/gofiber/fiber/v2"
)

func (h *TypeMailHandlerImpl) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	result, err := h.getUseCase.GetByID(id)
	if err != nil {
		err = pkgres.Err(err).BadReq()
		return h.log.Log(c, err)
	}
	return c.JSON(pkgres.ResData(result).SetMessage("Lấy thông tin loại mail thành công"))
}
