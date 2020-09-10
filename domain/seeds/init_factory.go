package seeds

import (
	"go-rest-skeleton/domain/entity"

	"github.com/jinzhu/gorm"

	"github.com/google/uuid"
)

type InitFactory struct {
	seeders []Seed
}

var (
	user = &entity.User{
		UUID:      uuid.New().String(),
		FirstName: "Trisna",
		LastName:  "Ashari",
		Email:     "trisna.x22@gmail.com",
		Phone:     "01234567890",
		Password:  "123456",
	}
	role = &entity.Role{
		UUID: uuid.New().String(),
		Name: "Super Administrator",
	}
	permissions = []*entity.Permission{
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
		{UUID: uuid.New().String(), ModuleKey: "role", PermissionKey: "detail"},
	}
	userRole = &entity.UserRole{
		UUID:     uuid.New().String(),
		UserUUID: user.UUID,
		RoleUUID: role.UUID,
	}
)

func newInitFactory() *InitFactory {
	return &InitFactory{seeders: make([]Seed, 0)}
}

func (is *InitFactory) GenerateUserSeeder() *InitFactory {
	is.seeders = append(is.seeders, Seed{
		Name: "Create initial user",
		Run: func(db *gorm.DB) error {
			_, errDB := createUser(db, user)
			return errDB
		},
	})

	return is
}

func (is *InitFactory) generateRoleSeeder() *InitFactory {
	is.seeders = append(is.seeders, Seed{
		Name: "Create initial role",
		Run: func(db *gorm.DB) error {
			_, errDB := createRole(db, role)
			return errDB
		},
	})

	return is
}

func (is *InitFactory) generateRolePermissionsSeeder() *InitFactory {
	for _, p := range permissions {
		cp := p
		is.seeders = append(is.seeders, Seed{
			Name: "Create initial permission",
			Run: func(db *gorm.DB) error {
				_, errDB := createPermission(db, cp)
				return errDB
			},
		})
	}

	return is
}

func (is *InitFactory) generateUserRoleSeeder() *InitFactory {
	is.seeders = append(is.seeders, Seed{
		Name: "Assign initial role to user",
		Run: func(db *gorm.DB) error {
			_, errDB := createUserRole(db, userRole)
			return errDB
		},
	})

	return is
}

func initFactory() []Seed {
	initialSeeds := newInitFactory()
	initialSeeds.GenerateUserSeeder()
	initialSeeds.generateRoleSeeder()
	initialSeeds.generateRolePermissionsSeeder()
	initialSeeds.generateUserRoleSeeder()

	return initialSeeds.seeders
}
