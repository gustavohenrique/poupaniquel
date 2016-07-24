package nubank

import (
	"log"
	"fmt"
	"errors"
	"strings"
	"time"
	"encoding/json"

	"github.com/parnurzeal/gorequest"
)

type ApiImporter interface {
	Authenticate(string, string, string)  (error, map[string]string)
	GetBillsSummary(string, string)       (error, []map[string]interface{})
	GetBillItems(string, string)          (error, []map[string]interface{})
	GetTransactionDetails(string, string) (error, map[string]interface{})
}

type Service struct {
	Origin string
}

type Auth struct {
	AccessToken string                 `json:"access_token"`
	Links map[string]map[string]string `json:"_links"`
}

type Summary struct {
	Bills []Bill `json:"bills"`
}

type Bill struct {
	Id string                          `json:"id"`
	State string                       `json:"state"`
	Links map[string]map[string]string `json:"_links"`
	Summary map[string]interface{}     `json:"summary"`
	Items []map[string]interface{}     `json:"line_items"`
}

type Detail struct {
	Amount float64                `json:"amount"`
	MerchantName string           `json:"merchant_name"`
	Date time.Time                `json:"pulled_at"`
	Tags []map[string]interface{} `json:"tags"`
}

const (
	AuthUrl = "https://prod-auth.nubank.com.br/api/token"
	TransactionDetailsUrlBase = "https://prod-s0-feed.nubank.com.br/api/transactions"
	Origin = "https://conta.nubank.com.br"
	ClientId = "other.legacy"
	ClientSecret = "1iHY2WHAXj25GFSHTx9lyaTYnb4uB-v6"
	GrantType = "password"
	Nonce = "NOT-RANDOM-YET"
)

func NewService(origin string) ApiImporter {
	return &Service{
		Origin: origin,
	}
}

func (this *Service) Authenticate(url string, username string, password string) (error, map[string]string) {
	requestData := fmt.Sprintf(`{
		"username": "%s",
		"password": "%s",
		"client_id": "%s",
		"client_secret": "%s",
		"grant_type": "%s",
		"nonce": "%s"}`,
		username, password, ClientId, ClientSecret, GrantType, Nonce)

	request := gorequest.New()
	response, body, _ := request.Post(url).
		Set("Origin", this.Origin).
		Set("Content-Type", "application/json").
		Send(requestData).
		End()

	var err error
	if response.StatusCode != 200 {
		err = errors.New("Invalid credentials.")
	}

	if err != nil {
		log.Println("Error connecting to Nubank.", err)
		return err, nil
	}

	b := []byte(body)
	var auth Auth
	json.Unmarshal(b, &auth)
	links := auth.Links
	summaryUrl := links["bills_summary"]["href"]
	result := map[string]string{
		"token": auth.AccessToken,
		"url": summaryUrl,
	}
	return err, result
}

func (this *Service) GetBillsSummary(url string, token string) (error, []map[string]interface{}) {
	request := gorequest.New()
	response, body, _ := request.Get(url).
		Set("Origin", this.Origin).
		Set("Content-Type", "application/json").
		Set("Authorization", "Bearer " + token).
		End()

	var summary Summary
	if err := json.Unmarshal([]byte(body), &summary); err != nil || response.StatusCode != 200 {
		return errors.New("Error fetching bills summary."), nil
	}
	
	var result []map[string]interface{}
	for _, bill := range summary.Bills {
		if bill.Id != "" {
			b := map[string]interface{}{
				"id": bill.Id,
				"state": bill.State,
				"paid": bill.Summary["paid"].(float64) / 100,
				"closeDate": bill.Summary["close_date"],
				"dueDate": bill.Summary["due_date"],
				"link": bill.Links["self"]["href"],
			}
			result = append(result, b)
		}
	}
	return nil, result
}

func (this *Service) GetBillItems(url string, token string) (error, []map[string]interface{}) {
	request := gorequest.New()
	response, body, _ := request.Get(url).
		Set("Origin", this.Origin).
		Set("Content-Type", "application/json").
		Set("Authorization", "Bearer " + token).
		End()

	var items map[string]Bill
	err := json.Unmarshal([]byte(body), &items);
	if response.StatusCode != 200 || err != nil {
		log.Println("Error fetching bill's items.", err, response.StatusCode)
		return err, nil
	}
	
	var result []map[string]interface{}
	for _, item := range items["bill"].Items {
		amount := item["amount"].(float64) / 100
		if amount > 0 {
			href := item["href"].(string)
			b := map[string]interface{}{
				"id": item["id"],
				"date": item["post_date"],
				"amount": amount,
				"title": item["title"],
				"transactionId": strings.Replace(href, "nuapp://transaction/", "", -1),
			}
			result = append(result, b)
		}
	}
	return nil, result
}

func (this *Service) GetTransactionDetails(url string, token string) (error, map[string]interface{}) {
	request := gorequest.New()
	response, body, _ := request.Get(url).
		Set("Origin", this.Origin).
		Set("Content-Type", "application/json").
		Set("Authorization", "Bearer " + token).
		End()

	if response.StatusCode == 429 {
		log.Println("Too many requests fetching details.")
		return errors.New("Too many requests fetching details."), nil
	}

	var details map[string]Detail
	err := json.Unmarshal([]byte(body), &details);
	if response.StatusCode > 400 || err != nil {
		log.Printf("Error fetching transaction details. Status=%d. %s", response.StatusCode, err)
		return err, nil
	}
	
	detail := details["transaction"]
	amount := detail.Amount / 100
	if amount > 0 {
		var tags []string
		for _, tag := range detail.Tags {
			t := fmt.Sprintf("|%s|", tag["description"])
			tags = append(tags, t)
		}
		return nil, map[string]interface{}{
			"date": detail.Date,
			"title": detail.MerchantName,
			"amount": amount,
			"tags": tags,
		}
	}
	return nil, nil
}