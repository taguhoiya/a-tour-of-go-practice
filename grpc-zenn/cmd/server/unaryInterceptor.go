package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
)

func myUnaryServerInterceptor1(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("[pre] my unary server interceptor 1: ", info.FullMethod) // ハンドラの前に割り込ませる前処理
	res, err := handler(ctx, req)                                         // 本来の処理
	log.Println("[post] my unary server interceptor 1: ", m)              // ハンドラの後に割り込ませる後処理
	return res, err
}
