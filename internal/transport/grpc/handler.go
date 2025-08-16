package grpc

import (
	"context"

	userpb "github.com/AleksKAG/project-protos/proto/user"
	"github.com/AleksKAG/users-service/internal/user"
)

type Handler struct {
	svc *user.Service
	userpb.UnimplementedUserServiceServer
}

func NewHandler(svc *user.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	u, err := h.svc.CreateUser(req.Email)
	if err != nil {
		return nil, err
	}
	return &userpb.CreateUserResponse{
		User: &userpb.User{Id: u.ID, Email: u.Email},
	}, nil
}

func (h *Handler) GetUser(ctx context.Context, req *userpb.User) (*userpb.User, error) {
	u, err := h.svc.GetUserByID(req.Id)
	if err != nil {
		return nil, err
	}
	return &userpb.User{Id: u.ID, Email: u.Email}, nil
}

func (h *Handler) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	u, err := h.svc.UpdateUserByID(req.Id, &user.User{Email: req.Email})
	if err != nil {
		return nil, err
	}
	return &userpb.UpdateUserResponse{
		User: &userpb.User{Id: u.ID, Email: u.Email},
	}, nil
}

func (h *Handler) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	if err := h.svc.DeleteUserByID(req.Id); err != nil {
		return nil, err
	}
	return &userpb.DeleteUserResponse{Success: true}, nil
}

func (h *Handler) ListUsers(ctx context.Context, req *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	users, err := h.svc.GetAllUsers(req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}
	pbUsers := make([]*userpb.User, 0, len(users))
	for _, u := range users {
		pbUsers = append(pbUsers, &userpb.User{Id: u.ID, Email: u.Email})
	}
	return &userpb.ListUsersResponse{Users: pbUsers}, nil
}
