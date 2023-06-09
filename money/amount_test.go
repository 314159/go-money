package money

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type want struct {
		a *amount
		s string
	}
	tests := []struct {
		name    string
		arg     string
		wantAmt *amount
		wantStr string
		wantErr bool
	}{
		{
			name:    "Big negative number with decimal digits",
			arg:     "-999999999999999999.999999999999999999",
			wantAmt: &amount{sign: "-", intPart: "999999999999999999", fracPart: "999999999999999999"},
			wantStr: "-999999999999999999.999999999999999999",
			wantErr: false,
		},
		{
			name:    "Big negative number",
			arg:     "-999999999999999999",
			wantAmt: &amount{sign: "-", intPart: "999999999999999999"},
			wantStr: "-999999999999999999",
			wantErr: false,
		},
		{
			name:    "Big number",
			arg:     "999999999999999999",
			wantAmt: &amount{intPart: "999999999999999999"},
			wantStr: "999999999999999999",
			wantErr: false,
		},
		{
			name:    "Too many decimals",
			arg:     "12.45.6",
			wantAmt: nil,
			wantStr: "",
			wantErr: true,
		},
		{
			name:    "Bad decimal value",
			arg:     "123.oo",
			wantAmt: nil,
			wantStr: "",
			wantErr: true,
		},
		{
			name:    "Bad integer value",
			arg:     "A.12",
			wantAmt: nil,
			wantStr: "",
			wantErr: true,
		},
		{
			name:    "+1",
			arg:     "+1",
			wantAmt: &amount{intPart: "1"},
			wantStr: "1",
			wantErr: false,
		},
		{
			name:    "+",
			arg:     "+",
			wantAmt: &amount{intPart: "0"},
			wantStr: "0",
			wantErr: false,
		},
		{
			name:    "-",
			arg:     "-",
			wantAmt: &amount{intPart: "0"}, 
			wantStr: "0",
			wantErr: false,
		},
		{
			name:    "Fred",
			arg:     "Fred",
			wantAmt: nil,
			wantStr: "",
			wantErr: true,
		},
		{
			name:    "Empty string",
			arg:     "",
			wantAmt: &amount{intPart: "0"}, 
			wantStr: "0",
			wantErr: false,
		},
		{
			name:    "0",
			arg:     "0",
			wantAmt: &amount{intPart: "0"}, 
			wantStr: "0",
			wantErr: false,
		},
		{
			name:    ".0",
			arg:     ".0",
			wantAmt: &amount{intPart: "0", fracPart: "0"}, 
			wantStr: "0.0",
			wantErr: false,
		},
		{
			name:    "-1",
			arg:     "-1",
			wantAmt: &amount{sign: "-", intPart: "1"}, 
			wantStr: "-1",
			wantErr: false,
		},
		{
			name:    "-.1",
			arg:     "-.1",
			wantAmt: &amount{sign: "-", intPart: "0", fracPart: "1"}, 
			wantStr: "-0.1",
			wantErr: false,
		},
		{
			name:    "+.1",
			arg:     "+.1",
			wantAmt: &amount{intPart: "0", fracPart: "1"}, 
			wantStr: "0.1",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil {
				if !reflect.DeepEqual(got, tt.wantAmt) {
					t.Errorf("New() = %v, want %v", got, tt.wantAmt)
				}

				s := got.String()
				if s != tt.wantStr {
					t.Errorf("New().String() = %v, want %v", got, tt.wantStr)
				}
			}
		})
	}
}

func Test_amount_String(t *testing.T) {
	tests := []struct {
		name  string
		given *amount
		want  string
	}{
		{
			name:  "nil",
			given: nil,
			want:  "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.given.String(); got != tt.want {
				t.Errorf("amount.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_amount_SetDecimalDigits(t *testing.T) {
	type args struct {
		a  *amount
		dd int
	}
	tests := []struct {
		name  string
		given args
		want  *amount
	}{
		{
			name: ".1 Scaleup",
			given: args{
				a:  &amount{intPart: "0", fracPart: "1"},
				dd: -1,
			},
			want: &amount{intPart: "1"},
		},
		{
			name: "nil",
			given: args{
				a:  nil,
				dd: 12,
			},
			want: nil,
		},
		{
			name: "ScaleUp",
			given: args{
				a:  &amount{intPart: "12", fracPart: "1"},
				dd: -2,
			},
			want: &amount{intPart: "1210"},
		},
		{
			name: "'12.1', -1",
			given: args{
				a:  &amount{intPart: "12", fracPart: "1"},
				dd: -1,
			},
			want: &amount{intPart: "121"},
		},
		{
			name: "'12.1', 0",
			given: args{
				a:  &amount{intPart: "12", fracPart: "1"},
				dd: 0,
			},
			want: &amount{intPart: "12"},
		},
		{
			name: "'12.1', 1",
			given: args{
				a:  &amount{intPart: "12", fracPart: "1"},
				dd: 1,
			},
			want: &amount{intPart: "12", fracPart: "1"},
		},
		{
			name: "'12', 1",
			given: args{
				a:  &amount{intPart: "12"},
				dd: 1,
			},
			want: &amount{intPart: "12", fracPart: "0"},
		},
		{
			name: "'12', 0",
			given: args{
				a:  &amount{intPart: "12"},
				dd: 0,
			},
			want: &amount{intPart: "12"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.given.a.SetDecimalDigits(tt.given.dd); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("amount.SetDecimalDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
