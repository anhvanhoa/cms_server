package authhandler

import (
	modelauth "cms-server/internal/model/auth"
	"cms-server/internal/repository"
	"cms-server/internal/usecase/auth"
	pkglog "cms-server/pkg/logger"
	pkgres "cms-server/pkg/response"
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
)

type LoginHandler interface {
	Login(c *fiber.Ctx) error
}

type loginHandler struct {
	loginUsecase auth.LoginUsecase
	log          pkglog.Logger
}

func (lh *loginHandler) Login(c *fiber.Ctx) error {
	var body modelauth.LoginReq

	if err := c.BodyParser(&body); err != nil {
		res := pkgres.NewErr("Dữ liệu không hợp lệ").BadReq()
		return c.Status(res.GetCode()).JSON(res)
	}

	if _, err := govalidator.ValidateStruct(body); err != nil {
		err := pkgres.Err(err).UnprocessableEntity()
		return lh.log.Log(c, err)
	}

	user, err := lh.loginUsecase.GetUserByEmailOrPhone(body.Identifier)
	if errors.Is(err, pg.ErrNoRows) {
		err := pkgres.NewErr("Tài khoản không hợp lệ, hãy kiểm tra lại").Code(fiber.StatusBadRequest)
		return lh.log.Log(c, err)
	}

	if err != nil {
		return lh.log.Log(c, err)
	}

	return c.JSON(pkgres.ResData(user.GetInfor()).SetMessage("Đăng nhập thành công"))
}

func NewLoginHandler(loginUsecase auth.LoginUsecase, log pkglog.Logger) LoginHandler {
	return &loginHandler{
		loginUsecase: loginUsecase,
		log:          log,
	}
}

func NewRouteLoginHandler(db *pg.DB, log pkglog.Logger) LoginHandler {
	loginUsecase := repository.NewUserRepository(db)
	return NewLoginHandler(loginUsecase, log)
}
