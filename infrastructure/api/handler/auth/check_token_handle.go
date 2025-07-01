package authhandler

import (
	"cms-server/infrastructure/repo"
	pkglog "cms-server/infrastructure/service/logger"
	pkgres "cms-server/infrastructure/service/response"
	authUC "cms-server/internal/usecase/auth"

	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
)

type CheckTokenHandler interface {
	CheckToken(c *fiber.Ctx) error
}

type checkTokenHandlerImpl struct {
	log              pkglog.Logger
	checkCodeUsecase authUC.CheckTokenUsecase
}

func NewCheckTokenHandler(
	checkCodeUsecase authUC.CheckTokenUsecase,
	log pkglog.Logger,
) CheckTokenHandler {
	return &checkTokenHandlerImpl{
		log,
		checkCodeUsecase,
	}
}

func NewRouterTokenHandler(
	db *pg.DB,
	log pkglog.Logger,
) CheckTokenHandler {
	checkCodeUsecase := authUC.NewCheckTokenUsecase(
		repo.NewSessionRepository(db),
	)
	return NewCheckTokenHandler(checkCodeUsecase, log)
}

func (h *checkTokenHandlerImpl) CheckToken(c *fiber.Ctx) error {
	token := c.Query("token", "")
	if token == "" {
		err := pkgres.NewErr("Dữ liệu không hợp lệ").BadReq()
		return h.log.Log(c, err)
	}

	ok, err := h.checkCodeUsecase.CheckToken(token)
	if err != nil {
		err := pkgres.Err(err).BadReq()
		return h.log.Log(c, err)
	}

	if !ok {
		err := pkgres.NewErr("Phiên làm việc không hợp lệ").BadReq()
		return h.log.Log(c, err)
	}
	return c.JSON(pkgres.NewRes("Phiên làm việc hợp lệ"))
}
