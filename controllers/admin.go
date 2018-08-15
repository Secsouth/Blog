package controllers

import (
	"blog/models"
	"fmt"
)

type AdminController struct {
	baseController
}

//login in back-end
func (c *AdminController) Login() {
	if c.Ctx.Request.Method == "POST" {
		username := c.GetString("username")
		password := c.GetString("password")
		user := models.User{UserName: username}
		c.o.Read(&user, "username")
		fmt.Println("username,password : ", username, password)
		fmt.Println("real username: ", user.UserName, "--real password: ", user.Password)

		// if user.Password == "" {
		c.History("account is not existed", "/admin/main.html")
		// }

		// if util.Md5(password) != strings.Trim(user.Password, " ") {
		// 	c.History("password is not correct", "/admin/main.html")
		// }
		// user.LastIp = c.getClientIp()
		// user.LoginCount = user.LoginCount + 1
		// if _, err := c.o.Update(&user); err != nil {
		// 	c.History("login abnormal", "/admin/main.html")
		// } else {
		// 	c.History("login success", "/admin/main.html")
		// }
		// c.SetSession("user", user)
	}
	c.TplName = c.controllerName + "/login.html"
}

func (c *AdminController) Main() {
	//c.TplName相当于http.Handle(http.FileServer())是用来寻找html的
	c.TplName = c.controllerName + "/main.html"
}
