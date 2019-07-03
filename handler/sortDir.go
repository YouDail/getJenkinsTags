package handler

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/YouDail/getJenkinsTags/common"
	"io/ioutil"
)

func SortDir(str string) ([]string, error) {
	stre, err := ioutil.ReadDir(str)
	if err != nil {
		glog.Errorln("目录不存在:", err)
		return nil, err
	}
	var (
		timeSlice []int64
		dirMap    []map[int64]string
	)
	for _, v := range stre {
		if v.IsDir() {
			timeSlice = append(timeSlice, v.ModTime().Unix())
			p := make(map[int64]string)
			p[v.ModTime().Unix()] = v.Name()
			dirMap = append(dirMap, p)
		}
	}

	sortRes := common.IntArray(timeSlice)
	common.Sort(sortRes)
	fmt.Println(sortRes)

	dat := []string{}
	for i, _ := range sortRes {
		for _, node := range dirMap {
			if node[sortRes[len(sortRes)-i-1]] != "" {
				fmt.Println("node[a[len(a)-i]: ", node[sortRes[len(sortRes)-i-1]])
				dat = append(dat, node[sortRes[len(sortRes)-i-1]])
			}
		}
	}
	fmt.Println("排序完成,返回结果: ", dat)
	return dat, nil
}
