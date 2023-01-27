package apiserver

import (
	"context"
	"errors"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func AuthInterceptor(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "basic")
	if err != nil {
		return nil, err
	}
	tokenInfo, err := parseToken(token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, " %v", err)
	}
	//使用context.WithValue添加了值后，可以用Value(key)方法获取值
	newCtx := context.WithValue(ctx, tokenInfo.ID, tokenInfo)
	//log.Println(newCtx.Value(tokenInfo.ID))
	return newCtx, nil
}

func parseToken(token string) (TokenInfo, error) {
	var tokenInfo TokenInfo
	if token == "grpc.auth.token" {
		tokenInfo.ID = "1"
		tokenInfo.Roles = []string{"admin"}
		return tokenInfo, nil
	}
	return tokenInfo, errors.New("Token无效: bearer " + token)
}

func userClaimFromToken(tokenInfo TokenInfo) string {
	return tokenInfo.ID
}

type TokenInfo struct {
	ID    string
	Roles []string
}
