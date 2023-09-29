package services

import (
	"Society-Synergy/base/models"
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	sendinblue "github.com/sendinblue/APIv3-go-library/v2/lib"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (us *ServiceUserImpl) SendOTP(userID string) error {
	objectid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	query := bson.D{bson.E{Key: "_id", Value: objectid}}
	var userFound *models.User
	err = us.usercollection.FindOne(us.ctx, query).Decode(&userFound)
	if err != nil {
		return err
	}

	rand.Seed(time.Now().UnixNano())
	otpChars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	otpLength := 8
	otp := generateOTP(otpChars, otpLength)

	userFound.OTP = otp
	// kolkata timezone + 5 minutes
	userFound.OTPExpiry = time.Now().UTC().Add(time.Hour * 5).Add(time.Minute * 35)
	_, err = us.usercollection.UpdateOne(us.ctx, query, bson.D{bson.E{Key: "$set", Value: userFound}})
	if err != nil {
		return err
	}

	err = godotenv.Load()
	if err != nil {
		return err
	}

	apiKey := os.Getenv("API_KEY")
	apisecret := os.Getenv("API_SECRET")

	var ctx context.Context
	cfg := sendinblue.NewConfiguration()
	//Configure API key authorization: api-key
	cfg.AddDefaultHeader("api-key", apiKey)
	//Configure API key authorization: partner-key
	cfg.AddDefaultHeader("partner-key", apisecret)

	sib := sendinblue.NewAPIClient(cfg)
	_, resp, err := sib.TransactionalEmailsApi.SendTransacEmail(ctx, sendinblue.SendSmtpEmail{
		Sender: &sendinblue.SendSmtpEmailSender{
			Name:  "Society Synergy",
			Email: "sahil.sandhu@societysynergy.com",
		},
		To: []sendinblue.SendSmtpEmailTo{
			{
				Name:  userFound.UserName,
				Email: userFound.Email,
			},
		},
		HtmlContent: generateHtmlContent(otp),
		Subject:     "OTP verification, Society Synergy",
		ReplyTo: &sendinblue.SendSmtpEmailReplyTo{
			Name:  "Society Synergy",
			Email: "9921103131@mail.jiit.ac.in",
		},
	})
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated {
		return nil
	} else {
		return fmt.Errorf("unexpected status code returned: %d", resp.StatusCode)
	}
}

func generateOTP(otpChars string, length int) string {
	otp := make([]byte, length)
	for i := 0; i < length; i++ {
		otp[i] = otpChars[rand.Intn(len(otpChars))]
	}
	return string(otp)
}

func (us *ServiceUserImpl) VerifyOTP(userID string, otp string) error {
	objectid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	query := bson.D{bson.E{Key: "_id", Value: objectid}}
	var userFound *models.User
	err = us.usercollection.FindOne(us.ctx, query).Decode(&userFound)
	if err != nil {
		return err
	}

	if userFound.OTP != otp {
		return fmt.Errorf("invalid OTP")
	}

	if time.Now().UTC().Add(time.Hour * 5).Add(time.Minute * 30).After(userFound.OTPExpiry) {
		return fmt.Errorf("OTP expired")
	}

	if !userFound.Varified {
		userFound.Varified = true
		_, err = us.usercollection.UpdateOne(us.ctx, query, bson.D{bson.E{Key: "$set", Value: userFound}})
		if err != nil {
			return err
		}
	}

	return nil
}
