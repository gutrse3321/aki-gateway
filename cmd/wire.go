//+build wireinject

package main

import (
	"akigate"
	"github.com/google/wire"
	"github.com/gutrse3321/aki/pkg/app"
	"github.com/gutrse3321/aki/pkg/config"
	"github.com/gutrse3321/aki/pkg/log"
	"github.com/gutrse3321/aki/pkg/transports/http"
)

/**
 * @Author: Tomonori
 * @Date: 2020/6/18 17:17
 * @Title:
 * --- --- ---
 * @Desc:
 */

var wireSet = wire.NewSet(
	log.WireSet,
	config.WireSet,
	http.WireSet,
	akigate.WireSet,
)

func CreateApp(configPath string) (*app.Application, error) {
	panic(wire.Build(wireSet))
}
