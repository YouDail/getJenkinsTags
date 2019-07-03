package handler

import (
	"context"
	"github.com/golang/glog"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	proto "github.com/YouDail/getJenkinsTags/proto"
	"strings"
)

type ProjUrl struct{}

func (g *ProjUrl) GetProjTags(ctx context.Context, req *proto.ProjTagsRequest, rsp *proto.ProjTagsResponse) error {
	glog.Infoln("GetMaxClass client request: ", req.ProjUrl)

	if len(req.ProjUrl) == 0 {
		return errors.New("请求参数有误!")
	}

	var (
		tags []string
	)

	BasePath := viper.GetString("BasePath")
	for _, url := range req.ProjUrl {
		if strings.Contains(url, "/") {
			glog.Infoln("req.ProjUrl url:", url)

			var (
				groupName   string
				projectName string
			)

			if strings.Contains(url, "gitlab.hyphen.com/") {
				NewUrl := strings.Split(url, "/")
				groupName = NewUrl[len(NewUrl)-2]
				projectName = NewUrl[len(NewUrl)-1]
			} else if strings.Contains(url, "/") && len(strings.Split(url, "/")) == 2 {
				groupName = strings.Split(url, "/")[0]
				projectName = strings.Split(url, "/")[1]
			} else {
				glog.Error("请求参数有误, ", url)
				rsp.TagsList = nil
				continue
			}

			newPath := BasePath + groupName + "/" + projectName

			node := proto.TagList{
				ProjectName: projectName,
			}
			tagList, err := SortDir(newPath)
			if err != nil {
				glog.Errorln("调用SortDir失败! ", err)
				node.Tags = nil
				continue
			}

			if len(tagList) > 20 {
				node.Tags = tagList[len(tagList)-20:]
			} else {
				node.Tags = tagList
			}

			rsp.TagsList = append(rsp.TagsList, &node)
		} else {
			return errors.New("请求参数有误!")
		}
	}
	glog.Infoln("tags:", tags)

	return nil
}

type Nods struct {
	ProjectName string   `json:"projectName"`
	Tags        []string `json:"tags"`
}
