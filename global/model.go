package global

type Model struct {
	ID         int    `gorm:"primary_key" json:"id"`
	CreatedOn  string `json:"created_on"`
	ModifiedOn string `json:"modified_on"`
}
