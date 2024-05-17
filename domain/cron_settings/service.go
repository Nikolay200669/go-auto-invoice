package cron_settings

import "errors"

type CronSettingsService struct {
	repository CronSettingsRepository
}

func NewCronSettingsService(repo CronSettingsRepository) *CronSettingsService {
	return &CronSettingsService{repository: repo}
}

func (s *CronSettingsService) CreateCronSettings(settings *CronSettings) error {
	if settings.CronExpression == "" {
		return errors.New("cron expression is required")
	}

	return s.repository.CreateCronSettings(settings)
}

func (s *CronSettingsService) GetCronSettingsByOrganizationID(orgID uint) (*CronSettings, error) {
	return s.repository.GetCronSettingsByOrganizationID(orgID)
}

func (s *CronSettingsService) UpdateCronSettings(settings *CronSettings) error {
	if settings.CronExpression == "" {
		return errors.New("cron expression is required")
	}

	return s.repository.UpdateCronSettings(settings)
}

func (s *CronSettingsService) DeleteCronSettings(id uint) error {
	return s.repository.DeleteCronSettings(id)
}
