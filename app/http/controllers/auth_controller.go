package controllers

import (
	"goblog/pkg/view"
	"net/http"
)

/**
 * AuthController 结构
 */
type AuthController struct{}

/**
 * Register 注册页面
 */
func (*AuthController) Register(w http.ResponseWriter, r *http.Request) {
	view.RenderSimple(w, view.D{}, "auth.register")
}

/**
 * DoRegister 处理注册逻辑
 */
func (*AuthController) DoRegister(w http.ResponseWriter, r *http.Request) {
	//
}