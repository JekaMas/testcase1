package domain

import (
	"encoding/json"
	"testing"
	"github.com/mailru/easyjson"
)

var jsonCampaignCollection = []byte(`
	[
		{
			"compaign_name": "campaign870",
			"price":3.25,
			"target_list" : [
       			{"target":"attr_A", "attr_list":["A0","A1","A2","A3","A4","A5","A99"]},
       			{"target":"attr_B", "attr_list":["B0","B1","B2","B3","B4","B5","B6","B36"]},
       			{"target":"attr_D", "attr_list":["D0","D1","D2","D3","D4","D5","D99"]},
       			{"target":"attr_E", "attr_list":["E0","E1","E2","E3","E4","E5","E99"]},
       			{"target":"attr_F", "attr_list":["F0","F1","F2","F3","F4","F5","F6","F36"]}
			]
		},
		{
			"compaign_name": "campaign881",
			"price":5.25,
			"target_list" : [
       			{"target":"attr_C", "attr_list":["C0","C1","C2","C3","C4","C5","C99"]},
       			{"target":"attr_B", "attr_list":["B0","B1","B2","B3","B4","B5","B6","B36"]}
			]
		},
		{
			"compaign_name": "campaign870",
			"price":3.25,
			"target_list" : [
       			{"target":"attr_A", "attr_list":["A0","A1","A2","A3","A4","A5","A99"]},
       			{"target":"attr_B", "attr_list":["B0","B1","B2","B3","B4","B5","B6","B36"]}
			]
		},
		{
			"compaign_name": "campaign881",
			"price":5.25,
			"target_list" : [
       			{"target":"attr_C", "attr_list":["C0","C1","C2","C3","C4","C5","C99"]},
       			{"target":"attr_B", "attr_list":["B0","B1","B2","B3","B4","B5","B6","B36"]},
       			{"target":"attr_D", "attr_list":["D0","D1","D2","D3","D4","D5","D99"]},
       			{"target":"attr_E", "attr_list":["E0","E1","E2","E3","E4","E5","E99"]},
       			{"target":"attr_F", "attr_list":["F0","F1","F2","F3","F4","F5","F6","F36"]}
			]
		},
		{
			"compaign_name": "campaign870",
			"price":3.25,
			"target_list" : [
				{"target":"attr_A", "attr_list":["A0","A1","A2","A3","A4","A5","A99"]},
       			{"target":"attr_C", "attr_list":["C0","C1","C2","C3","C4","C5","C99"]},
       			{"target":"attr_B", "attr_list":["B0","B1","B2","B3","B4","B5","B6","B36"]}
			]
		},
		{
			"compaign_name": "campaign881",
			"price":5.25,
			"target_list" : [
       			{"target":"attr_C", "attr_list":["C0","C1","C2","C3","C4","C5","C99"]},
       			{"target":"attr_B", "attr_list":["B0","B1","B2","B3","B4","B5","B6","B36"]},
       			{"target":"attr_D", "attr_list":["D0","D1","D2","D3","D4","D5","D99"]},
       			{"target":"attr_E", "attr_list":["E0","E1","E2","E3","E4","E5","E99"]},
       			{"target":"attr_F", "attr_list":["F0","F1","F2","F3","F4","F5","F6","F36"]}
			]
		},
		{
			"compaign_name": "campaign870",
			"price":3.25,
			"target_list" : [
				{"target":"attr_A", "attr_list":["A0","A1","A2","A3","A4","A5","A99"]},
       			{"target":"attr_C", "attr_list":["C0","C1","C2","C3","C4","C5","C99"]},
       			{"target":"attr_B", "attr_list":["B0","B1","B2","B3","B4","B5","B6","B36"]}
			]
		},
		{
			"compaign_name": "campaign881",
			"price":5.25,
			"target_list" : [
				{"target":"attr_A", "attr_list":["A0","A1","A2","A3","A4","A5","A99"]},
       			{"target":"attr_C", "attr_list":["C0","C1","C2","C3","C4","C5","C99"]},
       			{"target":"attr_B", "attr_list":["B0","B1","B2","B3","B4","B5","B6","B36"]},
       			{"target":"attr_D", "attr_list":["D0","D1","D2","D3","D4","D5","D99"]},
       			{"target":"attr_E", "attr_list":["E0","E1","E2","E3","E4","E5","E99"]},
       			{"target":"attr_F", "attr_list":["F0","F1","F2","F3","F4","F5","F6","F36"]}
			]
		}
	]`)

var res []byte
func TestMarshall(t *testing.T) {

	c := CampaignCollection{}
	easyjson.Unmarshal(jsonCampaignCollection, &c)

	_, err := easyjson.Marshal(c)
	if err != nil {
		t.Fatal(err)
	}
}

func BenchmarkMarshall(b *testing.B) {
	var err error
	b.StartTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		c := CampaignCollection{}
		easyjson.Unmarshal(jsonCampaignCollection, &c)

		res, err = easyjson.Marshal(c)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshallJSON(b *testing.B) {
	var err error
	b.StartTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		c := CampaignCollection{}
		json.Unmarshal(jsonCampaignCollection, &c)

		res, err = json.Marshal(c)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshallSingle(b *testing.B) {
	json := []byte(`
		{
			"compaign_name": "campaign870",
			"price":3.25,
			"target_list" : [
       			{"target":"attr_A", "attr_list":["A0","A99"]},
       			{"target":"attr_B", "attr_list":["B0","B36"]}
			]
		}`)

	var err error
	b.StartTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		c := Campaign{}
		easyjson.Unmarshal(json, &c)

		res, err = easyjson.Marshal(c)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshallJSONSingle(b *testing.B) {
	jsonRaw := []byte(`
		{
			"compaign_name": "campaign870",
			"price":3.25,
			"target_list" : [
       			{"target":"attr_A", "attr_list":["A0","A99"]},
       			{"target":"attr_B", "attr_list":["B0","B36"]}
			]
		}`)

	var err error
	b.StartTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		c := Campaign{}
		json.Unmarshal(jsonRaw, &c)

		res, err = json.Marshal(c)
		if err != nil {
			b.Fatal(err)
		}
	}
}
