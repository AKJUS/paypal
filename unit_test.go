package paypal

import (
	"encoding/json"
	"testing"
)

func TestNewClient(t *testing.T) {
	c, err := NewClient("", "", "")
	if err == nil {
		t.Errorf("Expected error for NewClient('','','')")
	}
	if c != nil {
		t.Errorf("Expected nil Client for NewClient('','',''), got %v", c)
	}

	c, err = NewClient("1", "2", "3")
	if err != nil {
		t.Errorf("Not expected error for NewClient(1, 2, 3), got %v", err)
	}
	if c == nil {
		t.Errorf("Expected non-nil Client for NewClient(1, 2, 3)")
	}
}

func TestTypeUserInfo(t *testing.T) {
	response := `{
    "user_id": "https://www.paypal.com/webapps/auth/server/64ghr894040044",
    "name": "Peter Pepper",
    "given_name": "Peter",
    "family_name": "Pepper",
    "email": "ppuser@example.com"
    }`

	u := &UserInfo{}
	err := json.Unmarshal([]byte(response), u)
	if err != nil {
		t.Errorf("UserInfo Unmarshal failed")
	}

	if u.ID != "https://www.paypal.com/webapps/auth/server/64ghr894040044" ||
		u.Name != "Peter Pepper" ||
		u.GivenName != "Peter" ||
		u.FamilyName != "Pepper" ||
		u.Email != "ppuser@example.com" {
		t.Errorf("UserInfo decoded result is incorrect, Given: %v", u)
	}
}

func TestTypeItem(t *testing.T) {
	response := `{
    "name":"Item",
    "quantity":"1"
}`

	i := &Item{}
	err := json.Unmarshal([]byte(response), i)
	if err != nil {
		t.Errorf("Item Unmarshal failed")
	}

	if i.Name != "Item" ||
		i.Quantity != "1" {
		t.Errorf("Item decoded result is incorrect, Given: %v", i)
	}
}

func TestTypeErrorResponseOne(t *testing.T) {
	response := `{
		"name":"USER_BUSINESS_ERROR",
		"message":"User business error.",
		"debug_id":"f05063556a338",
		"information_link":"https://developer.paypal.com/docs/api/payments.payouts-batch/#errors",
		"details":[
			{
				"field":"SENDER_BATCH_ID",
				"issue":"Batch with given sender_batch_id already exists",
				"link":[
					{
						"href":"https://api.sandbox.paypal.com/v1/payments/payouts/CR9VS2K4X4846",
						"rel":"self",
						"method":"GET"
					}
				]
			}
		]
	}`

	i := &ErrorResponse{}
	err := json.Unmarshal([]byte(response), i)
	if err != nil {
		t.Errorf("ErrorResponse Unmarshal failed")
	}

	if i.Name != "USER_BUSINESS_ERROR" ||
		i.Message != "User business error." ||
		i.DebugID != "f05063556a338" ||
		i.InformationLink != "https://developer.paypal.com/docs/api/payments.payouts-batch/#errors" ||
		len(i.Details) != 1 ||
		i.Details[0].Field != "SENDER_BATCH_ID" ||
		i.Details[0].Issue != "Batch with given sender_batch_id already exists" ||
		len(i.Details[0].Links) != 1 ||
		i.Details[0].Links[0].Href != "https://api.sandbox.paypal.com/v1/payments/payouts/CR9VS2K4X4846" {
		t.Errorf("ErrorResponse decoded result is incorrect, Given: %v", i)
	}
}

