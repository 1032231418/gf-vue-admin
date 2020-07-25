package v1

import (
	"fmt"
	"server/app/api/request"
	"server/app/api/response"
	"server/app/service"
	"server/library/global"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

// @Tags authorityAndMenu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body string true "可以什么都不填"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /menu/getMenu [post]
func GetMenu(r *ghttp.Request) {
	authorityId := gconv.String(r.GetParam("authority_id"))
	menus, err := service.GetMenuTree(authorityId)
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("获取失败，%v", err))
		r.Exit()
	}
	global.OkWithData(r, response.AuthorityMenu{Menus: menus})
}

// @Tags menus
// @Summary 分页获取基础menu列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取基础menu列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/getMenuList [post]
func GetMenuList(r *ghttp.Request) {
	var pageInfo request.PageInfo
	if err := r.Parse(&pageInfo); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	list, total, err := service.GetMenuList()
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("获取数据失败，err:%v", err))
		r.Exit()
	}
	global.OkWithData(r, response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	})
}

// @Tags authorityAndMenu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body string true "可以什么都不填"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /menu/getBaseMenuTree [post]
func GetBaseMenuTree(r *ghttp.Request) {
	menus, err := service.GetBaseMenuTree()
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("获取失败，%v", err))
		r.Exit()
	}
	global.OkWithData(r, response.BaseMenus{Menus: menus})
}

// @Tags authorityAndMenu
// @Summary 增加menu和角色关联关系
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AddMenuAuthorityInfo true "增加menu和角色关联关系"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/addMenuAuthority [post]
func AddMenuAuthority(r *ghttp.Request) {
	var addMenuAuthorityInfo request.AddMenuAuthorityInfo
	err := service.AddMenuAuthority(&addMenuAuthorityInfo)
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("添加失败，%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "添加成功")
}

// @Tags authorityAndMenu
// @Summary 获取指定角色menu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AuthorityIdInfo true "获取指定角色menu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/GetMenuAuthority [post]
func GetMenuAuthority(r *ghttp.Request) {
	var authorityIdInfo request.AuthorityIdInfo
	if err := r.Parse(&authorityIdInfo); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	menus, err := service.GetMenuAuthority(&authorityIdInfo)
	if err != nil {
		global.FailWithDetailed(r, global.SUCCESS, response.AuthorityMenu{Menus: menus}, fmt.Sprintf("添加失败，%v", err))
		r.Exit()
	}
	global.Result(r, global.SUCCESS, response.AuthorityMenu{Menus: menus}, "获取成功")
}

// @Tags menus
// @Summary 新增菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CreateBaseMenu true "新增菜单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/addBaseMenu [post]
func CreateBaseMenu(r *ghttp.Request) {
	var create request.CreateBaseMenu
	if err := r.Parse(&create); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if err := service.CreateBaseMenu(&create); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("添加失败，%v", err))
		r.Exit()
	}
	global.FailWithMessage(r, "添加成功")
}

// @Tags menu
// @Summary 删除菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "删除菜单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/deleteBaseMenu [post]
func DeleteBaseMenu(r *ghttp.Request) {
	var deleteInfo request.GetById
	if err := r.Parse(&deleteInfo); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if err := service.DeleteBaseMenu(&deleteInfo); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("删除菜单失败，err:%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "删除成功")
}

// @Tags menus
// @Summary 更新菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.UpdateBaseMenu true "更新菜单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/updateBaseMenu [post]
func UpdateBaseMenu(r *ghttp.Request) {
	var updateInfo request.UpdateBaseMenu
	if err := r.Parse(&updateInfo); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if err := service.UpdateBaseMenu(&updateInfo); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("删除菜单失败，err:%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "修改成功")
}

// @Tags menus
// @Summary 根据id获取菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "根据id获取菜单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/getBaseMenuById [post]
func GetBaseMenuById(r *ghttp.Request) {
	var idInfo request.GetById
	if err := r.Parse(&idInfo); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	menu, err := service.GetBaseMenuById(&idInfo)
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("查询失败，err:%v", err))
		r.Exit()
	}
	global.OkWithData(r, response.BaseMenu{Menu: menu})
}
