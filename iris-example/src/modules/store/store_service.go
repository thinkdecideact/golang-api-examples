package store

import (
	"errors"
	"thinkdecideact/src/newtime"

	"thinkdecideact/src/utils"

	"gorm.io/gorm"
)

type StoreService struct {
	DB *gorm.DB
}

func NewStoreService(db *gorm.DB) StoreService {
	return StoreService{DB: db}
}

// The first solution to getting many records by the given condition
func (service *StoreService) GetManyByPage(pageCondition PageConditionDTO) (map[string]interface{}, error) {
	returnMap := map[string]interface{}{}
	var rowList []map[string]interface{}
	whereCondition := map[string]interface{}{"is_del": 0, "is_active": 1}
	selectedFields := []string{"id", "ctime", "mtime", "priority", "is_active", "name", "address"}
	rowStartIndex := pageCondition.PageIndex * pageCondition.RowCountPerPage

	// Get rows
	err := service.DB.Model(&StoreEntity{}).Select(selectedFields).Where(whereCondition).Limit(pageCondition.RowCountPerPage).Offset(rowStartIndex).Order("mtime DESC, id DESC").Find(&rowList).Error
	if err != nil {
		return returnMap, err
	}

	// Get row count
	rowCount, err := service.getRowCount(whereCondition)
	if err != nil {
		return returnMap, err
	}

	returnMap["pageIndex"] = pageCondition.PageIndex
	returnMap["pageCount"] = utils.CalPageCount(rowCount, int64(pageCondition.RowCountPerPage))
	returnMap["rowCount"] = rowCount
	returnMap["rowCountPerPage"] = pageCondition.RowCountPerPage
	returnMap["rows"] = rowList
	returnMap["rowStartIndex"] = rowStartIndex
	return returnMap, nil
}

// The second solution to getting many records by the given condition
func (service *StoreService) GetManyByPage2(pageCondition PageConditionDTO) (map[string]interface{}, error) {
	returnMap := map[string]interface{}{}
	var rowList []ResultDTO
	whereCondition := map[string]interface{}{"is_del": 0, "is_active": 1}
	rowStartIndex := pageCondition.PageIndex * pageCondition.RowCountPerPage

	// Get rows
	err := service.DB.Model(&StoreEntity{}).Where(whereCondition).Limit(pageCondition.RowCountPerPage).Offset(rowStartIndex).Order("mtime DESC, id DESC").Find(&rowList).Error
	if err != nil {
		return returnMap, err
	}

	// Get row count
	rowCount, err := service.getRowCount(whereCondition)
	if err != nil {
		return returnMap, err
	}

	returnMap["pageIndex"] = pageCondition.PageIndex
	returnMap["pageCount"] = utils.CalPageCount(rowCount, int64(pageCondition.RowCountPerPage))
	returnMap["rowCount"] = rowCount
	returnMap["rowCountPerPage"] = pageCondition.RowCountPerPage
	returnMap["rows"] = rowList
	returnMap["rowStartIndex"] = rowStartIndex
	return returnMap, nil
}

// The third solution to getting many records by the given condition
func (service *StoreService) GetManyByPage3(pageCondition PageConditionDTO) (map[string]interface{}, error) {
	returnMap := map[string]interface{}{}
	var rowList []StoreEntity
	whereCondition := map[string]interface{}{"is_del": 0, "is_active": 1}
	rowStartIndex := pageCondition.PageIndex * pageCondition.RowCountPerPage

	// Get rows
	err := service.DB.Where(whereCondition).Limit(pageCondition.RowCountPerPage).Offset(rowStartIndex).Order("mtime DESC, id DESC").Find(&rowList).Error
	if err != nil {
		return returnMap, err
	}

	// Get row count
	rowCount, err := service.getRowCount(whereCondition)
	if err != nil {
		return returnMap, err
	}

	returnMap["pageIndex"] = pageCondition.PageIndex
	returnMap["pageCount"] = utils.CalPageCount(rowCount, int64(pageCondition.RowCountPerPage))
	returnMap["rowCount"] = rowCount
	returnMap["rowCountPerPage"] = pageCondition.RowCountPerPage
	returnMap["rows"] = rowList
	returnMap["rowStartIndex"] = rowStartIndex
	return returnMap, nil
}

