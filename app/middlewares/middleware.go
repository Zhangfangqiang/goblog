package middlewares

import "net/http"

/**
 * 定义个数据类型 HttpHandlerFunc 是一个 方法
 */
type HttpHandlerFunc func(http.ResponseWriter, *http.Request)
