//#########################################################################
// Author: Johnny Shi
// Created Time: 2014-04-30 14:26:02
// File Name: http_test.go
// Description: 
//#########################################################################
package main
 
import (
    "fmt"
    "os"
    "runtime"
    "path"
    "strings"
    "time"
    //"io"
    "io/ioutil"
    "net/http"
    "net/http/cookiejar"
    "net/url"
    "bytes"
)
 
 
/***************************************************************************************************
    Global Variables
***************************************************************************************************/
var gCurCookies []*http.Cookie;
var gCurCookieJar *cookiejar.Jar;
var gLogger log4go.Logger;
 
/***************************************************************************************************
    Functions
***************************************************************************************************/
//do init before all others
func initAll(){
    gCurCookies = nil
    //var err error;
    gCurCookieJar,_ = cookiejar.New(nil)
    gLogger = nil
     
    //......
}
 
//get url response html
func getUrlRespHtml(strUrl string, postDict map[string]string) string{
    gLogger.Info("getUrlRespHtml, strUrl=%s", strUrl)
    gLogger.Info("postDict=%s", postDict)
     
    var respHtml string = "";
     
    httpClient := &http.Client{
        //Transport:nil,
        //CheckRedirect: nil,
        Jar:gCurCookieJar,
    }
 
    var httpReq *http.Request
    //var newReqErr error
    if nil == postDict {
        gLogger.Info("is GET")
        //httpReq, newReqErr = http.NewRequest("GET", strUrl, nil)
        httpReq, _ = http.NewRequest("GET", strUrl, nil)
        // ...
        //httpReq.Header.Add("If-None-Match", `W/"wyzzy"`)
    } else {
        gLogger.Info("is POST")
        postValues := url.Values{}
        for postKey, PostValue := range postDict{
            postValues.Set(postKey, PostValue)
        }
        gLogger.Info("postValues=%s", postValues)
        postDataStr := postValues.Encode()
        gLogger.Info("postDataStr=%s", postDataStr)
        postDataBytes := []byte(postDataStr)
        gLogger.Info("postDataBytes=%s", postDataBytes)
        postBytesReader := bytes.NewReader(postDataBytes)
        //httpReq, newReqErr = http.NewRequest("POST", strUrl, postBytesReader)
        httpReq, _ = http.NewRequest("POST", strUrl, postBytesReader)
        //httpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
        httpReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    }
     
    httpResp, err := httpClient.Do(httpReq)
    // ...
     
    //httpResp, err := http.Get(strUrl)
    //gLogger.Info("http.Get done")
    if err != nil {
        gLogger.Warn("http get strUrl=%s response error=%s\n", strUrl, err.Error())
    }
    gLogger.Info("httpResp.Header=%s", httpResp.Header)
    gLogger.Debug("httpResp.Status=%s", httpResp.Status)
 
    defer httpResp.Body.Close()
    // gLogger.Info("defer httpResp.Body.Close done")
     
    body, errReadAll := ioutil.ReadAll(httpResp.Body)
    //gLogger.Info("ioutil.ReadAll done")
    if errReadAll != nil {
        gLogger.Warn("get response for strUrl=%s got error=%s\n", strUrl, errReadAll.Error())
    }
    //gLogger.Debug("body=%s\n", body)
 
    //gCurCookies = httpResp.Cookies()
    //gCurCookieJar = httpClient.Jar;
    gCurCookies = gCurCookieJar.Cookies(httpReq.URL);
    //gLogger.Info("httpResp.Cookies done")
     
    //respHtml = "just for test log ok or not"
    respHtml = string(body)
    //gLogger.Info("httpResp body []byte to string done")
 
    return respHtml
}
 
func dbgPrintCurCookies() {
    var cookieNum int = len(gCurCookies);
    gLogger.Info("cookieNum=%d", cookieNum)
    for i := 0; i < cookieNum; i++ {
        var curCk *http.Cookie = gCurCookies[i];
        //gLogger.Info("curCk.Raw=%s", curCk.Raw)
        gLogger.Info("------ Cookie [%d]------", i)
        gLogger.Info("Name\t\t=%s", curCk.Name)
        gLogger.Info("Value\t=%s", curCk.Value)
        gLogger.Info("Path\t\t=%s", curCk.Path)
        gLogger.Info("Domain\t=%s", curCk.Domain)
        gLogger.Info("Expires\t=%s", curCk.Expires)
        gLogger.Info("RawExpires\t=%s", curCk.RawExpires)
        gLogger.Info("MaxAge\t=%d", curCk.MaxAge)
        gLogger.Info("Secure\t=%t", curCk.Secure)
        gLogger.Info("HttpOnly\t=%t", curCk.HttpOnly)
        gLogger.Info("Raw\t\t=%s", curCk.Raw)
        gLogger.Info("Unparsed\t=%s", curCk.Unparsed)
    }
}
 
func main() {
    initAll()
     
    //......
     
    //step3: verify returned cookies
    if bGotCookieBaiduid && bExtractTokenValueOK {
        gLogger.Info("======步骤3：登陆百度并检验返回的Cookie ======");
        staticPageUrl := "http://www.baidu.com/cache/user/html/jump.html";
         
        postDict := map[string]string{}
        //postDict["ppui_logintime"] = ""
        postDict["charset"] = "utf-8"
        //postDict["codestring"] = ""
        postDict["token"] = strLoginToken
        postDict["isPhone"] = "false"
        postDict["index"] = "0"
        //postDict["u"] = ""
        //postDict["safeflg"] = "0"
        postDict["staticpage"] = staticPageUrl
        postDict["loginType"] = "1"
        postDict["tpl"] = "mn"
        postDict["callback"] = "parent.bdPass.api.login._postCallback"
 
        strBaiduUsername := ""
        strBaiduPassword := ""
        // stdinReader := bufio.NewReader(os.Stdin)
        // inputBytes, _ := stdinReader.ReadString('\n')
        // fmt.Printf("Input Char Is : %v", string([]byte(input)[0]))
        //_, err1 := fmt.Scanf("%s", &strBaiduUsername)
        //fmt.Println("Plese input:")
        //fmt.Println("Baidu Username:")
        gLogger.Info("Plese input:")
        gLogger.Info("Baidu Username:")
        _, err1 := fmt.Scanln(&strBaiduUsername)
        if nil == err1 {
            gLogger.Info("strBaiduUsername=%s", strBaiduUsername)
        }
        //fmt.Println("Baidu Password:")
        gLogger.Info("Baidu Password:")
        //_, err2 := fmt.Scanf("%s", &strBaiduPassword)
        _, err2 := fmt.Scanln(&strBaiduPassword)
        if nil == err2 {
            gLogger.Info("strBaiduPassword=%s", strBaiduPassword)
        }
         
        postDict["username"] = strBaiduUsername
        postDict["password"] = strBaiduPassword
        postDict["verifycode"] = ""
        postDict["mem_pass"] = "on"
         
        gLogger.Debug("postDict=%s", postDict)
         
        baiduMainLoginUrl := "https://passport.baidu.com/v2/api/?login";
        loginBaiduRespHtml := getUrlRespHtml(baiduMainLoginUrl, postDict);
        gLogger.Debug("loginBaiduRespHtml=%s", loginBaiduRespHtml)
    }
 
 }

// vim: set noexpandtab ts=4 sts=4 sw=4 :
