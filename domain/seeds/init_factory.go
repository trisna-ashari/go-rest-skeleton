package seeds

import (
	"go-rest-skeleton/domain/entity"

	"github.com/jinzhu/gorm"

	"github.com/bxcodec/faker"

	"github.com/google/uuid"
)

func initFactory() []Seed {
	userFaker := entity.UserFaker{}
	_ = faker.FakeData(&userFaker)
	user := &entity.User{
		UUID:      uuid.New().String(),
		FirstName: userFaker.FirstName,
		LastName:  userFaker.LastName,
		Email:     "trisna.x2@gmail.com",
		Phone:     userFaker.Phone,
		Password:  "123456",
	}

	role := &entity.Role{
		UUID: uuid.New().String(),
		Name: "Super Administrator",
	}

	permissions := []*entity.Permission{
		{UUID: uuid.New().String(), ModuleKey: "user", PermissionKey: "read"},
		{UUID: uuid.New().String(), ModuleKey: "user", PermissionKey: "create"},
		{UUID: uuid.New().String(), ModuleKey: "user", PermissionKey: "update"},
		{UUID: uuid.New().String(), ModuleKey: "user", PermissionKey: "delete"},
		{UUID: uuid.New().String(), ModuleKey: "user", PermissionKey: "bulk_delete"},
		{UUID: uuid.New().String(), ModuleKey: "user", PermissionKey: "detail"},
		{UUID: uuid.New().String(), ModuleKey: "role", PermissionKey: "read"},
		{UUID: uuid.New().String(), ModuleKey: "role", PermissionKey: "create"},
		{UUID: uuid.New().String(), ModuleKey: "role", PermissionKey: "update"},
		{UUID: uuid.New().String(), ModuleKey: "role", PermissionKey: "delete"},
		{UUID: uuid.New().String(), ModuleKey: "role", PermissionKey: "bulk_delete"},
	}

	rolePermissions := make([]*entity.RolePermission, len(permissions))
	for i, permission := range permissions {
		rolePermissions[i] = &entity.RolePermission{
			UUID:           uuid.New().String(),
			RoleUUID:       role.UUID,
			PermissionUUID: permission.UUID,
		}
	}

	userRole := &entity.UserRole{
		UUID:     uuid.New().String(),
		UserUUID: user.UUID,
		RoleUUID: role.UUID,
	}

	initSeeds := make([]Seed, len(permissions)+len(rolePermissions)+3)
	initSeeds[0] = Seed{
		Name: "Create initial user",
		Run: func(db *gorm.DB) error {
			_, errDB := createUser(db, user)
			return errDB
		},
	}

	initSeeds[1] = Seed{
		Name: "Create initial role",
		Run: func(db *gorm.DB) error {
			_, errDB := createRole(db, role)
			return errDB
		},
	}

	initSeeds[2] = Seed{
		Name: "Create initial user role",
		Run: func(db *gorm.DB) error {
			_, errDB := createUserRole(db, userRole)
			return errDB
		},
	}

	for i, p := range permissions {
		cp := p
		initSeeds[i+3] = Seed{
			Name: "Create initial permission",
			Run: func(db *gorm.DB) error {
				_, errDB := createPermission(db, cp)
				return errDB
			},
		}
	}

	for i, rp := range rolePermissions {
		crp := rp
		initSeeds[i+3+len(permissions)] = Seed{
			Name: "Create initial role permission",
			Run: func(db *gorm.DB) error {
				_, errDB := createRolePermission(db, crp)
				return errDB
			},
		}
	}

	return initSeeds
}
