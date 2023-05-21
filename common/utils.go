package common

import (
	"crypto/aes"
	"crypto/cipher"

	crypto_rand "crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type RespCode int
type Role int

const (
	SUCCEED RespCode = 0
	FAILED           = -1
)
const APP_VERSION = "0.2.1"
const EN_KEY1 = "5c072181d566c72ff169"
const EN_KEY2 = "a624d7bbb012c6490df2653f5f0f4037c4d407d7291e"

func GetRespResult(code int, message string, dataList interface{}, totalCount int) interface{} {
	return GetRespResultWithPage(code, message, dataList, 1, 1, totalCount)
}

func GetRespResultWithPage(code int, message string, dataList interface{}, pageIndex int, totalPages int, totalCount int) interface{} {
	return gin.H{
		"code":       code,
		"dataList":   dataList,
		"pageIndex":  pageIndex,
		"totalPages": totalPages,
		"totalCount": totalCount,
		"message":    message,
	}
}

func GetEnvValue(envKey string, defaultValue string) (envValue string) {
	envValue = os.Getenv(envKey)
	if envValue == "" {
		envValue = defaultValue
	}
	return
}

const (
	PLATFORM_ADMIN Role = iota + 1 // value --> 1
	NORMAL_ADMIN
	NORMAL_MEMBER
	GROUP_ADMIN
	GROUP_MASTER
	GROUP_DEVELOP
)

func EncryptString(stringToEncrypt string) string {
	bytes := make([]byte, 32) //generate a random 32 byte key for AES-256
	if _, err := rand.Read(bytes); err != nil {
		log.Error(err)
		return stringToEncrypt
	}

	//Since the key is in string, we need to convert decode it to bytes
	key, _ := hex.DecodeString(EN_KEY1 + EN_KEY2)

	plaintext := []byte(stringToEncrypt)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Error(err)
		return stringToEncrypt
	}

	//Create a new GCM - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	//https://golang.org/pkg/crypto/cipher/#NewGCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		log.Error(err)
		return stringToEncrypt
	}

	//Create a nonce. Nonce should be from GCM
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(crypto_rand.Reader, nonce); err != nil {
		log.Error(err)
		return stringToEncrypt
	}

	//Encrypt the data using aesGCM.Seal
	//Since we don't want to save the nonce somewhere else in this case, we add it as a prefix to the encrypted data. The first nonce argument in Seal is the prefix.
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x", ciphertext)
}

func DecryptString(encryptedString string) string {
	if len(encryptedString) == 0 {
		return encryptedString
	}

	key, _ := hex.DecodeString(EN_KEY1 + EN_KEY2)
	enc, _ := hex.DecodeString(encryptedString)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Error(err)
		return encryptedString
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		log.Error(err)
		return encryptedString
	}

	//Get the nonce size
	nonceSize := aesGCM.NonceSize()

	//Extract the nonce from the encrypted data
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	//Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		log.Error(err)
		return encryptedString
	}

	return fmt.Sprintf("%s", plaintext)
}

func GetUtcNowTimeStampSec() int64 {
	//set timezone,
	return time.Now().Unix()
}
