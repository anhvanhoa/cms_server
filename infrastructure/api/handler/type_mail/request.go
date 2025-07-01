package typeMail

type createReq struct {
	Name string `validate:"required"`
}

type updateReq struct {
	Name string `validate:"required"`
}
