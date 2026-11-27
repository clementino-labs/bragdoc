package mcp

import (
	"context"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/vagnerclementino/bragdoc/internal/domain"
)

// handleUserGet retrieves a user profile by ID.
func (s *Server) handleUserGet(ctx context.Context, _ *mcp.CallToolRequest, params UserGetParams) (*mcp.CallToolResult, any, error) {
	user, err := s.userService.GetByID(ctx, params.ID)
	if err != nil {
		return toolError(err), nil, nil
	}

	return marshalResult(toUserResponse(user))
}

// handleUserList lists all users.
func (s *Server) handleUserList(ctx context.Context, _ *mcp.CallToolRequest, _ UserListParams) (*mcp.CallToolResult, any, error) {
	users, err := s.userService.List(ctx)
	if err != nil {
		return toolError(err), nil, nil
	}

	responses := make([]UserResponse, 0, len(users))
	for _, u := range users {
		responses = append(responses, toUserResponse(u))
	}

	return marshalResult(responses)
}

// handleUserGetByEmail retrieves a user profile by email address.
func (s *Server) handleUserGetByEmail(ctx context.Context, _ *mcp.CallToolRequest, params UserGetByEmailParams) (*mcp.CallToolResult, any, error) {
	user, err := s.userService.GetByEmail(ctx, params.Email)
	if err != nil {
		return toolError(err), nil, nil
	}

	return marshalResult(toUserResponse(user))
}

// toUserResponse converts a domain.User to a UserResponse.
func toUserResponse(u *domain.User) UserResponse {
	return UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		JobTitle:  u.JobTitle,
		Company:   u.Company,
		Locale:    string(u.Locale),
		CreatedAt: u.CreatedAt.UTC().Format(time.RFC3339),
		UpdatedAt: u.UpdatedAt.UTC().Format(time.RFC3339),
	}
}
