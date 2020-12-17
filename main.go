package main

import (
    "github.com/hollson/goddd/config"
    "github.com/hollson/goddd/proxy"
    "golang.org/x/sync/errgroup"
)

func main() {
    config.StartUp()

    var group errgroup.Group
    group.Go(func() error {
        return proxy.NewGinSerer().Run(":8080")
    })

    group.Go(func() error {
        return proxy.NewGrpcServer().Run(":8082")
    })

    if err := group.Wait(); err != nil {
        panic(err)
    }
}
