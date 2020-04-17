package cron

//
//import (
//	"fmt"
//	"log"
//	"sync"
//
//	"github.com/robfig/cron/v3"
//)
//
//type TimerTrigger struct {
//	sync.Map
//
//	*cron.Cron
//}
//
//func NewTimerTrigger() *TimerTrigger {
//	return &TimerTrigger{
//		Cron: cron.New(),
//	}
//}
//
//func (tt *TimerTrigger) InstallFunc(serviceName, functionName, spec string) error {
//	cmd := func() {
//		resp, err := DoFunc(NewContext(serviceName, functionName))
//		if err != nil {
//			log.Println(err)
//			return
//		}
//
//		log.Println(resp.StatusCode(), resp.Result())
//	}
//
//	entryID, err := tt.AddFunc(spec, cmd)
//	if err != nil {
//		return err
//	}
//
//	tt.Store(serviceName+functionName, entryID)
//	return nil
//}
//
//func (tt *TimerTrigger) UninstallFunc(service, name string) error {
//	v, ok := tt.Load(service + name)
//	if !ok {
//		return fmt.Errorf("func %s not exist", name)
//	}
//
//	tt.Remove(v.(cron.EntryID))
//	tt.Delete(service + name)
//	return nil
//}
//
//// 由Leader进行调度（指定执行者）
//
//// 触发器产生事件，事件转换成HTTP请求，HTTP请求发送给目标服务
//
//// 上传bin文件或者image，上传bin文件本地运行，上传image在k8s上运行
//type Server interface {
//}
//
//// 本地运行环境
//type LocalServer struct {
//	Port     uint16
//	Filepath string
//}
//
//// K8S运行环境
//type K8SServer struct {
//}
