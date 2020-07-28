package entity

// Module represent schema of table module.
type Module struct {
	UUID string `gorm:"size:36;not null;unique_index;primary_key" json:"uuid"`
	Key  string `gorm:"size:100;not null;" json:"module_key"`
}
