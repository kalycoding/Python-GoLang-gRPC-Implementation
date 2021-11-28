package grpc_go

func MsgToUserResponse(in *UserRequest) *UserResponse {
	return &UserResponse{
		ID:        in.ID,
		Username:  in.Username,
		Firstname: in.Firstname,
		Lastname:  in.Lastname,
		DOB:       in.DOB,
		Age:       in.Age,
		IsAdmin:   in.IsAdmin,
	}
}
