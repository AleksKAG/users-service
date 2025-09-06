package grpc

import (
	"context"
	"log"

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
	if err := ctx.Err(); err != nil {
		log.Printf("CreateUser: context error: %v", err)
		return nil, err
	}

	u, err := h.svc.CreateUser(req.Email)
	if err != nil {
		log.Printf("CreateUser: failed to create user with email %s: %v", req.Email, err)
		return nil, err
	}

	return &userpb.CreateUserResponse{
		User: &userpb.User{
			Id:    uint32(u.ID), // uint -> uint32
			Email: u.Email,
		},
	}, nil
}

func (h *Handler) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	if err := ctx.Err(); err != nil {
		log.Printf("GetUser: context error: %v", err)
		return nil, err
	}

	log.Printf("GetUser: called with ID: %d", req.Id)
	u, err := h.svc.GetUser(uint(req.Id)) // req.Id (uint32) -> uint
	if err != nil {
		log.Printf("GetUser: error for ID %d: %v", req.Id, err)
		return nil, err
	}

	return &userpb.GetUserResponse{
		User: &userpb.User{
			Id:    uint32(u.ID), // uint -> uint32
			Email: u.Email,
		},
	}, nil
}

func (h *Handler) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	if err := ctx.Err(); err != nil {
		log.Printf("UpdateUser: context error: %v", err)
		return nil, err
	}

	u, err := h.svc.UpdateUser(uint(req.Id), req.Email) // req.Id (uint32) -> uint
	if err != nil {
		log.Printf("UpdateUser: failed to update user with ID %d: %v", req.Id, err)
		return nil, err
	}

	return &userpb.UpdateUserResponse{
		User: &userpb.User{
			Id:    uint32(u.ID), // uint -> uint32
			Email: u.Email,
		},
	}, nil
}

func (h *Handler) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	if err := ctx.Err(); err != nil {
		log.Printf("DeleteUser: context error: %v", err)
		return nil, err
	}

	if err := h.svc.DeleteUser(uint(req.Id)); err != nil { // req.Id (uint32) -> uint
		log.Printf("DeleteUser: failed to delete user with ID %d: %v", req.Id, err)
		return nil, err
	}

	return &userpb.DeleteUserResponse{Success: true}, nil
}

func (h *Handler) ListUsers(ctx context.Context, req *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	if err := ctx.Err(); err != nil {
		log.Printf("ListUsers: context error: %v", err)
		return nil, err
	}

	users, err := h.svc.ListUsers(int(req.Page), int(req.PageSize))
	if err != nil {
		log.Printf("ListUsers: failed to list users: %v", err)
		return nil, err
	}

	pbUsers := make([]*userpb.User, 0, len(users))
	for _, u := range users {
		pbUsers = append(pbUsers, &userpb.User{
			Id:    uint32(u.ID), // uint -> uint32
			Email: u.Email,
		})
	}

	return &userpb.ListUsersResponse{Users: pbUsers}, nil
}
