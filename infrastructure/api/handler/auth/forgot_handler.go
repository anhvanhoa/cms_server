package authhandler

import (
	"cms-server/bootstrap"
	authModel "cms-server/infrastructure/model/auth"
	"cms-server/infrastructure/repo"
	pkgjwt "cms-server/infrastructure/service/jwt"
	pkglog "cms-server/infrastructure/service/logger"
	pkgres "cms-server/infrastructure/service/response"
	"cms-server/internal/service/cache"
	"cms-server/internal/service/queue"
	authUC "cms-server/internal/usecase/auth"
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
)

type ForgotPasswordHandler interface {
	Forgot(c *fiber.Ctx) error
}
type forgotPasswordHandler struct {
	forgotUsecase authUC.ForgotPasswordUsecase
	log           pkglog.Logger
	env           *bootstrap.Env
}

func NewForgotPasswordHandler(
	forgotUsecase authUC.ForgotPasswordUsecase,
	log pkglog.Logger,
	env *bootstrap.Env,
) ForgotPasswordHandler {
	return &forgotPasswordHandler{
		forgotUsecase,
		log,
		env,
	}
}

func NewRouteForgotHandler(
	db *pg.DB,
	log pkglog.Logger,
	env *bootstrap.Env,
	qc queue.QueueClient,
	cache cache.RedisConfigImpl,
) ForgotPasswordHandler {
	forgotUsecase := authUC.NewForgotPasswordUsecase(
		repo.NewUserRepository(db),
		repo.NewSessionRepository(db),
		repo.NewMailTplRepository(db),
		repo.NewStatusHistoryRepository(db),
		repo.NewMailHistoryRepository(db),
		repo.NewManagerTransaction(db),
		pkgjwt.NewJWT(env.JWT_SECRET.Forgot),
		qc,
		cache,
	)
	return NewForgotPasswordHandler(forgotUsecase, log, env)
}

func (h *forgotPasswordHandler) Forgot(c *fiber.Ctx) error {
	var body authModel.ForgotPasswordReq
	if err := c.BodyParser(&body); err != nil {
		err := pkgres.NewErr("Dữ liệu không hợp lệ").BadReq()
		return h.log.Log(c, err)
	}

	if _, err := govalidator.ValidateStruct(body); err != nil {
		err := pkgres.Err(err).UnprocessableEntity()
		return h.log.Log(c, err)
	}

	os := c.Get("User-Agent")
	resForpass, err := h.forgotUsecase.ForgotPassword(body.Email, os, body.Type)

	if errors.Is(err, authUC.ErrValidateForgotPassword) {
		err := pkgres.Err(err).BadReq()
		return h.log.Log(c, err)
	} else if err != nil {
		err := pkgres.NewErr("Không tìm thấy tài khoản, hãy kiểm tra lại").BadReq()
		return h.log.Log(c, err)
	}

	var link string
	if body.Type == authUC.ForgotByToken {
		link = h.env.FRONTEND_URL + "/auth/forgot-password?code=" + resForpass.Token
	}
	if err := h.forgotUsecase.SendEmailForgotPassword(resForpass.User, resForpass.Code, link); err != nil {
		err := pkgres.Err(err).Code(fiber.StatusInternalServerError)
		return h.log.Log(c, err)
	}
	return c.JSON(pkgres.NewRes("Yêu cầu đặt lại mật khẩu đã được gửi đến email của bạn"))
}
