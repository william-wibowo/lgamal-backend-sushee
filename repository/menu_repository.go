package repository

import (
	"final-project-backend/entity"
	"strings"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MenuRepository interface {
	GetMenu(entity.MenuQuery) (*[]entity.Menu, error)
	GetMenuCount(q entity.MenuQuery) (int, error)
	GetPromotionMenu() (*[]entity.Promotion, error)
	ValidatePromoMenu(menuId, promoId int) (int, error)
}

type MenuRepositoryImpl struct {
	db *gorm.DB
}

type MenuRepositoryConfig struct {
	DB *gorm.DB
}

func NewMenuRepository(c MenuRepositoryConfig) MenuRepository {
	return &MenuRepositoryImpl{
		db: c.DB,
	}
}

func (r *MenuRepositoryImpl) GetMenuCount(q entity.MenuQuery) (int, error) {
	var rows int64

	menuCategorySQ := r.db.
		Select("id").
		Where("category_name IN (?)", strings.Split(q.FilterByCategory, ",")).
		Or("'' = ?", q.FilterByCategory).
		Table("categories")
	query := r.db.
		Joins("menus").
		Where("category_id IN (?)", menuCategorySQ).
		Table("menus")
	query.Count(&rows)
	err := query.Error
	return int(rows), err
}

func (r *MenuRepositoryImpl) GetMenu(q entity.MenuQuery) (*[]entity.Menu, error) {
	var menus []entity.Menu

	menuCategorySQ := r.db.
		Select("id").
		Where("category_name IN (?)", strings.Split(q.FilterByCategory, ",")).
		Or("'' = ?", q.FilterByCategory).
		Table("categories")
	menuSQ := r.db.
		Joins("menus").
		Where("category_id IN (?)", menuCategorySQ).
		Table("menus")
	query := r.db.
		Table("(?) as th", menuSQ).
		Where("menu_name ilike ?", "%"+q.Search+"%").
		Order(clause.OrderByColumn{
			Column: clause.Column{
				Name: q.SortBy,
			},
			Desc: strings.ToLower(q.Sort) == "desc",
		}).
		Limit(q.Limit).
		Offset(q.Page*q.Limit - q.Limit).
		Find(&menus)
	err := query.Error
	return &menus, err
}

func (r *MenuRepositoryImpl) GetPromotionMenu() (*[]entity.Promotion, error) {
	var promotions []entity.Promotion
	err := r.db.
		Model(&entity.Promotion{}).
		Preload("PromoMenus.Menu").
		Where("? between started_at and expired_at", time.Now()).
		Find(&promotions).Error
	return &promotions, err
}

func (r *MenuRepositoryImpl) ValidatePromoMenu(menuId, promoId int) (int, error) {
	var c int64
	err := r.db.
		Model(entity.PromoMenu{}).
		Where("menu_id = ? AND promotion_id = ?", menuId, promoId).
		Count(&c).
		Table("promo_menus").
		Error
	return int(c), err
}