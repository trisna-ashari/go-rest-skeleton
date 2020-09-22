package seeds

import (
	"go-rest-skeleton/domain/entity"

	"gorm.io/gorm"

	"github.com/google/uuid"
)

type InitFactory struct {
	seeders []Seed
}

var (
	user = &entity.User{
		UUID:     uuid.New().String(),
		Name:     "Trisna Novi Ashari",
		Email:    "trisna.x2@gmail.com",
		Phone:    "01234567890",
		Password: "123456",
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
	storageCategory = []*entity.StorageCategory{
		{
			UUID:      uuid.New().String(),
			Slug:      "avatar",
			Path:      "avatar",
			Name:      "Avatar",
			MimeTypes: "image/jpg,image/jpeg,image/png,image/bmp,image/gif",
		},
		{
			UUID:      uuid.New().String(),
			Slug:      "document",
			Path:      "document",
			Name:      "Document",
			MimeTypes: "application/pdf",
		},
		{
			UUID:      uuid.New().String(),
			Slug:      "file",
			Path:      "file",
			Name:      "File",
			MimeTypes: "application/pdf",
		},
		{
			UUID:      uuid.New().String(),
			Slug:      "thumbnail",
			Path:      "thumbnail",
			Name:      "Thumbnail",
			MimeTypes: "image/png",
		},
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

func (is *InitFactory) generatePermissionsSeeder() *InitFactory {
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

func (is *InitFactory) generateRolePermissionsSeeder() *InitFactory {
	for _, p := range permissions {
		cp := p
		crp := &entity.RolePermission{
			UUID:           uuid.New().String(),
			RoleUUID:       role.UUID,
			PermissionUUID: cp.UUID,
		}
		is.seeders = append(is.seeders, Seed{
			Name: "Create initial permission",
			Run: func(db *gorm.DB) error {
				_, errDB := createRolePermission(db, crp)
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

func (is *InitFactory) generateStorageCategorySeeder() *InitFactory {
	for _, sc := range storageCategory {
		csc := sc
		is.seeders = append(is.seeders, Seed{
			Name: "Create initial storage category",
			Run: func(db *gorm.DB) error {
				_, errDB := createStorageCategory(db, csc)
				return errDB
			},
		})
	}

	return is
}

func initFactory() []Seed {
	initialSeeds := newInitFactory()
	initialSeeds.GenerateUserSeeder()
	initialSeeds.generateRoleSeeder()
	initialSeeds.generatePermissionsSeeder()
	initialSeeds.generateRolePermissionsSeeder()
	initialSeeds.generateUserRoleSeeder()
	initialSeeds.generateStorageCategorySeeder()

	return initialSeeds.seeders
}
