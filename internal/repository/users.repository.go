package repository

import (
	"context"

	"github.com/darwinuzcategui1973/qrtomfoto-produ/internal/entity"
)

const (
	qryInsertUser = `
		insert into USERS ( name, password)
		values ( ?, ?);`

	qryGetUserByName = `
		select
			id,
			name,
			password
		from USERS
		where name = ?;`

	qryInsertUserRole = `
		insert into USER_ROLES (user_id, role_id) values (:user_id, :role_id);`

	qryRemoveUserRole = `
		delete from USER_ROLES where user_id = :user_id and role_id = :role_id;`

	qryListaDeUsuario = `
		select id,name from USERS;`

	qryLista = `
		select id,nombre as name from BdEmpleado.empleados e ;`

	//from BdEmpleado.empleados e  ;
	// aqui voy un query prueba
	qryListap = `
		select pcode as id, preferencia as descripcion ,pdescribe as name from d3xd_dbase.productos p ;`
)

func (r *repo) SaveUser(ctx context.Context, name, password string) error {
	_, err := r.db.ExecContext(ctx, qryInsertUser, name, password)
	return err
}

func (r *repo) GetUserByName(ctx context.Context, name string) (*entity.User, error) {
	u := &entity.User{}
	err := r.db.GetContext(ctx, u, qryGetUserByName, name)
	if err != nil {
		return nil, err
	}

	//return u, nil
	return u, err
}

func (r *repo) SaveUserRole(ctx context.Context, userID, roleID int64) error {
	data := entity.UserRole{
		UserID: userID,
		RoleID: roleID,
	}

	_, err := r.db.NamedExecContext(ctx, qryInsertUserRole, data)
	return err
}

func (r *repo) RemoveUserRole(ctx context.Context, userID, roleID int64) error {
	data := entity.UserRole{
		UserID: userID,
		RoleID: roleID,
	}

	_, err := r.db.NamedExecContext(ctx, qryRemoveUserRole, data)

	return err
}

func (r *repo) GetUserRoles(ctx context.Context, userID int64) ([]entity.UserRole, error) {
	roles := []entity.UserRole{}

	err := r.db.SelectContext(ctx, &roles, "select user_id, role_id from USER_ROLES where user_id = ?", userID)
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (r *repo) GetListaUsuario(ctx context.Context) ([]entity.User, error) {
	users := []entity.User{}
	//err := r.db.SelectContext(ctx, &users, qryListaDeUsuario)
	err := r.db.SelectContext(ctx, &users, qryListap)
	if err != nil {
		return nil, err
	}

	return users, nil

}

func (r *repo) GetLista(ctx context.Context) ([]entity.User, error) {
	users := []entity.User{}
	//err := r.db.SelectContext(ctx, &users, qryLista)
	err := r.db.SelectContext(ctx, &users, qryListap)
	if err != nil {
		return nil, err
	}

	return users, nil

}
