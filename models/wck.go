package models

import (
   "encoding/base64"
   "encoding/json"
   "net/http"
   "net/url"
   "strconv"
   "strings"


   "github.com/beego/beego/v2/client/httplib"
   "github.com/beego/beego/v2/core/logs"
   "github.com/buger/jsonparser"
)


func appjmp(urls string, tokenKey string) (string, error) {
   ua := cloudInfo(urls)
   v := url.Values{}
   v.Add("tokenKey", tokenKey)
   v.Add("to", `https://plogin.m.jd.com/jd-mlogin/static/html/appjmp_blank.html`)
   v.Add("client_type", "android")
   v.Add("appid", "879")
   v.Add("appup_type", "1")
   req := httplib.Get(`https://un.m.jd.com/cgi-bin/app/appjmp?` + v.Encode())
   req.Header("User-Agent", ua)
   req.Header("accept", `text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3`)
   req.SetCheckRedirect(func(req *http.Request, via []*http.Request) error {
      return http.ErrUseLastResponse
   })
   rsp, err := req.Response()
   if err != nil {
      return "", err
   }
   cookies := strings.Join(rsp.Header.Values("Set-Cookie"), " ")
   // ptKey := FetchJdCookieValue("pt_key", cookies)
   return cookies, nil
}
func GetKey(WSCK string) (string, error) {
   url := checkCloud()
   return getToken(url, WSCK)
}

func checkCloud() string {
   urlList := []string{"aHR0cDovLzQzLjEzNS45MC4yMy8=", "aHR0cHM6Ly9zaGl6dWt1Lm1sLw==", "aHR0cHM6Ly9jZi5zaGl6dWt1Lm1sLw=="}
   for i := range urlList {
      decodeString, _ := base64.StdEncoding.DecodeString(urlList[i])
      _, err := httplib.Get(string(decodeString)).String()
      if err == nil {
         return string(decodeString)
      }
   }
   return ""
}

type T struct {
   Code      int    `json:"code"`
   Update    int    `json:"update"`
   Jdurl     string `json:"jdurl"`
   UserAgent string `json:"User-Agent"`
}
type T2 struct {
   FunctionId    string `json:"functionId"`
   ClientVersion string `json:"clientVersion"`
   Build         string `json:"build"`
   Client        string `json:"client"`
   Partner       string `json:"partner"`
   Oaid          string `json:"oaid"`
   SdkVersion    string `json:"sdkVersion"`
   Lang          string `json:"lang"`
   HarmonyOs     string `json:"harmonyOs"`
   NetworkType   string `json:"networkType"`
   Uemps         string `json:"uemps"`
   Ext           string `json:"ext"`
   Ef            string `json:"ef"`
   Ep            string `json:"ep"`
   St            int64  `json:"st"`
   Sign          string `json:"sign"`
   Sv            string `json:"sv"`
}

func cloudInfo(url string) string {
   req := httplib.Get(url + "check_api")
   req.Header("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36 Edg/100.0.1185.50")
   req.Header("authorization", "Bearer Shizuku")
   s, _ := req.Bytes()
   t := T{}
   json.Unmarshal(s, &t)
   return t.UserAgent
}
func getToken(urls string, WSCK string) (string, error) {
   ua := cloudInfo(urls)
   req1 := httplib.Get(urls + "genToken")
   req1.Header("User-Agent", ua)
   s, _ := req1.Bytes()
   t := T2{}
   json.Unmarshal(s, &t)
   v := url.Values{}
   v.Add("functionId", t.FunctionId)
   v.Add("clientVersion", t.ClientVersion)
   v.Add("build", t.Build)
   v.Add("client", t.Client)
   v.Add("partner", t.Partner)
   v.Add("oaid", t.Oaid)
   v.Add("sdkVersion", t.SdkVersion)
   v.Add("lang", t.Lang)
   v.Add("harmonyOs", t.HarmonyOs)
   v.Add("networkType", t.NetworkType)
   v.Add("uemps", t.Uemps)
   v.Add("ext", t.Ext)
   v.Add("ef", t.Ef)
   v.Add("ep", t.Ep)
   v.Add("st", strconv.FormatInt(t.St, 10))
   v.Add("sign", t.Sign)
   v.Add("sv", t.Sv)
   req := httplib.Post(`https://api.m.jd.com/client.action?` + v.Encode())
   req.Header("cookie", WSCK)
   req.Header("User-Agent", ua)
   req.Header("content-type", `application/x-www-form-urlencoded; charset=UTF-8`)
   req.Header("charset", `UTF-8`)
   req.Header("accept-encoding", `br,gzip,deflate`)
   req.Body(`body=%7B%22to%22%3A%22https%253a%252f%252fplogin.m.jd.com%252fjd-mlogin%252fstatic%252fhtml%252fappjmp_blank.html%22%7D&`)
   data, err := req.Bytes()
   if err != nil {
      return "", err
   }
   tokenKey, _ := jsonparser.GetString(data, "tokenKey")
   cookie, err := appjmp(urls, tokenKey)
   logs.Info(cookie)
   if err != nil {
      return "", err
   }
   return cookie, nil
}