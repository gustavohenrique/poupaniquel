package nubank_test

import (
	"testing"
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"

	"github.com/gustavohenrique/poupaniquel/api/importers/nubank"
)

const (
	ACCESS_TOKEN = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IjIwMTUtMTItMDRUMTc6MzY6MjIuNjY0LXU5ZC1ldWN1Ri1zQUFBRlJiaER3aUEifQ.eyJpc3MiOiJodHRwczpcL1wvd3d3Lm51YmFuay5jb20uYnIiLCJhdWQiOiJvdGhlci5sZWdhY3kiLCJzdWIiOiI1NGM5NDViMi03ZTVmLTQzZDEtOTc1YS1lYjUxMWY2N2ExOTciLCJleHAiOjE0Njk4MjM4NjEsInNjb3BlIjoiYXV0aFwvdXNlciB1c2VyIiwidmVyc2lvbiI6IjIiLCJpYXQiOjE0NjkyMTkwNjEsImp0aSI6IjE5ekJrY1oxeVhRQUFBRldGRWU5cHcifQ.MxWrJdHo3HHLzsSe4VMgzIDYEycnlSgPFIigjVbBgfsScc7PMzT2OAs5SlGw3v_n-e61Sj8-ucgeAN61wt26Qo7RLLwuHrB8y0mOC61hCpz8LFOCRAghTH64msQ54MPj7fwSGHI6PIhHMha-ggCuZwwzGbr9EJ9PPflN8iBX9FFl7EWa6tVo9z2CvFsnyE_fjc0b69z1Fi5yGvX4hAI1a5ObvkwL7GHiT5gounrH5VrMrLW1tCEozn_QInpq7AhUB4o2qXHcb3-uF2XuVQT0KVeIvE-r3c7jdVdqfRJl2i4Fm2P3UZASFTKe1EfaWUJ2EA9h_56AZQ5ASDQmECKaFA"
)

func TestAuthenticate(t *testing.T) {
	username, password := "12345678955", "passwd"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check method is GET before going to check other features
		if r.Method != "POST" {
			t.Errorf("Expected method %q; got %q", "POST", r.Method)
		}
		if r.Header == nil {
			t.Errorf("Expected non-nil request Header")
		}

		defer r.Body.Close()
		/* body, _ := ioutil.ReadAll(r.Body)
		expected := fmt.Sprintf(`{
			"username": "%s",
			"password": "%s",
			"client_id": "%s",
			"client_secret": "%s",
			"grant_type": "%s",
			"nonce": "%s"}`,
			username, password, nubank.ClientId, nubank.ClientSecret, nubank.GrantType, nubank.Nonce)
		if string(body) != expected {
			t.Error("Expected Body", expected, " but got ", string(body))
		} */
		w.Header().Set("Allow", "GET, POST, PUT, OPTIONS")
		// w.WriteHeader(200)
		response := `{
		    "_links": {
		        "account": {
		            "href": "https://prod-s0-credit-card-accounts.nubank.com.br/api/accounts/54c95e6b-f1d1-4af1-b8dd-2580c541bd92"
		        },
		        "bills_summary": {
		            "href": "https://prod-s0-billing.nubank.com.br/api/accounts/54c95e6b-f1d1-4af1-b8dd-2580c541bd92/bills/summary"
		        },
		        "change_password": {
		            "href": "https://prod-s0-auth.nubank.com.br/api/change-password"
		        },
		        "customer": {
		            "href": "https://prod-s0-customers.nubank.com.br/api/customers/54c945b2-7e5f-43d1-975a-eb511f67a197"
		        },
		        "events": {
		            "href": "https://prod-s0-notification.nubank.com.br/api/contacts/54c945b2-7e5f-43d1-975a-eb511f67a197/feed"
		        },
		        "purchases": {
		            "href": "https://prod-s0-feed.nubank.com.br/api/accounts/54c95e6b-f1d1-4af1-b8dd-2580c541bd92/transactions"
		        },
		        "revoke_all": {
		            "href": "https://prod-s0-auth.nubank.com.br/api/revoke/all"
		        },
		        "revoke_token": {
		            "href": "https://prod-s0-auth.nubank.com.br/api/revoke"
		        },
		        "user_change_password": {
		            "href": "https://prod-s0-auth.nubank.com.br/api/user/54c945b2-7e5f-43d1-975a-eb511f67a197/password"
		        },
		        "userinfo": {
		            "href": "https://prod-s0-auth.nubank.com.br/api/userinfo"
		        }
		    },
		    "access_token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IjIwMTUtMTItMDRUMTc6MzY6MjIuNjY0LXU5ZC1ldWN1Ri1zQUFBRlJiaER3aUEifQ.eyJpc3MiOiJodHRwczpcL1wvd3d3Lm51YmFuay5jb20uYnIiLCJhdWQiOiJvdGhlci5sZWdhY3kiLCJzdWIiOiI1NGM5NDViMi03ZTVmLTQzZDEtOTc1YS1lYjUxMWY2N2ExOTciLCJleHAiOjE0Njk4MjM4NjEsInNjb3BlIjoiYXV0aFwvdXNlciB1c2VyIiwidmVyc2lvbiI6IjIiLCJpYXQiOjE0NjkyMTkwNjEsImp0aSI6IjE5ekJrY1oxeVhRQUFBRldGRWU5cHcifQ.MxWrJdHo3HHLzsSe4VMgzIDYEycnlSgPFIigjVbBgfsScc7PMzT2OAs5SlGw3v_n-e61Sj8-ucgeAN61wt26Qo7RLLwuHrB8y0mOC61hCpz8LFOCRAghTH64msQ54MPj7fwSGHI6PIhHMha-ggCuZwwzGbr9EJ9PPflN8iBX9FFl7EWa6tVo9z2CvFsnyE_fjc0b69z1Fi5yGvX4hAI1a5ObvkwL7GHiT5gounrH5VrMrLW1tCEozn_QInpq7AhUB4o2qXHcb3-uF2XuVQT0KVeIvE-r3c7jdVdqfRJl2i4Fm2P3UZASFTKe1EfaWUJ2EA9h_56AZQ5ASDQmECKaFA",
		    "refresh_before": "2016-07-29T20:24:21Z",
		    "refresh_token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IjIwMTUtMTItMDRUMTc6MzY6MjIuNjY0LXU5ZC1ldWN1Ri1zQUFBRlJiaER3aUEifQ.eyJpc3MiOiJodHRwczpcL1wvd3d3Lm51YmFuay5jb20uYnIiLCJhdWQiOiJvdGhlci5sZWdhY3kiLCJzdWIiOiI1NGM5NDViMi03ZTVmLTQzZDEtOTc1YS1lYjUxMWY2N2ExOTciLCJleHAiOjE0NzE4MTEwNjEsInNjb3BlIjoiYXV0aFwvcmVmcmVzaCIsInZlcnNpb24iOiIyIiwiaWF0IjoxNDY5MjE5MDYxLCJqdGkiOiJLSG9hM01RbklQUUFBQUZXRkVlOW9BIn0.AhcKoCEx7uxRx8mcIR-28lABolxi0FhQgfAQC99hBcuo9h2pPGf1SvbkPIjxrJ6hC2yYFu6aNOCQTTkj2wMuYPZA5lnh13uEslXcJz3iuE6i3sR2PIJvpIVVBdmebxhLqEXgyxHHY0b7DAb8nOV5kRmb-34kt1qz3A5MHrWQ6yIEoErcdJdQk5Ks7knWFkCwnUP0w2CEA5sUET-oaIi2GTfhlEUzHyiEga7eTjVWc3rh6LWnk9YgC1Ekd941CJaUlmJIa2LCjSarkKsE5GF2-1N6LdjcULZ6tXYyp29Pd23RWoWgs0RSssBpEouHymomQ-pY3l4tZBbYpEqI9955BA",
		    "token_type": "bearer"
		}`
		w.Write([]byte(response))
	}))

	defer ts.Close()

	service := nubank.NewService(ts.URL)
	err, auth := service.Authenticate(ts.URL, username, password)

	assert.Nil(t, err)
	assert.Equal(t, auth["url"], "https://prod-s0-billing.nubank.com.br/api/accounts/54c95e6b-f1d1-4af1-b8dd-2580c541bd92/bills/summary")
	assert.Equal(t, auth["token"], ACCESS_TOKEN)
}

