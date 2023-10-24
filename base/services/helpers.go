package services

import (
	"Society-Synergy/base/models"
	"context"
	"unicode"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	LogCollection *mongo.Collection
	Ctx           context.Context
)

func IsPasswordValid(s string) bool {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(s) >= 7 {
		hasMinLen = true
	}
	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}

func RegisterLogS(log *models.AuditLogs) (string, error) {
	_, err := LogCollection.InsertOne(Ctx, log)
	if err != nil {
		return "", err
	}
	return "", nil
}