// Get how many rows there are
func (service *StoreService) getRowCount(whereCondition map[string]interface{}) (int64, error) {
	var rowCount int64
	if err := service.DB.Model(&StoreEntity{}).Where(whereCondition).Count(&rowCount).Error; err != nil {
		return rowCount, err
	}
	return rowCount, nil
}

// The first solution to getting a record by its ID
func (service *StoreService) GetOne(id int) (map[string]interface{}, error) {
	var resultMap map[string]interface{}
	selectedFields := []string{"id", "ctime", "mtime", "priority", "is_active", "name", "address"}
	err := service.DB.Model(&StoreEntity{}).Select(selectedFields).Where("id=? AND is_del=? AND is_active=?", id, 0, 1).First(&resultMap).Error
	if err != nil {
		return resultMap, err
	}
	if len(resultMap) == 0 {
		return resultMap, errors.New("fail to find a record")
	}
	return resultMap, nil
}

// The second solution to getting a record by its ID
func (service *StoreService) GetOne2(id int) (ResultDTO, error) {
	var resultDTO ResultDTO
	err := service.DB.Model(&StoreEntity{}).Where("id=? AND is_del=? AND is_active=?", id, 0, 1).First(&resultDTO).Error
	if err != nil {
		return resultDTO, err
	}
	if resultDTO.ID == 0 {
		return resultDTO, errors.New("fail to find a record")
	}
	return resultDTO, nil
}

// The third solution to getting a record by its ID
func (service *StoreService) GetOne3(id int) (ResultDTO, error) {
	var store StoreEntity
	err := service.DB.Where("id=? AND is_del=? AND is_active=?", id, 0, 1).First(&store).Error
	if err != nil {
		return ResultDTO{}, err
	}
	if store.ID == 0 {
		return ResultDTO{}, errors.New("fail to find a record")
	}
	return store.Serializer(), nil
}

func (service *StoreService) Create(store StoreEntity) (uint, error) {
	store.Ctime = newtime.GetNowTime()
	store.Mtime = newtime.GetNowTime()
	store.IsActive = 1
	store.IsDel = 0
	// if err := service.DB.Create(&store).Error; err != nil {
	if err := service.DB.Save(&store).Error; err != nil {
		return 0, err
	}
	if store.ID == 0 {
		return 0, errors.New("failed to create a record")
	}
	return store.ID, nil
}

func (service *StoreService) Delete(id uint) error {
	store := StoreEntity{}
	if err := service.DB.Where("id = ?", id).First(&store).Error; err != nil {
		return err
	}
	result := service.DB.Delete(&store)
	if err := result.Error; err != nil {
		return err
	}
	if result.RowsAffected == 0 {
		return errors.New("failed to delete a record")
	}
	return nil
}

// Update all the fields of a record
func (service *StoreService) Update(id uint, updateData StoreEntity) error {
	existedStore := StoreEntity{}
	if err := service.DB.Where("id = ?", id).First(&existedStore).Error; err != nil {
		return err
	}
	updateData.Mtime = newtime.GetNowTime()
	result := service.DB.Model(&existedStore).Select("*").Omit("id").Updates(updateData)
	if err := result.Error; err != nil {
		return err
	}
	if result.RowsAffected == 0 {
		return errors.New("failed to update a record")
	}
	return nil
}

// Partially update one or more fields of a record
func (service *StoreService) PartialUpdate(id uint, updateDTO UpdateDTO) error {
	store := StoreEntity{}
	if err := service.DB.Where("id = ?", id).First(&store).Error; err != nil {
		return err
	}

	store.Name = updateDTO.Name
	store.Address = updateDTO.Address
	store.Mtime = newtime.GetNowTime()
	result := service.DB.Save(&store)
	if err := result.Error; err != nil {
		return err
	}
	// if result.RowsAffected == 0 {
	// 	return errors.New("failed to update a record")
	// }
	return nil
}
