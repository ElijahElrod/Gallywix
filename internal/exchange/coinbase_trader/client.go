package coinbase_trader

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/elijahelrod/vespene/config"
	"github.com/elijahelrod/vespene/pkg/logger"
	"github.com/elijahelrod/vespene/pkg/model"
)

const Post = "POST"
const OrderPath = "/orders"
const CancelOrderPath = OrderPath + "/batch_cancel"

type ExchangeService struct {
	exchangeCfg config.ExchangeConfig
	logger      logger.Logger
	httpClient  http.Client
}

func NewExchangeService(exCfg config.ExchangeConfig, logger logger.Logger) *ExchangeService {
	return &ExchangeService{
		exchangeCfg: exCfg,
		logger:      logger,
		httpClient:  http.Client{},
	}
}

// PlaceOrder takes the productId, side (Buy/Sell), size, and price to place an order to coinbase; also
// generates custom headers off the [config.ExchangeConfig]
func (es *ExchangeService) PlaceOrder(productId, side, size, price string) {
	var accessKey = es.exchangeCfg.AccessKey
	var accessPassphrase = es.exchangeCfg.AccessPassphrase
	var accessSecret = es.exchangeCfg.AccessSecret
	var timestamp = strconv.Itoa(int(time.Now().UnixNano()))

	orderBody, err := json.Marshal(model.OrderBody{
		ProductId: productId,
		Side:      side,
		Size:      size,
		Price:     price,
	})

	if err != nil {
		es.logger.Error(err)
		return
	}

	// Create pre-hashed string
	var message = timestamp + Post + OrderPath + string(orderBody[:])

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

	// Create Reader for sending POST Request to place the order
	bodyReader := bytes.NewReader(orderBody)
	req, err := http.NewRequest(Post, es.exchangeCfg.Url+OrderPath, bodyReader)

	if err != nil {
		es.logger.Error(err)
		return
	}

	// [Required Coinbase Headers]: https://docs.cloud.coinbase.com/exchange/docs/rest-auth
	req.Header.Add("CB-ACCESS-KEY", accessKey)
	req.Header.Add("CB-ACCESS-SIGN", signedAccessStr)
	req.Header.Add("CB-ACCESS-TIMESTAMP", accessKey)
	req.Header.Add("CB-ACCESS-PASSPHRASE", accessPassphrase)

	// Send order request
	res, err := es.httpClient.Do(req)
	if err != nil {
		es.logger.Error(err)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			es.logger.Error(err)
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		es.logger.Error(err)
		return
	}

	// TODO: Add DB Write here or something for order tracking and cancelling later
	es.logger.Info("Made request: " + string(body))
}

// CancelOrder sends a POST request to cancel one of more unfilled orders
// it generates headers off the [config.ExchangeConfig]
func (es *ExchangeService) CancelOrder() error {

	var timestamp = strconv.Itoa(int(time.Now().UnixNano()))

	orderBody, err := json.Marshal(model.OrderBody{
		ProductId: "",
		Side:      "",
		Size:      "",
		Price:     "",
	})

	if err != nil {
		es.logger.Error(err)
		return err
	}

	// Create pre-hashed string
	_ = timestamp + Post + CancelOrderPath + string(orderBody[:])
	return nil
}
