package util

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/elastic/go-elasticsearch/v9"
	"gopkg.in/ini.v1"
)

var ESClient *elasticsearch.Client

func init() {
	config, iniErr := ini.Load("./config/app.ini")
	if iniErr != nil {
		fmt.Printf("Fail to read file: %v", iniErr)
		os.Exit(1)
	}

	enable, _ := config.Section("elasticsearch").Key("enable").Bool()
	if !enable {
		fmt.Println("Elasticsearch未启用")
		return
	}

	addr := config.Section("elasticsearch").Key("addr").String()
	username := config.Section("elasticsearch").Key("username").String()
	password := config.Section("elasticsearch").Key("password").String()

	cfg := elasticsearch.Config{
		Addresses: []string{addr},
		Username:  username,
		Password:  password,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
			ForceAttemptHTTP2:     false,
			DialContext:           (&net.Dialer{Timeout: 10 * time.Second}).DialContext,
			TLSHandshakeTimeout:   10 * time.Second,
			ResponseHeaderTimeout: 10 * time.Second,
		},
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		fmt.Printf("Elasticsearch客户端创建失败: %v\n", err)
		return
	}

	info, err := client.Info()
	if err != nil {
		fmt.Printf("Elasticsearch连接失败: %v\n", err)
		return
	}
	defer info.Body.Close()

	fmt.Println("Elasticsearch连接成功!")

	ESClient = client
}

type OrderItemDoc struct {
	OrderId      int     `json:"order_id"`
	Uid          int     `json:"uid"`
	ProductTitle string  `json:"product_title"`
	ProductPrice float64 `json:"product_price"`
}

func IndexOrderItem(doc OrderItemDoc) {
	if ESClient == nil {
		return
	}
	body, _ := json.Marshal(doc)
	docId := fmt.Sprintf("%d_%d", doc.OrderId, doc.Uid)
	res, err := ESClient.Index(
		"order_items",
		bytes.NewReader(body),
		ESClient.Index.WithDocumentID(docId),
		ESClient.Index.WithContext(context.Background()),
	)
	if err == nil && res != nil {
		res.Body.Close()
	}
}

type esOrderHit struct {
	Source OrderItemDoc `json:"_source"`
}
type esOrderHits struct {
	Hits []esOrderHit `json:"hits"`
}
type esOrderResponse struct {
	Hits esOrderHits `json:"hits"`
}

func SearchOrderItems(userId int, keywords string) []int {
	if ESClient == nil {
		return nil
	}

	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					{"term": map[string]interface{}{"uid": userId}},
					{"match": map[string]interface{}{"product_title": keywords}},
				},
			},
		},
		"size": 100,
	}
	json.NewEncoder(&buf).Encode(query)

	res, err := ESClient.Search(
		ESClient.Search.WithContext(context.Background()),
		ESClient.Search.WithIndex("order_items"),
		ESClient.Search.WithBody(&buf),
	)
	if err != nil {
		return nil
	}
	defer res.Body.Close()

	var result esOrderResponse
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil
	}

	seen := make(map[int]bool)
	var orderIds []int
	for _, hit := range result.Hits.Hits {
		if !seen[hit.Source.OrderId] {
			orderIds = append(orderIds, hit.Source.OrderId)
			seen[hit.Source.OrderId] = true
		}
	}
	return orderIds
}
