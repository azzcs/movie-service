package commons

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)
// ViewPrefix 视图前缀
const ViewPrefix string = "viewPrefix"

func IsMobileDevice(requestHeader string) bool {
	deviceArray := [3]string{"android","mac os","windows phone"}
	if requestHeader==""{
		return false
	}
	requestHeader = strings.ToLower(requestHeader)
	for _,device:=range deviceArray{
		if strings.Contains(requestHeader,device){
			return true
		}
	}
	return false
}


func SplitByDollar(index int,str string) string{
	return strings.Split(str, "$")[index]
}

type MovieContent struct {
	Kuyun []string `json:"kuyun"`
	Ckm3u8 []string `json:"ckm3u8"`
}
func AnalysisMovieContentJson(str string) MovieContent{
	var m MovieContent
	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		fmt.Printf("err was %v", err)
	}
	return m
}
func CheckDevice(c *gin.Context)  {
	viewPrefix := "pc/"
	userAgent := c.Request.Header.Get("user-agent")
	if IsMobileDevice(userAgent){
		viewPrefix = "phone/"
	}
	c.Request.Header.Add(ViewPrefix,viewPrefix)
}