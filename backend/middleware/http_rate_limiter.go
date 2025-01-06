package middleware

import (
	"net/http"
	"sync"

	// 你的 protobuf 定义

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

// 使用 map 存储每个方法的限流器
var limiters sync.Map

func rateLimitMiddleware(grpcServer *grpc.Server, next http.Handler) echo.MiddlewareFunc {
	// methodInfos := make(map[string]*descriptorpb.MethodDescriptorProto)

	// services := grpcServer.GetServiceInfo()

	// for name, info := range services {
	// 	methods := info.Methods

	// 	for _, method := range methods {
	// 		methodInfos[name+"/"+method.Name] = method.Name
	// 	}
	// }

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			return next(c)
		}
	}

	// return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	// Find the matching route
	// 	var method *runtime.MethodDescriptor
	// 	// ... (查找匹配路由的代码，与之前的版本相同)

	// 	if method == nil {
	// 		log.Println("No matching route found for", r.URL.Path, r.Method)
	// 		next.ServeHTTP(w, r)
	// 		return
	// 	}

	// 	methodDescriptorProto := method.Method

	// 	opts := methodDescriptorProto.GetOptions()
	// 	me, ok := proto.GetExtension(opts, v1pb.E_MethodExtend).(*v1pb.MethodExtend)
	// 	if !ok {
	// 		next.ServeHTTP(w, r)
	// 		return
	// 	}

	// 	if me.RateLimit <= 0 {
	// 		next.ServeHTTP(w, r)
	// 		return
	// 	}

	// 	// 使用方法名作为 key
	// 	methodName := methodDescriptorProto.GetName()

	// 	// 从 map 中获取限流器，如果不存在则创建
	// 	limiter, ok := limiters.Load(methodName)
	// 	if !ok {
	// 		newLimiter := rate.NewLimiter(rate.Limit(rateLimit), int(rateLimit))
	// 		limiter, _ = limiters.LoadOrStore(methodName, newLimiter)
	// 	}

	// 	// 进行限流
	// 	if !limiter.(*rate.Limiter).Allow() {
	// 		http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
	// 		return
	// 	}

	// 	next.ServeHTTP(w, r)
	// })
}
