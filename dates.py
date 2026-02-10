
def process_events(infosec_events, datalake_events, walmart_events):
    # given the following events input as each a list of a tuple of 2 (event_id, time)
    # with event_id as an integer and time as a string
    # infosec time is Europe/Berlin in m/d/yyyy hh:mm format
    # walmart time is America/Los_Angeles m/d/yy h:mm am/pm format
    # datalake time is UTC dd/mm/yyyy hh:mm:ss format
    # return a uniform events list comprised of the source, event_id, and time
    # output time format must be mm/dd/yyyy hh:mm:ss in UTC time
    # first datalake, then infosec, and finally walmart

    return []


def main():
    infosec_events = [
        (929217, "3/2/2021 04:22"),
        (28393, "12/22/2022 18:57"),
        (18929, "11/1/2023 20:21")
    ]
    walmart_events = [
        (823, "1/1/22 4:22 AM"),
        (29, "3/9/22 11:59 PM")
    ]
    datalake_events = [
        (91234867, "03/02/2021 04:22:29"),
        (234789, "22/12/2022 18:57:28"),
        (494123978, "30/01/2023 20:21:59"),
        (2082347, "15/01/2025 20:21:48")
    ]
    events = process_events(infosec_events, datalake_events, walmart_events)
    expected_output = [
        ("datalake", 91234867, "02/03/2021 04:22:29"),
        ("datalake", 234789, "12/22/2022 18:57:28"),
        ("datalake", 494123978, "01/30/2023 20:21:59"),
        ("datalake", 2082347, "01/15/2025 20:21:48"),
        ("infosec", 929217, "03/02/2021 03:22:00"),
        ("infosec", 28393, "12/22/2022 17:57:00"),
        ("infosec", 18929, "11/01/2023 19:21:00"),
        ("walmart", 823, "01/01/2022 12:22:00"),
        ("walmart", 29, "03/10/2022 07:59:00"),
    ]
    for i in range(len(events)):
        if events[i] != expected_output[i]:
            print(f"error! expected={expected_output[i]} actual={events[i]}")
            
    print("success!!!")

if __name__ == "__main__":
    main()
