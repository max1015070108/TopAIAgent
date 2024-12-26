package database

// Model 数据库模型定义
type Model struct {
	ID   string `gorm:"primaryKey"`
	Name string
}

// NetworkModel 关联表模型
type NetworkModel struct {
	ModelID   string `gorm:"primaryKey"`
	NetworkID string `gorm:"primaryKey"`
}

// TableName 设置 Model 表名
func (Model) TableName() string {
	return "models"
}

// TableName 设置 NetworkModel 表名
func (NetworkModel) TableName() string {
	return "network_models"
}
