package usecase

import (
	"praktikum/models"
	"praktikum/models/payload"
	"praktikum/repository/database"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

type ItemUsecase interface {
	GetAllitems() (item []models.Item, err error)
	GetItemById(id string) (item *models.Item, err error)
	GetItemByCategoryId(id int) (item []models.Item, err error)
	GetItemByName(keyword string) (item []models.Item, err error)
	CreateItem(req *payload.CreateItem) (item *models.Item, err error)
	UpdateItemById(id string, req *payload.UpdateItem) (res payload.UpdateItemResponse, err error)
	DeleteItemById(id string) (item *models.Item, err error)
}

type itemUsecase struct {
	itemRepository database.ItemRepository
}

func NewItemUsecase(itemRepository database.ItemRepository) *itemUsecase {
	return &itemUsecase{itemRepository}
}

func (i *itemUsecase) GetAllitems() (item []models.Item, err error) {
	item, err = i.itemRepository.GetAllItem()
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (i *itemUsecase) GetItemById(id string) (item *models.Item, err error) {
	item, err = i.itemRepository.GetItemById(id)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (i *itemUsecase) GetItemByCategoryId(id int) (item []models.Item, err error) {
	item, err = i.itemRepository.GetItemByCategoryId(id)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (i *itemUsecase) GetItemByName(keyword string) (item []models.Item, err error) {
	item, err = i.itemRepository.GetItemByName(keyword)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (i *itemUsecase) CreateItem(req *payload.CreateItem) (item *models.Item, err error) {
	id, _ := uuid.NewRandom()
	itemReq := &models.Item{
		ID:           id,
		Name:         req.Name,
		Description:  req.Description,
		Stock:        req.Stock,
		Price:        req.Price,
		CategoryID: req.CategoryID,
	}

	err = i.itemRepository.CreateItem(itemReq)
	if err != nil {
		return nil, echo.NewHTTPError(400, err.Error())
	}

	return itemReq, nil
}

func (i *itemUsecase) UpdateItemById(id string, req *payload.UpdateItem) (res payload.UpdateItemResponse, err error) {
	_, err = i.itemRepository.GetItemById(id)
	if err != nil {
		return res, echo.NewHTTPError(400, "Item Not Found!")
	}

	itemReq := &models.Item{
		Name:        req.Name,
		Description: req.Description,
		Stock:       req.Stock,
		Price:       req.Price,
	}

	err = i.itemRepository.UpdateItemById(id, itemReq)
	if err != nil {
		return res, echo.NewHTTPError(400, err.Error())
	}

	res = payload.UpdateItemResponse{
		Name:        itemReq.Name,
		Description: itemReq.Description,
		Stock:       itemReq.Stock,
		Price:       itemReq.Price,
	}

	return res, nil
}

func (i *itemUsecase) DeleteItemById(id string) (item *models.Item, err error) {
	err = i.itemRepository.DeleteItemById(id, item)
	if err != nil {
		return nil, echo.NewHTTPError(400, "Cannot delete item")
	}

	return item, nil
}
