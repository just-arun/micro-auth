package util

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
)

type serviceMap struct {
	Key   string `json:"key" gorm:"uniqueIndex"`
	Value string `json:"value" gorm:"uniqueIndex"`
	Auth  bool   `json:"bool" gorm:"default:false"`
}

func GetPath(e *echo.Echo) []serviceMap {
	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		panic(err)
	}

	dat := []serviceMap{}

	da := []map[string]string{}

	err = json.Unmarshal(data, &da)
	if err != nil {
		panic(err)
	}

	for _, v := range da {
		val := strings.ReplaceAll(v["name"], "github.com/just-arun/micro-auth/", "")
		val = strings.ReplaceAll(val, "handler.", "auth.")
		for i := 0; i < 10; i++ {
			val = strings.ReplaceAll(val, fmt.Sprintf(".func%v", i+1), "")
		}
		wa := &serviceMap{
			Key:   val,
			Value: strings.ReplaceAll(val, "auth ", ""),
			Auth:  true,
		}
		fmt.Println(wa)
		dat = append(dat, *wa)
	}

	jData, err := json.Marshal(dat)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jData))

	err = os.WriteFile("routes.json", jData, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println(da)
	return dat
}
