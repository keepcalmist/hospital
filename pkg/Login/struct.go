package Login

type Login struct {
	Name     string `json:"name" schema:"name, required"`
	Password string `json:"password" schema:"password, required"`
}
