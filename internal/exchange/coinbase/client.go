package coinbase

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"github.com/elijahelrod/vespene/config"
	"github.com/elijahelrod/vespene/pkg/logger"
	"github.com/elijahelrod/vespene/pkg/models"
	"net/http"
	"strconv"
	"time"
)

const POST = "POST"
const ORDER_PATH = "/orders"

type ExchangeService struct {
	exchangeCfg config.ExchangeConfig
	logger      logger.Logger
	httpClient  http.Client
}

func (es *ExchangeService) PlaceOrder(productId, side, size, price string) {
	var accessKey = es.exchangeCfg.AccessKey
	var accessPassphrase = es.exchangeCfg.AccessPassphrase
	var accessSecret = es.exchangeCfg.AccessSecret // Switch this to AccessSecret later
	var timestamp = strconv.Itoa(int(time.Now().UnixNano()))

	orderBody, err := json.Marshal(&models.OrderBody{
		productId,
		side,
		size,price
	})
	if err != nil {
		es.logger.Error(err)
		return
	}

	// Create pre-hashed string
	var message = timestamp + POST + ORDER_PATH + string(orderBody[:])

	// Decode the base64 access secret
	var decodedAccessSecret []byte
	_, err = base64.StdEncoding.Decode(decodedAccessSecret, []byte(accessSecret))
	if err != nil {
		es.logger.Error(err)
		return
	}

	// Create a SHA256 Hmac with the decodedAccessSecret
	hmacKey := hmac.New(sha256.New, decodedAccessSecret)

	// Sign the message with the hmac and base64 encode the result
	var signedAccess []byte
	base64.StdEncoding.Encode(signedAccess, hmacKey.Sum([]byte(message)))
	var signedAccessStr = string(signedAccess[:])

	bodyReader := bytes.NewReader(orderBody)

	req, err := http.NewRequest(POST, "http://localhost:8080", bodyReader)

	if err != nil {
		es.logger.Error(err)
		return
	}

	// Required Coinbase Headers
	req.Header.Add("CB-ACCESS-KEY", accessKey)
	req.Header.Add("CB-ACCESS-SIGN", signedAccessStr)
	req.Header.Add("CB-ACCESS-TIMESTAMP", accessKey)
	req.Header.Add("CB-ACCESS-PASSPHRASE", accessPassphrase)
	_, err = es.httpClient.Do(req)
	if err != nil {
		es.logger.Error(err)
		return
	}

}
