package middleware

import (
	"farmacare/shared/config"
	"farmacare/shared/dto"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/golang-module/carbon"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Middleware struct {
	Env *config.EnvConfig
	log *logrus.Logger
	DB  *gorm.DB
}

func (m *Middleware) AuthMiddleware(c *fiber.Ctx) error {
	claims, err := m.getToken(m.Env.SecretKey, c)
	if err != nil {
		c.Locals("context", dto.SessionContext{})
		return c.Next()
	}

	userId := uint(claims["id"].(float64))

	context, err := m.getSessionContext(userId)
	if err != nil {
		return err
	}

	c.Locals("context", context)

	return c.Next()
}

func (m *Middleware) getSessionContext(userId uint) (dto.SessionContext, error) {
	var (
		user    dto.User
		context dto.SessionContext
	)
	err := m.DB.Where("id = ?", userId).Find(&user).Error

	if err != nil {
		return context, err
	}

	context = dto.SessionContext{
		User: user,
	}

	// if dto.GetSubscriptionStatus(user) {
	// 	context.IsActiveSubscriber = true
	// }

	return context, nil
}

func (m *Middleware) UserMiddleware(c *fiber.Ctx) error {
	claims, err := m.getToken(m.Env.SecretKey, c)

	if err != nil {
		return err
	}

	if claims["role"] != "user" {
		return fiber.ErrForbidden
	}

	return c.Next()
}

func (m *Middleware) DoctorMiddleware(c *fiber.Ctx) error {
	claims, err := m.getToken(m.Env.SecretKey, c)

	if err != nil {
		return err
	}

	if claims["role"] != "doctor" {
		return fiber.ErrForbidden
	}

	return c.Next()
}

func (m *Middleware) PharmacistMiddleware(c *fiber.Ctx) error {
	claims, err := m.getToken(m.Env.SecretKey, c)
	if err != nil {
		return err
	}

	if claims["role"] != "pharmacist" {
		return fiber.ErrForbidden
	}

	return c.Next()
}

func (m *Middleware) AdminMiddleware(c *fiber.Ctx) error {
	claims, err := m.getToken(m.Env.SecretKey, c)

	if err != nil {
		return err
	}

	if claims["role"] != "admin" {
		return fiber.ErrForbidden
	}

	return c.Next()
}

func (m *Middleware) getToken(secret string, c *fiber.Ctx) (jwt.MapClaims, error) {
	header := c.Get("Authorization", "")

	if header == "" {
		return nil, fiber.ErrUnauthorized
	}

	token, err := jwt.Parse(header, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fiber.ErrUnauthorized
	}

	m.log.Infof("Attempting auth check with header claims: %s", claims)

	exp, ok := claims["exp"].(float64)
	if !ok || carbon.Now().Timestamp() > int64(exp) {
		return nil, fiber.ErrUnauthorized
	}

	return claims, nil
}

func NewMiddleware(env *config.EnvConfig, log *logrus.Logger, db *gorm.DB) *Middleware {
	return &Middleware{
		Env: env,
		log: log,
		DB:  db,
	}
}