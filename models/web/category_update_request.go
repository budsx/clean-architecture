package web

type CategoryUpdateRequest struct {
	Id int `json:"id"`
	Name string `json:"name"`
}
