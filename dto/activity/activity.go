package activitydto

type ActivityRequest struct {
	Title string `json:"title" validate:"required"`
	Email string `json:"email" validate:"required"`
}

type ActivityUpdate struct {
	Title string `json:"title"`
}
