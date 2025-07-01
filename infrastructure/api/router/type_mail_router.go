package router

import (
	typeMail "cms-server/domain/usecase/type_mail"
	handler "cms-server/infrastructure/api/handler/type_mail"
	"cms-server/infrastructure/repo"
	goidS "cms-server/infrastructure/service/goid"
)

func (r *Router) initTypeMailRouter() {
	tmRouter := r.app.Group("/type-mails")
	typeMailRepo := repo.NewTypeMailRepository(r.db)
	goid := goidS.NewGoId()
	h := handler.NewTypeMailHandler(
		typeMail.NewGetUseCase(typeMailRepo),
		typeMail.NewGetAllUseCase(typeMailRepo),
		typeMail.NewCreateUseCase(typeMailRepo, goid),
		typeMail.NewUpdateUseCase(typeMailRepo),
		typeMail.NewDeleteUseCase(typeMailRepo),
		r.log,
	)
	tmRouter.Get("/:id", h.GetByID)
	tmRouter.Get("/", h.GetAll)
	tmRouter.Post("/", h.Create)
	tmRouter.Put("/:id", h.Update)
	tmRouter.Delete("/:id", h.Delete)
}
