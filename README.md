# go-rachio
Golang API client for [Rachio Sprinkler Systems](https://rachio.com)

## Example Usage
 
```
r,err := rachio.NewClient("<token>")
if err != nil {
    log.Fatal(err)
}

person err := r.Self()
if err != nil {
    log.Fatal(err)
}

fmt.Printf("       ID:%s\n", person.ID)
fmt.Printf("User Name:%s\n", person.Username)
fmt.Printf("Full Name:%s\n", person.FullName)
fmt.Printf("   E-Mail:%s\n", person.Email)
``` 

## Example CLI tool

**cmd/rachio-cli**


### List devices
```
$ ./rachio-cli --token=27be2185-2cc0-4bb8-8b47-5e76150579b4 device list
ID                                    Name  Status  Model               Serial Number  MAC Addess
<deviceid>                            Home  ONLINE  GENERATION2_16ZONE  VRXXXXXXX      123456789000
```

### List Events

$ ./rachio-cli --token=<redacted> device events --device.id=<deviceid>
```
Date                           Type             Subtype             Topic     Summary
2019-07-17 17:01:33 -0600 MDT  DEVICE_ADD       ADD_DEVICE          DEVICE    Welcome to your Rachio Controller! Learn more about your new smart watering assistant.
2019-07-17 17:02:04 -0600 MDT  SCHEDULE_STATUS  SCHEDULE_STARTED    WATERING  Quick Run will run for 1 minutes.
2019-07-17 17:02:06 -0600 MDT  ZONE_STATUS      ZONE_STARTED        WATERING  Zone 1 began watering at 05:02 PM (MDT).
2019-07-17 17:02:33 -0600 MDT  ZONE_STATUS      ZONE_STOPPED        WATERING  Zone 1 stopped watering at 05:02 PM (MDT) for 30 seconds.
2019-07-17 17:02:35 -0600 MDT  SCHEDULE_STATUS  SCHEDULE_STOPPED    WATERING  Quick Run was manually stopped after 30 seconds.
2019-07-17 17:04:08 -0600 MDT  SCHEDULE_STATUS  SCHEDULE_STARTED    WATERING  Quick Run will run for 1 minutes.
2019-07-17 17:04:10 -0600 MDT  ZONE_STATUS      ZONE_STARTED        WATERING  Zone 2 began watering at 05:04 PM (MDT).
2019-07-17 17:04:19 -0600 MDT  ZONE_STATUS      ZONE_STOPPED        WATERING  Zone 2 stopped watering at 05:04 PM (MDT) for 12 seconds.
2019-07-17 17:04:20 -0600 MDT  SCHEDULE_STATUS  SCHEDULE_STOPPED    WATERING  Quick Run was manually stopped after 12 seconds.
2019-07-17 17:05:27 -0600 MDT  SCHEDULE_STATUS  SCHEDULE_STARTED    WATERING  Quick Run will run for 1 minutes.
2019-07-17 17:05:29 -0600 MDT  ZONE_STATUS      ZONE_STARTED        WATERING  Zone 3 began watering at 05:05 PM (MDT).
2019-07-17 17:05:33 -0600 MDT  ZONE_STATUS      ZONE_STOPPED        WATERING  Zone 3 stopped watering at 05:05 PM (MDT) for 7 seconds.
2019-07-17 17:05:35 -0600 MDT  SCHEDULE_STATUS  SCHEDULE_STOPPED    WATERING  Quick Run was manually stopped after 7 seconds.
2019-07-17 17:06:31 -0600 MDT  SCHEDULE_STATUS  SCHEDULE_STARTED    WATERING  Quick Run will run for 1 minutes.
2019-07-17 17:06:33 -0600 MDT  ZONE_STATUS      ZONE_STARTED        WATERING  Zone 4 began watering at 05:06 PM (MDT).
2019-07-17 17:06:47 -0600 MDT  ZONE_STATUS      ZONE_STOPPED        WATERING  Zone 4 stopped watering at 05:06 PM (MDT) for 18 seconds.
2019-07-17 17:06:49 -0600 MDT  SCHEDULE_STATUS  SCHEDULE_STOPPED    WATERING  Quick Run was manually stopped after 18 seconds.
2019-07-17 17:07:49 -0600 MDT  SCHEDULE_STATUS  SCHEDULE_STARTED    WATERING  Quick Run will run for 1 minutes.
```

# Notes
This is a work in progress and not affiliated with Rachio. It is a hobby project based on the documentation provided at https://rachio.readme.io/v1.0/docs.

