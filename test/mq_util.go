package main

import (
	"encoding/json"
	"fmt"
	"github.com/volcengine/volc-sdk-golang/base"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	jsoniter "github.com/json-iterator/go"
)

type Topics map[int64]ResourceTopics

func (t Topics) GetTopics(accountId, resourceType int64) []Topic {
	resourceTopics, exist := t[accountId]
	if !exist || resourceTopics == nil || len(resourceTopics) == 0 {
		return nil
	}

	topics, exist := resourceTopics[resourceType]
	if !exist || topics == nil || len(topics) == 0 {
		return nil
	}
	return topics
}

type ResourceTopics map[int64][]Topic

type Topic struct {
	Cluster string `json:"cluster"`
	Topic   string `json:"topic"`
}

type BmqHandler struct {
	accessKey       string
	secretAccessKey string
	region          string
	service         string
	apiHost         string
}

func NewBmqHandler(accessKey, secretAccessKey, region, service, apiHost string) BmqHandler {
	return BmqHandler{
		accessKey:       accessKey,
		secretAccessKey: secretAccessKey,
		region:          region,
		service:         service,
		apiHost:         apiHost,
	}
}

func (b BmqHandler) getPartitionNum(topic Topic) (int, error) {
	url := fmt.Sprintf("https://%s/?Action=GetGMSMetaTopicDetail&Version=2022-06-01&FullName=%s&TopicName=%s", b.apiHost, topic.Cluster, topic.Topic)
	resp, err := b.send(url)
	if err != nil {
		return 0, err
	}

	if resp.Result != nil {
		return Interface2Int(resp.Result["PartitionCount"], 0), nil
	} else {
		return 0, fmt.Errorf("result is empty")
	}
}

type Response struct {
	ResponseMetadata map[string]interface{} `json:"ResponseMetadata"`
	Result           map[string]interface{} `json:"Result"`
}

func (b BmqHandler) send(url string) (*Response, error) {
	credentials := base.Credentials{
		AccessKeyID:     b.accessKey,
		SecretAccessKey: b.secretAccessKey,
		Region:          b.region,
		Service:         b.service,
	}

	req, _ := http.NewRequest("GET", url, nil)

	timeout := 10 * time.Second
	httpClient := http.Client{Timeout: timeout}
	signRequest := credentials.Sign(req) // 签名
	httpResponse, err := httpClient.Do(signRequest)
	if err != nil {
		return nil, err
	}

	defer func() {
		if httpResponse != nil && httpResponse.Body != nil {
			httpResponse.Body.Close()
		}
	}()

	if httpResponse.StatusCode != http.StatusOK {
		response, err := ioutil.ReadAll(httpResponse.Body)
		if err != nil {
			return nil, err
		}

		fmt.Println(string(response))

		err = fmt.Errorf("failed to get status OK response (status code: %d)", httpResponse.StatusCode)
		return nil, err
	}

	response, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}

	var resp Response
	if err = json.Unmarshal(response, &resp); err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	return &resp, err
}

func Interface2Int(iInt interface{}, defaultValue int) int {
	switch iInt.(type) {
	case float32:
		return int(iInt.(float32))
	case float64:
		return int(iInt.(float64))
	case int:
		return int(iInt.(int))
	case int32:
		return int(iInt.(int32))
	case int64:
		return int(iInt.(int64))
	case string:
		if iInt.(string) == "" {
			return defaultValue
		}
		num, _ := strconv.ParseInt(iInt.(string), 10, 64)
		return int(num)
	case json.Number:
		num, _ := iInt.(json.Number).Int64()
		return int(num)
	case jsoniter.Number:
		num, _ := iInt.(jsoniter.Number).Int64()
		return int(num)
	default:
		return defaultValue
	}
}
