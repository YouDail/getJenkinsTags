package common

import (
	log "github.com/golang/glog"
	"strings"
)

func HandlerStr(src string) string {

	var sdt string
	log.Infof("HandlerStr 参数内容: ", src)
	if strings.Contains(src, `\n`) {
		sdt = strings.ReplaceAll(src, `\n`, "AAAAAAAA")
		log.Infof("HandlerStr 参数包含换行符 , 返回处理后的字符串: ", sdt)
		return sdt
	} else if strings.Contains(src, "AAAAAAAA") {
		sdt = strings.ReplaceAll(src, "AAAAAAAA", "\n")
		log.Infof("HandlerStr 参数包含AAAAAAAA , 返回处理后的字符串: ", sdt)
		return sdt
	} else {
		log.Infof("HandlerStr 参数不包含 AAAAAAAA , 返回原字符串: ", src)
		return src
	}

}
