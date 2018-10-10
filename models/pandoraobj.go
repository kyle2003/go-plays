package models

import (
	"pandora/constants"
)

type PandoraObj struct {
	// 自增ID
	ID uint64 `gorm:"primary_key;column:F_id;" json:"-"`
	// 拼音名
	Name string `gorm:"column:F_name;type:text" json:"name"`
	// 中文名
	Title string `gorm:"column:F_title;type:text" json:"title"`
	// URL地址
	URL string `gorm:"column:F_url;type:varchar(32)" json:"url"`
	// 是否采集完成
	ReapStatus constants.ReapStatus `gorm:"column:F_reap_status;default:0" json:"reap_status"`
	// 下载完成状态
	DownloadStatus constants.DownloadStatus `gorm:"column:F_download_status;type:int;default:0" json:"download_status"`
	// 创建时间戳
	Created int64
	// 更新时间戳
	Updated int64
	// 软删除
	Enabled uint8
}
