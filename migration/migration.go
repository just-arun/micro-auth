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

	pDb := boot.PostgresDB(env.DB.Uri)

	// register enums
	sql := model.RegisterUserType()
	_ = pDb.Exec(sql).Error
	sql = model.RegisterTokenPlacement()
	_ = pDb.Exec(sql).Error

	err := pDb.AutoMigrate(
		&model.User{},
		&model.App{},
		&model.Profile{},
		&model.Role{},
		&model.Access{},
		&model.General{},
		&model.ServiceMap{},
		&model.Mail{},
	)

	if err != nil {
		panic(err)
	}

	user := &model.User{
		Email:    env.Admin.Email,
		UserName: env.Admin.UserName,
		Password: env.Admin.Password,
		Roles: []model.Role{},
		Apps: []model.App{},
	}

	// registering super user user
	_, _ = service.User().CreateOne(pDb, user)

	generalData := &model.General{
		CanLogin:                env.General.CanLogin,
		CanRegister:             env.General.CanRegister,
		HttpOnlyCookie:          env.General.HttpOnlyCookie,
		AccessTokenExpiryTime:   env.General.AccessTokenExpiryTime,
		RefreshTokenExpiryTime:  env.General.RefreshTOkenExpiryTime,
		OrganizationEmailDomain: env.General.OrganizationEmailDomain,
	}

	fmt.Println(generalData)

	_ = service.General().Create(pDb, generalData)

	for _, v := range env.ServiceMap {
		_ = service.ServiceMap().Add(pDb, &v)
	}
}
