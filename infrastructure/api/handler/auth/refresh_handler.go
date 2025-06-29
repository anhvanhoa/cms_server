package authhandler

import (
	"cms-server/bootstrap"
	"cms-server/constants"
	"cms-server/infrastructure/repo"
	pkgjwt "cms-server/infrastructure/service/jwt"
	pkglog "cms-server/infrastructure/service/logger"
	pkgres "cms-server/infrastructure/service/response"
	authUC "cms-server/internal/usecase/auth"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
)

type RefreshHandler interface {
	Refresh(c *fiber.Ctx) error
}

type refreshHandler struct {
	refreshUsecase authUC.RefreshUsecase
	log            pkglog.Logger
	env            *bootstrap.Env
}

func NewRefreshHandler(
	refreshUsecase authUC.RefreshUsecase,
	log pkglog.Logger,
	env *bootstrap.Env,
) RefreshHandler {
	return &refreshHandler{
		refreshUsecase: refreshUsecase,
		log:            log,
		env:            env,
	}
}

func NewRouteRefreshHandler(
	db *pg.DB,
	log pkglog.Logger,
	env *bootstrap.Env,
) RefreshHandler {
	refreshUsecase := authUC.NewRefreshUsecase(
		repo.NewSessionRepository(db),
		pkgjwt.NewJWT(env.JWT_SECRET.Access),
		pkgjwt.NewJWT(env.JWT_SECRET.Refresh),
	)
	return NewRefreshHandler(refreshUsecase, log, env)
}

func (rh *refreshHandler) Refresh(c *fiber.Ctx) error {
	refresh := c.Cookies(constants.KeyCookieRefreshToken, "")
	session, err := rh.refreshUsecase.GetSessionByToken(refresh)
	if err != nil {
		err := pkgres.NewErr("Phiên làm việc không hợp lệ, hãy đăng nhập lại").Unauthorized()
		return rh.log.Log(c, err)
	}

	if err := rh.refreshUsecase.ClearSessionExpired(); err != nil {
		err := pkgres.NewErr("Không thể làm mới phiên làm việc").InternalServerError()
		return rh.log.Log(c, err)
	}

	claims, err := rh.refreshUsecase.VerifyToken(refresh)
	if err != nil {
		err := pkgres.NewErr("Phiên làm việc không hợp lệ, hãy đăng nhập lại").Unauthorized()
		return rh.log.Log(c, err)
	}
	expAccess := time.Now().Add(constants.AccessExpiredAt * time.Second)
	access, err := rh.refreshUsecase.GengerateAccessToken(session.UserID, claims.FullName, expAccess)
	if err != nil {
		err := pkgres.NewErr("Không thể tạo token mới").InternalServerError()
		return rh.log.Log(c, err)
	}
	expRefresh := time.Now().Add(constants.RefreshExpiredAt * time.Second)
	os := c.Get("User-Agent")
	refreshToken, err := rh.refreshUsecase.GengerateRefreshToken(session.UserID, claims.FullName, expRefresh, os)
	if err != nil {
		err := pkgres.NewErr("Không thể tạo token mới").InternalServerError()
		return rh.log.Log(c, err)
	}
	expR := time.Now().Add(constants.RefreshExpiredAt * time.Second)
	expA := time.Now().Add(constants.AccessExpiredAt * time.Second)
	c.Cookie(&fiber.Cookie{
		Name:     constants.KeyCookieRefreshToken,
		Value:    refreshToken,
		Path:     "/",
		Domain:   rh.env.HOST_APP,
		Secure:   rh.env.IsProduction(),
		HTTPOnly: true,
		Expires:  expR,
	})
	c.Cookie(&fiber.Cookie{
		Name:     constants.KeyCookieAccessToken,
		Value:    access,
		Path:     "/",
		Domain:   rh.env.HOST_APP,
		Secure:   rh.env.IsProduction(),
		HTTPOnly: true,
		Expires:  expA,
	})
	return c.JSON(pkgres.ResData(nil).SetMessage("Làm mới phiên làm việc thành công"))
}
