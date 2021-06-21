package clients

type Client struct {
	ID    string `json:"id" param:"id" query:"id"`
	Name  string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`
}
