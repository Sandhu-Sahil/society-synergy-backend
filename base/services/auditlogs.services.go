package services

import "Society-Synergy/base/models"

func (service *ServiceLogsImpl) RegisterLog(log *models.AuditLogs) (string, error) {
	// _, err := service.logcollection.InsertOne(service.ctx, log)
	// if err != nil {
	// 	return "", err
	// }
	return "", nil
}

func (service *ServiceLogsImpl) GetLog(*models.AuditLogs) (string, error) {
	return "", nil
}
