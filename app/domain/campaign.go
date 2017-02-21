package domain

import (
	"encoding/json"
	"strings"

	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

type Campaign struct {
	CampaignName string           `json:"compaign_name"`
	Price        float64          `json:"price"`
	TargetList   TargetCollection `json:"target_list"`
}

func (c *Campaign) Verify() bool {
	return compose(
		c.IsPricePositive,
		c.IsCorrectName,
		c.VerifyTargetList,
	)()
}

const (
	//FIXME: should use either math.Bigint or env dependant params
	PriceThreshold         = 0.001
	CampaignPrefix         = "campaign"
	campaignNumberPosition = len(CampaignPrefix)
)

func (c *Campaign) IsPricePositive() bool {
	return c.Price > PriceThreshold
}

func (c *Campaign) IsCorrectName() bool {
	var ok bool

	if ok = strings.HasPrefix(c.CampaignName, CampaignPrefix); !ok {
		return false
	}

	if ok = IsLatinNumber(c.CampaignName, campaignNumberPosition); !ok {
		return false
	}

	return true
}

func (c *Campaign) VerifyTargetList() bool {
	return c.TargetList.Verify()
}

type CampaignCollection []Campaign

func (c CampaignCollection) Verify() bool {
	var ok bool
	for _, campaign := range c {
		if ok = campaign.Verify(); !ok {
			return false
		}
	}

	return true

}

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonFf86ed8bDecodeGeneratorAppDomain(in *jlexer.Lexer, out *CampaignCollection) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}

	in.Delim('[')
	if !in.IsDelim(']') {
		*out = make(CampaignCollection, 0, 5)
	} else {
		*out = CampaignCollection{}
	}
	for !in.IsDelim(']') {
		var v1 Campaign
		(v1).UnmarshalEasyJSON(in)
		*out = append(*out, v1)
		in.WantComma()
	}
	in.Delim(']')

	if isTopLevel {
		in.Consumed()
	}
}

func easyjsonFf86ed8bEncodeGeneratorAppDomain(out *jwriter.Writer, in CampaignCollection) {
	out.RawByte('[')
	for v2, v3 := range in {
		if v2 > 0 {
			out.RawByte(',')
		}
		(v3).MarshalEasyJSON(out)
	}
	out.RawByte(']')
}

// MarshalJSON supports json.Marshaler interface
func (v CampaignCollection) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFf86ed8bEncodeGeneratorAppDomain(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CampaignCollection) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFf86ed8bEncodeGeneratorAppDomain(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CampaignCollection) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFf86ed8bDecodeGeneratorAppDomain(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CampaignCollection) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFf86ed8bDecodeGeneratorAppDomain(l, v)
}
