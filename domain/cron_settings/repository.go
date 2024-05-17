package cron_settings

type CronSettingsRepository interface {
	CreateCronSettings(settings *CronSettings) error
	GetCronSettingsByID(id uint) (*CronSettings, error)
	GetCronSettingsByOrganizationID(orgID uint) (*CronSettings, error)
	UpdateCronSettings(settings *CronSettings) error
	DeleteCronSettings(id uint) error
}
