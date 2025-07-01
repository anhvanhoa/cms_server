package typeMail

import (
	pkgres "cms-server/infrastructure/service/response"

	"github.com/gofiber/fiber/v2"
)

func (h *TypeMailHandlerImpl) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.deleteUseCase.Delete(id); err != nil {
		err = pkgres.Err(err).BadReq()
		return h.log.Log(c, err)
	}
	return c.JSON(pkgres.NewRes("Xóa loại mail thành công"))
}
