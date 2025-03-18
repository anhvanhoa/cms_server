package authhandler

import (
	"cms-server/bootstrap"
	"cms-server/internal/repository"
	"cms-server/internal/usecase/auth"
	pkgjwt "cms-server/pkg/jwt"
	pkglog "cms-server/pkg/logger"
	pkgres "cms-server/pkg/response"

	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
)

type VerifyAccountHandler interface {
	VerifyAccount(c *fiber.Ctx) error
}

type verifyAccountHandler struct {
	verifyAccountUsecase auth.VerifyAccountUsecase
	log                  pkglog.Logger
}

func (vah *verifyAccountHandler) VerifyAccount(c *fiber.Ctx) error {
	t := c.Params("t")
	if t == "" {
		err := pkgres.NewErr("Dữ liệu không hợp lệ").BadReq()
		return vah.log.Log(c, err)
	}

	claims, err := vah.verifyAccountUsecase.VerifyRegister(t)
	if err == pkgjwt.ErrParseToken {
		err := pkgres.NewErr("Không thể lấy thông tin")
		return vah.log.Log(c, err)
	} else if err != nil {
		err := pkgres.NewErr("Token không hợp lệ").BadReq()
		return vah.log.Log(c, err)
	}
	user, err := vah.verifyAccountUsecase.GetUserById(claims.Id)
	if err == pg.ErrNoRows {
		err := pkgres.NewErr("Tài khoản không tồn tại").NotFound()
		return vah.log.Log(c, err)
	} else if err != nil {
		return vah.log.Log(c, err) // internal error
	} else if user.Veryfied != nil {
		err := pkgres.NewErr("Tài khoản đã được xác thực").BadReq()
		return vah.log.Log(c, err)
	} else if user.CodeVerify != claims.Code {
		err := pkgres.NewErr("Mã xác thực không hợp lệ").BadReq()
		return vah.log.Log(c, err)
	}

	if err := vah.verifyAccountUsecase.VerifyAccount(claims.Id); err != nil {
		return vah.log.Log(c, err)
	}
	res := pkgres.NewRes("Xác thực tài khoản thành công").Code(fiber.StatusOK)
	return c.Status(res.GetCode()).JSON(res)
}

func NewVerifyAccountHandler(
	db *pg.DB,
	log pkglog.Logger,
	env *bootstrap.Env,
) VerifyAccountHandler {
	vau := auth.NewVerifyAccountUsecase(
		repository.NewUserRepository(db),
		pkgjwt.NewJWT(env.JWT_SECRET),
	)
	return &verifyAccountHandler{
		vau,
		log,
	}
}
