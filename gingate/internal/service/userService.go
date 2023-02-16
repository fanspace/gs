package service

import (
	log "gingate/core"
	"gingate/internal/model"
)

func TestXorm() (*model.User, error) {
	user := new(model.User)
	_, err := log.Orm.ID(1).Get(user)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return user, nil
}
