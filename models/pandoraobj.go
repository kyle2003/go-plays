package models

import (
	"pandora/constants"
	"pandora/utils"
	"strconv"
)

type PandoraObj struct {
	// 自增ID
	ID uint64 `gorm:"primary_key;column:F_id;" json:"-"`
	// 拼音名
	Name string `gorm:"column:F_name;type:text;not null" json:"name"`
	// 中文名
	Title string `gorm:"column:F_title;type:text;not null" json:"title"`
	// URL地址
	URL string `gorm:"column:F_url;type:varchar(32);unique;not null" json:"url"`
	// 是否采集完成
	ReapStatus constants.ReapStatus `gorm:"column:F_reap_status;default:2;index" json:"reap_status"`
	// 下载完成状态
	DownloadStatus constants.DownloadStatus `gorm:"column:F_download_status;type:int;default:2;index" json:"download_status"`
	// 创建时间戳
	Created int64 `gorm:"column:F_created;type:int;default:0" json:"created"`
	// 更新时间戳
	Updated int64 `gorm:"column:F_updated;type:int;default:0" json:"updaeted"`
	// 软删除
	Enabled uint8 `gorm:"column:F_enabled;type:int;default:1;index" json:"enabled"`
}

// GetHtml content of the category page
func (p *PandoraObj) GetHtml(limit int) string {
	html := utils.GetHtml(p.URL)
	i := 2
	if limit == 0 {
		limit = p.GetPageLimit()
	}
	for i <= limit {
		index := "index" + strconv.Itoa(i) + ".html"
		href := p.URL + index
		html += utils.GetHtml(href)
		i++
	}
	return html
}

// GetPageLimit get the limit of page
func (p *PandoraObj) GetPageLimit() int {
	html := utils.GetHtml(p.URL)
	return utils.GetPageLimit(html)
}
