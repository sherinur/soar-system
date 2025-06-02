package grpcserver

import (
	"context"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func AuthInterceptor(secretKey string) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		// Skip auth for these methods
		if strings.Contains(info.FullMethod, "Login") ||
			strings.Contains(info.FullMethod, "Register") ||
			strings.Contains(info.FullMethod, "RefreshToken") ||
			strings.Contains(info.FullMethod, "GetAllUsers") ||
			strings.Contains(info.FullMethod, "ChangeUserRole") {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "metadata missing")
		}

		authHeader := md["authorization"]
		if len(authHeader) == 0 {
			return nil, status.Error(codes.Unauthenticated, "authorization missing")
		}

		tokenStr := strings.TrimPrefix(authHeader[0], "Bearer ")
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "invalid token")
		}

		if !token.Valid {
			return nil, status.Error(codes.Unauthenticated, "invalid token")
		}

		userID, ok := claims["user_id"].(float64)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "user_id missing in token")
		}

		role, ok := claims["role"].(string)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "role missing in token")
		}

		ctx = context.WithValue(ctx, "user_id", int64(userID))
		ctx = context.WithValue(ctx, "role", role)
		ctx = context.WithValue(ctx, "token", tokenStr)

		return handler(ctx, req)
	}
}

func loggingInterceptor(log *zap.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		start := time.Now()
		resp, err := handler(ctx, req)

		log.Info("gRPC request completed",
			zap.String("method", info.FullMethod),
			zap.Duration("duration", time.Since(start)),
		)

		log.Debug("Request details", zap.Any("request", req))

		return resp, err
	}
}

func errorInterceptor(log *zap.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		resp, err := handler(ctx, req)
		if err != nil {
			log.Error("gRPC request error",
				zap.Error(err),
			)
		}

		return resp, err
	}
}

func recoveryInterceptor(log *zap.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		defer func() {
			if r := recover(); r != nil {
				log.Error("gRPC service panic recovered",
					zap.Any("panic", r),
					zap.Stack("stack"),
				)
				// err = status.Errorf(codes.Internal, "internal server error")
			}
		}()

		return handler(ctx, req)
	}
}

// Deprecated
// // func otelInterceptor(log *zap.Logger) grpc.UnaryServerInterceptor {
// // 	return otelgrpc.UnaryServerInterceptor()
// // }
