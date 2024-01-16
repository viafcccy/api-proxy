package forward

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// forwardRequest 处理传入的HTTP和HTTPS请求，并将它们转发到目标服务器。
func forwardRequest(target string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 解析目标服务器地址
		url, err := url.Parse(target)
		if err != nil {
			http.Error(w, "Error parsing the target URL.", http.StatusInternalServerError)
			return
		}

		// 创建一个反向代理
		proxy := httputil.NewSingleHostReverseProxy(url)

		// 下面的传输设置可以让代理服务器支持HTTPS
		proxy.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //如果要验证SSL证书，则需要删除此行或设置为false
		}

		// 更新header中的Host信息
		r.Host = url.Host
		r.URL.Host = url.Host
		r.URL.Scheme = url.Scheme
		r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
		r.Host = url.Host

		// 转发请求
		proxy.ServeHTTP(w, r)
	}
}

func Run(target string) {
	// 目标服务器地址，请替换为实际地址
	// 测试地址：https://jsonplaceholder.typicode.com/todos/1

	// 设置HTTP代理服务器
	http.HandleFunc("/", forwardRequest(target))

	// 启动HTTP服务器
	log.Println("Starting HTTP server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
