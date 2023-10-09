package modules

import (
	"context"
	"errors"
	"log"

	usrv1 "buf.build/gen/go/digibear/digibear/protocolbuffers/go/usr/v1"
	"connectrpc.com/connect"
	pkgerr "github.com/digibearapp/digibear-api/pkg/err"
	pkgtypes "github.com/digibearapp/digibear-api/pkg/types"
	"github.com/google/uuid"
)

// UsrServer represents the gRPC server that provides (account, profile) services.
type UsrServer struct {
	Store UsrStore
}

// NewUsrServer creates a new UsrServer.
func NewUsrServer(store UsrStore) *UsrServer {
	return &UsrServer{
		Store: store,
	}
}

// CreateProfile is a unary RPC that creates a new profile.
func (s *UsrServer) CreateProfile(
	ctx context.Context,
	req *connect.Request[usrv1.CreateProfileRequest],
) (*connect.Response[usrv1.CreateProfileResponse], error) {
	profile := req.Msg.GetProfile()
	log.Printf("Received a CreateProfile RPC request with ID : %s", profile.Id)
	res, err := s.Store.CreateProfile(profile)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(&usrv1.CreateProfileResponse{
		Id: res.String(),
	}), nil
}

// GetProfile is a unary RPC that returns a profile.
func (s *UsrServer) GetProfile(
	ctx context.Context,
	req *connect.Request[usrv1.GetProfileRequest],
) (*connect.Response[usrv1.GetProfileResponse], error) {
	id, err := uuid.Parse(req.Msg.GetId())
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	log.Printf("Received a GetProfile RPC request with ID : %s", id)
	res, err := s.Store.GetProfile(id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(&usrv1.GetProfileResponse{
		Profile: res,
	}), nil
}

// CheckEmailInUse is a unary RPC that checks if an email is in use.
func (s *UsrServer) CheckEmailInUse(
	ctx context.Context,
	req *connect.Request[usrv1.CheckEmailInUseRequest],
) (*connect.Response[usrv1.CheckEmailInUseResponse], error) {
	log.Printf("Received a CheckEmailInUse RPC request with email : %s", req.Msg.GetEmail())
	email, err := pkgtypes.NewEmail(req.Msg.GetEmail())
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	log.Printf("Received a CheckEmailInUse RPC request with email : %s", email.String)
	res, err := s.Store.CheckEmailInUse(*email)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(&usrv1.CheckEmailInUseResponse{
		IsEmailInUse: res,
	}), nil
}

// GetAccountForProfileSelection is a unary RPC that returns a profile and all the necessary information to select a profile.
func (s *UsrServer) GetAccountForProfileSelection(
	ctx context.Context,
	req *connect.Request[usrv1.GetAccountForProfileSelectionRequest],
) (*connect.Response[usrv1.GetAccountForProfileSelectionResponse], error) {
	id, err := uuid.Parse(req.Msg.GetId())
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	log.Printf("Received a GetAccountForProfileSelection RPC request with ID : %s", id)
	res, err := s.Store.GetAccountForProfileSelection(id)
	if err != nil {
		log.Printf("Error : %s", err)
		var pkgErr *pkgerr.Error
		if errors.As(err, &pkgErr) {
			if pkgErr.Status == pkgerr.NotFound {
				return nil, connect.NewError(connect.CodeNotFound, err)
			}
		}
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(&usrv1.GetAccountForProfileSelectionResponse{
		AccountForProfileSelection: res,
	}), nil
}

// UpdateProfile is a unary RPC that updates a profile.
func (s *UsrServer) UpdateProfile(
	ctx context.Context,
	req *connect.Request[usrv1.UpdateProfileRequest],
) (*connect.Response[usrv1.UpdateProfileResponse], error) {
	profile := req.Msg.GetProfile()
	log.Printf("Received an UpdateProfile RPC request with ID : %s", profile.Id)
	res, err := s.Store.UpdateProfile(profile)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(&usrv1.UpdateProfileResponse{
		Id: res.String(),
	}), nil
}

// DeleteProfile is a unary RPC that deletes a profile.
func (s *UsrServer) DeleteProfile(
	ctx context.Context,
	req *connect.Request[usrv1.DeleteProfileRequest],
) (*connect.Response[usrv1.DeleteProfileResponse], error) {
	id, err := uuid.Parse(req.Msg.GetId())
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	log.Printf("Received a DeleteProfile RPC request with ID : %s", id)
	res, err := s.Store.DeleteProfile(id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(&usrv1.DeleteProfileResponse{
		Id: res.String(),
	}), nil
}
