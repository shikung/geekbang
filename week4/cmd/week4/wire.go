// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import "week4/internal/biz"

func initApp() {
	biz.NewUserCase()
}