func TestTypeErrorResponseTwo(t *testing.T) {
	response := `{
		"name":"VALIDATION_ERROR",
		"message":"Invalid request - see details.",
		"debug_id":"662121ee369c0",
		"information_link":"https://developer.paypal.com/docs/api/payments.payouts-batch/#errors",
		"details":[
			{
				"field":"items[0].recipient_type",
				"issue":"Value is invalid (must be EMAILor PAYPAL_ID or PHONE)"
			}
		]
	}`

	i := &ErrorResponse{}
	err := json.Unmarshal([]byte(response), i)
	if err != nil {
		t.Errorf("ErrorResponse Unmarshal failed")
	}

	if i.Name != "VALIDATION_ERROR" ||
		i.Message != "Invalid request - see details." ||
		len(i.Details) != 1 ||
		i.Details[0].Field != "items[0].recipient_type" {
		t.Errorf("ErrorResponse decoded result is incorrect, Given: %v", i)
	}
}

func TestTypeErrorResponseThree(t *testing.T) {

	response := `{
            "name": "DUPLICATE_TRACKER",
            "message": "Business error",
            "debug_id": "[REDACTED]",
            "details": [
                {
                    "field": "#/trackers/0/tracking_number",
                    "value": "[TRACKING_NUMBER]",
                    "location": "body",
                    "issue": "TRACKING_NUMBER_ALREADY_EXIST",
                    "description": "Tracking information already exists with same Tracking number",
                    "links": [

                    ]
                }
            ]
        }`

	i := &ErrorResponse{}
	err := json.Unmarshal([]byte(response), i)
	if err != nil {
		t.Errorf("ErrorResponse Unmarshal failed")
	}

	if i.Name != "DUPLICATE_TRACKER" ||
		i.Message != "Business error" ||
		len(i.Details) != 1 ||
		i.Details[0].Issue != "TRACKING_NUMBER_ALREADY_EXIST" ||
		i.Details[0].Value != "[TRACKING_NUMBER]" ||
		i.Details[0].Description != "Tracking information already exists with same Tracking number" {
		t.Errorf("ErrorResponse decoded result is incorrect, Given: %v", i)
	}
}

func TestTypePayoutResponse(t *testing.T) {
	response := `{
		"batch_header":{
			"payout_batch_id":"G4E6WJE6Y4853",
			"batch_status":"SUCCESS",
			"time_created":"2017-11-01T23:08:25Z",
			"time_completed":"2017-11-01T23:08:46Z",
			"sender_batch_header":{
				"sender_batch_id":"2017110109",
				"email_subject":"Payment"
			},
			"amount":{
				"currency":"USD",
				"value":"6.37"
			},
			"fees":{
				"currency":"USD",
				"value":"0.25"
			}
		},
		"items":[
			{
				"payout_item_id":"9T35G83YA546X",
				"transaction_id":"4T328230B1D337285",
				"transaction_status":"UNCLAIMED",
				"payout_item_fee":{
					"currency":"USD",
					"value":"0.25"
				},
				"payout_batch_id":"G4E6WJE6Y4853",
				"payout_item":{
					"recipient_type":"EMAIL",
					"amount":{
						"currency":"USD",
						"value":"6.37"
					},
					"note":"Optional note",
					"receiver":"ppuser@example.com",
					"sender_item_id":"1"
				},
				"time_processed":"2017-11-01T23:08:43Z",
				"errors":{
					"name":"RECEIVER_UNREGISTERED",
					"message":"Receiver is unregistered",
					"information_link":"https://developer.paypal.com/docs/api/payments.payouts-batch/#errors",
					"details":[]
				},
				"links":[
					{
						"href":"https://api.sandbox.paypal.com/v1/payments/payouts-item/9T35G83YA546X",
						"rel":"item",
						"method":"GET",
						"encType":"application/json"
					}
				]
			}
		],
		"links":[
			{
				"href":"https://api.sandbox.paypal.com/v1/payments/payouts/G4E6WJE6Y4853?page_size=1000&page=1",
				"rel":"self",
				"method":"GET",
				"encType":"application/json"
			}
		]
	}`

	pr := &PayoutResponse{}
	err := json.Unmarshal([]byte(response), pr)
	if err != nil {
		t.Errorf("PayoutResponse Unmarshal failed")
	}

	if pr.BatchHeader.BatchStatus != "SUCCESS" ||
		pr.BatchHeader.PayoutBatchID != "G4E6WJE6Y4853" ||
		len(pr.Items) != 1 ||
		pr.Items[0].PayoutItemID != "9T35G83YA546X" ||
		pr.Items[0].TransactionID != "4T328230B1D337285" ||
		pr.Items[0].TransactionStatus != "UNCLAIMED" ||
		pr.Items[0].Error.Name != "RECEIVER_UNREGISTERED" {
		t.Errorf("PayoutResponse decoded result is incorrect, Given: %v", pr)
	}
}

