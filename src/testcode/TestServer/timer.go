// 定时器逻辑
// linux crontab

package testserver

import (
	cron "gopkg.in/robfig/cron.v2"
)

type Timer struct {
	c *cron.Cron
}

var t *Timer

func init() {
	t = &Timer{c: cron.New()}
	t.c.Start()
}

func (t *Timer) AddFunc(spec string, cmd func()) (id int, err error) {
	eneryId, err := t.c.AddFunc(spec, cmd)
	if err != nil {
		return
	}
	id = int(eneryId)
	return
}

func (t *Timer) Remove(id int) {
	t.c.Remove(cron.EntryID(id))
}

func (t *Timer) Cron() *cron.Cron {
	return t.c
}

func AddFunc(spec string, cmd func()) (id int, err error) {
	return t.AddFunc(spec, cmd)
}

func Remove(id int) {
	t.Remove(id)
}

func Cron() *cron.Cron {
	return t.Cron()
}
