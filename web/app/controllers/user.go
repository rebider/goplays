// Copyright 2015 lisijie. All Rights Reserved.

package controllers

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"goplays/web/app/libs"
	"goplays/web/app/service"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	BaseController
}

// 帐号列表
func (this *UserController) List() {
	page, _ := strconv.Atoi(this.GetString("page"))
	if page < 1 {
		page = 1
	}

	count, _ := service.UserService.GetTotal()
	users, _ := service.UserService.GetUserList(page, this.pageSize, true)

	this.Data["pageTitle"] = "帐号管理"
	this.Data["count"] = count
	this.Data["list"] = users
	this.Data["pageBar"] = libs.NewPager(page, int(count), this.pageSize, beego.URLFor("UserController.List"), true).ToString()
	this.display()
}

// 添加帐号
func (this *UserController) Add() {
	if this.isPost() {
		valid := validation.Validation{}

		username := this.GetString("username")
		//email := this.GetString("email")
		agent := this.GetString("agent")
		sex, _ := this.GetInt("sex")
		password1 := this.GetString("password1")
		password2 := this.GetString("password2")

		valid.Required(username, "username").Message("请输入用户名")
		//valid.Required(email, "email").Message("请输入Email")
		//valid.Email(email, "email").Message("Email无效")
		//valid.Required(agent, "agent").Message("请输入游戏ID")
		valid.Required(password1, "password1").Message("请输入密码")
		valid.Required(password2, "password2").Message("请输入确认密码")
		valid.MinSize(password1, 6, "password1").Message("密码长度不能小于6个字符")
		valid.Match(password1, regexp.MustCompile(`^`+regexp.QuoteMeta(password2)+`$`), "password2").Message("两次输入的密码不一致")
		if valid.HasErrors() {
			for _, err := range valid.Errors {
				this.showMsg(err.Message, MSG_ERR)
			}
		}

		user, err := service.UserService.AddUser(username, agent, password1, sex)
		if err == nil {
			service.ActionService.AddUser(this.auth.GetUser().UserName, username)
		}
		this.checkError(err)

		// 更新角色
		roleIds := make([]string, 0)
		for _, v := range this.GetStrings("role_ids") {
			//if roleId, _ := strconv.Atoi(v); roleId > 0 {
			//	roleIds = append(roleIds, roleId)
			//}
			roleIds = append(roleIds, v)
		}
		service.UserService.UpdateUserRoles(user.Id, roleIds)

		this.redirect(beego.URLFor("UserController.List"))
	}

	roleList, _ := service.RoleService.GetAllRoles()
	this.Data["pageTitle"] = "添加帐号"
	this.Data["roleList"] = roleList
	this.display()
}

func (this *UserController) Edit() {
	id := this.GetString("id")
	user, err := service.UserService.GetUser(id, true)
	this.checkError(err)

	if this.isPost() {
		valid := validation.Validation{}

		//email := this.GetString("email")
		agent := this.GetString("agent")
		atype, _ := this.GetInt("atype")
		belong, _ := this.GetInt("belong")
		sex, _ := this.GetInt("sex")
		status, _ := this.GetInt("status")
		password1 := this.GetString("password1")
		password2 := this.GetString("password2")

		//valid.Required(agent, "agent").Message("请输入游戏ID")
		//valid.Required(email, "email").Message("请输入Email")
		//valid.Email(email, "email").Message("Email无效")
		if password1 != "" {
			valid.Required(password1, "password1").Message("请输入密码")
			valid.Required(password2, "password2").Message("请输入确认密码")
			valid.MinSize(password1, 6, "password1").Message("密码长度不能小于6个字符")
			valid.Match(password1, regexp.MustCompile(`^`+regexp.QuoteMeta(password2)+`$`), "password2").Message("两次输入的密码不一致")
		}

		if valid.HasErrors() {
			for _, err := range valid.Errors {
				this.showMsg(err.Message, MSG_ERR)
			}
		}

		user.Sex = sex
		user.Status = status
		//user.Email = email
		user.Agent = agent
		fileds := bson.M{"sex": user.Sex, "status": user.Status, "agent": user.Agent, "update_time": bson.Now()}
		if atype > 0 {
			user2, err2 := service.UserService.GetUserByAtype(uint32(atype))
			if err2 == nil && user2.Atype > 0 {
				this.checkError(errors.New("分包类型已经存在"))
			} else {
				fileds["atype"] = uint32(atype)
			}
		}
		if belong > 0 {
			fileds["belong"] = uint32(belong)
		}
		service.UserService.UpdateUser(user, fileds)

		if password1 != "" {
			service.UserService.ModifyPassword(user.Id, password1)
		}
		service.ActionService.UpdateUser(this.auth.GetUser().UserName, id)

		// 更新角色
		roleIds := make([]string, 0)
		for _, v := range this.GetStrings("role_ids") {
			//if roleId, _ := strconv.Atoi(v); roleId > 0 {
			//	roleIds = append(roleIds, roleId)
			//}
			roleIds = append(roleIds, v)
		}
		service.UserService.UpdateUserRoles(user.Id, roleIds)

		this.redirect(beego.URLFor("UserController.List"))
	}

	chkmap := make(map[string]string)
	for _, v := range user.RoleList {
		chkmap[v.Id] = "selected"
	}

	roleList, _ := service.RoleService.GetAllRoles()
	this.Data["pageTitle"] = "修改帐号"
	this.Data["user"] = user
	this.Data["roleList"] = roleList
	this.Data["chkmap"] = chkmap
	this.Data[fmt.Sprintf("sex%d", user.Sex)] = "checked"
	this.display()
}

func (this *UserController) Del() {
	id := this.GetString("id")

	if id == "1" {
		this.showMsg("不能删除ID为1的帐号。", MSG_ERR)
	}

	err := service.UserService.DeleteUser(id)
	if err == nil {
		service.ActionService.DelUser(this.auth.GetUser().UserName, id)
	}
	this.checkError(err)

	this.redirect(beego.URLFor("UserController.List"))
}
