package web

type CategoryCreateRequest struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
