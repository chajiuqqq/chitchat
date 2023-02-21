package data

func (p *Post) CreatedAtFormat() string {
	return p.CreatedAt.Format("2006.01.02 15:04:05")
}
