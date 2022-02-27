package services

import (
	"fmt"

	mwd "github.com/aditya109/library-system/internal/middlewares/errors"
	model "github.com/aditya109/library-system/internal/models"
	logger "github.com/sirupsen/logrus"
)

var items = []model.Item{
	{
		ID:   66,
		Name: "Heidt",
	},
	{
		ID:   12,
		Name: "Bertine",
	},
	{
		ID:   59,
		Name: "Vastah",
	},
	{
		ID:   39,
		Name: "Frendel",
	},
}

// GetAllItems gets all items present in the system
func GetAllItems() ([]model.Item, error) {
	return items, nil
}

// GetItemByID gets item by ID
func GetItemByID(id int64) (model.Item, error) {
	items, err := GetAllItems()
	if err != nil {
		logger.Error(err)
	}
	for _, v := range items {
		if v.ID == id {
			return v, nil
		}
	}
	return model.Item{}, &mwd.AppError{
		Cause: fmt.Sprintf("item with ID=%d not present", id),
	}
}

// GetItemsByIDAndName gets a list of items filtered by ID and Name
func GetItemsByIDAndName(id int64, name string) ([]model.Item, error) {
	items, err := GetAllItems()
	if err != nil {
		logger.Error(err)
	}
	for _, v := range items {
		if v.ID == id && v.Name == name {
			return []model.Item{v}, nil
		}
	}
	return []model.Item{}, &mwd.AppError{
		Cause: fmt.Sprintf("item with ID=%d not present", id),
	}
}
