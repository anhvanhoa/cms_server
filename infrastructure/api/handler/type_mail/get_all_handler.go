package typeMail

import (
	pkgres "cms-server/infrastructure/service/response"

	"github.com/gofiber/fiber/v2"
)

func (h *TypeMailHandlerImpl) GetAll(c *fiber.Ctx) error {
	pageSize := c.QueryInt("pageSize", 10)
	page := c.QueryInt("page", 1)
	result, err := h.getAllUseCase.GetAll(pageSize, page)
	if err != nil {
		err = pkgres.Err(err).BadReq()
		return h.log.Log(c, err)
	}
	return c.JSON(pkgres.ResData(result).SetMessage("Lấy danh sách loại mail thành công"))
}
