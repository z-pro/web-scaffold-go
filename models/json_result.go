package models

type JsonResult struct{
	errcode int "json:errcode"
	errmsg string "json:errmsg"
	data interface{} "json:data"
}

func (this *JsonResult)Success()  JsonResult  {
	var jsonResult JsonResult
	jsonResult.errcode=0
	return jsonResult
}

func (this *JsonResult)SuccessWithData(data interface{})  JsonResult  {
	var jsonResult JsonResult
	jsonResult.errcode=0
	jsonResult.errmsg="success"
	jsonResult.data=data
	return jsonResult
}

func (this *JsonResult)Failed(errcode int,errmsg string)  JsonResult  {
	var jsonResult JsonResult
	jsonResult.errcode= errcode
	jsonResult.errmsg =errmsg
	return jsonResult
}
