package router

import authhandler "cms-server/internal/api/handler/auth"

func (r *Router) initAuthRouter() {
	authRouter := r.app.Group("/auth")
	lh := authhandler.NewRouteLoginHandler(r.db, r.log)
	authRouter.Post("/login", lh.Login)
	// authRouter.Post("/register", r.Register)
	// authRouter.Post("/forgot-password", r.ForgotPassword)
	// authRouter.Post("/reset-password", r.ResetPassword)
	// authRouter.Post("/refresh-token", r.RefreshToken)
	// authRouter.Post("/logout", r.Logout)
	// authRouter.Post("/me", r.Me)
}
