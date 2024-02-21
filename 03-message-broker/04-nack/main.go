package main

import (
	"context"
	"fmt"

	"github.com/ThreeDotsLabs/watermill/message"
)

type AlarmClient interface {
	StartAlarm() error
	StopAlarm() error
}

func ConsumeMessages(sub message.Subscriber, alarmClient AlarmClient) {
	messages, err := sub.Subscribe(context.Background(), "smoke_sensor")
	if err != nil {
		panic(err)
	}

	for msg := range messages {
		smokeValue := string(msg.Payload)

		if smokeValue == "0" {
			err := alarmClient.StopAlarm()
			if err != nil {
				fmt.Println("Error with alarm state:", err)
				msg.Nack()
				continue
			}
		}

		if smokeValue == "1" {
			err := alarmClient.StartAlarm()
			if err != nil {
				fmt.Println("Error with alarm state:", err)
				msg.Nack()
				continue
			}
		}

		msg.Ack()
	}
}
