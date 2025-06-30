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

type ResetCodeHandler interface {
	ResetPassword(c *fiber.Ctx) error
}

type resetCodeHandler struct {
	resetCodeUsecase authUC.ResetPasswordByCodeUsecase
	log              pkglog.Logger
}

func NewResetCodeHandler(resetCodeUsecase authUC.ResetPasswordByCodeUsecase) ResetCodeHandler {
	return &resetCodeHandler{
		resetCodeUsecase: resetCodeUsecase,
	}
}

func NewRouteResetByCodeHandler(
	db *pg.DB,
	log pkglog.Logger,
	cache cache.RedisConfigImpl,
	env *bootstrap.Env,
) ResetCodeHandler {
	resetCodeUsecase := authUC.NewResetPasswordCodeUsecase(
		repo.NewUserRepository(db),
		repo.NewSessionRepository(db),
		cache,
		pkgjwt.NewJWT(env.JWT_SECRET.Forgot),
		argonS.NewArgon().SetSaltLength(constants.SaltLength),
	)
	return &resetCodeHandler{
		resetCodeUsecase: resetCodeUsecase,
		log:              log,
	}
}

func (rth *resetCodeHandler) ResetPassword(c *fiber.Ctx) error {
	var req authModel.ResetPasswordByCodeRequest
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

	userID, err := rth.resetCodeUsecase.VerifySession(req.Code, req.Email)
	if err != nil {
		err = pkgres.Err(err).BadReq()
		return rth.log.Log(c, err)
	}

	if err := rth.resetCodeUsecase.ResetPass(userID, req.Password, req.ConfirmPassword); err != nil {
		err = pkgres.Err(err).BadReq()
		return rth.log.Log(c, err)
	}

	return c.JSON(pkgres.NewRes("Cập nhật mật khẩu thành công"))
}
