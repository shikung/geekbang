// +build wireinject

package main

import "github.com/google/wire"

func InitializeShop() shop {
	wire.Build(NewS, NewA, NewB)
	return shop{}
}