func TestGetBillsSummary(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected method %q; got %q", "GET", r.Method)
		}
		if r.Header == nil {
			t.Errorf("Expected non-nil request Header")
		}

		defer r.Body.Close()
		
		w.Header().Set("Allow", "GET, POST, PUT, OPTIONS")
		response := `{
		    "_links": {
		        "future": {
		            "href": "https://prod-s0-billing.nubank.com.br/api/accounts/54c95e6b-f1d1-4af1-b8dd-2580c541bd92/bills/future"
		        },
		        "open": {
		            "href": "https://prod-s0-billing.nubank.com.br/api/accounts/54c95e6b-f1d1-4af1-b8dd-2580c541bd92/bills/open"
		        }
		    },
		    "bills": [
		        {
		            "state": "future",
		            "summary": {
		                "close_date": "2017-01-08",
		                "due_date": "2017-01-15",
		                "effective_due_date": "2017-01-16",
		                "interest": 0,
		                "interest_rate": "0.0775",
		                "minimum_payment": 195,
		                "open_date": "2016-12-08",
		                "paid": 0,
		                "past_balance": 0,
		                "precise_minimum_payment": "1.950",
		                "precise_total_balance": "13.0",
		                "total_balance": 1300,
		                "total_cumulative": 1300,
		                "total_prior_bill": "0.0"
		            }
		        },
		        {
		            "state": "future",
		            "summary": {
		                "close_date": "2016-12-08",
		                "due_date": "2016-12-15",
		                "effective_due_date": "2016-12-15",
		                "interest": 0,
		                "interest_rate": "0.0775",
		                "minimum_payment": 195,
		                "open_date": "2016-11-08",
		                "paid": 0,
		                "past_balance": 0,
		                "precise_minimum_payment": "1.950",
		                "precise_total_balance": "13.0",
		                "total_balance": 1300,
		                "total_cumulative": 1300,
		                "total_prior_bill": "0.0"
		            }
		        },
		        {
		            "state": "future",
		            "summary": {
		                "close_date": "2016-11-08",
		                "due_date": "2016-11-15",
		                "effective_due_date": "2016-11-16",
		                "interest": 0,
		                "interest_rate": "0.0775",
		                "minimum_payment": 195,
		                "open_date": "2016-10-08",
		                "paid": 0,
		                "past_balance": 0,
		                "precise_minimum_payment": "1.950",
		                "precise_total_balance": "13.0",
		                "total_balance": 1300,
		                "total_cumulative": 1300,
		                "total_prior_bill": "0.0"
		            }
		        },
		        {
		            "state": "future",
		            "summary": {
		                "close_date": "2016-10-08",
		                "due_date": "2016-10-15",
		                "effective_due_date": "2016-10-17",
		                "interest": 0,
		                "interest_rate": "0.0775",
		                "minimum_payment": 195,
		                "open_date": "2016-09-08",
		                "paid": 0,
		                "past_balance": 0,
		                "precise_minimum_payment": "1.950",
		                "precise_total_balance": "13.0",
		                "total_balance": 1300,
		                "total_cumulative": 1300,
		                "total_prior_bill": "0.0"
		            }
		        },
		        {
		            "state": "future",
		            "summary": {
		                "close_date": "2016-09-08",
		                "due_date": "2016-09-15",
		                "effective_due_date": "2016-09-15",
		                "interest": 0,
		                "interest_rate": "0.0775",
		                "minimum_payment": 195,
		                "open_date": "2016-08-08",
		                "paid": 0,
		                "past_balance": 0,
		                "precise_minimum_payment": "1.950",
		                "precise_total_balance": "13.0",
		                "total_balance": 1300,
		                "total_cumulative": 1300,
		                "total_prior_bill": "0.0"
		            }
		        },
		        {
		            "state": "open",
		            "summary": {
		                "close_date": "2016-08-08",
		                "due_date": "2016-08-15",
		                "effective_due_date": "2016-08-15",
		                "interest": 0,
		                "interest_rate": "0.0775",
		                "minimum_payment": 25791,
		                "open_date": "2016-07-08",
		                "paid": 0,
		                "past_balance": 0,
		                "precise_minimum_payment": "257.912700",
		                "precise_total_balance": "1719.4180",
		                "total_balance": 171942,
		                "total_cumulative": 171942,
		                "total_prior_bill": "1727.7080"
		            }
		        },
		        {
		            "_links": {
		                "self": {
		                    "href": "https://prod-s0-billing.nubank.com.br/api/bills/57802736-6645-4dfe-bdc4-d46cbdd3e699"
		                }
		            },
		            "href": "nuapp://bill/57802736-6645-4dfe-bdc4-d46cbdd3e699",
		            "id": "57802736-6645-4dfe-bdc4-d46cbdd3e699",
		            "state": "overdue",
		            "summary": {
		                "close_date": "2016-07-08",
		                "due_date": "2016-07-15",
		                "effective_due_date": "2016-07-15",
		                "interest": 0,
		                "interest_rate": "0.0775",
		                "minimum_payment": 25916,
		                "open_date": "2016-06-08",
		                "paid": 172771,
		                "past_balance": -44523,
		                "precise_minimum_payment": "259.156200",
		                "precise_total_balance": "1727.7080",
		                "total_balance": 172771,
		                "total_cumulative": 217294,
		                "total_financed": "0",
		                "total_international": "0",
		                "total_national": "2172.94",
		                "total_payments": "0"
		            }
		        },
		        {
		            "_links": {
		                "self": {
		                    "href": "https://prod-s0-billing.nubank.com.br/api/bills/57583184-28fb-4565-80e2-dedbdaa5cd13"
		                }
		            },
		            "href": "nuapp://bill/57583184-28fb-4565-80e2-dedbdaa5cd13",
		            "id": "57583184-28fb-4565-80e2-dedbdaa5cd13",
		            "state": "overdue",
		            "summary": {
		                "close_date": "2016-06-08",
		                "due_date": "2016-06-15",
		                "effective_due_date": "2016-06-15",
		                "interest": 0,
		                "interest_rate": "0.0775",
		                "minimum_payment": 0,
		                "open_date": "2016-05-05",
		                "paid": 0,
		                "past_balance": 0,
		                "precise_minimum_payment": "0",
		                "precise_total_balance": "-445.2320",
		                "total_balance": -44523,
		                "total_cumulative": -44523,
		                "total_financed": "0",
		                "total_international": "0",
		                "total_national": "2521.1200",
		                "total_payments": "-3189.1000"
		            }
		        },
		        {
		            "_links": {
		                "self": {
		                    "href": "https://prod-s0-billing.nubank.com.br/api/bills/572b5e2a-8c33-48f8-884a-6b4e0bcbb713"
		                }
		            },
		            "href": "nuapp://bill/572b5e2a-8c33-48f8-884a-6b4e0bcbb713",
		            "id": "572b5e2a-8c33-48f8-884a-6b4e0bcbb713",
		            "state": "overdue",
		            "summary": {
		                "close_date": "2016-05-05",
		                "due_date": "2016-05-15",
		                "effective_due_date": "2016-05-16",
		                "interest": 0,
		                "interest_rate": "0.0775",
		                "minimum_payment": 3341,
		                "open_date": "2016-04-05",
		                "paid": 22275,
		                "past_balance": 0,
		                "precise_minimum_payment": "33.412200",
		                "precise_total_balance": "222.7480",
		                "total_balance": 22275,
		                "total_cumulative": 22275,
		                "total_financed": "0",
		                "total_international": "0",
		                "total_national": "377.0400",
		                "total_payments": "-3572.1320"
		            }
		        },
		        {
		            "_links": {
		                "self": {
		                    "href": "https://prod-s0-billing.nubank.com.br/api/bills/5703bceb-738d-487b-a0f2-b90f3f6c5d87"
		                }
		            },
		            "href": "nuapp://bill/5703bceb-738d-487b-a0f2-b90f3f6c5d87",
		            "id": "5703bceb-738d-487b-a0f2-b90f3f6c5d87",
		            "state": "overdue",
		            "summary": {
		                "close_date": "2016-04-05",
		                "due_date": "2016-04-15",
		                "effective_due_date": "2016-04-15",
		                "interest": 0,
		                "interest_rate": "0.0775",
		                "minimum_payment": 51268,
		                "open_date": "2016-03-05",
		                "paid": 341784,
		                "past_balance": 0,
		                "precise_minimum_payment": "512.676000",
		                "precise_total_balance": "3417.8400",
		                "total_balance": 341784,
		                "total_cumulative": 341784,
		                "total_financed": "0",
		                "total_international": "2966.3500",
		                "total_national": "451.4900",
		                "total_payments": "-681.0200"
		            }
		        },
		        {
		            "_links": {
		                "self": {
		                    "href": "https://prod-s0-billing.nubank.com.br/api/bills/56daecd6-859f-452c-96ff-df17643a2ecf"
		                }
		            },
		            "href": "nuapp://bill/56daecd6-859f-452c-96ff-df17643a2ecf",
		            "id": "56daecd6-859f-452c-96ff-df17643a2ecf",
		            "state": "overdue",
		            "summary": {
		                "close_date": "2016-03-05",
		                "due_date": "2016-03-15",
		                "effective_due_date": "2016-03-15",
		                "interest": 0,
		                "interest_rate": "0.0775",
		                "minimum_payment": 10215,
		                "open_date": "2016-02-05",
		                "paid": 68102,
		                "past_balance": 0,
		                "precise_minimum_payment": "102.153000",
		                "precise_total_balance": "681.0200",
		                "total_balance": 68102,
		                "total_cumulative": 68102,
		                "total_financed": "0",
		                "total_international": "0",
		                "total_national": "681.0200",
		                "total_payments": "-803.5400"
		            }
		        },
		        {
		            "_links": {
		                "self": {
		                    "href": "https://prod-s0-billing.nubank.com.br/api/bills/56b4a16f-677c-412b-a8f9-d3f1f69e9126"
		                }
		            },
		            "href": "nuapp://bill/56b4a16f-677c-412b-a8f9-d3f1f69e9126",
		            "id": "56b4a16f-677c-412b-a8f9-d3f1f69e9126",
		            "state": "overdue",
		            "summary": {
		                "close_date": "2016-02-05",
		                "due_date": "2016-02-15",
		                "effective_due_date": "2016-02-15",
		                "interest": 0,
		                "interest_rate": "0.0775",
		                "minimum_payment": 12053,
		                "open_date": "2016-01-05",
		                "paid": 80354,
		                "past_balance": 0,
		                "precise_minimum_payment": "120.531000",
		                "precise_total_balance": "803.5400",
		                "total_balance": 80354,
		                "total_cumulative": 80354,
		                "total_financed": "0",
		                "total_international": "0",
		                "total_national": "803.5400",
		                "total_payments": "-404.5500"
		            }
		        },
		        {
		            "_links": {
		                "self": {
		                    "href": "https://prod-s0-billing.nubank.com.br/api/bills/568c15fe-3fdd-4071-9eea-3d313f635707"
		                }
		            },
		            "href": "nuapp://bill/568c15fe-3fdd-4071-9eea-3d313f635707",
		            "id": "568c15fe-3fdd-4071-9eea-3d313f635707",
		            "state": "overdue",
		            "summary": {
		                "close_date": "2016-01-05",
		                "due_date": "2016-01-15",
		                "effective_due_date": "2016-01-15",
		                "interest": 0,
		                "interest_rate": "0.0775",
		                "minimum_payment": 6068,
		                "open_date": "2015-12-05",
		                "paid": 40455,
		                "past_balance": 0,
		                "precise_minimum_payment": "60.682500",
		                "precise_total_balance": "404.5500",
		                "total_balance": 40455,
		                "total_cumulative": 40455,
		                "total_financed": "0",
		                "total_international": "0",
		                "total_national": "404.5500",
		                "total_payments": "-2710.6600"
		            }
		        },
		        {
		            "_links": {
		                "self": {
		                    "href": "https://prod-s0-billing.nubank.com.br/api/bills/566301ed-f74a-46cd-96ff-10efd13a3245"
		                }
		            },
		            "href": "nuapp://bill/566301ed-f74a-46cd-96ff-10efd13a3245",
		            "id": "566301ed-f74a-46cd-96ff-10efd13a3245",
		            "state": "overdue",
		            "summary": {
		                "close_date": "2015-12-05",
		                "due_date": "2015-12-15",
		                "effective_due_date": "2015-12-15",
		                "interest": 0,
		                "interest_rate": "0.0775",
		                "minimum_payment": 40660,
		                "open_date": "2015-11-03",
		                "paid": 271066,
		                "past_balance": 0,
		                "precise_minimum_payment": "406.599000",
		                "precise_total_balance": "2710.6600",
		                "total_balance": 271066,
		                "total_cumulative": 271066,
		                "total_financed": "0",
		                "total_international": "0",
		                "total_national": "2710.6600",
		                "total_payments": "-1871.0700"
		            }
		        },
		        {
		            "_links": {
		                "self": {
		                    "href": "https://prod-s0-billing.nubank.com.br/api/bills/563a49c3-e23a-4bca-ad93-a5b737f8ac0d"
		                }
		            },
		            "href": "nuapp://bill/563a49c3-e23a-4bca-ad93-a5b737f8ac0d",
		            "id": "563a49c3-e23a-4bca-ad93-a5b737f8ac0d",
		            "state": "overdue",
		            "summary": {
		                "close_date": "2015-11-03",
		                "due_date": "2015-11-15",
		                "effective_due_date": "2015-11-16",
		                "interest": 0,
		                "interest_rate": "0.0775",
		                "minimum_payment": 28066,
		                "open_date": "2015-10-02",
		                "paid": 187107,
		                "past_balance": 0,
		                "precise_minimum_payment": "280.660500",
		                "precise_total_balance": "1871.0700",
		                "total_balance": 187107,
		                "total_cumulative": 187107,
		                "total_financed": "0",
		                "total_international": "0",
		                "total_national": "1871.0700",
		                "total_payments": "-2697.8000"
		            }
		        },
		        {
		            "_links": {
		                "self": {
		                    "href": "https://prod-s0-billing.nubank.com.br/api/bills/560ed8d7-faf4-452a-8399-803a40d66ae8"
		                }
		            },
		            "href": "nuapp://bill/560ed8d7-faf4-452a-8399-803a40d66ae8",
		            "id": "560ed8d7-faf4-452a-8399-803a40d66ae8",
		            "state": "overdue",
		            "summary": {
		                "close_date": "2015-10-02",
		                "due_date": "2015-10-15",
		                "effective_due_date": "2015-10-15",
		                "interest": 0,
		                "interest_rate": "0.0775",
		                "minimum_payment": 40467,
		                "open_date": "2015-09-02",
		                "paid": 269780,
		                "past_balance": 0,
		                "precise_minimum_payment": "404.670000",
		                "precise_total_balance": "2697.8000",
		                "total_balance": 269780,
		                "total_cumulative": 269780,
		                "total_financed": "0",
		                "total_international": "0",
		                "total_national": "2697.8000",
		                "total_payments": "-1688.2000"
		            }
		        },
		        {
		            "_links": {
		                "self": {
		                    "href": "https://prod-s0-billing.nubank.com.br/api/bills/55e872f3-528e-4248-b55d-bd967ff21a75"
		                }
		            },
		            "href": "nuapp://bill/55e872f3-528e-4248-b55d-bd967ff21a75",
		            "id": "55e872f3-528e-4248-b55d-bd967ff21a75",
		            "state": "overdue",
		            "summary": {
		                "close_date": "2015-09-02",
		                "due_date": "2015-09-15",
		                "effective_due_date": "2015-09-15",
		                "interest": 0,
		                "interest_rate": "0.0775",
		                "minimum_payment": 25323,
		                "open_date": "2015-08-03",
		                "paid": 168820,
		                "past_balance": 0,
		                "precise_minimum_payment": "253.2300",
		                "precise_total_balance": "1688.2000",
		                "total_balance": 168820,
		                "total_cumulative": 168820
		            }
		        },
		        {
		            "_links": {
		                "self": {
		                    "href": "https://prod-s0-billing.nubank.com.br/api/bills/55c0ef21-f547-4ddb-8bd8-7de9ccc8292b"
		                }
		            },
		            "href": "nuapp://bill/55c0ef21-f547-4ddb-8bd8-7de9ccc8292b",
		            "id": "55c0ef21-f547-4ddb-8bd8-7de9ccc8292b",
		            "state": "overdue",
		            "summary": {
		                "close_date": "2015-08-03",
		                "due_date": "2015-08-15",
		                "effective_due_date": "2015-08-17",
		                "interest": 0,
		                "interest_rate": "0.0775",
		                "minimum_payment": 11323,
		                "open_date": "2015-07-02",
		                "paid": 75489,
		                "past_balance": 0,
		                "precise_minimum_payment": "113.2300",
		                "precise_total_balance": "754.8900",
		                "total_balance": 75489,
		                "total_cumulative": 75489
		            }
		        },
		        {
		            "_links": {
		                "self": {
		                    "href": "https://prod-s0-billing.nubank.com.br/api/bills/559699da-d628-4781-9a25-66efad82fa0e"
		                }
		            },
		            "href": "nuapp://bill/559699da-d628-4781-9a25-66efad82fa0e",
		            "id": "559699da-d628-4781-9a25-66efad82fa0e",
		            "state": "overdue",
		            "summary": {
		                "close_date": "2015-07-02",
		                "due_date": "2015-07-15",
		                "effective_due_date": "2015-07-15",
		                "interest": 0,
		                "interest_rate": "0.0775",
		                "minimum_payment": 23745,
		                "open_date": "2015-06-02",
		                "paid": 158299,
		                "past_balance": 0,
		                "precise_minimum_payment": "237.4500",
		                "precise_total_balance": "1582.9900",
		                "total_balance": 158299,
		                "total_cumulative": 158299
		            }
		        },
		        {
		            "_links": {
		                "self": {
		                    "href": "https://prod-s0-billing.nubank.com.br/api/bills/556f0d84-810a-4676-a747-66d09f5c12f9"
		                }
		            },
		            "href": "nuapp://bill/556f0d84-810a-4676-a747-66d09f5c12f9",
		            "id": "556f0d84-810a-4676-a747-66d09f5c12f9",
		            "state": "overdue",
		            "summary": {
		                "close_date": "2015-06-02",
		                "due_date": "2015-06-15",
		                "effective_due_date": "2015-06-15",
		                "interest": 0,
		                "interest_rate": "0.0775",
		                "minimum_payment": 20697,
		                "open_date": "2015-05-04",
		                "paid": 137983,
		                "past_balance": 0,
		                "precise_minimum_payment": "206.9700",
		                "precise_total_balance": "1379.8300",
		                "total_balance": 137983,
		                "total_cumulative": 137983
		            }
		        },
		        {
		            "_links": {
		                "self": {
		                    "href": "https://prod-s0-billing.nubank.com.br/api/bills/554911e8-a060-4c0e-bd0e-48c791835111"
		                }
		            },
		            "href": "nuapp://bill/554911e8-a060-4c0e-bd0e-48c791835111",
		            "id": "554911e8-a060-4c0e-bd0e-48c791835111",
		            "state": "overdue",
		            "summary": {
		                "close_date": "2015-05-04",
		                "due_date": "2015-05-15",
		                "effective_due_date": "2015-05-15",
		                "interest": 0,
		                "interest_rate": "0.0775",
		                "minimum_payment": 23758,
		                "open_date": "2015-04-01",
		                "paid": 158389,
		                "past_balance": 0,
		                "precise_minimum_payment": "237.5800",
		                "precise_total_balance": "1583.8900",
		                "total_balance": 158389,
		                "total_cumulative": 158389
		            }
		        },
		        {
		            "_links": {
		                "self": {
		                    "href": "https://prod-s0-billing.nubank.com.br/api/bills/551dc206-098b-4fce-a4e6-f10d8dfae0f9"
		                }
		            },
		            "href": "nuapp://bill/551dc206-098b-4fce-a4e6-f10d8dfae0f9",
		            "id": "551dc206-098b-4fce-a4e6-f10d8dfae0f9",
		            "state": "overdue",
		            "summary": {
		                "close_date": "2015-04-01",
		                "due_date": "2015-04-15",
		                "effective_due_date": "2015-04-15",
		                "interest": 0,
		                "interest_rate": "0.0775",
		                "minimum_payment": 22677,
		                "open_date": "2015-03-02",
		                "paid": 151183,
		                "past_balance": 0,
		                "precise_minimum_payment": "226.7700",
		                "precise_total_balance": "1511.8300",
		                "total_balance": 151183,
		                "total_cumulative": 151183
		            }
		        },
		        {
		            "_links": {
		                "self": {
		                    "href": "https://prod-s0-billing.nubank.com.br/api/bills/54f5e7ee-81c8-451b-8d94-9affa7f5505d"
		                }
		            },
		            "href": "nuapp://bill/54f5e7ee-81c8-451b-8d94-9affa7f5505d",
		            "id": "54f5e7ee-81c8-451b-8d94-9affa7f5505d",
		            "state": "overdue",
		            "summary": {
		                "close_date": "2015-03-02",
		                "due_date": "2015-03-15",
		                "effective_due_date": "2015-03-16",
		                "interest": 0,
		                "interest_rate": "0.0775",
		                "minimum_payment": 5451,
		                "open_date": "2015-01-28",
		                "paid": 36338,
		                "past_balance": 0,
		                "precise_minimum_payment": "54.5100",
		                "precise_total_balance": "363.3800",
		                "total_balance": 36338,
		                "total_cumulative": 36338
		            }
		        }
		    ]
		}`
		w.Write([]byte(response))
	}))
	defer ts.Close()

	service := nubank.NewService(ts.URL)
	
	err, bills := service.GetBillsSummary(ts.URL, ACCESS_TOKEN)

	assert.Nil(t, err)
	assert.Equal(t, 17, len(bills))

	bill := bills[16]
	assert.Equal(t, bill["id"].(string), "54f5e7ee-81c8-451b-8d94-9affa7f5505d")
	assert.Equal(t, bill["state"].(string), "overdue")
	assert.Equal(t, bill["paid"].(float64), 363.38)
	assert.Equal(t, bill["closeDate"].(string), "2015-03-02")
	assert.Equal(t, bill["dueDate"].(string), "2015-03-15")
	assert.Equal(t, bill["link"].(string), "https://prod-s0-billing.nubank.com.br/api/bills/54f5e7ee-81c8-451b-8d94-9affa7f5505d")
}

