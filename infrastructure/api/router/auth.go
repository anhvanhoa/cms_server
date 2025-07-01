package router

import authhandler "cms-server/infrastructure/api/handler/auth"

func (r *Router) initAuthRouter() {
	authRouter := r.app.Group("/auth")
	lh := authhandler.NewRouteLoginHandler(r.db, r.log, r.env, r.cache)
	rh := authhandler.NewRouteRegisterHandler(r.db, r.log, r.qc, r.env, r.cache)
	vah := authhandler.NewVerifyAccountHandler(r.db, r.log, r.env, r.cache)
	rfh := authhandler.NewRouteRefreshHandler(r.db, r.log, r.env, r.cache)
	fh := authhandler.NewRouteForgotHandler(r.db, r.log, r.env, r.qc, r.cache)
	rth := authhandler.NewRouteResetByTokenHandler(r.db, r.log, r.cache, r.env)
	rch := authhandler.NewRouteResetByCodeHandler(r.db, r.log, r.cache, r.env)
	cth := authhandler.NewRouterTokenHandler(r.db, r.log)
	cch := authhandler.NewRouterCodeHandler(r.db, r.log)
	logoutH := authhandler.NewRouterLogoutHandler(r.db, r.log, r.env, r.cache)
	authRouter.Post("/login", lh.Login)
	authRouter.Post("/register", rh.Register)
	authRouter.Post("/verify/:t", vah.VerifyAccount)
	authRouter.Post("/forgot-password", fh.Forgot)
	authRouter.Get("/forgot-password", cth.CheckToken) //?token=...
	authRouter.Post("/reset-password", rth.ResetPassword)
	authRouter.Post("/check-code/forgot-password", cch.CheckCode) //body: {code:"...", email:"..."}
	authRouter.Post("/reset-password/code", rch.ResetPassword)
	authRouter.Post("/refresh", rfh.Refresh)
	authRouter.Post("/logout", logoutH.Logout)
}
