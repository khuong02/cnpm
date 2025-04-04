package api

import (
	"cnpm/pkg/model"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

func (r Route) sso(c *gin.Context) {
	var payload model.Login

	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Sai định dạng dữ liệu đầu vào!",
		})
		return
	}

	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: jar,
	}

	loginURL := "https://sso.hcmut.edu.vn/cas/login?service=https://mybk.hcmut.edu.vn/my/homeSSO.action"

	// B1: Gửi GET request để lấy trang login
	resp, err := client.Get(loginURL)
	if err != nil {
		log.Fatal("GET failed:", err)
	}
	defer resp.Body.Close()

	// Debug: In HTML response
	body, _ := io.ReadAll(resp.Body)
	html := string(body)

	// Đọc lại response từ string (vì resp.Body chỉ đọc được 1 lần)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatal("goquery parse error:", err)
	}

	lt, exists1 := doc.Find("input[name=lt]").Attr("value")
	execution, exists2 := doc.Find("input[name=execution]").Attr("value")

	if !exists1 || !exists2 {
		c.AbortWithStatusJSON(401, gin.H{
			"status":  "error",
			"message": "Đăng nhập thất bại! Có thể sai tài khoản hoặc mật khẩu.",
		})
		return
	}

	// B2: Gửi POST để login
	form := url.Values{}
	form.Add("username", payload.Username) // <- thay bằng mã số sinh viên
	form.Add("password", payload.Password) // <- thay bằng mật khẩu
	form.Add("lt", lt)
	form.Add("execution", execution)
	form.Add("_eventId", "submit")
	form.Add("submit", "LOGIN")

	req, err := http.NewRequest("POST", loginURL, strings.NewReader(form.Encode()))
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{
			"status":  "error",
			"message": "Đăng nhập thất bại! Có thể sai tài khoản hoặc mật khẩu.",
		})
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err = client.Do(req)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{
			"status":  "error",
			"message": "Đăng nhập thất bại! Có thể sai tài khoản hoặc mật khẩu.",
		})
		return
	}
	defer resp.Body.Close()

	finalURL := resp.Request.URL.String()
	fmt.Println("Redirect đến:", finalURL)
	if finalURL != "https://mybk.hcmut.edu.vn/my/homeSSO.action" {
		c.AbortWithStatusJSON(401, gin.H{
			"status":  "error",
			"message": "Đăng nhập thất bại! Có thể sai tài khoản hoặc mật khẩu.",
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Đăng nhập thành công!",
	})
}
