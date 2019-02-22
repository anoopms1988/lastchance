package controllers

import (
	"github.com/astaxie/beego"
)

type BankingController struct {
	beego.Controller
}

// @router /banking [get]
func (c *BankingController) ShowAccounts() {
	c.TplName = "banking.tpl"
}
