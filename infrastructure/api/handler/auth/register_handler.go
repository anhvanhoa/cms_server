package authhandler

import (
	"cms-server/bootstrap"
	"cms-server/constants"
	authModel "cms-server/infrastructure/model/auth"
	"cms-server/infrastructure/repo"
	pkgjwt "cms-server/infrastructure/service/jwt"
	pkglog "cms-server/infrastructure/service/logger"
	pkgres "cms-server/infrastructure/service/response"
	"cms-server/internal/service/queue"
	authUC "cms-server/internal/usecase/auth"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
)

type RegisterHandler interface {
	Register(c *fiber.Ctx) error
}

type registerHandler struct {
	registerUsecase authUC.RegisterUsecase
	log             pkglog.Logger
	env             *bootstrap.Env
}

func (rh *registerHandler) Register(c *fiber.Ctx) error {
	var body authModel.RegisterReq

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

	if u, err := rh.registerUsecase.CheckUserExist(body.Email); err != nil && err != pg.ErrNoRows {
		return rh.log.Log(c, err)
	} else if u.ID != "" && u.Veryfied != nil {
		err := pkgres.NewErr("Tài khoản đã tồn tại, vui lòng thử lại !").Code(fiber.StatusBadRequest)
		return rh.log.Log(c, err)
	}

	expAt := time.Now().Add(time.Second * constants.VerifyExpiredAt)
	if opt, err := rh.registerUsecase.GengerateCode(expAt, rh.env.SECRET_OTP); err != nil {
		return rh.log.Log(c, err)
	} else {
		body.Code = opt
	}
	os := c.Get("User-Agent")
	dataRegister := authUC.RegisterReq{
		Email:           body.Email,
		FullName:        body.FullName,
		Password:        body.Password,
		ConfirmPassword: body.ConfirmPassword,
		Code:            body.Code,
	}
	res, err := rh.registerUsecase.Register(dataRegister, os, expAt)
	if pg.ErrNoRows == err {
		err := pkgres.NewErr("Không tìm thấy mẫu email").NotFound()
		return rh.log.Log(c, err)
	} else if err != nil {
		return rh.log.Log(c, err)
	}

	err = rh.registerUsecase.SendMail(res.MailTpl, res.UserInfor, rh.env.FRONTEND_URL+"/auth/verify/"+res.Token)
	if pg.ErrNoRows == err {
		err := pkgres.NewErr("Không tìm thấy mẫu email").NotFound()
		return rh.log.Log(c, err)
	} else if err != nil {
		return rh.log.Log(c, err)
	}
	return c.JSON(pkgres.ResData(res.UserInfor).SetMessage("Đăng ký thành công"))
}

func NewRegisterHandler(registerUsecase authUC.RegisterUsecase, log pkglog.Logger, env *bootstrap.Env) RegisterHandler {
	return &registerHandler{
		registerUsecase: registerUsecase,
		log:             log,
		env:             env,
	}
}

func NewRouteRegisterHandler(
	db *pg.DB,
	log pkglog.Logger,
	qc queue.QueueClient,
	env *bootstrap.Env,
) RegisterHandler {
	registerUsecase := authUC.NewRegisterUsecase(
		repo.NewUserRepository(db),
		repo.NewMailTplRepository(db),
		repo.NewMailHistoryRepository(db),
		repo.NewStatusHistoryRepository(db),
		repo.NewSessionRepository(db),
		pkgjwt.NewJWT(env.JWT_SECRET.Verify),
		qc,
		repo.NewManagerTransaction(db),
	)
	return NewRegisterHandler(registerUsecase, log, env)
}
