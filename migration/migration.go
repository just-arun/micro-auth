package boot

import (
	"fmt"

	"github.com/just-arun/micro-auth/boot"
	"github.com/just-arun/micro-auth/model"
	"github.com/just-arun/micro-auth/service"
	"github.com/just-arun/micro-auth/util"
)

func Run(context, envi string) {

	env := &model.Env{}
	util.GetEnv(fmt.Sprintf(".env.%v", envi), context, &env)

	seed := &model.Seed{}
	util.GetEnv(fmt.Sprintf(".seed.%v", envi), context, &seed)

	fmt.Println(seed)

	pDb := boot.PostgresDB(env.DB.Uri)
	err := pDb.AutoMigrate(
		&model.User{},
		&model.App{},
		&model.Profile{},
		&model.Role{},
		&model.Access{},
	)

	if err != nil {
		panic(err)
	}

	user := &model.User{
		Email:    seed.Admin.Email,
		UserName: seed.Admin.UserName,
		Password: seed.Admin.Password,
	}

	userID, _ := service.User().CreateOne(pDb, user)

	fmt.Println("User ID: ", userID)
}
