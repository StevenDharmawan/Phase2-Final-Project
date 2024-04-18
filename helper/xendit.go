package helper

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"phase2-final-project/model/domain"
	"phase2-final-project/model/web/response"
)

func CreateInvoice(user response.UserResponse, totalPrice float64) (*domain.Invoice, error) {
	apiKey := os.Getenv("XENDIT_API_KEY")
	apiUrl := "https://api.xendit.co/v2/invoices"

	bodyRequest := map[string]interface{}{
		"external_id":      "1",
		"amount":           totalPrice,
		"description":      "Dummy Invoice RMT003",
		"invoice_duration": 86400,
		"customer": domain.User{
			FullName: user.FullName,
			Email:    user.Email,
		},
		"currency": "IDR",
	}
	reqBody, err := json.Marshal(bodyRequest)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	request, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	request.SetBasicAuth(apiKey, "")
	request.Header.Set("Content-Type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	var resInvoice domain.Invoice
	if err := json.NewDecoder(response.Body).Decode(&resInvoice); err != nil {
		return nil, err
	}
	return &resInvoice, nil
}
