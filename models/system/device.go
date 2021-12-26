package system

import (
	"mwp3000/database/models"
	"time"
)

// 设备表
type ModelDeviceInfo struct {
	DeviceId    string `json:"device_id" gorm:"size:128;comment:设备Id-唯一标识"`
	ParentId    string `json:"parent_id" gorm:"size:128;comment:父节点设备Id"`
	DeviceName  string `json:"device_name" gorm:"size:128;comment:设备名字"`
	DeviceIP    string `json:"device_ip" gorm:"size:128;comment:设备Ip地址"`
	DevicePwd   string `json:"device_pwd" gorm:"size:128;comment:设备访问密码"`
	DeviceType  string `json:"device_type" gorm:"size:128;comment:设备类型"`
	DeviceModel string `json:"device_model" gorm:"size:128;comment:设备型号"`
	DeviceSN    string `json:"device_serial" gorm:"size:128;comment:设备序列号"`

	LoginId         string    `json:"login_id" gorm:"size:128;comment:会话Id"`
	Manufacture     string    `json:"manufacture" gorm:"size:128;comment:厂家"`
	Category        string    `json:"category" gorm:"size:128;comment:分类"`
	Firmware        string    `json:"firmware" gorm:"size:128;comment:固件版本"`
	Status          string    `json:"status" gorm:"size:128;comment:在线状态"`
	PingRate        float64   `json:"ping_rate" gorm:"size:128;comment:通讯状况"`
	LastContactedAt time.Time `json:"last_contacted_at" gorm:"comment:最后通讯时间"`
}

type TDeviceInfo struct {
	models.Model

	ModelDeviceInfo

	models.ControlBy
	models.ModelTime
}

func (TDeviceInfo) TableName() string {
	return "t_device_info"
}

func (e *TDeviceInfo) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *TDeviceInfo) GetId() interface{} {
	return e.Id
}
