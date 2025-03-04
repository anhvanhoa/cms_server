package authhandler

import (
	"cms-server/bootstrap"
	modelauth "cms-server/internal/model/auth"
	"cms-server/internal/repository"
	"cms-server/internal/usecase/auth"
	pkglog "cms-server/pkg/logger"
	pkgres "cms-server/pkg/response"

	"github.com/asaskevich/govalidator"
	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
)

type RegisterHandler interface {
	Register(c *fiber.Ctx) error
}

type registerHandler struct {
	registerUsecase auth.RegisterUsecase
	log             pkglog.Logger
}

func (rh *registerHandler) Register(c *fiber.Ctx) error {
	var body modelauth.RegisterReq

	if err := c.BodyParser(&body); err != nil {
		err := pkgres.NewErr("Dữ liệu không hợp lệ").BadReq()
		return rh.log.Log(c, err)
	}

	if _, err := govalidator.ValidateStruct(body); err != nil {
		err := pkgres.Err(err).UnprocessableEntity()
		return rh.log.Log(c, err)
	}

	if body.Password != body.ConfirmPassword {
		err := pkgres.NewErr("Mật khẩu không khớp").BadReq()
		return rh.log.Log(c, err)
	}

	if isExist, err := rh.registerUsecase.CheckUserExist(body.Email); err != nil {
		return rh.log.Log(c, err)
	} else if isExist {
		err := pkgres.NewErr("Tài khoản đã tồn tại, vui lòng thử lại !").Code(fiber.StatusBadRequest)
		return rh.log.Log(c, err)
	}

	user, tpl, err := rh.registerUsecase.CreateUser(body)
	if pg.ErrNoRows == err {
		err := pkgres.NewErr("Không tìm thấy mẫu email").NotFound()
		return rh.log.Log(c, err)
	} else if err != nil {
		return rh.log.Log(c, err)
	}

	err = rh.registerUsecase.SendMail(tpl, "register", user)
	if pg.ErrNoRows == err {
		err := pkgres.NewErr("Không tìm thấy mẫu email").NotFound()
		return rh.log.Log(c, err)
	} else if err != nil {
		return rh.log.Log(c, err)
	}
	return c.JSON(pkgres.ResData(user).SetMessage("Đăng ký thành công"))
}

func NewRegisterHandler(registerUsecase auth.RegisterUsecase, log pkglog.Logger) RegisterHandler {
	return &registerHandler{
		registerUsecase: registerUsecase,
		log:             log,
	}
}

func NewRouteRegisterHandler(db *pg.DB, log pkglog.Logger, qc bootstrap.QueueClient) RegisterHandler {
	registerUsecase := auth.NewRegisterUsecase(repository.NewUserRepository(db), repository.NewMailTplRepository(db), qc)
	return NewRegisterHandler(registerUsecase, log)
}
