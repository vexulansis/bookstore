package server

import (
	. "project/internal"
	"project/internal/model"
)

// разобраться с транзакциями

func (s *Server) NewUser(u *model.User) {
	query := `--sql do$$
	declare
	users_user_id integer;
	begin
	insert into users default values returning user_id into users_user_id;
	insert into auth (user_id, login, password) values (users_user_id, $1, $2);
	end$$;`
	_, err := s.Store.DB.Query(query, u.Login, u.Password)
	if err != nil {
		ErrorLog(err.Error())
	}
}
