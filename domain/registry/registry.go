package registry

import "go-rest-skeleton/domain/entity"

type entities struct {
	Entity interface{}
}

type table struct {
	Name interface{}
}

func CollectEntities() []entities {
	return []entities{
		{Entity: entity.Module{}},
		{Entity: entity.Permission{}},
		{Entity: entity.Role{}},
		{Entity: entity.RolePermission{}},
		{Entity: entity.StorageCategory{}},
		{Entity: entity.StorageFile{}},
		{Entity: entity.User{}},
		{Entity: entity.UserForgotPassword{}},
		{Entity: entity.UserLogin{}},
		{Entity: entity.UserPreference{}},
		{Entity: entity.UserRole{}},
	}
}

func CollectTableNames() []table {
	var module entity.Module
	var permission entity.Permission
	var role entity.Role
	var rolePermission entity.RolePermission
	var storageCategory entity.StorageCategory
	var storageFile entity.StorageFile
	var user entity.User
	var userForgotPassword entity.UserForgotPassword
	var userLogin entity.UserLogin
	var userPreference entity.UserPreference
	var userRole entity.UserRole

	return []table{
		{Name: module.TableName()},
		{Name: permission.TableName()},
		{Name: role.TableName()},
		{Name: rolePermission.TableName()},
		{Name: storageCategory.TableName()},
		{Name: storageFile.TableName()},
		{Name: user.TableName()},
		{Name: userForgotPassword.TableName()},
		{Name: userLogin.TableName()},
		{Name: userPreference.TableName()},
		{Name: userRole.TableName()},
	}
}
