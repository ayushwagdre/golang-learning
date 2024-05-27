package internal_authentication_learning

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/tuvistavie/securerandom"

	"ayush/config"
)

type Headers map[string]string

func GenerateHmacSignature() (string, string) {
	cfg := config.NewConfig()
	serviceID := cfg.Service.ID
	serviceKey := cfg.Service.Key
	nonce, _ := securerandom.Hex(16)
	key := fmt.Sprintf("%s-%s", nonce, serviceID)
	signature := hmac.New(sha1.New, []byte(serviceKey))
	signature.Write([]byte(key))
	serviceSignature := hex.EncodeToString(signature.Sum(nil))
	return nonce, serviceSignature
}

// headers from other service
func GetHeadersForInternalRequest() Headers {
	cfg := config.NewConfig()
	nonce, serviceSignature := GenerateHmacSignature()
	headers := Headers{
		"SIMPL-SERVICE-ID":        cfg.Service.ID,
		"SIMPL-SERVICE-NONCE":     nonce,
		"SIMPL-SERVICE-SIGNATURE": serviceSignature,
	}
	return headers
}

func CalculateHmacSignature(key, data string) string {
	hash := hmac.New(sha1.New, []byte(key))
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}

// validate hmac digest from other service
func validateHmacDigest(serviceID, nonce, serviceSignature string) bool {
	if serviceID == "" || nonce == "" || serviceSignature == "" {
		return false
	}

	serviceConfig := map[string]string{
		"shopify-backend-7197183c": "8be73de034dd4e5fa48bcaa183aed4bf",
	}
	message := strings.Join([]string{nonce, serviceID}, "-")
	digest := CalculateHmacSignature(serviceConfig[serviceID], message)

	return digest == serviceSignature
}

func InternalAuthenticationInit() {
	headers := GetHeadersForInternalRequest()
	fmt.Println(validateHmacDigest("shopify-backend-7197183c", headers["SIMPL-SERVICE-NONCE"], headers["SIMPL-SERVICE-SIGNATURE"]))
}
