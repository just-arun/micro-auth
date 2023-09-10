package boot

import (
	"fmt"
	"log"

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
		&model.App{},
		&model.User{},
	)

	if err != nil {
		panic(err)
	}

	userID, _ := service.User().CreateOne(pDb, &model.User{
		Email: seed.Admin.Email,
		UserName: seed.Admin.UserName,
		Password: seed.Admin.Password,
	})

	fmt.Println("User ID: ", userID)

	appID, _ := service.App().CreateOne(pDb, &model.App{
		Name: "App 1",
	})

	fmt.Println("App ID: ", appID)

	err = service.User().AddApp(pDb, 21, 1)

	if err != nil {
		panic(err)
	}

	user, err := service.User().GetOne(pDb, &model.User{ID: 23},)

	if err != nil {
		log.Fatalf("ERROR: %v; \n", err.Error())
	}

	fmt.Println("User Data", user)
}
