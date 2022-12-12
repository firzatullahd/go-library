package correlation

import (
	"context"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

const (
	CorrelationIDKey string = "X-Correlation-ID"
)

func Correlation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cid := c.Request().Header.Get(CorrelationIDKey)
		if len(cid) <= 0 {
			cid = uuid.New().String()
		}
		c.Set(CorrelationIDKey, cid)
		ctx := context.WithValue(c.Request().Context(), CorrelationIDKey, cid)
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}

func SetCorrelationIDToContext(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, CorrelationIDKey, id)
}

func GetCorrelationIDFromContext(ctx context.Context) string {
	id, ok := ctx.Value(CorrelationIDKey).(string)
	if !ok {
		return ""
	}
	return id
}
