package domain

type Cron interface {
	AddFunc(task func())
	Start()
	Stop()
}
