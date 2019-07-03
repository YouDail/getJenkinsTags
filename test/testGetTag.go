package main

import (
	"context"
	"flag"
	log "github.com/golang/glog"
	"github.com/ha666/golibs"
	"github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	proto "github.com/YouDail/getJenkinsTags/proto"
	"os"
)

func init() {

	//initial log
	_ = os.MkdirAll("logs", 0766)
	_ = flag.Set("alsologtostderr", "true")
	_ = flag.Set("stderrthreshold", "INFO")
	_ = flag.Set("log_dir", "logs")
	flag.Parse()

	defer log.Flush()
}

func main() {

	TagCli = GetTaglient("10.52.26.3:8500", "com.hyphen.sops.srv.getPorjTags.prod")
	if TagCli == nil {
		panic("init DeptUserClient initialize error")
	}

	a := []string{"infr/sops-rbac", "infr/sops-message", "infr/sops-login", "infr/sops-job"}

	rsp, err := TagCli.GetProjTags(context.TODO(), &proto.ProjTagsRequest{
		ProjUrl: a})

	if err != nil {
		log.Errorf("err:", err)
	}
	log.Infof("jobV2  GetProjectList  GetListRequest 请求返回数据: ", golibs.ToJson(rsp))

}

var (
	TagCli proto.ProjTagsService
)

func GetTaglient(registryName, serveName string) proto.ProjTagsService {
	os.Setenv("MICRO_CLIENT_REQUEST_TIMEOUT", "1m")
	srv := grpc.NewService(
		micro.Name("getPorjTags.client"),
		micro.Registry(
			consul.NewRegistry(
				registry.Addrs(registryName),
			),
		),
	)
	srv.Init()
	return proto.NewProjTagsService(serveName, srv.Client())
}
