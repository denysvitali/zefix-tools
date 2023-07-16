package zefix

type Company struct {
	LegalName string `gorm:"primaryKey,index:idx_legaln_name" json:"legalName"`
	Name      string `gorm:"index:idx_legaln_name" json:"name"`
	Uri       string `json:"uri"`
	Locality  string `json:"locality"`
	Type      string `gorm:"index" json:"type"`
	Address   string `json:"address"`
}
