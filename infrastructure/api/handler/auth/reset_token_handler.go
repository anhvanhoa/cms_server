package authhandler

import (
	"cms-server/bootstrap"
	"cms-server/constants"
	authModel "cms-server/infrastructure/model/auth"
	"cms-server/infrastructure/repo"
	argonS "cms-server/infrastructure/service/argon"
	pkgjwt "cms-server/infrastructure/service/jwt"
	pkglog "cms-server/infrastructure/service/logger"
	pkgres "cms-server/infrastructure/service/response"
	"cms-server/internal/service/cache"
	authUC "cms-server/internal/usecase/auth"

	"github.com/asaskevich/govalidator"
	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
)

type ResetTokenHandler interface {
	ResetPassword(c *fiber.Ctx) error
}

type resetTokenHandler struct {
	resetTokenUsecase authUC.ResetPasswordByTokenUsecase
	log               pkglog.Logger
}

func NewResetTokenHandler(resetTokenUsecase authUC.ResetPasswordByTokenUsecase) ResetTokenHandler {
	return &resetTokenHandler{
		resetTokenUsecase: resetTokenUsecase,
	}
}

func NewRouteResetByTokenHandler(
	db *pg.DB,
	log pkglog.Logger,
	cache cache.RedisConfigImpl,
	env *bootstrap.Env,
) ResetTokenHandler {
	resetTokenUsecase := authUC.NewResetPasswordUsecase(
		repo.NewUserRepository(db),
		repo.NewSessionRepository(db),
		cache,
		pkgjwt.NewJWT(env.JWT_SECRET.Forgot),
		argonS.NewArgon().SetSaltLength(constants.SaltLength),
	)
	return &resetTokenHandler{
		resetTokenUsecase: resetTokenUsecase,
		log:               log,
	}
}

func (rth *resetTokenHandler) ResetPassword(c *fiber.Ctx) error {
	var req authModel.ResetPasswordByTokenRequest
	if err := c.BodyParser(&req); err != nil {
		err := pkgres.NewErr("Dữ liệu không hợp lệ").BadReq()
		return rth.log.Log(c, err)
	}

	if _, err := govalidator.ValidateStruct(req); err != nil {
		err := pkgres.Err(err).UnprocessableEntity()
		return rth.log.Log(c, err)
	}

	if req.Password != req.ConfirmPassword {
		err := pkgres.NewErr("Mật khẩu không khớp").BadReq()
		return rth.log.Log(c, err)
	}

	userID, err := rth.resetTokenUsecase.VerifySession(req.Token)
	if err != nil {
		err = pkgres.Err(err).BadReq()
		return rth.log.Log(c, err)
	}

	if err := rth.resetTokenUsecase.ResetPass(userID, req.Password, req.ConfirmPassword); err != nil {
		err = pkgres.Err(err).BadReq()
		return rth.log.Log(c, err)
	}

	return c.JSON(pkgres.NewRes("Cập nhật mật khẩu thành công"))
}
