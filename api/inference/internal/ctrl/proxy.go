package ctrl

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/0glabs/0g-serving-broker/common/errors"
	"github.com/0glabs/0g-serving-broker/common/util"
	constant "github.com/0glabs/0g-serving-broker/inference/const"
	"github.com/0glabs/0g-serving-broker/inference/model"
)



func (c *Ctrl) PrepareHTTPRequest(ctx *gin.Context, targetURL string, reqBody []byte) (*http.Request, error) {
	req, err := http.NewRequest(ctx.Request.Method, targetURL, io.NopCloser(bytes.NewBuffer(reqBody)))
	if err != nil {
		return nil, err
	}

	for k, v := range ctx.Request.Header {
		if _, ok := constant.RequestMetaData[k]; !ok {
			req.Header.Set(k, v[0])
			continue
		}
	}

	// may need additional secret to access the target service
	if additionalSecret := c.Service.AdditionalSecret; additionalSecret != nil {
		for k, v := range additionalSecret {
			req.Header.Set(k, v)
		}
	}

	return req, nil
}

func (c *Ctrl) ProcessHTTPRequest(ctx *gin.Context, svcType string, req *http.Request, reqModel model.Request, fee string, outputPrice int64, charing bool) {
	client := &http.Client{}

	// back up body for other usage
	body, err := io.ReadAll(req.Body)
	if err != nil {
		handleBrokerError(ctx, err, "failed to read request body")
		return
	}
	req.Body = io.NopCloser(bytes.NewBuffer(body))

	resp, err := client.Do(req)
	if err != nil {
		handleBrokerError(ctx, err, "call proxied service")
		return
	}
	defer resp.Body.Close()

	for k, v := range resp.Header {
		if k == "Content-Length" {
			continue
		}
		ctx.Writer.Header()[k] = v
	}

	if resp.StatusCode != http.StatusOK {
		ctx.Writer.WriteHeader(resp.StatusCode)
		handleServiceError(ctx, resp.Body)
		return
	}

	ctx.Writer.Header().Add("provider", c.contract.ProviderAddress)
	c.addExposeHeaders(ctx)

	ctx.Status(resp.StatusCode)

	if !charing {
		c.handleResponse(ctx, resp)
		return
	}

	oldAccount, err := c.GetOrCreateAccount(ctx, reqModel.UserAddress)
	if err != nil {
		handleBrokerError(ctx, err, "")
		return
	}
	unsettledFee, err := util.Add(fee, oldAccount.UnsettledFee)
	if err != nil {
		handleBrokerError(ctx, err, "add unsettled fee")
		return
	}

	account := model.User{
		User:             reqModel.UserAddress,
		LastRequestNonce: &reqModel.Nonce,
		UnsettledFee:     model.PtrOf(unsettledFee.String()),
	}

	switch svcType {
	case "chatbot":
		c.handleChatbotResponse(ctx, resp, account, outputPrice, body)
	default:
		handleBrokerError(ctx, errors.New("unknown service type"), "prepare request extractor")
	}
}

func (c *Ctrl) handleResponse(ctx *gin.Context, resp *http.Response) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		handleBrokerError(ctx, err, "read from body")
		return
	}
	if _, err := ctx.Writer.Write(body); err != nil {
		handleBrokerError(ctx, err, "write response body")
	}
}

func (c *Ctrl) addExposeHeaders(ctx *gin.Context) {
	// Set 'Access-Control-Expose-Headers' for CORS
	exposeHeaders := []string{"Provider", "content-encoding"}
	existing := ctx.Writer.Header().Get("Access-Control-Expose-Headers")
	var newHeaders string
	if existing != "" {
		headerSet := make(map[string]struct{})
		for _, header := range strings.Split(existing, ",") {
			headerSet[strings.TrimSpace(header)] = struct{}{}
		}

		for _, header := range exposeHeaders {
			if _, exists := headerSet[header]; !exists {
				existing += "," + header
			}
		}

		newHeaders = existing
	} else {
		newHeaders = strings.Join(exposeHeaders, ",")
	}
	ctx.Writer.Header().Set("Access-Control-Expose-Headers", newHeaders)
}

func handleBrokerError(ctx *gin.Context, err error, context string) {
	// TODO: recorded to log system
	info := "Provider proxy: handle proxied service response"
	if context != "" {
		info += (", " + context)
	}
	errors.Response(ctx, errors.Wrap(err, info))
}

func handleServiceError(ctx *gin.Context, body io.ReadCloser) {
	respBody, err := io.ReadAll(body)
	if err != nil {
		// TODO: recorded to log system
		log.Println(err)
		return
	}
	if _, err := ctx.Writer.Write(respBody); err != nil {
		// TODO: recorded to log system
		log.Println(err)
	}
}
