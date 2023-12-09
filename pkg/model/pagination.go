package model

import (
	"chenxi/pkg/utils"
	"errors"
	"time"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

// Param 分页参数
type Param struct {
	DB       *gorm.DB
	PageNum  int
	PageSize int
	OrderBy  []string
	ShowSQL  bool
}

// Paginator 分页返回
type Paginator struct {
	Total   int64       `json:"total_record"`
	Records interface{} `json:"records"`
}

// PaginatorToApi 将分页器转换为API所需的数据类型
func PaginatorToApi[T any](p *Paginator, toValue T) T {
	// 使用copier.CopyWithOption将分页器中的记录复制到目标值toValue中
	copier.CopyWithOption(&toValue, p.Records, copier.Option{
		IgnoreEmpty: true, // 忽略空值
		DeepCopy:    true, // 深度复制
		Converters: []copier.TypeConverter{ // 自定义类型转换器
			{
				SrcType: time.Time{},   // 源类型为time.Time
				DstType: copier.String, // 目标类型为copier.String
				Fn: func(src interface{}) (dst interface{}, err error) {
					s, ok := src.(time.Time) // 将源类型转换为time.Time
					if !ok {
						return nil, errors.New("src type not matching") // 如果类型不匹配则返回错误
					}
					return s.Local().Format(utils.DateTime), nil // 返回本地时间格式化后的字符串
				},
			},
		},
	})
	return toValue // 返回转换后的值
}

// Paging 分页
func Paging(p *Param, result interface{}) *Paginator {
	db := p.DB

	if p.ShowSQL {
		db = db.Debug()
	}
	if p.PageNum < 1 {
		p.PageNum = 1
	}
	if p.PageSize == 0 || p.PageSize > 100 {
		p.PageSize = 10
	}
	if len(p.OrderBy) > 0 {
		for _, o := range p.OrderBy {
			db = db.Order(o)
		}
	}

	var paginator Paginator
	var count int64
	var offset int

	countRecords(db, result, &count)

	if p.PageNum == 1 {
		offset = 0
	} else {
		offset = (p.PageNum - 1) * p.PageSize
	}

	db.Limit(p.PageSize).Offset(offset).Find(result)

	paginator.Total = count
	paginator.Records = result

	return &paginator
}

func countRecords(db *gorm.DB, anyType interface{}, count *int64) {
	db.Count(count)
}
