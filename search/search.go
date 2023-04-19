package search

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/1005281342/user-manager/models"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"strconv"
)

var client *elasticsearch.Client

func Connect() error {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://192.168.8.42:9200",
		},
	}

	var err error
	client, err = elasticsearch.NewClient(cfg)
	if err != nil {
		return err
	}

	_, err = client.Info()
	if err != nil {
		return err
	}

	return nil
}

func GetClient() *elasticsearch.Client {
	return client
}

func SearchUsersWithKeywords(keywords []string, perPage int, page int) ([]models.User, error) {
	var users []models.User

	offset := (page - 1) * perPage

	// 创建一个新的搜索查询
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"should": make([]map[string]interface{}, 0),
			},
		},
		"size": perPage,
		"from": offset,
		"sort": []map[string]interface{}{{"id": "asc"}},
		"_source": []string{
			"id",
			"first_name",
			"last_name",
			"email",
		},
	}
	for _, keyword := range keywords {
		// 向搜索查询中添加关键词查询条件
		query["query"].(map[string]interface{})["bool"].(map[string]interface{})["should"] = append(
			query["query"].(map[string]interface{})["bool"].(map[string]interface{})["should"].([]map[string]interface{}),
			map[string]interface{}{
				"match": map[string]interface{}{
					"first_name": map[string]interface{}{
						"query":     keyword,
						"fuzziness": "AUTO",
					},
				},
			},
			map[string]interface{}{
				"match": map[string]interface{}{
					"last_name": map[string]interface{}{
						"query":     keyword,
						"fuzziness": "AUTO",
					},
				},
			},
			map[string]interface{}{
				"match": map[string]interface{}{
					"email": map[string]interface{}{
						"query":     keyword,
						"fuzziness": "AUTO",
					},
				},
			},
		)
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Printf("Encode Error:%+v", err)
		return nil, err
	}

	// 发送搜索请求并解析响应
	res, err := esapi.SearchRequest{
		Index: []string{"users"},
		Body:  &buf,
	}.Do(context.Background(), GetClient())
	if err != nil {
		log.Printf("Do Error:%+v", err)
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("res Error:%s", res.String())
		return nil, fmt.Errorf("failed to search users: %s", res.String())
	}

	var searchResult map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&searchResult); err != nil {
		log.Printf("Decode Error:%+v", err)
		return nil, err
	}

	hits := searchResult["hits"].(map[string]interface{})["hits"].([]interface{})
	for _, hit := range hits {
		source := hit.(map[string]interface{})["_source"].(map[string]interface{})
		user := models.User{
			ID:        int(source["id"].(float64)),
			FirstName: source["first_name"].(string),
			LastName:  source["last_name"].(string),
			Email:     source["email"].(string),
		}
		users = append(users, user)
	}

	return users, nil
}

func SaveUser(user models.User) error {
	jsonUser, err := json.Marshal(user)
	if err != nil {
		log.Printf("Marshal Error:%+v", err)
		return err
	}

	_, err = esapi.IndexRequest{
		Index:      "users",
		DocumentID: strconv.Itoa(user.ID),
		Body:       bytes.NewReader(jsonUser),
	}.Do(context.Background(), GetClient())
	if err != nil {
		log.Printf("IndexRequest Error:%+v", err)
		return err
	}

	return nil
}
