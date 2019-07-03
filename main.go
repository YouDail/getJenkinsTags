package main

import (
	"flag"
	"fmt"
	log "github.com/golang/glog"
	"github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	"github.com/spf13/viper"
	"github.com/YouDail/getJenkinsTags/handler"
	proto "github.com/YouDail/getJenkinsTags/proto"
	"os"
	"time"
)

//检查配置项
func VaildConf(s string) {

	if viper.GetString(s) == "" {
		log.Errorln("init", s, "  config is null")
		panic("启动失败！ 缺少配置参数:" + s)
	} else {
		log.Infoln("init 配置参数 ", s, "的值是: ", viper.GetString(s))
	}

}

func init() {

	log.Infoln("initial log config")
	_ = os.MkdirAll("logs", 0766)
	_ = flag.Set("alsologtostderr", "true")
	_ = flag.Set("stderrthreshold", "INFO")
	_ = flag.Set("log_dir", "logs")
	flag.Parse()

	defer log.Flush()

	curPath, err := os.Getwd()
	if err != nil {
		log.Errorln(err)
	}
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // or viper.SetConfigType("YAML")
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	viper.AddConfigPath(curPath)
	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		log.Errorln(fmt.Errorf("Fatal error config file config.yaml : %s \n", err))
		panic("请检查配置文件config.yaml是否正确")
	}

	confs := []string{
		"BasePath",
		"serviceName",
		"registry.type",
		"registry.addr",
	}

	for _, v := range confs {
		VaildConf(v)
	}
}

func main() {

	log.Infoln("create n new service")

	os.Setenv("MICRO_REGISTRY", viper.GetString("registry.type"))
	os.Setenv("MICRO_REGISTRY_ADDRESS", viper.GetString("registry.addr"))
	os.Setenv("MICRO_SERVER_NAME", viper.GetString("serviceName"))

	service := grpc.NewService(
		micro.RegisterTTL(time.Second*60),
		micro.RegisterInterval(time.Second*60),
	)
	service.Init()

	log.Infoln("register handler to consul")
	proto.RegisterProjTagsHandler(service.Server(), new(handler.ProjUrl))

	log.Infoln("run the server")
	err := service.Run()
	if err != nil {
		panic(err)
	}
}