func TestOrderUnmarshal(t *testing.T) {
	response := `{
		"id": "5O190127TN364715T",
		"status": "CREATED",
		"links": [
		  {
			"href": "https://api.paypal.com/v2/checkout/orders/5O190127TN364715T",
			"rel": "self",
			"method": "GET"
		  },
		  {
			"href": "https://api.sandbox.paypal.com/checkoutnow?token=5O190127TN364715T",
			"rel": "approve",
			"method": "GET"
		  },
		  {
			"href": "https://api.paypal.com/v2/checkout/orders/5O190127TN364715T/capture",
			"rel": "capture",
			"method": "POST"
		  }
		]
	}`

	order := &Order{}
	err := json.Unmarshal([]byte(response), order)
	if err != nil {
		t.Errorf("Order Unmarshal failed")
	}

	if order.ID != "5O190127TN364715T" ||
		order.Status != "CREATED" ||
		order.Links[0].Href != "https://api.paypal.com/v2/checkout/orders/5O190127TN364715T" {
		t.Errorf("Order decoded result is incorrect, Given: %+v", order)
	}
}

func TestOrderCompletedUnmarshal(t *testing.T) {
	response := `{
		"id": "1K412082HD5737736",
		"status": "COMPLETED",
		"purchase_units": [
			{
				"reference_id": "default",
				"amount": {
					"currency_code": "EUR",
					"value": "99.99"
				},
				"payee": {
					"email_address": "payee@business.example.com",
					"merchant_id": "7DVPP5Q2RZJQY"
				},
				"custom_id": "123456",
				"soft_descriptor": "PAYPAL *TEST STORE",
				"shipping": {
					"name": {
						"full_name": "John Doe"
					},
					"address": {
						"address_line_1": "Address, Country",
						"admin_area_2": "Area2",
						"admin_area_1": "Area1",
						"postal_code": "123456",
						"country_code": "US"
					}
				},
				"payments": {
					"captures": [
						{
							"id": "6V864560EH247264J",
							"status": "COMPLETED",
							"amount": {
								"currency_code": "EUR",
								"value": "99.99"
							},
							"final_capture": true,
							"custom_id": "123456",
							"create_time": "2021-07-27T09:39:17Z",
							"update_time": "2021-07-27T09:39:17Z"
						}
					]
				}
			}
		],
		"payer": {
			"name": {
				"given_name": "John",
				"surname": "Doe"
			},
			"email_address": "payer@personal.example.com",
			"payer_id": "7D36CJQ2TUEUU",
			"address": {
				"address_line_1": "City, Country",
				"admin_area_2": "Area2",
				"admin_area_1": "Area1",
				"postal_code": "123456",
				"country_code": "US"
			}
		},
		"create_time": "2021-07-27T09:38:37Z",
		"update_time": "2021-07-27T09:39:17Z",
		"links": [
			{
				"href": "https://api.sandbox.paypal.com/v2/checkout/orders/1K412082HD5737736",
				"rel": "self",
				"method": "GET"
			}
		]
	}`

	order := &Order{}
	err := json.Unmarshal([]byte(response), order)
	if err != nil {
		t.Errorf("Order Unmarshal failed")
	}

	if order.ID != "1K412082HD5737736" ||
		order.Status != "COMPLETED" ||
		order.PurchaseUnits[0].Payee.EmailAddress != "payee@business.example.com" ||
		order.PurchaseUnits[0].CustomID != "123456" ||
		order.PurchaseUnits[0].Shipping.Name.FullName != "John Doe" ||
		order.PurchaseUnits[0].Shipping.Address.AdminArea1 != "Area1" ||
		order.Payer.Name.GivenName != "John" ||
		order.Payer.Address.AddressLine1 != "City, Country" ||
		order.Links[0].Href != "https://api.sandbox.paypal.com/v2/checkout/orders/1K412082HD5737736" {
		t.Errorf("Order decoded result is incorrect, Given: %+v", order)
	}
}

