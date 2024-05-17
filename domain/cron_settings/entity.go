package cron_settings

// CronSettings представляет собой сущность настроек крона
type CronSettings struct {
	ID             uint
	OrganizationID uint
	CronExpression string
}
