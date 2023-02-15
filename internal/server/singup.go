package server

import "project/internal/model"

func (s *Server) SignUp(u *model.User) *model.Respond {
	res := new(model.Respond)
	req := &model.RequestUser{
		By:   "login",
		Data: u.Login,
	}
	if s.SearchUser(req) != nil {
		res.UserExist = true
	}
	res.UserValid, res.Message = u.Valid()
	return res
}
