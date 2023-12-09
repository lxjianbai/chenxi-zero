func ({{.upperStartCamelObject}}) TableName() string {
    return {{.table}}
}

func new{{.upperStartCamelObject}}Model(conn *gorm.DB{{if .withCache}}, c cache.CacheConf{{end}}) *default{{.upperStartCamelObject}}Model {
	return &default{{.upperStartCamelObject}}Model{
		{{if .withCache}}CachedConn: mysql.NewConn(conn, c){{else}}conn:conn{{end}},
		table: {{.table}},
	}
}

func(m *default{{.upperStartCamelObject}}Model)Sharding(sharding model.ISharding) *default{{.upperStartCamelObject}}Model {
	m.sharding = sharding
	return m
}

func(m *default{{.upperStartCamelObject}}Model)Builder(tx *gorm.DB) *gorm.DB{
	return m.scopes(tx).Model(&{{.upperStartCamelObject}}{})
}

func (m *default{{.upperStartCamelObject}}Model)scopes(tx *gorm.DB) *gorm.DB{
	var db = m.conn
	if tx != nil {
		db = tx	
	}
	if m.sharding == nil {
		return db
	}
	return db.Scopes(func(d *gorm.DB) *gorm.DB {
		return d.Table(strings.Trim(m.table,"`") + "_" + m.sharding.GetTableSuffix())
	})
}