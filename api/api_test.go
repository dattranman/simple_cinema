package api

import (
	"testing"

	"github.com/dattranman/simple_cinema/app"
	"github.com/gin-gonic/gin"
)

func TestAPI_Run(t *testing.T) {
	type fields struct {
		App         *app.App
		BaseRouters *Routers
		RoomRouter  *gin.RouterGroup
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "success case",
			fields: fields{
				App:         &app.App{},
				BaseRouters: &Routers{},
				RoomRouter:  &gin.RouterGroup{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &API{
				App:         tt.fields.App,
				BaseRouters: tt.fields.BaseRouters,
				RoomRouter:  tt.fields.RoomRouter,
			}
			if err := a.Run(); (err != nil) != tt.wantErr {
				t.Errorf("API.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
