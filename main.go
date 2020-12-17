package main

import (
	"github.com/hollson/goddd/crosscutting"
	"github.com/hollson/goddd/presentation"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"
)

func main() {
	crosscutting.StartUp()
	g := new(errgroup.Group)
	g.Go(func() error {
		router := presentation.InitRouter()
		return router.Run(":" + viper.GetString("APP_PORT"))
	})

	// g.Go(func() error {
	// 	err := presentation.InitRPC()
	// 	return err
	// })

	if err := g.Wait(); err != nil {
		panic(err)
	}
}
