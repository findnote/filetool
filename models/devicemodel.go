package models

import (
	"mwp3000/pkg/utils"

	log "github.com/pion/ion-log"
	"github.com/tal-tech/go-zero/core/logx"
	"gorm.io/gorm"
)

type TDeviceInfoModel struct {
	db *gorm.DB
}

func NewTDeviceINfoModel() *TDeviceInfoModel {
	return &TDeviceInfoModel{
		db: GetDB(),
	}
}

// 新增设备信息; 要保证deviceId与deviceSerial（设备ID与产品序列号唯一）
func (m *TDeviceInfoModel) Insert(data *DeviceInfo) error {
	log.Infof(">>(m *TDeviceInfoModel) Insert db: %v", m.db)

	deviceId := data.DeviceId
	deviceSn := data.DeviceSN

	log.Infof(">>deviceId: %v, deviceSn: %v", deviceId, deviceSn)

	var queryInfo DeviceInfo
	var info DeviceInfo

	results := m.db.Table(data.TableName()).Where("device_id = ? and device_sn = ?", deviceId, deviceSn).First(&queryInfo)
	if results.Error != nil {
		if results.Error == gorm.ErrRecordNotFound {
			log.Infof(">>(m *TDeviceInfoModel) Insert Create")

			data.CreatedAt = utils.GetTimeNow()

			err := m.db.Table(data.TableName()).Create(&data).Error
			if err != nil {
				logx.Errorf("insert by %s error, reason: %s", data.TableName(), err)
				return err
			}
		}
	} else {
		log.Infof(">>(m *TDeviceInfoModel) Insert Updates, Id: %v", queryInfo.Id)

		info = *data
		info.Id = queryInfo.Id // 需保证使用查询到记录的Id
		info.UpdatedAt = utils.GetTimeNow()

		err := m.db.Table(data.TableName()).Save(&info).Error

		if err != nil {
			logx.Errorf("insert by %s error, reason: %s", data.TableName(), err)
			return err
		}
	}

	return nil
}

// 软删除设备信息
func (m *TDeviceInfoModel) Delete(id int) error {
	var info DeviceInfo

	db := m.db.Model(&info).Where("id = ?", id).Delete(&info)
	if db.Error != nil {
		logx.Errorf("insert by %s error, reason: %s", info.TableName(), db.Error)
		return db.Error
	}
	return nil
}

// 获取单条设备信息
func (m *TDeviceInfoModel) Select(id int) (*DeviceInfo, error) {
	var info DeviceInfo

	db := m.db.Model(&info).Where("id = ?", id).First(&info)
	if db.Error != nil {
		logx.Errorf("FindOne by %s error, reason: %s", info.TableName(), db.Error)
		return &info, db.Error
	}

	return &info, nil
}

// 更新设备信息
func (m *TDeviceInfoModel) Update(c *DeviceInfo) error {
	db := m.db.Model(&c).Updates(c)
	if db.Error != nil {
		logx.Errorf("Update by %s error, reason: %s", c.TableName(), db.Error)
		return db.Error
	}
	return nil
}

// 根据条件获取设备信息
func (m *TDeviceInfoModel) List(fields map[string]string, page int, size int) (*[]DeviceInfo, int64, error) {
	var resp []DeviceInfo
	var info DeviceInfo

	fieldsStr := utils.ListMap2Str(fields)
	var offset = (page - 1) * size
	if offset < 0 {
		offset = 0
	}

	err := m.db.Where(fieldsStr).Limit(size).Offset(offset).Find(&resp).Error
	if err != nil {
		logx.Errorf("List by `%s` error, reason: %s", info.TableName(), err)
		return &resp, 0, err
	}

	// 查询符合条件的总条数
	var count int64
	var resps []DeviceInfo
	m.db.Where(fieldsStr).Find(&resps).Count(&count)
	return &resp, count, nil
}
