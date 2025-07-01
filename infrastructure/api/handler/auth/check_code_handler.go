package authhandler

import (
	authModel "cms-server/infrastructure/model/auth"
	"cms-server/infrastructure/repo"
	pkglog "cms-server/infrastructure/service/logger"
	pkgres "cms-server/infrastructure/service/response"
	authUC "cms-server/internal/usecase/auth"

	"github.com/asaskevich/govalidator"
	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
)

type CheckCodeHandler interface {
	CheckCode(c *fiber.Ctx) error
}

type checkCodeHandlerImpl struct {
	log              pkglog.Logger
	checkCodeUsecase authUC.CheckCodeUsecase
}

func NewCheckCodeHandler(
	checkCodeUsecase authUC.CheckCodeUsecase,
	log pkglog.Logger,
) CheckCodeHandler {
	return &checkCodeHandlerImpl{
		log,
		checkCodeUsecase,
	}
}

func NewRouterCodeHandler(
	db *pg.DB,
	log pkglog.Logger,
) CheckCodeHandler {
	checkCodeUsecase := authUC.NewCheckCodeUsecase(
		repo.NewUserRepository(db),
		repo.NewSessionRepository(db),
	)
	return NewCheckCodeHandler(checkCodeUsecase, log)
}

func (h *checkCodeHandlerImpl) CheckCode(c *fiber.Ctx) error {
	var req authModel.CheckCodeReq
	if err := c.BodyParser(&req); err != nil {
		err := pkgres.NewErr("Dữ liệu không hợp lệ").BadReq()
		return h.log.Log(c, err)
	}

	if _, err := govalidator.ValidateStruct(req); err != nil {
		err := pkgres.NewErr("Dữ liệu không hợp lệ").UnprocessableEntity()
		return h.log.Log(c, err)
	}

	ok, err := h.checkCodeUsecase.CheckCode(req.Code, req.Email)
	if err != nil {
		err := pkgres.Err(err).BadReq()
		return h.log.Log(c, err)
	}

	if !ok {
		err := pkgres.NewErr("Mã xác thực không hợp lệ").BadReq()
		return h.log.Log(c, err)
	}
	return c.JSON(pkgres.NewRes("Mã xác thực hợp lệ"))
}
