
type (
	{{.lowerStartCamelObject}}Model interface{
		{{.method}}
		Sharding(Sharding model.ISharding) *default{{.upperStartCamelObject}}Model 
		Builder(tx *gorm.DB) *gorm.DB
	}

	default{{.upperStartCamelObject}}Model struct {
		{{if .withCache}}mysql.CachedConn{{else}}conn *gorm.DB{{end}}
		table string
		sharding model.ISharding	
	}

	{{.upperStartCamelObject}} struct {
		{{.fields}}
	}
)
