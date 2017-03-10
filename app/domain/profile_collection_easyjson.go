// AUTOGENERATED FILE: easyjson marshaler/unmarshalers.

package domain

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson3704f49aDecodeGeneratorAppDomain(in *jlexer.Lexer, out *ProfileCollection) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
	} else {
		in.Delim('{')
		if !in.IsDelim('}') {
			*out = make(ProfileCollection)
		} else {
			*out = nil
		}
		for !in.IsDelim('}') {
			key := string(in.String())
			in.WantColon()
			var v1 Attribute
			v1 = Attribute(in.String())
			(*out)[key] = v1
			in.WantComma()
		}
		in.Delim('}')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson3704f49aEncodeGeneratorAppDomain(out *jwriter.Writer, in ProfileCollection) {
	if in == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v2First := true
		for v2Name, v2Value := range in {
			if !v2First {
				out.RawByte(',')
			}
			v2First = false
			out.String(string(v2Name))
			out.RawByte(':')
			out.String(string(v2Value))
		}
		out.RawByte('}')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v ProfileCollection) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3704f49aEncodeGeneratorAppDomain(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ProfileCollection) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3704f49aEncodeGeneratorAppDomain(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ProfileCollection) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3704f49aDecodeGeneratorAppDomain(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ProfileCollection) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3704f49aDecodeGeneratorAppDomain(l, v)
}