func TestGetBillItems(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected method %q; got %q", "GET", r.Method)
		}
		if r.Header == nil {
			t.Errorf("Expected non-nil request Header")
		}

		defer r.Body.Close()
		
		w.Header().Set("Allow", "GET, POST, PUT, OPTIONS")
		response := `{
		    "bill": {
		        "_links": {
		            "barcode": {
		                "href": "https://prod-s0-billing.nubank.com.br/api/bills/54f5e7ee-81c8-451b-8d94-9affa7f5505d/boleto/barcode"
		            },
		            "boleto_email": {
		                "href": "https://prod-s0-billing.nubank.com.br/api/bills/54f5e7ee-81c8-451b-8d94-9affa7f5505d/boleto/email"
		            },
		            "invoice_email": {
		                "href": "https://prod-s0-billing.nubank.com.br/api/bills/54f5e7ee-81c8-451b-8d94-9affa7f5505d/invoice/email"
		            },
		            "self": {
		                "href": "https://prod-s0-billing.nubank.com.br/api/bills/54f5e7ee-81c8-451b-8d94-9affa7f5505d"
		            }
		        },
		        "auto_debit_failed": false,
		        "barcode": "",
		        "id": "54f5e7ee-81c8-451b-8d94-9affa7f5505d",
		        "line_items": [
		            {
		                "amount": 791,
		                "charges": 1,
		                "href": "nuapp://transaction/54e60e7d-9bfd-4510-a0f1-8eea8ed3fa7e",
		                "id": "54e86ddd-60eb-422f-b630-b6730cc04b14",
		                "index": 0,
		                "post_date": "2015-02-21",
		                "title": "Rest Max Grill"
		            },
		            {
		                "amount": 903,
		                "charges": 1,
		                "href": "nuapp://transaction/54e63d98-4da2-45f1-8362-c292688d9bbb",
		                "id": "54e86dd8-a8c7-4250-ba41-390403776e29",
		                "index": 0,
		                "post_date": "2015-02-21",
		                "title": "Laticinios Ramalho"
		            },
		            {
		                "amount": 1701,
		                "charges": 1,
		                "href": "nuapp://transaction/54e7d661-0c95-44ac-a8e6-981221ebf4c9",
		                "id": "54ec7b87-415a-4eaf-8e6a-a4b58959e58d",
		                "index": 0,
		                "post_date": "2015-02-24",
		                "title": "Drogarias Pacheco"
		            },
		            {
		                "amount": 1263,
		                "charges": 1,
		                "href": "nuapp://transaction/54e75f82-e97c-4277-8771-abcf80eec28f",
		                "id": "54ec7b8d-37a8-49ec-85bb-1e8dbd706fc7",
		                "index": 0,
		                "post_date": "2015-02-24",
		                "title": "Restaurante Atelier"
		            },
		            {
		                "amount": 3337,
		                "charges": 1,
		                "href": "nuapp://transaction/54ea01af-7b97-44ef-a6e2-75c236272abe",
		                "id": "54ec7b6e-21dd-48f3-8fd3-6e467c15eadc",
		                "index": 0,
		                "post_date": "2015-02-24",
		                "title": "Adega Du Dudu"
		            },
		            {
		                "amount": 13280,
		                "charges": 1,
		                "href": "nuapp://transaction/54e8ea3b-cea2-4865-b315-ade8443eaa2e",
		                "id": "54ec7b77-16f8-4fc4-9e31-c9f3e745c79c",
		                "index": 0,
		                "post_date": "2015-02-24",
		                "title": "Karol Pet Shop"
		            },
		            {
		                "amount": 8546,
		                "charges": 1,
		                "href": "nuapp://transaction/54e913c3-2086-4b58-8953-73bfcb243dc7",
		                "id": "54ec7b81-bd4c-4796-9dda-2c47a8676f0a",
		                "index": 0,
		                "post_date": "2015-02-24",
		                "title": "Cencosud Brasil"
		            },
		            {
		                "amount": 249,
		                "charges": 1,
		                "href": "nuapp://transaction/54e9f401-d179-4319-870e-f8da62c404b8",
		                "id": "54ec7b66-c804-44fd-9490-9800cfed2331",
		                "index": 0,
		                "post_date": "2015-02-24",
		                "title": "Cencosud Brasil"
		            },
		            {
		                "amount": 1378,
		                "charges": 1,
		                "href": "nuapp://transaction/54ec9599-658e-414a-8b83-800f6b10298e",
		                "id": "54ef00da-3f2b-447d-b62e-e1304c40c7aa",
		                "index": 0,
		                "post_date": "2015-02-26",
		                "title": "Yaki"
		            },
		            {
		                "amount": 2990,
		                "charges": 1,
		                "href": "nuapp://transaction/54f0a426-1134-4c94-b1db-627a21747cb0",
		                "id": "54f5e7e5-a716-4930-b3b1-e8ae5715338d",
		                "index": 0,
		                "post_date": "2015-03-01",
		                "title": "Creps Metropolitano Ri"
		            },
		            {
		                "amount": 1900,
		                "charges": 1,
		                "href": "nuapp://transaction/54f2305f-3752-4957-a82b-c0d346832942",
		                "id": "54f5e7e5-7d41-4c4e-992c-6a870d90d4f7",
		                "index": 0,
		                "post_date": "2015-03-01",
		                "title": "Posto Rio Lima Ri"
		            }
		        ],
		        "linha_digitavel": "",
		        "payment_method": "boleto",
		        "state": "overdue",
		        "summary": {
		            "close_date": "2015-03-02",
		            "due_date": "2015-03-15",
		            "effective_due_date": "2015-03-16",
		            "interest": 0,
		            "interest_rate": "0.0775",
		            "minimum_payment": 5451,
		            "open_date": "2015-01-28",
		            "paid": 36338,
		            "past_balance": 0,
		            "precise_minimum_payment": "54.5100",
		            "precise_total_balance": "363.3800",
		            "total_balance": 36338,
		            "total_cumulative": 36338
		        }
		    }
		}`
		w.Write([]byte(response))
	}))
	defer ts.Close()

	service := nubank.NewService(ts.URL)
	
	err, items := service.GetBillItems(ts.URL, ACCESS_TOKEN)

	assert.Nil(t, err)
	assert.Equal(t, 11, len(items))
}

