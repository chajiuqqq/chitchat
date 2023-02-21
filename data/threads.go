package data

func Threads() (threads []Thread, err error) {
	//要获取thread里的post，用预加载Preload+字段名，这里就是“Posts”。
	//再获取Post里的User需要用嵌套预加载："Posts.User"
	err = Db.Preload("Posts").Preload("Posts.User").Preload("User").Find(&threads).Error
	return
}

func (t *Thread) NumReplies() int {
	nums := Db.Model(t).Association("Posts").Count()
	return int(nums)
}

func (t *Thread) CreatedAtFormat() string {
	return t.CreatedAt.Format("2006.01.02 15:04:05")
}
