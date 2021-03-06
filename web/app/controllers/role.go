package controllers

import (
	"strings"

	"goplays/web/app/entity"
	"goplays/web/app/service"

	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
)

type RoleController struct {
	BaseController
}

func (this *RoleController) List() {
	roleList, err := service.RoleService.GetAllRoles()
	this.checkError(err)
	for k, role := range roleList {
		roleList[k].UserList, _ = service.UserService.GetUserListByRoleId(role.Id)
	}
	this.Data["pageTitle"] = "角色管理"
	this.Data["list"] = roleList
	this.display()
}

func (this *RoleController) Add() {
	if this.isPost() {
		role := &entity.Role{}
		role.RoleName = this.GetString("role_name")
		role.Description = this.GetString("description")
		if role.RoleName == "" {
			this.showMsg("角色名不能为空", MSG_ERR)
		}
		err := service.RoleService.AddRole(role)
		if err == nil {
			service.ActionService.AddRole(this.auth.GetUser().UserName, role.RoleName)
		}
		this.checkError(err)
		this.redirect(beego.URLFor("RoleController.List"))
	}
	this.Data["pageTitle"] = "创建角色"
	this.display()
}

func (this *RoleController) Edit() {
	id := this.GetString("id")
	role, err := service.RoleService.GetRole(id)
	this.checkError(err)

	if this.isPost() {
		role.RoleName = this.GetString("role_name")
		role.Description = this.GetString("description")
		fileds := bson.M{"RoleName": role.RoleName,
			"Description": role.Description}
		err := service.RoleService.UpdateRole(role, fileds)
		if err == nil {
			service.ActionService.UpdateRole(this.auth.GetUser().UserName, role.RoleName)
		}
		this.checkError(err)
		this.redirect(beego.URLFor("RoleController.List"))
	}

	this.Data["pageTitle"] = "编辑角色"
	this.Data["role"] = role
	this.display()
}

func (this *RoleController) Del() {
	id := this.GetString("id")

	err := service.RoleService.DeleteRole(id)
	if err == nil {
		service.ActionService.DelRole(this.auth.GetUser().UserName, id)
	}
	this.checkError(err)

	this.redirect(beego.URLFor("RoleController.List"))
}

func (this *RoleController) Perm() {
	id := this.GetString("id")
	role, err := service.RoleService.GetRole(id)
	this.checkError(err)

	if this.isPost() {
		pids := this.GetStrings("pids")
		perms := this.GetStrings("perms")
		if len(pids) == 0 {
			role.ProjectIds = ""
		} else {
			role.ProjectIds = strings.Join(pids, ",")
		}
		fileds := bson.M{"ProjectIds": role.ProjectIds}
		err := service.RoleService.UpdateRole(role, fileds)
		this.checkError(err)
		err = service.RoleService.SetPerm(role.Id, perms)
		if err == nil {
			service.ActionService.PermRole(this.auth.GetUser().UserName, role.RoleName)
		}
		this.checkError(err)
		this.redirect(beego.URLFor("RoleController.List"))
	}

	//projectList, _ := service.ProjectService.GetAllProject()
	permList := service.SystemService.GetPermList()

	chkmap := make(map[string]string)
	for _, v := range role.PermList {
		chkmap[v.Key] = "checked"
	}
	if role.ProjectIds != "" {
		pids := strings.Split(role.ProjectIds, ",")
		for _, pid := range pids {
			chkmap[pid] = "checked"
		}
	}

	this.Data["pageTitle"] = "编辑权限"
	this.Data["permList"] = permList
	//this.Data["projectList"] = projectList
	this.Data["role"] = role
	this.Data["chkmap"] = chkmap
	this.display()
}
