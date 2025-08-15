package otp

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/md-asharaf/go-fiber-boilerplate/internal/services/email"
)

type OtpService struct {
	emailService *email.EmailService
}

func NewOtpService(emailService *email.EmailService) *OtpService {
	return &OtpService{
		emailService: emailService,
	}
}

func GenerateOtp() string {
	rand.Seed(time.Now().UnixNano())
	otp := rand.Intn(1000000) // Generate a random 6-digit number
	return strconv.Itoa(otp)
}

func (o *OtpService) SendOtp(email, subject string) error {
	otp := GenerateOtp()
	if subject == "" {
		subject = "Your OTP Code"
	}
	body := "Your OTP code is: " + otp
	return o.emailService.SendEmail(email, subject, body)
}
