package server

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
)

func (s *Server) ConfigHttpRoutes(ctx context.Context) {
	s.echoServer.Use(recoverMiddleware)

	s.echoServer.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogMethod: true,
		LogStatus: true,
		LogError:  true,
		LogValuesFunc: func(_ echo.Context, values middleware.RequestLoggerValues) error {
			if values.Error != nil {
				slog.Error("echo request logger", "method", values.Method, "uri", values.URI, "status", values.Status)
			}
			return nil
		},
	}))

	s.echoServer.HideBanner = true
	s.echoServer.HidePort = false

	grpcSkipper := func(c echo.Context) bool {
		// Skip grpc and webhook calls.
		return false
	}

	s.echoServer.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper: grpcSkipper,
		Timeout: 0, // unlimited
	}))
	s.echoServer.Use(middleware.RateLimiterWithConfig(middleware.RateLimiterConfig{
		Skipper: grpcSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{Rate: 30, Burst: 60, ExpiresIn: 3 * time.Minute},
		),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			id := ctx.RealIP()
			return id, nil
		},
		ErrorHandler: func(context echo.Context, _ error) error {
			return context.JSON(http.StatusForbidden, nil)
		},
		DenyHandler: func(context echo.Context, _ string, _ error) error {
			return context.JSON(http.StatusTooManyRequests, nil)
		},
	}))

	s.echoServer.Any("/*", echo.WrapHandler(s.grpcMux))
}

func recoverMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			if r := recover(); r != nil {
				err, ok := r.(error)
				if !ok {
					err = errors.Errorf("%v", r)
				}

				c.Error(err)
			}
		}()
		return next(c)
	}
}