func TestTypePayoutItemResponse(t *testing.T) {
	response := `{
		"payout_item_id":"9T35G83YA546X",
		"transaction_id":"4T328230B1D337285",
		"transaction_status":"UNCLAIMED",
		"payout_item_fee":{
			"currency":"USD",
			"value":"0.25"
		},
		"payout_batch_id":"G4E6WJE6Y4853",
		"payout_item":{
			"recipient_type":"EMAIL",
			"amount":{
				"currency":"USD",
				"value":"6.37"
			},
			"note":"Optional note",
			"receiver":"ppuser@example.com",
			"sender_item_id":"1"
		},
		"time_processed":"2017-11-01T23:08:43Z",
		"errors":{
			"name":"RECEIVER_UNREGISTERED",
			"message":"Receiver is unregistered",
			"information_link":"https://developer.paypal.com/docs/api/payments.payouts-batch/#errors",
			"details":[]
		},
		"links":[
			{
				"href":"https://api.sandbox.paypal.com/v1/payments/payouts-item/3YA546X9T35G8",
				"rel":"self",
				"method":"GET",
				"encType":"application/json"
			},
			{
				"href":"https://api.sandbox.paypal.com/v1/payments/payouts/6Y4853G4E6WJE",
				"rel":"batch",
				"method":"GET",
				"encType":"application/json"
			}
		]
	}`

	pir := &PayoutItemResponse{}
	err := json.Unmarshal([]byte(response), pir)
	if err != nil {
		t.Errorf("PayoutItemResponse Unmarshal failed")
	}

	if pir.PayoutItemID != "9T35G83YA546X" ||
		pir.PayoutBatchID != "G4E6WJE6Y4853" ||
		pir.TransactionID != "4T328230B1D337285" ||
		pir.TransactionStatus != "UNCLAIMED" ||
		pir.Error.Name != "RECEIVER_UNREGISTERED" {
		t.Errorf("PayoutItemResponse decoded result is incorrect, Given: %+v", pir)
	}
}

func TestTypePaymentPatch(t *testing.T) {
	// test unmarshaling
	response := `{
		"op": "replace",
		"path": "/transactions/0/amount",
		"value": "5"
	}`
	pp := &PaymentPatch{}
	err := json.Unmarshal([]byte(response), pp)
	if err != nil {
		t.Errorf("TestTypePaymentPatch Unmarshal failed")
	}
	if pp.Operation != "replace" ||
		pp.Path != "/transactions/0/amount" ||
		pp.Value != "5" {
		t.Errorf("PaymentPatch decoded result is incorrect, Given: %+v", pp)
	}
}

func TestTypePaymentPatchMarshal(t *testing.T) {
	// test marshalling
	p2 := &PaymentPatch{
		Operation: "add",
		Path:      "/transactions/0/amount",
		Value: map[string]interface{}{
			"total":    "18.37",
			"currency": "EUR",
			"details": map[string]interface{}{
				"subtotal": "13.37",
				"shipping": "5.00",
			},
		},
	}
	p2expectedresponse := `{"op":"add","path":"/transactions/0/amount","value":{"currency":"EUR","details":{"shipping":"5.00","subtotal":"13.37"},"total":"18.37"}}`
	response2, _ := json.Marshal(p2)
	if string(response2) != string(p2expectedresponse) {
		t.Errorf("PaymentPatch response2 is incorrect,\n Given:    %+v\n Expected: %+v", string(response2), string(p2expectedresponse))
	}
}
