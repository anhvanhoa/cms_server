package authhandler

import (
	"cms-server/bootstrap"
	"cms-server/constants"
	"cms-server/infrastructure/repo"
	pkgjwt "cms-server/infrastructure/service/jwt"
	pkglog "cms-server/infrastructure/service/logger"
	pkgres "cms-server/infrastructure/service/response"
	"cms-server/internal/service/cache"
	serviceError "cms-server/internal/service/error"
	authUC "cms-server/internal/usecase/auth"
	"errors"

	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
)

type LogoutHandler interface {
	Logout(c *fiber.Ctx) error
}

type logoutHandlerImpl struct {
	log pkglog.Logger
	uc  authUC.LogoutUsecase
}

func NewLogoutHandler(
	log pkglog.Logger,
	uc authUC.LogoutUsecase,
) LogoutHandler {
	return &logoutHandlerImpl{
		log,
		uc,
	}
}

func NewRouterLogoutHandler(
	db *pg.DB,
	log pkglog.Logger,
	env *bootstrap.Env,
	cache cache.RedisConfigImpl,
) LogoutHandler {
	uc := authUC.NewLogoutUsecase(
		repo.NewSessionRepository(db),
		pkgjwt.NewJWT(env.JWT_SECRET.Access),
		cache,
	)
	return NewLogoutHandler(log, uc)
}

func (l *logoutHandlerImpl) Logout(c *fiber.Ctx) error {
	token := c.Cookies(constants.KeyCookieAccessToken)
	if token == "" {
		err := pkgres.NewErr("Phiên làm việc đã hết hạn, vui lòng đăng nhập").Unauthorized()
		return l.log.Log(c, err)
	}

	if err := l.uc.VerifyToken(token); err != nil {
		if errors.Is(err, serviceError.ErrNotFoundSession) {
			err := pkgres.Err(err).Unauthorized()
			return l.log.Log(c, err)
		}
		err := pkgres.NewErr("Phiên làm việc không hợp lệ").Unauthorized()
		return l.log.Log(c, err)
	}

	if err := l.uc.Logout(token); err != nil {
		err := pkgres.Err(err).InternalServerError()
		return l.log.Log(c, err)
	}

	return c.JSON(pkgres.NewErr("Đăng xuất thành công"))
}
