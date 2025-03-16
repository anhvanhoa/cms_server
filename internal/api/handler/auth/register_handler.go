package authhandler

import (
	"cms-server/bootstrap"
	modelauth "cms-server/internal/model/auth"
	"cms-server/internal/repository"
	"cms-server/internal/usecase/auth"
	pkgjwt "cms-server/pkg/jwt"
	pkglog "cms-server/pkg/logger"
	pkgres "cms-server/pkg/response"
	"time"

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

	time := time.Now().Add(time.Minute * 5)
	if opt, err := rh.registerUsecase.GengerateCode(time); err != nil {
		return rh.log.Log(c, err)
	} else {
		body.Code = opt
	}

	user, tpl, err := rh.registerUsecase.CreateUser(body)
	if pg.ErrNoRows == err {
		err := pkgres.NewErr("Không tìm thấy mẫu email").NotFound()
		return rh.log.Log(c, err)
	} else if err != nil {
		return rh.log.Log(c, err)
	}

	claims := pkgjwt.NewRegisterClaims(body.Code, time)
	token, err := rh.registerUsecase.GengerateToken(claims)
	if err != nil {
		return rh.log.Log(c, err)
	}
	err = rh.registerUsecase.SendMail(tpl, token, user)
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

func NewRouteRegisterHandler(
	db *pg.DB,
	log pkglog.Logger,
	qc bootstrap.QueueClient,
	env *bootstrap.Env,
) RegisterHandler {
	registerUsecase := auth.NewRegisterUsecase(
		repository.NewUserRepository(db),
		repository.NewMailTplRepository(db),
		repository.NewMailHistoryRepository(db),
		repository.NewStatusHistoryRepository(db),
		pkgjwt.NewJWT(env.JWT_SECRET),
		qc,
		repository.NewManagerTransaction(db),
		env,
	)
	return NewRegisterHandler(registerUsecase, log)
}
