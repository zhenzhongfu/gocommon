package cron

import (
	"os"
	"sync"

	"github.com/zhenzhongfu/gocommon/logging"

	"github.com/robfig/cron"
)

type CronSrv struct {
	c  *cron.Cron
	ch chan os.Signal
	mu sync.Mutex
}

func NewCronSrv() *CronSrv {
	return &CronSrv{
		c:  cron.New(),
		ch: make(chan os.Signal),
	}
}

func (srv *CronSrv) Start() error {
	logging.Infoln("Starting...")

	srv.mu.Lock()
	srv.c.Start()
	srv.mu.Unlock()

	select {
	case <-srv.ch:
		logging.Infoln("cron stop.")
		return nil
	}
	return nil
}

func (srv *CronSrv) Stop() error {
	logging.Infoln("cron server stoping...")

	srv.mu.Lock()
	srv.c.Stop()
	srv.mu.Unlock()

	srv.ch <- os.Interrupt
	return nil
}

func (srv *CronSrv) AddJob(cmd string, command func()) error {
	srv.mu.Lock()
	srv.c.AddFunc(cmd, command)
	srv.mu.Unlock()
	return nil
}
