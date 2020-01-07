package controllers

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"ssnbee/models"
)

type AboutController struct {
	beego.Controller
}

func (c *AboutController) Get() {

	subject := models.GetSubject()
	c.Data["price"] = subject.Price
	c.TplName = "about.tpl"
	const (
		name = iota
		_
		_
		age = iota
	)
	fmt.Printf("%d", age)

	err := func() error {
		id, err := c.GetInt("id")
		beego.Info(id)
		if err != nil {
			id = 1
		}

		if err != nil {
			return errors.New("not exist")
		}
		return nil
	}

	if err != nil {
		c.Ctx.WriteString("wrong param")
	}
	var options map[string]string;
	print(options)

	/*if err:=json.Unmarshal([]byte(subject),&options);err != nil {
	c.Ctx.WriteString("wrong param,json decode failed")
	}
	*/
}

