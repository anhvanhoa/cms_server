package typeMail

import (
	pkgres "cms-server/infrastructure/service/response"

	"github.com/gofiber/fiber/v2"
)

func (h *TypeMailHandlerImpl) GetAll(c *fiber.Ctx) error {
	limit := c.QueryInt("pageSize", 10)
	offset := c.QueryInt("offset", 0)
	result, err := h.getAllUseCase.GetAll(limit, offset)
	if err != nil {
		err = pkgres.Err(err).BadReq()
		return h.log.Log(c, err)
	}
	return c.JSON(pkgres.ResData(result).SetMessage("Lấy danh sách loại mail thành công"))
}
