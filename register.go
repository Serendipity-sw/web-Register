package main

import (
	"fmt"
	"github.com/smtc/glog"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func register(method string, urlPath string, parms url.Values, cookieArray []*http.Cookie) (string, *[]*http.Cookie, string, error) {
	httpClient := http.Client{}
	httpRespone, err := http.NewRequest(method, fmt.Sprintf("%s/%s", forwardingDomain, urlPath), strings.NewReader(parms.Encode()))
	if err != nil {
		glog.Error("register NewRequest run err! method: %s  urlPath: %s  parms: %v  cookie: %v  err: %s \n", method, urlPath, parms, cookieArray, err.Error())
		return "", nil, "00001", err
	}
	httpRespone.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for _, item := range cookieArray {
		httpRespone.AddCookie(item)
	}
	result, err := httpClient.Do(httpRespone)
	if err != nil {
		glog.Error("register Do run err! method: %s  urlPath: %s  parms: %v  cookie: %v  err: %s \n", method, urlPath, parms, cookieArray, err.Error())
		return "", nil, "00002", err
	}
	defer result.Body.Close()
	resultByte, err := ioutil.ReadAll(result.Body)
	if err != nil {
		glog.Error("register ReadAll run err! method: %s  urlPath: %s  parms: %v  cookie: %v  err: %s \n", method, urlPath, parms, cookieArray, err.Error())
		return "", nil, "00003", err
	}
	resultStr := string(resultByte)
	resultCookieArray := result.Cookies()
	glog.Info("register run success! method: %s  urlPath: %s  parms: %v  cookie: %v  result: %s \n", method, urlPath, parms, cookieArray, resultStr)
	return resultStr, &resultCookieArray, "00000", nil
}
