package mocker

import "encoding/json"

type NewContractTx struct {
	Data                         Data    `json:"data"`
	Signature                    string  `json:"signature"`
	Fee                          *int    `json:"fee"`
	Nid                          int     `json:"nid"`
	Block_number                 int     `json:"block_number"`
	Transaction_index            int     `json:"transaction_index"`
	Type                         string  `json:"type"`
	Receipt_step_price           int     `json:"receipt_stop_price"`
	Block_timestamp              int     `json:"block_timestamp"`
	Step_limit                   int     `json:"step_limit"`
	Receipt_step_used            int     `json:"receipt_step_used"`
	From_address                 string  `json:"from_address"`
	Value                        int     `json:"value"`
	Timestamp                    string  `json:"timestamp"`
	Receipt_status               int     `json:"receipt_status"`
	Item_id                      string  `json:"item_id"`
	Receipt_logs                 *string `json:"receipt_logs"`
	Block_hash                   string  `json:"block_hash"`
	To_address                   string  `json:"to_address"`
	Version                      string  `json:"version"`
	Nonce                        *string `json:"nonce"`
	Receipt_cumulative_step_used int     `json:"receipt_cumulative_step_used"`
	Receipt_score_address        string  `json:"receipt_score_address"`
	Data_type                    string  `json:"data_type"`
	Item_timestamp               string  `json:"first_name"`
	Hash                         string  `json:"hash"`
}

type Data struct {
	Method string `json:"method"`
	Params Params `json:"params"`
}

type Params struct {
	Txhash string `json:"txhash"`
}

func getSampleNewContractTx() *NewContractTx {
	newContract := &NewContractTx{
		Data: Data{
			Method: "acceptScore",
			Params: Params{
				Txhash: "0x7a384f5dcea149d8dab5652c255ae1650c8d8c2ac419d962c1fba977ffda08d7",
			},
		},
		Signature:                    "7flU0KJ/O5X/I0wSr7UhfjVWTly7pYUbTgbchYDzJM8P9OFns8Gpz9xAknD1PCj6Fizx7PqZbN57P6XNm2WVBQE=",
		Nid:                          1,
		Block_number:                 31069871,
		Transaction_index:            1,
		Type:                         "transaction",
		Receipt_step_price:           12500000000,
		Block_timestamp:              1614307356753694,
		Step_limit:                   25000000,
		Receipt_step_used:            165740,
		From_address:                 "hxa2a3ad042ce6f1d2d41469115b597a7a0e20f11c",
		Value:                        0,
		Timestamp:                    "0x5bc34387957d8",
		Receipt_status:               1,
		Item_id:                      "transaction_0xe5ce2f2f61ddf7c8ab1b30ebde0c2bcd918d8f1c7f74264c465284f0e33612bd",
		Block_hash:                   "fb8dcd026f8cd77f8157fe6c620a0797a40273064e368cb292e59ddcb8e700d1",
		To_address:                   "cx0000000000000000000000000000000000000001",
		Version:                      "0x3",
		Receipt_cumulative_step_used: 165740,
		Data_type:                    "call",
		Item_timestamp:               "2021-02-26T02:42:36Z",
		Hash:                         "0xe5ce2f2f61ddf7c8ab1b30ebde0c2bcd918d8f1c7f74264c465284f0e33612bd",
	}
	return newContract
}

func GetSampleNewContractTxJsonString() string {
	newContract := getSampleNewContractTx()
	jsonbytes, _ := json.Marshal(newContract)
	jsonstr := string(jsonbytes)
	return jsonstr
}
