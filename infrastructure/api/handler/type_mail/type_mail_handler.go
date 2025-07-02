package typeMail

import (
	"cms-server/bootstrap"
	typeMailUC "cms-server/domain/usecase/type_mail"
	pkglog "cms-server/infrastructure/service/logger"

	"github.com/gofiber/fiber/v2"
)

type TypeMailHandler interface {
	GetByID(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type TypeMailHandlerImpl struct {
	getUseCase    typeMailUC.GetUseCase
	getAllUseCase typeMailUC.GetAllUseCase
	createUseCase typeMailUC.CreateUseCase
	updateUseCase typeMailUC.UpdateUseCase
	deleteUseCase typeMailUC.DeleteUseCase
	log           pkglog.Logger
	validate      bootstrap.IValidator
}

func NewTypeMailHandler(
	getUseCase typeMailUC.GetUseCase,
	getAllUseCase typeMailUC.GetAllUseCase,
	createUseCase typeMailUC.CreateUseCase,
	updateUseCase typeMailUC.UpdateUseCase,
	deleteUseCase typeMailUC.DeleteUseCase,
	log pkglog.Logger,
	validate bootstrap.IValidator,
) TypeMailHandler {
	return &TypeMailHandlerImpl{
		getUseCase:    getUseCase,
		getAllUseCase: getAllUseCase,
		createUseCase: createUseCase,
		updateUseCase: updateUseCase,
		deleteUseCase: deleteUseCase,
		log:           log,
		validate:      validate,
	}
}
