package restaurant

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"go-restaurant-app/internal/mocks"
	"go-restaurant-app/internal/model"
	"go-restaurant-app/internal/model/constant"
	"go-restaurant-app/internal/repository/menu"
	"go-restaurant-app/internal/repository/order"
	"go-restaurant-app/internal/repository/user"
	"reflect"
	"testing"
)

func Test_restaurantUsecase_GetMenuList(t *testing.T) {
	type fields struct {
		menuRepo  menu.Repository
		orderRepo order.Repository
		userRepo  user.Repository
	}
	type args struct {
		ctx      context.Context
		menuType string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.MenuItem
		wantErr bool
	}{
		{
			name: "success get menu list",
			fields: fields{
				menuRepo: func() menu.Repository {
					ctrl := gomock.NewController(t)
					mock := mocks.NewMockMenuRepository(ctrl)

					mock.EXPECT().GetMenuList(gomock.Any(), string(constant.MenuTypeFood)).
						Times(1).
						Return([]model.MenuItem{
							{
								OrderCode: "nasduk",
								Name:      "Nasi Uduk",
								Price:     15000,
								Type:      constant.MenuTypeFood,
							},
						}, nil)
					return mock
				}(),
			},
			args: args{
				ctx:      context.Background(),
				menuType: string(constant.MenuTypeFood),
			},
			want: []model.MenuItem{
				{
					OrderCode: "nasduk",
					Name:      "Nasi Uduk",
					Price:     15000,
					Type:      constant.MenuTypeFood,
				},
			},
			wantErr: false,
		},

		{
			name: "fail get menu list",
			fields: fields{
				menuRepo: func() menu.Repository {
					ctrl := gomock.NewController(t)
					mock := mocks.NewMockMenuRepository(ctrl)

					mock.EXPECT().GetMenuList(gomock.Any(), string(constant.MenuTypeFood)).
						Times(1).
						Return(nil, errors.New("mock error"))
					return mock
				}(),
			},
			args: args{
				ctx:      context.Background(),
				menuType: string(constant.MenuTypeFood),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &restaurantUsecase{
				menuRepo:  tt.fields.menuRepo,
				orderRepo: tt.fields.orderRepo,
				userRepo:  tt.fields.userRepo,
			}
			got, err := r.GetMenuList(tt.args.ctx, tt.args.menuType)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMenuList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMenuList() got = %v, want %v", got, tt.want)
			}
		})
	}
}
