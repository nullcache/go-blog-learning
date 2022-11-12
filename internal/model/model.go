package model

type Model struct {
	ID         uint32 `json:"id"`
	CreateOn   uint32 `json:"create_on"`
	ModifiedOn uint32 `json:"modified_on"`
	CreateBy   string `json:"create_by"`
	ModifiedBy string `json:"modified_by"`
	DeletedOn  uint32 `json:"deleted_on"`
	IdDel      uint8  `json:"id_del"`
}
