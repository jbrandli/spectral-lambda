package main

import (
	"context"
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func Test_handler(t *testing.T) {
	type args struct {
		ctx      context.Context
		sqsEvent events.SQSEvent
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := handler(tt.args.ctx, tt.args.sqsEvent); (err != nil) != tt.wantErr {
				t.Errorf("handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

var oneSpikeArray = []float32{-90, -90, -30, -30, -90, -90}
var twoSpikeArray = []float32{-90, -90, -30, -30, -90, -90, -30, -30, -90, -90}

func Test_calcAverage(t *testing.T) {
	got := calcAverage(oneSpikeArray)
	if got != -70 {
		t.Errorf("calcAverage() = %v, want %v", got, -70)
	}
}

func Test_findPowerSpikes_CountsTheSpikes(t *testing.T) {
	got := findPowerSpikes(oneSpikeArray)
	if len(got) != 1 {
		t.Errorf("findPowerSpikes() = %v, want %v", len(got), 1)
	}
	if len(got[0]) != 2 {
		t.Errorf("Spike length= %v, want %v", len(got[0]), 2)
	} else if got[0][0] != 2 {
		t.Errorf("Spike start= %v, want %v", got[0][0], 2)
	} else if got[0][1] != 3 {
		t.Errorf("Spike end= %v, want %v", got[0][1], 3)
	}
	got = findPowerSpikes(twoSpikeArray)
	want := [][]int{{2, 3}, {6, 7}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got spikes=%v, want %v", got, want)
	}
}

func Test_processFFT(t *testing.T) {

	// var arg = args{fftArray: testArray}
	signals := []signal{signal{centerFrequency: 1500, bandwidth: 20, power: -30, name: "Signal 1"}}
	type args struct {
		fftArray []float32
	}
	tests := []struct {
		name string
		args args
		want []signal
	}{
		{name: "test1",
			args: args{fftArray: oneSpikeArray},
			want: signals},
		// s := signal{centerFrequency: 1500, bandwidth: 20, power: 10, name: "test"}
		// expected := []signal{s}
		// tests[0] := {name: "returns signal",args: []float32, want: s}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processFFT(tt.args.fftArray); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("processFFT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
