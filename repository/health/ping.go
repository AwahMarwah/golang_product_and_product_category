package health

func (r *repo) Ping() (err error) {
	return r.db.Ping()
}
