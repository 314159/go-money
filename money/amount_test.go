package money

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		arg     string
		want    *amount
		wantErr bool
	}{
		{
			name:    "Big negative number with decimal digits",
			arg:     "-999999999999999999.999999999999999999",
			want:    &amount{s: "-999999999999999999.999999999999999999"},
			wantErr: false,
		},
		{
			name:    "Big negative number",
			arg:     "-999999999999999999",
			want:    &amount{s: "-999999999999999999"},
			wantErr: false,
		},
		{
			name:    "Big number",
			arg:     "999999999999999999",
			want:    &amount{s: "999999999999999999"},
			wantErr: false,
		},
		{
			name:    "Too many decimals",
			arg:     "12.45.6",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Bad decimal value",
			arg:     "123.oo",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Bad integer value",
			arg:     "A.12",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "+1",
			arg:     "+1",
			want:    &amount{s: "1"},
			wantErr: false,
		},
		{
			name:    "+",
			arg:     "+",
			want:    &amount{s: "0"},
			wantErr: false,
		},
		{
			name:    "-",
			arg:     "-",
			want:    &amount{s: "0"},
			wantErr: false,
		},
		{
			name:    "Fred",
			arg:     "Fred",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Empty string",
			arg:     "",
			want:    &amount{s: "0"},
			wantErr: false,
		},
		{
			name:    "0",
			arg:     "0",
			want:    &amount{s: "0"},
			wantErr: false,
		},
		{
			name:    ".0",
			arg:     ".0",
			want:    &amount{s: "0.0"},
			wantErr: false,
		},
		{
			name:    "-1",
			arg:     "-1",
			want:    &amount{s: "-1"},
			wantErr: false,
		},
		{
			name:    "-.1",
			arg:     "-.1",
			want:    &amount{s: "-0.1"},
			wantErr: false,
		},
		{
			name:    "+.1",
			arg:     "+.1",
			want:    &amount{s: "0.1"},
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
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("New() = %v, want %v", got, tt.want)
				}

				s := got.String()
				if s != tt.want.s {
					t.Errorf("New().String() = %v, want %v", got, tt.want.s)
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
				a:  &amount{s: "0.1"},
				dd: -1,
			},
			want: &amount{s: "1"},
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
			name: "invalid amount", // Just for code coverage; this shouldn't happen
			given: args{
				a:  &amount{s: "12.0.3"},
				dd: 12,
			},
			want: nil,
		},
		{
			name: "ScaleUp",
			given: args{
				a:  &amount{s: "12.1"},
				dd: -2,
			},
			want: &amount{s: "1210"},
		},
		{
			name: "'12.1', -1",
			given: args{
				a:  &amount{s: "12.1"},
				dd: -1,
			},
			want: &amount{s: "121"},
		},
		{
			name: "'12.1', 0",
			given: args{
				a:  &amount{s: "12.1"},
				dd: 0,
			},
			want: &amount{s: "12"},
		},
		{
			name: "'12.1', 1",
			given: args{
				a:  &amount{s: "12.1"},
				dd: 1,
			},
			want: &amount{s: "12.1"},
		},
		{
			name: "'12', 1",
			given: args{
				a:  &amount{s: "12"},
				dd: 1,
			},
			want: &amount{s: "12.0"},
		},
		{
			name: "'12', 0",
			given: args{
				a:  &amount{s: "12"},
				dd: 0,
			},
			want: &amount{s: "12"},
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
