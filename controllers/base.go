package controllers

import (
	"fmt"
	"strings"

	"blog/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type baseController struct {
	beego.Controller
	o              orm.Ormer
	controllerName string
	actionName     string
}

func (p *baseController) Prepare() {
	controllerName, actionName := p.GetControllerAndAction()
	p.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	p.actionName = strings.ToLower(actionName)
	p.o = orm.NewOrm()
	// if strings.ToLower(p.controllerName) == "admin" && strings.ToLower(p.actionName) != "login" {
	// 	if p.GetSession("user") == nil {
	// 		p.History("no login", "/admin/login")
	// 		p.Ctx.WriteString(p.controllerName + "====" + p.actionName)
	// 	}
	// }

	if strings.ToLower(p.controllerName) == "blog" {
		p.Data["actionName"] = strings.ToLower(actionName)
		var result []*models.User
		p.o.QueryTable(new(models.User).TableName()).All(&result)
		configs := make(map[string]string)
		for _, v := range result {
			configs[v.UserName] = v.UserName
			fmt.Println(v)
		}
		p.Data["config"] = configs
	}
}

func (p *baseController) History(msg string, url string) {
	if url == "" {
		p.Ctx.WriteString("<script>alert('" + msg + "');window.history.go(-1);</script>")
		p.StopRun()
	} else {
		fmt.Println("url:", url)
		p.Redirect(url, 302)
	}
}

func (p *baseController) getClientIp() string {
	s := strings.Split(p.Ctx.Request.RemoteAddr, ":")
	return s[0]
}
