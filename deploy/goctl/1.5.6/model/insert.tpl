
func (m *default{{.upperStartCamelObject}}Model) Insert(ctx context.Context, tx *gorm.DB, data *{{.upperStartCamelObject}}) error {
	{{if .withCache}}
    err := m.ExecCtx(ctx, func(conn *gorm.DB) error {
        db = m.scopes(tx)
        return db.Save(&data).Error
	}, m.getCacheKeys(data)...){{else}}db := m.scopes(tx)
        err:= db.WithContext(ctx).Save(&data).Error{{end}}
	return err
}
