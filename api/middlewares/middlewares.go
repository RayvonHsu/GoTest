package middlewares

import (
	"net/http"
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {			//将数据格式设置为json的中间件（相当于拦截器）
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}
