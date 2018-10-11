package models

import (
	"pandora/constants"
)

type PandoraObj struct {
	// 自增ID
	ID uint64 `gorm:"primary_key;column:F_id;" json:"-"`
	// 拼音名
	Name string `gorm:"column:F_name;type:text;unique;not null" json:"name"`
	// 中文名
	Title string `gorm:"column:F_title;type:text;not null" json:"title"`
	// URL地址
	URL string `gorm:"column:F_url;type:varchar(32);not null" json:"url"`
	// 是否采集完成
	ReapStatus constants.ReapStatus `gorm:"column:F_reap_status;default:0;index" json:"reap_status"`
	// 下载完成状态
	DownloadStatus constants.DownloadStatus `gorm:"column:F_download_status;type:int;default:0;index" json:"download_status"`
	// 创建时间戳
	Created int64 `gorm:"column:F_created;type:int;default:0" json:"created"`
	// 更新时间戳
	Updated int64 `gorm:"column:F_updated;type:int;default:0" json:"updaeted"`
	// 软删除
	Enabled uint8 `gorm:"column:F_enabled;type:int;default:1;index" json:"enabled"`
}
