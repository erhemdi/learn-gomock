package user

type User struct {
	ID   int64
	Name string
}

// Redis
type UserRedisIface interface {
	GetData(id int64) (User, error)
}

type UserRedis struct{}

func (u *UserRedis) GetData() (User, error) {
	return User{}, nil
}

// EOF Redis

// DB
type UserDBIface interface {
	GetData(id int64) (User, error)
}

type UserDB struct{}

func (u *UserDB) GetData() (User, error) {
	return User{}, nil
}

// EOF DB

// Service
type UserServiceIface interface {
	GetData(userID int64) (userData User, err error)
}

type UserService struct {
	Redis UserRedisIface
	DB    UserDBIface
}

func (u *UserService) GetData(userID int64) (userData User, err error) {
	userData, err = u.Redis.GetData(userID)
	if err != nil {
		return userData, err
	}

	if userData.ID > 0 {
		return userData, nil
	}

	userData, err = u.DB.GetData(userID)
	if err != nil {
		return userData, err
	}

	return userData, nil
}

// EOF Service
