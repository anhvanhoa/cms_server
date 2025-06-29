package authhandler

import (
	"cms-server/bootstrap"
	"cms-server/constants"
	authModel "cms-server/infrastructure/model/auth"
	"cms-server/infrastructure/repo"
	pkgjwt "cms-server/infrastructure/service/jwt"
	pkglog "cms-server/infrastructure/service/logger"
	pkgres "cms-server/infrastructure/service/response"
	authUC "cms-server/internal/usecase/auth"
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
)

type LoginHandler interface {
	Login(c *fiber.Ctx) error
}

type loginHandler struct {
	loginUsecase authUC.LoginUsecase
	log          pkglog.Logger
	env          *bootstrap.Env
}

func (lh *loginHandler) Login(c *fiber.Ctx) error {
	var body authModel.LoginReq

	if err := c.BodyParser(&body); err != nil {
		err := pkgres.NewErr("Dữ liệu không hợp lệ").BadReq()
		return lh.log.Log(c, err)
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

	if !lh.loginUsecase.CheckHashPassword(body.Password, user.Password) {
		err := pkgres.NewErr("Mật khẩu không chính xác, hãy kiểm tra lại").Code(fiber.StatusBadRequest)
		return lh.log.Log(c, err)
	}
	timeAccess := time.Now().Add(constants.AccessExpiredAt * time.Second)
	access, _ := lh.loginUsecase.GengerateAccessToken(user.ID, user.FullName, timeAccess)
	timeRefresh := time.Now().Add(constants.RefreshExpiredAt * time.Second)
	os := c.Get("User-Agent")
	refresh, _ := lh.loginUsecase.GengerateRefreshToken(user.ID, user.FullName, timeRefresh, os)

	c.Cookie(&fiber.Cookie{
		Name:     constants.KeyCookieAccessToken,
		Value:    access,
		Path:     "/",
		Domain:   lh.env.HOST_APP,
		Secure:   lh.env.IsProduction(),
		HTTPOnly: true,
		Expires:  timeAccess,
	})
	c.Cookie(&fiber.Cookie{
		Name:     constants.KeyCookieRefreshToken,
		Value:    refresh,
		Path:     "/",
		Domain:   lh.env.HOST_APP,
		Secure:   lh.env.IsProduction(),
		HTTPOnly: true,
		Expires:  timeRefresh,
	})

	return c.JSON(pkgres.ResData(user.GetInfor()).SetMessage("Đăng nhập thành công"))
}

func NewLoginHandler(loginUsecase authUC.LoginUsecase, log pkglog.Logger, env *bootstrap.Env) LoginHandler {
	return &loginHandler{
		loginUsecase: loginUsecase,
		log:          log,
		env:          env,
	}
}

func NewRouteLoginHandler(
	db *pg.DB,
	log pkglog.Logger,
	env *bootstrap.Env,
) LoginHandler {
	loginUsecase := authUC.NewLoginUsecase(
		repo.NewUserRepository(db),
		repo.NewSessionRepository(db),
		pkgjwt.NewJWT(env.JWT_SECRET.Access),
		pkgjwt.NewJWT(env.JWT_SECRET.Refresh),
	)
	return NewLoginHandler(loginUsecase, log, env)
}
