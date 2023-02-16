package server

import (
	"project/internal/model"
)

func (s *Server) SignUp(u *model.User) *model.AuthRespond {
	res := &model.AuthRespond{
		Valid: s.ValidateUser(u),
		Exist: s.UserExist(u),
	}
	switch {
	case res.Exist:
		res.Message = "This login is already taken"
		return res
	case !res.Valid:
		res.Message = "Invalid login or password"
		return res
	default:
		s.NewUser(u)
		res.Message = "Success"
		return res
	}

}
func (s *Server) SignIn(u *model.User) *model.AuthRespond {
	res := &model.AuthRespond{
		Valid: s.ValidateUser(u),
		Exist: s.UserExist(u),
	}
	if res.Exist {
		switch {
		case s.UserAccess(u):
			res.Message = "Access"
		default:
			res.Message = "Incorrect login or password"
		}
	}
	return res
}
