package main

import (
	"bytes"
	"testing"
)

func Test_burnRate(t *testing.T) {
	type args struct {
		MET float64
		kg  float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test mowing the lawn",
			args: args{MET: 4.5, kg: 94.25},
			want: 7.4221875,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := burnRate(tt.args.MET, tt.args.kg); got != tt.want {
				t.Errorf("burnRate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_totalBurn(t *testing.T) {
	type args struct {
		kg   float64
		MET  float64
		time float64
	}
	tests := []struct {
		name  string
		args  args
		wantW string
	}{
		{
			name:  "Test mowing the lawn",
			args:  args{kg: 94.25, MET: 4.5, time: 30.00},
			wantW: "You burned 222.67 Calories\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			totalBurn(w, tt.args.kg, tt.args.MET, tt.args.time)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("totalBurn() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func Test_lbTokg(t *testing.T) {
	type args struct {
		lb float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "Test my weight", args: args{lb: 208}, want: 94.3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lbTokg(tt.args.lb); got != tt.want {
				t.Errorf("lbTokg() = %v, want %v", got, tt.want)
			}
		})
	}
}
