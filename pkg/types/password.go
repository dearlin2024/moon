package types

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	cipher2 "github.com/aide-cloud/moon/pkg/utils/cipher"
	"google.golang.org/grpc/status"
)

var (
	defaultKey = "1234567890123456"
	defaultIv  = "1234567890123456"
)

var ErrValidatePassword = status.Error(http.StatusUnauthorized, "密码错误")

// SetDefaultKey 设置默认key
func SetDefaultKey(k string) {
	defaultKey = k
}

// SetDefaultIv 设置默认iv
func SetDefaultIv(iv string) {
	defaultIv = iv
}

func NewPassword(values ...string) Password {
	var value, salt string
	switch len(values) {
	case 1:
		value = values[0]
		salt = GenerateSalt()
	case 2:
		value = values[0]
		salt = values[1]
	default:
		salt = GenerateSalt()
		value = cipher2.GenerateRandomString(8, 0)
	}
	return &password{
		salt:  salt,
		value: value,
	}
}

type (
	Password interface {
		GetEncryptValue() (string, error)
		GetValue() string
		GetSalt() string
		Equal(old Password) bool
		fmt.Stringer
	}
	password struct {
		value, salt string
	}
)

func (p *password) GetSalt() string {
	return p.salt
}

func (p *password) String() string {
	v, _ := p.GetEncryptValue()
	return v
}

// GetValue 获取密码值
func (p *password) GetValue() string {
	return p.value
}

// GetEncryptValue 获取加密值
func (p *password) GetEncryptValue() (string, error) {
	return GeneratePassword(p.value, p.salt)
}

// Equal 密码是否相等
func (p *password) Equal(old Password) bool {
	return p.value == old.GetValue()
}

// GeneratePassword 生成密码
func GeneratePassword(password, salt string) (string, error) {
	aes, err := cipher2.NewAes(defaultKey, salt)
	if err != nil {
		return "", err
	}
	newPass, err := aes.EncryptBase64([]byte(password))
	if err != nil {
		return "", err
	}
	return newPass, nil
}

// GenerateSalt 生成盐
func GenerateSalt() string {
	return cipher2.MD5(strconv.FormatInt(time.Now().Unix(), 10))[0:16]
}

// DecryptPassword 解密密码
func DecryptPassword(password, salt string) (str string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = ErrValidatePassword
		}
	}()
	aes, err := cipher2.NewAes(defaultKey, salt)
	if err != nil {
		return "", err
	}
	newPass, err := aes.DecryptBase64(password)
	if err != nil {
		return "", err
	}
	return string(newPass), nil
}

// EncryptPassword 加密密码
func EncryptPassword(password, salt string) (str string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = ErrValidatePassword
		}
	}()
	aes, err := cipher2.NewAes(defaultKey, salt)
	if err != nil {
		return "", err
	}
	newPass, err := aes.EncryptBase64([]byte(password))
	if err != nil {
		return "", err
	}
	return newPass, nil
}

// ValidatePassword 验证密码
func ValidatePassword(p1, dePass, salt string) bool {
	newPass, err := DecryptPassword(dePass, salt)
	if err != nil {
		return false
	}

	return p1 == newPass
}

// ValidatePasswordErr 验证密码
func ValidatePasswordErr(p1, dePass, salt string) error {
	if ValidatePassword(p1, dePass, salt) {
		return nil
	}
	return ErrValidatePassword
}