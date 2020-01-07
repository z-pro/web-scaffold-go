package utils


type Pager struct {
	PageNum     int
	PageSize   int
	TotalPage  int
	Total int
	FirstPage  bool
	LastPage   bool
	List       interface{}
}


func PageUtil(count int, pageNo int, pageSize int, list interface{}) Pager {
	tp := count / pageSize
	if count % pageSize > 0 {
		tp = count / pageSize + 1
	}
	return Pager{
		PageNum: pageNo,
		PageSize: pageSize,
		TotalPage: tp,
		Total: count,
		FirstPage: pageNo == 1,
		LastPage: pageNo == tp,
		List: list,
	}
}
