package main

import (
	"fmt"
	"time"
)

type InputEvent struct {
	ID   int
	Time string
}

type Event struct {
	Source string
	ID     int
	Time   string
}

func convertEvents(src_events []InputEvent, src_layout string, src_tz *time.Location, src_name string, out_events []Event) (out []Event, err error) {
	out = out_events
	for _, ev := range src_events {
		var (
			dt       time.Time
			out_time string
		)
		dt, err = time.ParseInLocation(src_layout, ev.Time, src_tz)
		if err != nil {
			return
		}
		const out_layout = "01/02/2006 15:04:05"
		out_time = dt.UTC().Format(out_layout)
		out_ev := Event{
			Source: src_name,
			ID:     ev.ID,
			Time:   out_time,
		}
		out = append(out, out_ev)
	}
	return
}

// stub — replace with your real implementation
func processEvents(
	infosec []InputEvent,
	datalake []InputEvent,
	walmart []InputEvent,
) (events []Event, err error) {
	events = make([]Event, 0, len(infosec)+len(datalake)+len(walmart))
	infosec_tz, _ := time.LoadLocation("Europe/Berlin")
	walmart_tz, _ := time.LoadLocation("America/Los_Angeles")
	events, err = convertEvents(datalake, "02/01/2006 15:04:05", time.UTC, "datalake", events)
	if err != nil {
		return
	}
	events, err = convertEvents(infosec, "1/2/2006 15:04", infosec_tz, "infosec", events)
	if err != nil {
		return
	}
	events, err = convertEvents(walmart, "1/2/06 3:04 PM", walmart_tz, "walmart", events)
	return
}

func main() {
	infosecEvents := []InputEvent{
		{929217, "3/2/2021 04:22"},
		{28393, "12/22/2022 18:57"},
		{18929, "11/1/2023 20:21"},
	}

	walmartEvents := []InputEvent{
		{823, "1/1/22 4:22 AM"},
		{29, "3/9/22 11:59 PM"},
	}

	datalakeEvents := []InputEvent{
		{91234867, "03/02/2021 04:22:29"},
		{234789, "22/12/2022 18:57:28"},
		{494123978, "30/01/2023 20:21:59"},
		{2082347, "15/01/2025 20:21:48"},
	}

	events, err := processEvents(infosecEvents, datalakeEvents, walmartEvents)
	if err != nil {
		fmt.Printf("err: %s", err)
		return
	}
	expectedOutput := []Event{
		{"datalake", 91234867, "02/03/2021 04:22:29"},
		{"datalake", 234789, "12/22/2022 18:57:28"},
		{"datalake", 494123978, "01/30/2023 20:21:59"},
		{"datalake", 2082347, "01/15/2025 20:21:48"},
		{"infosec", 929217, "03/02/2021 03:22:00"},
		{"infosec", 28393, "12/22/2022 17:57:00"},
		{"infosec", 18929, "11/01/2023 19:21:00"},
		{"walmart", 823, "01/01/2022 12:22:00"},
		{"walmart", 29, "03/10/2022 07:59:00"},
	}

	flag := false
	if len(events) != len(expectedOutput) {
		fmt.Printf("❌len(events)=%d while len(expectedOutput)=%d", len(events), len(expectedOutput))
		return
	}
	for i := range events {
		if events[i] != expectedOutput[i] {
			fmt.Printf("❌expected=%v actual=%v\n", expectedOutput[i], events[i])
			flag = true
		} else {
			fmt.Printf("✅ expected=%v actual=%v\n", expectedOutput[i], events[i])
		}
	}

	if !flag {
		fmt.Println("success!!!")
	}
}
