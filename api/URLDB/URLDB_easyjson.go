// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package URLDB

import (
	"encoding/json"
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonE715c2a3DecodeUsersDorontabakmanGoSrc(in *jlexer.Lexer, out *URLDB) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Domains":
			if in.IsNull() {
				in.Skip()
				out.Domains = nil
			} else {
				in.Delim('[')
				if out.Domains == nil {
					if !in.IsDelim(']') {
						out.Domains = make([]string, 0, 4)
					} else {
						out.Domains = []string{}
					}
				} else {
					out.Domains = (out.Domains)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Domains = append(out.Domains, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonE715c2a3EncodeUsersDorontabakmanGoSrc(out *jwriter.Writer, in URLDB) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Domains\":"
		out.RawString(prefix[1:])
		if in.Domains == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Domains {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v URLDB) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE715c2a3EncodeUsersDorontabakmanGoSrc(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v URLDB) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE715c2a3EncodeUsersDorontabakmanGoSrc(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *URLDB) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE715c2a3DecodeUsersDorontabakmanGoSrc(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *URLDB) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE715c2a3DecodeUsersDorontabakmanGoSrc(l, v)
}