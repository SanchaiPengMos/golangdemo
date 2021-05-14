package route

// User
type User struct {
	Name  string `json:"name" xml:"name"`
	Email string `json:"email" xml:"email"`
}

type UserCheck struct {
	Id            uint   `sql:"primary_key" json:"id" form:"id" query:"id"`
	Email         string `json:"email" form:"email" query:"email"`
	Firstname     string `json:"firstname" form:"firstname" query:"firstname"`
	Tel           string `json:"tel" form:"tel" query:"tel"`
	Lastname      string `json:"lastname" form:"lastname" query:"lastname"`
	Birthday      string `json:"birthday" form:"birthday" query:"birthday"`
	Userpermis_id int    `json:"userpermis_id" form:"userpermis_id" query:"userpermis_id"`
}
