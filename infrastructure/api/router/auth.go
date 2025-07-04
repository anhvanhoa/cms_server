package router

import (
	authUC "cms-server/domain/usecase/auth"
	handler "cms-server/infrastructure/api/handler/auth"
	"cms-server/infrastructure/repo"
	argonS "cms-server/infrastructure/service/argon"
	goidS "cms-server/infrastructure/service/goid"
	pkgjwt "cms-server/infrastructure/service/jwt"
)

func (r *Router) initAuthRouter() {
	authR := r.app.Group("/auth")
	sessionRepo := repo.NewSessionRepository(r.db)
	userRepo := repo.NewUserRepository(r.db)
	mailTplRepo := repo.NewMailTplRepository(r.db)
	shRepo := repo.NewStatusHistoryRepository(r.db)
	mhRepo := repo.NewMailHistoryRepository(r.db)
	tx := repo.NewManagerTransaction(r.db)
	jwtForgot := pkgjwt.NewJWT(r.env.JWT_SECRET.Forgot)
	jwtAccess := pkgjwt.NewJWT(r.env.JWT_SECRET.Access)
	jwtRefresh := pkgjwt.NewJWT(r.env.JWT_SECRET.Refresh)
	jwtVerify := pkgjwt.NewJWT(r.env.JWT_SECRET.Verify)
	argon := argonS.NewArgon()
	goid := goidS.NewGoId()
	h := handler.NewAuthHandler(
		authUC.NewCheckTokenUsecase(sessionRepo),
		authUC.NewCheckCodeUsecase(userRepo, sessionRepo),
		authUC.NewForgotPasswordUsecase(userRepo, sessionRepo, mailTplRepo, shRepo, mhRepo, tx, jwtForgot, r.qc, r.cache),
		authUC.NewLoginUsecase(userRepo, sessionRepo, jwtAccess, jwtRefresh, argon, r.cache),
		authUC.NewLogoutUsecase(sessionRepo, jwtAccess, r.cache),
		authUC.NewRefreshUsecase(sessionRepo, jwtAccess, jwtRefresh, r.cache),
		authUC.NewRegisterUsecase(userRepo, mailTplRepo, mhRepo, shRepo, sessionRepo, jwtVerify, r.qc, tx, goid, argon, r.cache),
		authUC.NewResetPasswordCodeUsecase(userRepo, sessionRepo, r.cache, jwtForgot, argon),
		authUC.NewResetPasswordTokenUsecase(userRepo, sessionRepo, r.cache, jwtForgot, argon),
		authUC.NewVerifyAccountUsecase(userRepo, sessionRepo, jwtVerify, r.cache),
		r.log,
		r.env,
		r.valid,
	)
	authR.Post("/login", h.Login)
	authR.Post("/register", h.Register)
	authR.Post("/verify/:t", h.VerifyAccount)
	authR.Post("/forgot-password", h.Forgot)
	authR.Get("/forgot-password", h.CheckToken)
	authR.Post("/reset-password", h.ResetToken)
	authR.Post("/check-code/forgot-password", h.CheckCode)
	authR.Post("/reset-password/code", h.ResetCode)
	authR.Post("/refresh", h.Refresh)
	authR.Post("/logout", h.Logout)
}