func TestGetTransactionDetails(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected method %q; got %q", "GET", r.Method)
		}
		if r.Header == nil {
			t.Errorf("Expected non-nil request Header")
		}

		defer r.Body.Close()
		
		w.Header().Set("Allow", "GET, POST, PUT, OPTIONS")
		response := `{
		    "transaction": {
		        "_links": {
		            "categories": {
		                "href": "https://prod-s0-feed.nubank.com.br/api/transactions/578e8371-ba8d-435d-81d1-16b40fafe36d/category"
		            },
		            "category": {
		                "href": "https://prod-s0-feed.nubank.com.br/api/transactions/578e8371-ba8d-435d-81d1-16b40fafe36d/category"
		            },
		            "chargeback": {
		                "href": "https://prod-s0-feed.nubank.com.br/api/transactions/578e8371-ba8d-435d-81d1-16b40fafe36d/chargebacks"
		            },
		            "chargeback_reasons": {
		                "href": "https://prod-s0-feed.nubank.com.br/api/transactions/578e8371-ba8d-435d-81d1-16b40fafe36d/chargebacks/reasons"
		            },
		            "chargeback_reasons_v4": {
		                "href": "https://prod-s0-feed.nubank.com.br/api/transactions/578e8371-ba8d-435d-81d1-16b40fafe36d/chargebacks/reasons/v4"
		            },
		            "create_tag": {
		                "href": "https://prod-s0-feed.nubank.com.br/api/transactions/578e8371-ba8d-435d-81d1-16b40fafe36d/tags"
		            },
		            "merchant": {
		                "href": "https://prod-s0-feed.nubank.com.br/api/transactions/578e8371-ba8d-435d-81d1-16b40fafe36d/merchant"
		            },
		            "notify_geo": {
		                "href": "https://prod-s0-geo.nubank.com.br/api/waypoints/transaction/578e8371-ba8d-435d-81d1-16b40fafe36d"
		            },
		            "self": {
		                "href": "https://prod-s0-feed.nubank.com.br/api/transactions/578e8371-ba8d-435d-81d1-16b40fafe36d"
		            }
		        },
		        "account": "54c95e6b-f1d1-4af1-b8dd-2580c541bd92",
		        "amount": 5058,
		        "capture_mode": {
		            "entry_mode": "chip",
		            "pin_mode": "accepted"
		        },
		        "card": "54c95e6b-cc80-4ee2-8494-531b7cfe87e9",
		        "chargebacks": [],
		        "charges": 1,
		        "charges_list": [
		            {
		                "account_id": "54c95e6b-f1d1-4af1-b8dd-2580c541bd92",
		                "amount": 5058,
		                "extras": [],
		                "id": "5790b059-12b8-4b7b-a639-5494e7c03e61",
		                "index": 0,
		                "post_date": "2016-07-20",
		                "precise_amount": "50.58",
		                "status": "open",
		                "transaction_id": "578e8371-ba8d-435d-81d1-16b40fafe36d"
		            }
		        ],
		        "country": "BRA",
		        "event_type": "transaction_card_present",
		        "id": "578e8371-ba8d-435d-81d1-16b40fafe36d",
		        "mcc": "sa\u00fade",
		        "mcg": "21",
		        "merchant_name": "Drogarias Pacheco",
		        "original_merchant_name": "DROGARIAS PACHECO",
		        "postcode": "22631020",
		        "precise_amount": "50.58",
		        "pulled_at": "2016-07-19T19:45:53.07Z",
		        "pushed_at": "2016-07-19T19:45:52.737Z",
		        "source": "upfront_national",
		        "status": "settled",
		        "tags": [
		            {
		                "_links": {
		                    "remove": {
		                        "href": "https://prod-s0-feed.nubank.com.br/api/transactions/578e8371-ba8d-435d-81d1-16b40fafe36d/tags/5793c137-684a-4a25-807c-06b74d675325"
		                    }
		                },
		                "description": "farmacia"
		            },
		            {
		                "_links": {
		                    "remove": {
		                        "href": "https://prod-s0-feed.nubank.com.br/api/transactions/578e8371-ba8d-435d-81d1-16b40fafe36d/tags/5793c135-3a1c-48e2-979f-07df35c7712f"
		                    }
		                },
		                "description": "saude"
		            }
		        ],
		        "time": "2016-07-19T19:45:51Z",
		        "time_wallclock": "2016-07-19T19:45:51",
		        "type": "card_present"
		    }
		}`
		w.Write([]byte(response))
	}))
	defer ts.Close()

	service := nubank.NewService(ts.URL)
	
	err, details := service.GetTransactionDetails(ts.URL, ACCESS_TOKEN)

	assert.Nil(t, err)
	assert.Equal(t, "Drogarias Pacheco", details["title"].(string))
	assert.Equal(t, float64(50.58), details["amount"].(float64))
	assert.Equal(t, "|farmacia|", details["tags"].([]string)[0])
	assert.Equal(t, "|saude|", details["tags"].([]string)[1])
}