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
		{UUID: uuid.New().String(), ModuleKey: "tour", PermissionKey: "read"},
		{UUID: uuid.New().String(), ModuleKey: "tour", PermissionKey: "create"},
		{UUID: uuid.New().String(), ModuleKey: "tour", PermissionKey: "update"},
		{UUID: uuid.New().String(), ModuleKey: "tour", PermissionKey: "delete"},
		{UUID: uuid.New().String(), ModuleKey: "tour", PermissionKey: "bulk_delete"},
		{UUID: uuid.New().String(), ModuleKey: "tour", PermissionKey: "detail"},
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
	application = &entity.Application{
		UUID: uuid.New().String(),
		Name: "go-rest-skeleton",
	}
	applicationApiKey = &entity.ApplicationApiKey{
		UUID:            uuid.New().String(),
		ApplicationUUID: application.UUID,
		Name:            "go-rest-skeleton-api-key",
		ApiKey:          "9fde26c5-bc87-4081-a270-50b414d70fb6",
	}
	applicationOauth = &entity.ApplicationOauth{
		UUID:             uuid.New().String(),
		ApplicationUUID:  application.UUID,
		Name:             "go-rest-skeleton-oauth",
		SupportEmails:    "support_one@example.com,support_two@example.com",
		DeveloperEmails:  "dev@example.com",
		Logo:             "",
		HomePageURL:      "https://github.com/trisna-ashari/go-rest-skeleton",
		TosURL:           "https://github.com/trisna-ashari/go-rest-skeleton",
		PrivacyPolicyURL: "https://github.com/trisna-ashari/go-rest-skeleton",
		Domains:          "http://localhost:8181",
		Scopes:           "read write",
	}
	applicationOauthClient = &entity.ApplicationOauthClient{
		UUID:            uuid.New().String(),
		ApplicationUUID: application.UUID,
		Name:            "go-rest-skeleton-oauth-client",
		ClientID:        "10e207ab-79ec-42ed-85f2-3a10e3b3ddbb",
		ClientSecret:    "9e8c5bfe-a93e-4041-b404-0ae326a1e491",
		Referrers:       "http://localhost:8181",
		Callbacks:       "http://localhost:8181/oauth2/callback",
	}
)

func newInitFactory() *InitFactory {
	return &InitFactory{seeders: make([]Seed, 0)}
}

func (is *InitFactory) generateUserSeeder() *InitFactory {
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
	r := role
	for _, p := range permissions {
		csp := p
		crp := &entity.RolePermission{
			UUID: uuid.New().String(),
		}

		is.seeders = append(is.seeders, Seed{
			Name: "Create initial permission",
			Run: func(db *gorm.DB) error {
				_, errDB := createRolePermission(db, r, csp, crp)
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
			_, errDB := createUserRole(db, user, role, userRole)
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

func (is *InitFactory) generateApplicationSeeder() *InitFactory {
	is.seeders = append(is.seeders, Seed{
		Name: "Create initial application",
		Run: func(db *gorm.DB) error {
			_, errDB := createApplication(db, application)
			return errDB
		},
	})

	return is
}

func (is *InitFactory) generateApplicationApiKeySeeder() *InitFactory {
	is.seeders = append(is.seeders, Seed{
		Name: "Create initial api key for initial application",
		Run: func(db *gorm.DB) error {
			_, errDB := createApplicationApiKey(db, application, applicationApiKey)
			return errDB
		},
	})

	return is
}

func (is *InitFactory) generateApplicationOauthSeeder() *InitFactory {
	is.seeders = append(is.seeders, Seed{
		Name: "Create initial oauth for initial application",
		Run: func(db *gorm.DB) error {
			_, errDB := createApplicationOauth(db, application, applicationOauth)
			return errDB
		},
	})

	return is
}

func (is *InitFactory) generateApplicationOauthClientSeeder() *InitFactory {
	is.seeders = append(is.seeders, Seed{
		Name: "Create initial oauth clients for initial application",
		Run: func(db *gorm.DB) error {
			_, errDB := createApplicationOauthClient(db, application, applicationOauthClient)
			return errDB
		},
	})

	return is
}

func initFactory() []Seed {
	initialSeeds := newInitFactory()
	initialSeeds.generateUserSeeder()
	initialSeeds.generateRoleSeeder()
	initialSeeds.generatePermissionsSeeder()
	initialSeeds.generateRolePermissionsSeeder()
	initialSeeds.generateUserRoleSeeder()
	initialSeeds.generateStorageCategorySeeder()
	initialSeeds.generateApplicationSeeder()
	initialSeeds.generateApplicationApiKeySeeder()
	initialSeeds.generateApplicationOauthSeeder()
	initialSeeds.generateApplicationOauthClientSeeder()

	return initialSeeds.seeders
}
