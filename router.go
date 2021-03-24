package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/smtc/glog"
	"html/template"
	"net/http"
	"net/url"
	"path"
	"strings"
)

type resultData struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

var (
	// templates定义
	tmplDir  string = "templates"
	funcName        = template.FuncMap{
		"noescape": func(s string) template.HTML {
			return template.HTML(s)
		},
		"safeurl": func(s string) template.URL {
			return template.URL(s)
		},
	}
)

/**
设置gin路由规则
创建人:邵炜
创建时间:2017年2月9日13:51:48
输入参数:gin engine
*/
func setGinRouter(r *gin.Engine) {
	g := &r.RouterGroup
	if rootPrefix != "" {
		g = r.Group(rootPrefix)
	}
	{
		g.GET("/", func(c *gin.Context) { c.String(http.StatusOK, "ok") }) //确认接口服务程序是否健在
		g.GET("/api/*pth", htmlTemplates)
		g.GET("/api/:parms", apiRegisterGet)   //接口转发
		g.POST("/api/:parms", apiRegisterPost) //接口转发
	}
}

func htmlTemplates(c *gin.Context) {
	r := c.Request
	// path: [/rootprefix]/assets/....
	pth := c.Param("pth") //r.URL.Path
	if pth == "" {
		glog.Error("assetsFiles: path is empty: %s\n", r.URL.Path)
		c.Data(200, "text/plain", []byte(""))
		return
	}

	fp, err := getAssetFilePath(pth)
	if err != nil {
		glog.Error("assetsFiles: %s\n", err)
		c.Data(200, "text/plain", []byte(""))
		return
	}

	http.ServeFile(c.Writer, c.Request, fp)
}

func apiRegisterGet(c *gin.Context) {
	values := c.Request.URL.Query()
	httpProcess(c, values, http.MethodGet)
}

func apiRegisterPost(c *gin.Context) {
	values := c.Request.PostForm
	httpProcess(c, values, http.MethodGet)
}

func httpProcess(c *gin.Context, values url.Values, method string) {
	var (
		resultData resultData
	)
	parms := c.Param("parms")
	cookieIn := c.Request.Cookies()
	resopneStr, cookieArrayIn, resultCode, err := register(http.MethodGet, parms, values, cookieIn)
	resultData.Code = resultCode
	resultData.Message = resopneStr
	resultByte, _ := json.Marshal(resultData)
	defer c.String(http.StatusOK, string(resultByte))
	if err != nil {
		glog.Error("apiRegisterGet register run err! parms: %s values: %v cookieIn: %v \n", parms, values, cookieIn)
	}
	for _, item := range *cookieArrayIn {
		c.SetCookie(item.Name, item.Value, item.MaxAge, "/", c.Request.Host, item.Secure, item.HttpOnly)
	}
}

func getAssetFilePath(pth string) (string, error) {
	entrys := strings.Split(pth, "/")[1:]
	sentrys := []string{tmplDir}
	for _, s := range entrys {
		s = strings.TrimSpace(s)
		if s != "" {
			sentrys = append(sentrys, s)
		}
	}
	return path.Join(sentrys...), nil
}
