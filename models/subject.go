package models


type Subject struct {
	Author string
	Price float32
	Page int
	Title string
}

func GetSubject() Subject {
	m:=new(Subject)
	//x:=Subject{Price:12121}
	m.Author="zs"
	m.Price=111.1
	return *m
}