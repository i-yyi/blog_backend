package service

import (
	config "codehub.devcloud.cn-north-4.huaweicloud.com/taqse2020_mihoyo00001/blog_backend/conf"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	log "github.com/sirupsen/logrus"
)

func encryptPassword(password string) string {
	h := sha256.New()
	h.Write([]byte(password))
	return hex.EncodeToString(h.Sum(nil))
}

// generateToken 生成token
func generateToken(username string) (string, error) {
	return "testToken", nil
	//claims := &model.Claims{
	//	Username: username,
	//	StandardClaims: jwt.StandardClaims{
	//		ExpiresAt: time.Now().Add(time.Duration(config.GetInt64("jwt.expire")) * time.Second).Unix(),
	//	},
	//}
	//
	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//return token.SignedString([]byte(config.GetString("jwt.secret")))
}
func Login(username, password string) (string, error) {
	encodedPwd := encryptPassword(password) // 对输入的密码进行加密处理
	if username == config.GetConf().Admin.Username && encodedPwd == config.GetConf().Admin.Password {
		log.Info("Login Success")
		return generateToken(config.GetConf().Admin.Username)
	}
	log.Info("Login Failed, Pass %s, enc %s", config.GetConf().Admin.Password, encodedPwd)
	return "", errors.New("gen token failed")
}
