package repository

import (
	"hacktiv-digitalent-golang/assignment/assignment-2/model"

	"gorm.io/gorm"
)

type ordersRepository struct {
	db *gorm.DB
}

func NewOrdersRepository(db *gorm.DB) *ordersRepository {
	return &ordersRepository{
		db: db,
	}
}

func (or *ordersRepository) Create(newOrders model.Orders) (model.Orders, error) {
	err := or.db.Create(&newOrders).Error
	if err != nil {
		return model.Orders{}, err
	}
	return newOrders, nil
}

func (or *ordersRepository) GetAll() ([]model.Orders, error) {
	var orders []model.Orders
	err := or.db.Preload("Items").Find(&orders).Error
	if err != nil {
		return []model.Orders{}, err
	}
	return orders, nil
}


func (or *ordersRepository) Update(updatedOrders model.Orders) (model.Orders, error) {
    err := or.db.Model(&model.Items{}).Where("orders_id = ?", updatedOrders.Orders_id).Delete(&model.Items{}).Error
    if err != nil {
        return model.Orders{}, err
    }

    err = or.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&updatedOrders).Error
    if err != nil {
        return model.Orders{}, err
    }


    return updatedOrders, nil
}

func (or *ordersRepository) GetByID(orders_id int) (model.Orders, error) {
	var orders model.Orders
	err := or.db.Preload("Items").First(&orders, orders_id).Error
	if err != nil {
		return model.Orders{}, err
	}
	return orders, nil
}

func (or *ordersRepository) Delete(orders_id int) error {
	tx := or.db.Begin()

	if err := tx.Exec("DELETE FROM items WHERE orders_id = ?", orders_id).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Delete(&model.Orders{}, orders_id).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}