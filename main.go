package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type signal struct {
	centerFrequency int64
	bandwidth       int64
	power           float32
	name            string
}

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	for _, message := range sqsEvent.Records {
		fmt.Printf("The message %s for event source %s = %s \n", message.MessageId, message.EventSource, message.Body)
	}

	return nil
}

func processFFT(fftArray []float32) []signal {
	s := signal{centerFrequency: 1500, bandwidth: 20, power: 10, name: "test"}
	var signals []signal
	var avPower = calcAverage(fftArray)
	for idx, power := range fftArray {
		if idx < len(fftArray)-1 && power > avPower && fftArray[idx+1] < avPower {

		}
	}
	signals = append(signals, s)
	return signals
}

func findPowerSpikes(fftArray []float32) [][]int {
	var spike []int
	var spikes [][]int
	var count int
	count = 0
	var inSignal bool
	avPowr := calcAverage(fftArray)
	for idx, power := range fftArray {
		if power > avPowr && !inSignal{
			spike = append(spike, idx)
			
		} else if power > avPowr && !(fftArray[idx+1] > avPowr) && inSignal{
			spike = append(spike, idx)
			spikes[count] = append(spikes[count], spike)
			spike = []int
			count++;
		}
	}
	return spikes
}

func calcAverage(fftArray []float32) float32 {
	var sum float32
	for _, power := range fftArray {
		sum += power
	}
	return sum / float32(len(fftArray))
}

func main() {
	lambda.Start(handler)
}
