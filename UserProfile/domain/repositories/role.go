package repositories

import (
	"context"
	"fmt"
	"knovel/userprofile/domain/common"
	"knovel/userprofile/domain/entities"
)

type RoleRepository interface {
	GetRolesOfUser(ctx context.Context, userId string) ([]*entities.Role, error)
	GetPermissionByRoles(ctx context.Context, roleIds []int) ([]*entities.Permission, error)
	GetPermissionOfService(ctx context.Context, serviceName string) ([]*entities.Permission, error)
}

type RoleRepositoryInstance struct {
	dbContext       common.DbContext
	roleTable       string
	userRoleTable   string
	permissionTable string
	rbacTable       string
}

var _ RoleRepository = (*RoleRepositoryInstance)(nil)

func NewRbacRepository(dbContext common.DbContext) RoleRepository {
	return &RoleRepositoryInstance{
		dbContext:       dbContext,
		roleTable:       "role",
		userRoleTable:   "userrole",
		permissionTable: "permission",
		rbacTable:       "rbac",
	}
}

func (repo *RoleRepositoryInstance) GetRolesOfUser(ctx context.Context, userId string) ([]*entities.Role, error) {
	res := make([]*entities.Role, 0)

	//join role and userrole
	query := "SELECT %s.id, name ,%s.deleted_at FROM %s INNER JOIN %s on %s.id = %s.role_id WHERE %s.user_id = $1" //
	query = fmt.Sprintf(query, repo.roleTable, repo.roleTable, repo.roleTable, repo.userRoleTable, repo.roleTable, repo.userRoleTable, repo.userRoleTable)
	rows, err := repo.dbContext.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		item := &entities.Role{}
		err = rows.Scan(&item.Id, &item.Name, &item.DeletedAt)
		if err != nil {
			return nil, err
		}
		res = append(res, item)
	}
	return res, nil
}

func (repo *RoleRepositoryInstance) GetPermissionByRoles(ctx context.Context, roleIds []int) ([]*entities.Permission, error) {
	res := make([]*entities.Permission, 0)
	//join permission & rbac
	query := "SELECT  name ,%s.deleted_at FROM %s INNER JOIN %s on %s.id = %s.permission_id WHERE %s.role_id = ANY($1)"
	query = fmt.Sprintf(query, repo.permissionTable, repo.permissionTable, repo.rbacTable, repo.permissionTable, repo.rbacTable, repo.rbacTable)
	rows, err := repo.dbContext.QueryContext(ctx, query, roleIds)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		item := &entities.Permission{}
		err = rows.Scan(&item.Name, &item.DeletedAt)
		if err != nil {
			return nil, err
		}
		res = append(res, item)
	}
	return res, nil
}

// GetPermissionOfService implements RoleRepository.
func (repo *RoleRepositoryInstance) GetPermissionOfService(ctx context.Context, serviceName string) ([]*entities.Permission, error) {
	res := make([]*entities.Permission, 0)
	query := "SELECT name, func_name FROM %s  WHERE service_name = $1 and deleted_at = NULL"
	query = fmt.Sprintf(query, repo.permissionTable)
	rows, err := repo.dbContext.QueryContext(ctx, query, serviceName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		item := &entities.Permission{}
		err = rows.Scan(&item.Name, &item.FuncName, &item.DeletedAt)
		if err != nil {
			return nil, err
		}
		res = append(res, item)
	}
	return res, nil
}
