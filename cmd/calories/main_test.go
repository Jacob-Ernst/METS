package main

import (
	"testing"
)

func Test_convertWeight(t *testing.T) {
	type args struct {
		kg float64
		lb float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "valid kg and lbs",
			args: args{kg: 90, lb: 600}, want: 90, wantErr: false,
		},
		{
			name: "invalid kg, but valid lbs",
			args: args{kg: -99, lb: 200}, want: 90.72, wantErr: false,
		},
		{
			name: "invalid kg and lbs",
			args: args{kg: -1, lb: -1}, want: -1, wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convertWeight(tt.args.kg, tt.args.lb)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertWeight() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("convertWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMET(t *testing.T) {
	openDB("file::memory:?cache=shared")

	type args struct {
		activity string
		MET      float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "valid activity and METs",
			args: args{activity: "power mower", MET: 5}, want: 5, wantErr: false,
		},
		{
			name: "invalid activity, but valid METs",
			args: args{activity: "", MET: 7}, want: 7, wantErr: false,
		},
		{
			name: "invalid METs, but valid activity",
			args: args{activity: "power mower", MET: -1}, want: 4.5, wantErr: false,
		},
		{
			name: "invalid activity and METs",
			args: args{activity: "", MET: -1}, want: -1, wantErr: true,
		},
		{
			name: "activity not found",
			args: args{activity: "fake exercise", MET: -1}, want: -1, wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getMET(tt.args.activity, tt.args.MET)
			if (err != nil) != tt.wantErr {
				t.Errorf("getMET() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getMET() = %v, want %v", got, tt.want)
			}
		})
	}
}
