package data

func (sess *Session) Check() (ok bool, err error) {
	result := Db.Where("uuid=?", sess.Uuid).Find(sess)
	err = result.Error
	if err == nil && result.RowsAffected == 1 {
		ok = true
	}
	return
}
