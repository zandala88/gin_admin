package test

import (
	"encoding/json"
	"fmt"
	"gin_admin/helper"
	"github.com/gin-gonic/gin"
	"testing"
)

var (
	baseURL = "http://localhost:9090"
	token   = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6MSwiSWRlbnRpdHkiOiIxIiwiTmFtZSI6ImdldCIsIlJvbGVJZGVudGl0eSI6IjEiLCJJc0FkbWluIjpmYWxzZSwiZXhwIjoxNzA0MjA0Mzc1fQ.EKoj1FDH4PEZ3okXd-HCI_DJtPLuYcxCmISVq9gs93Q"
	header  []byte
)

func init() {
	header, _ = json.Marshal(gin.H{
		"AccessToken": token,
	})
}

// 登录
func TestLoginPassword(t *testing.T) {
	m := gin.H{
		"username": "get",
		"password": "123456",
	}
	data, _ := json.Marshal(m)
	resp, err := helper.HttpPost(baseURL+"/login/password", data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s\n", resp)
}

// 获取菜单列表
func TestMenus(t *testing.T) {
	resp, err := helper.HttpGet(baseURL+"/menus", header...)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s\n", resp)
}

// 获取用户信息
func TestUserInfo(t *testing.T) {
	resp, err := helper.HttpGet(baseURL+"/user/info", header...)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s\n", resp)
}

// 角色列表
func TestSetRoleList(t *testing.T) {
	resp, err := helper.HttpGet(baseURL+"/set/role/list", header...)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s\n", resp)
}
