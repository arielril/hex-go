package types

type Persistence interface {
	Create(model interface{}) Persistence
	Find(model interface{}, where ...interface{}) Persistence
	Update(model interface{}, where ...interface{}) Persistence
	Delete(model interface{}, where ...interface{}) Persistence
}

type UserModel interface {
	CreateUser(user *UserModel)
	FindUser(params map[string]interface{})
}
