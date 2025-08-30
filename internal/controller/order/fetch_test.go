package httporder

import (
	"testing"

	"github.com/RakhimovAns/L0/internal/controller/order/mocks"
	"github.com/gofiber/fiber/v3"
	"github.com/valyala/fasthttp"
)

func TestHandler_fetchByID(t *testing.T) {
	type args struct {
		ctx fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{{
		name: "success",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := mocks.NewService(t)
			h := &Handler{
				service: service,
			}
			if err := h.fetchByID(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("fetchByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_ping(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "success",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := mocks.NewService(t)

			service.On("Ping").Return(nil)

			h := &Handler{
				service: service,
			}
			app := fiber.New()
			ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
			defer app.ReleaseCtx(ctx)

			if err := h.ping(ctx); (err != nil) != tt.wantErr {
				t.Errorf("ping() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
