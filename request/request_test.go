package request

import (
	"net/url"
	"testing"
)

func TestRequest_Validate(t *testing.T) {
	type fields struct {
		Query               url.Values
		RequiredParams      []string
		RequiredOneOfParams [][]string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Test  #1: no required params",
			fields: fields{
				Query: url.Values{
					"k1": []string{"v1"},
				},
			},
			wantErr: false,
		},
		{
			name: "Test  #2: required param is missing",
			fields: fields{
				Query: url.Values{
					"k1": []string{"v1"},
				},
				RequiredParams: []string{"k1", "k2"},
			},
			wantErr: true,
		},
		{
			name: "Test  #3: all required param are present",
			fields: fields{
				Query: url.Values{
					"k1": []string{"v1"},
					"k2": []string{"v2"},
				},
				RequiredParams: []string{"k1", "k2"},
			},
			wantErr: false,
		},
		{
			name: "Test  #4: one of the `RequiredOneOfParams` param is present",
			fields: fields{
				Query: url.Values{
					"k1": []string{"v1"},
				},
				RequiredOneOfParams: [][]string{{"k1", "k2"}},
			},
			wantErr: false,
		},
		{
			name: "Test  #6: all of the `RequiredOneOfParams` params are missing",
			fields: fields{
				Query:               url.Values{},
				RequiredOneOfParams: [][]string{{"k1", "k2"}},
			},
			wantErr: true,
		},
		{
			name: "Test  #7: all required params are present",
			fields: fields{
				Query: url.Values{
					"k1": []string{"v1"},
					"k2": []string{"v2"},
				},
				RequiredParams:      []string{"k1"},
				RequiredOneOfParams: [][]string{{"k2", "k3"}},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Request{
				Query:               tt.fields.Query,
				RequiredParams:      tt.fields.RequiredParams,
				RequiredOneOfParams: tt.fields.RequiredOneOfParams,
			}
			if err := r.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Request.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
