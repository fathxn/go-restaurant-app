package restaurant

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"go-restaurant-app/internal/model"
	"go-restaurant-app/internal/model/constant"
	"go-restaurant-app/internal/repository/menu"
	"go-restaurant-app/internal/repository/order"
	"go-restaurant-app/internal/repository/user"
	"go-restaurant-app/internal/tracing"
)

type restaurantUsecase struct {
	menuRepo  menu.Repository
	orderRepo order.Repository
	userRepo  user.Repository
}

func GetUsecase(menuRepo menu.Repository, orderRepo order.Repository, userRepo user.Repository) Usecase {
	return &restaurantUsecase{
		menuRepo:  menuRepo,
		orderRepo: orderRepo,
		userRepo:  userRepo,
	}
}

func (r *restaurantUsecase) GetMenuList(ctx context.Context, menuType string) ([]model.MenuItem, error) {
	ctx, span := tracing.CreateSpan(ctx, "GetMenuList")
	defer span.End()
	return r.menuRepo.GetMenuList(ctx, menuType)
}

func (r *restaurantUsecase) Order(ctx context.Context, request model.OrderMenuRequest) (model.Order, error) {
	ctx, span := tracing.CreateSpan(ctx, "Order")
	defer span.End()

	productOrderData := make([]model.ProductOrder, len(request.OrderProducts))
	for i, orderProduct := range request.OrderProducts {
		menuData, err := r.menuRepo.GetMenu(ctx, orderProduct.OrderCode)
		if err != nil {
			return model.Order{}, err
		}

		productOrderData[i] = model.ProductOrder{
			ID:         uuid.New().String(),
			OrderCode:  menuData.OrderCode,
			Quantity:   orderProduct.Quantity,
			TotalPrice: menuData.Price * int64(orderProduct.Quantity),
			Status:     constant.ProductOrderStatusPreparing,
		}
	}

	orderData := model.Order{
		ID:            uuid.New().String(),
		UserID:        request.UserID,
		Status:        constant.OrderStatusProcessed,
		ProductOrders: productOrderData,
		ReferenceID:   request.ReferenceID,
	}

	createdOrderData, err := r.orderRepo.CreateOrder(ctx, orderData)
	if err != nil {
		return model.Order{}, err
	}
	return createdOrderData, nil
}
func (r *restaurantUsecase) GetOrderInfo(ctx context.Context, request model.GetOrderInfoRequest) (model.Order, error) {
	ctx, span := tracing.CreateSpan(ctx, "GetOrderInfo")
	defer span.End()

	orderData, err := r.orderRepo.GetOrderInfo(ctx, request.OrderID)
	if err != nil {
		return orderData, err
	}

	if orderData.UserID != request.UserID {
		return model.Order{}, errors.New("unauthorized")
	}

	return orderData, nil
}

func (r *restaurantUsecase) RegisterUser(ctx context.Context, request model.RegisterRequest) (model.User, error) {
	ctx, span := tracing.CreateSpan(ctx, "RegisterUser")
	defer span.End()

	registeredUser, err := r.userRepo.CheckRegistered(ctx, request.Username)
	if err != nil {
		return model.User{}, err
	}
	if registeredUser {
		return model.User{}, errors.New("username already registered")
	}

	passwordHash, err := r.userRepo.GenerateUserHash(ctx, request.Password)
	if err != nil {
		return model.User{}, err
	}

	userData, err := r.userRepo.RegisterUser(ctx, model.User{
		ID:       uuid.New().String(),
		Username: request.Username,
		Hash:     passwordHash,
	})

	if err != nil {
		return model.User{}, err
	}

	return userData, nil
}

func (r *restaurantUsecase) LoginUser(ctx context.Context, request model.LoginRequest) (model.UserSession, error) {
	ctx, span := tracing.CreateSpan(ctx, "LoginUser")
	defer span.End()

	userData, err := r.userRepo.GetUserData(ctx, request.Username)
	if err != nil {
		return model.UserSession{}, err
	}

	verified, err := r.userRepo.VerifyLogin(ctx, request.Username, request.Password, userData)
	if err != nil {
		return model.UserSession{}, err
	}
	if !verified {
		return model.UserSession{}, errors.New("invalid username or password")
	}

	userSession, err := r.userRepo.CreateUserSession(ctx, userData.ID)
	if err != nil {
		return model.UserSession{}, err
	}

	return userSession, nil
}

func (r *restaurantUsecase) CheckSession(ctx context.Context, data model.UserSession) (userID string, err error) {
	ctx, span := tracing.CreateSpan(ctx, "CheckSession")
	defer span.End()

	userID, err = r.userRepo.CheckSession(ctx, data)
	if err != nil {
		return "", err
	}

	return userID, nil
}
