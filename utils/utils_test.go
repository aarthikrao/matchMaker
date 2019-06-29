package utils

import (
	"testing"
)

func TestAbs(t *testing.T) {
	type args struct {
		x float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{"Negative", args{-32}, 32},
		{"Positive", args{32}, 32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Abs(tt.args.x); got != tt.want {
				t.Errorf("Abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrToInt(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name      string
		args      args
		wantValue int
		wantErr   bool
	}{
		{"Postive case", args{"10"}, 10, false},
		{"Negative case", args{"abcd"}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, err := StrToInt(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("StrToInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotValue != tt.wantValue {
				t.Errorf("StrToInt() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestStrToFloat32(t *testing.T) {
	type args struct {
		f string
	}
	tests := []struct {
		name            string
		args            args
		wantFloatString float32
		wantErr         bool
	}{
		{"Postive case", args{"10.123"}, 10.123, false},
		{"Negative case", args{"abcd"}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFloatString, err := StrToFloat32(tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("StrToFloat32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotFloatString != tt.wantFloatString {
				t.Errorf("StrToFloat32() = %v, want %v", gotFloatString, tt.wantFloatString)
			}
		})
	}
}
