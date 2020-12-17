package proxy

import (
    "log"
    "net"

    "github.com/hollson/goddd/proxy/middleware"
    "github.com/spf13/viper"
    "google.golang.org/grpc"
)

// func InitRPC() error {
//     if viper.GetString("APP_MODE") == "prod" {
//         middleware.IsProdMod = true
//     }
//     port := viper.GetString("GRPC_PORT")
//     port = fmt.Sprintf(":%s", port)
//     lis, err := net.Listen("tcp", port)
//     if err != nil {
//         log.Fatalf("failed to listen: %v", err)
//     }
//     // 添加拦截器，用以日志\认证等
//     var opts []grpc.ServerOption
//     opts = append(opts, grpc.UnaryInterceptor(middleware.GrpcInterceptor))
//     grpcServer := grpc.NewServer(opts...)
//     // envoyAuth.RegisterAuthorizationServer(grpcServer, &application.AuthServiceApp{})
//     // base_srv.RegisterBaseSvcServer(grpcServer, &application.BaseServiceApp{})
//     println("gRPC Srv start ...", port)
//     if err := grpcServer.Serve(lis); err != nil {
//         return err
//     }
//
//     return nil
// }

type rpcServer struct{}

func NewGrpcServer() *rpcServer {
    return &rpcServer{}

}

func (g *rpcServer) Run(port ...string) error {
    var p string = ":8081"
    if len(port) > 0 {
        p = port[0]
    }

    // 创建监听器
    if viper.GetString("APP_MODE") == "prod" {
        middleware.IsProdMod = true
    }
    lis, err := net.Listen("tcp", p)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    // 添加拦截器，用以日志\认证等
    var opts []grpc.ServerOption
    opts = append(opts, grpc.UnaryInterceptor(middleware.GrpcInterceptor))
    grpcServer := grpc.NewServer(opts...)
    // envoyAuth.RegisterAuthorizationServer(grpcServer, &application.AuthServiceApp{})
    // base_srv.RegisterBaseSvcServer(grpcServer, &application.BaseServiceApp{})
    log.Printf(" Rpc server is running on %s", p)
    return grpcServer.Serve(lis)
}
