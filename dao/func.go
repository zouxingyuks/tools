package dao

import (
	"errors"
	"gorm.io/gorm/clause"
)

// FuzzyMatch 模糊搜索，传入查询字段，以及结构值
func FuzzyMatch[T any](key string, values *T, preload ...string) error {
	db := DB
	//加载外键
	for _, v := range preload {
		db = db.Preload(v)
	}
	return db.Where("name LIKE ?", "%"+key+"%").Find(&values).Error
}

// PerfectMatch 精准搜索，传入查询结构体
func PerfectMatch[T any](key *T, values *[]T, preload ...string) error {
	db := DB
	//加载外键
	for _, v := range preload {
		db = db.Preload(v)
	}
	return db.Find(&values, key).Error
}

func Add[T any](m *T) error {
	result := DB.Create(&m)
	return result.Error
}
func Update[T any](m *T) error {
	result := DB.Updates(&m)
	return result.Error
}

// List 列出所有数据
func List[T any](values *T, preload ...string) (err error) {
	db := DB
	//加载外键
	for _, v := range preload {
		db = db.Preload(v)
	}
	return db.Find(&values).Error
}
func Del[T any](value *T, mode int) (err error) {
	db := DB
	switch mode {
	//软删除
	case 0:
		err = db.Delete(&value).Error
	//硬删除
	case 1:
		err = db.Unscoped().Delete(&value).Error
	//级联软删除
	case 2:
		err = db.Select(clause.Associations).Delete(&value).Error
	//级联硬删除
	case 3:
		err = db.Select(clause.Associations).Unscoped().Delete(&value).Error
	default:
		err = errors.New("delete mode choose wrong")
	}
	return err
}
