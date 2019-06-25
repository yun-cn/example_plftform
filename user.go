package noonde

// UserData data represents all the fields we keep about user.
type UserData struct {
}

// User is a user of the noonde platform
type User interface {
	GetByID(id int) (UserData, error)
	StoreByID(id int, data *UserData) error
}
