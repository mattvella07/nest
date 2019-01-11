package nest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
)

type thermostatTestData struct {
	Abc Thermostat `json:"abc"`
}

type smokeCOAlarmTestData struct {
	Abc SmokeCOAlarm `json:"abc"`
}

type cameraTestData struct {
	Abc Camera `json:"abc"`
}

type structureTestData struct {
	Abc Structure `json:"abc"`
}

func createTestConnection(scenario int) (Connection, *httptest.Server) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if scenario == 1 {
			if r.Method == "GET" {
				returnData := generateTestData(r.URL.String())
				w.Write(returnData)
			} else {
				w.Write([]byte("Succes"))
			}

		} else {
			w.Write(nil)
		}
	}))

	return Connection{
		AccessToken: "TEST",
		testURL:     fmt.Sprintf("%s/devices", server.URL),
	}, server
}

func generateTestData(url string) []byte {
	returnData := []byte{}
	var err error

	switch url {
	case "/devices/thermostats":
		data := thermostatTestData{
			Abc: Thermostat{
				Humidity:                  30,
				Locale:                    "en-US",
				TemperatureScale:          "F",
				IsUsingEmergencyHeat:      false,
				HasFan:                    true,
				SoftwareVersion:           "1.0",
				HasLeaf:                   false,
				WhereID:                   "location",
				DeviceID:                  "abc",
				Name:                      "test thermostat",
				CanHeat:                   true,
				CanCool:                   true,
				TargetTemperatureC:        20,
				TargetTemperatureF:        68,
				TargetTemperatureHighC:    23,
				TargetTemperatureHighF:    74,
				TargetTemperatureLowC:     21.5,
				TargetTemperatureLowF:     72,
				AmbientTemperatureC:       23,
				AmbientTemperatureF:       74,
				AwayTemperatureHighC:      25.5,
				AwayTemperatureHighF:      78,
				AwayTemperatureLowC:       15.5,
				AwayTemperatureLowF:       60,
				EcoTemperatureHighC:       25.5,
				EcoTemperatureHighF:       78,
				EcoTemperatureLowC:        15.5,
				EcoTemperatureLowF:        60,
				IsLocked:                  false,
				LockedTempMinC:            20,
				LockedTempMinF:            68,
				LockedTempMaxC:            22,
				LockedTempMaxF:            72,
				SunlightCorrectionActive:  false,
				SunlightCorrectionEnabled: true,
				StructureID:               "abc123",
				FanTimerActive:            false,
				FanTimerTimeout:           "1970-01-01T00:00:00.000Z",
				FanTimerDuration:          15,
				PreviousHVACMode:          "",
				HVACMode:                  "heat-cool",
				TimeToTarget:              "~0",
				TimeToTargetTraining:      "ready",
				WhereName:                 "location",
				Label:                     "",
				NameLong:                  "test thermostat",
				IsOnline:                  true,
				LastConnection:            "2019-01-02T14:27:53.729Z",
				HVACState:                 "off",
			},
		}

		returnData, err = json.Marshal(data)
		if err != nil {
			fmt.Println("ERR: ", err)
		}
	case "/devices/thermostats/abc":
		data := Thermostat{
			Humidity:                  30,
			Locale:                    "en-US",
			TemperatureScale:          "F",
			IsUsingEmergencyHeat:      false,
			HasFan:                    true,
			SoftwareVersion:           "1.0",
			HasLeaf:                   false,
			WhereID:                   "location",
			DeviceID:                  "abc",
			Name:                      "test thermostat",
			CanHeat:                   true,
			CanCool:                   true,
			TargetTemperatureC:        20,
			TargetTemperatureF:        68,
			TargetTemperatureHighC:    23,
			TargetTemperatureHighF:    74,
			TargetTemperatureLowC:     21.5,
			TargetTemperatureLowF:     72,
			AmbientTemperatureC:       23,
			AmbientTemperatureF:       74,
			AwayTemperatureHighC:      25.5,
			AwayTemperatureHighF:      78,
			AwayTemperatureLowC:       15.5,
			AwayTemperatureLowF:       60,
			EcoTemperatureHighC:       25.5,
			EcoTemperatureHighF:       78,
			EcoTemperatureLowC:        15.5,
			EcoTemperatureLowF:        60,
			IsLocked:                  false,
			LockedTempMinC:            20,
			LockedTempMinF:            68,
			LockedTempMaxC:            22,
			LockedTempMaxF:            72,
			SunlightCorrectionActive:  false,
			SunlightCorrectionEnabled: true,
			StructureID:               "abc123",
			FanTimerActive:            false,
			FanTimerTimeout:           "1970-01-01T00:00:00.000Z",
			FanTimerDuration:          15,
			PreviousHVACMode:          "",
			HVACMode:                  "heat-cool",
			TimeToTarget:              "~0",
			TimeToTargetTraining:      "ready",
			WhereName:                 "location",
			Label:                     "",
			NameLong:                  "test thermostat",
			IsOnline:                  true,
			LastConnection:            "2019-01-02T14:27:53.729Z",
			HVACState:                 "off",
		}

		returnData, err = json.Marshal(data)
		if err != nil {
			fmt.Println("ERR: ", err)
		}
	case "/devices/thermostats/abc/locale":
		returnData = []byte("en-US")
	case "/devices/thermostats/abc/software_version":
		returnData = []byte("1.0")
	case "/devices/thermostats/abc/name":
		returnData = []byte("test thermostat")
	case "/devices/thermostats/abc/last_connection":
		returnData = []byte("2016-12-31T23:59:59.000Z")
	case "/devices/thermostats/abc/is_online":
		returnData = []byte("true")
	case "/devices/thermostats/abc/temperature_scale", "/devices/thermostats/def/temperature_scale":
		returnData = []byte("F")
	case "/devices/thermostats/abc/target_temperature_f":
		returnData = []byte("68")
	case "/devices/thermostats/abc/target_temperature_high_f":
		returnData = []byte("72")
	case "/devices/thermostats/abc/target_temperature_low_f":
		returnData = []byte("70")
	case "/devices/thermostats/abc/hvac_mode":
		returnData = []byte("heat-cool")
	case "/devices/thermostats/abc/label":
		returnData = []byte("test thermostat label")
	case "/devices/smoke_co_alarms":
		data := smokeCOAlarmTestData{
			Abc: SmokeCOAlarm{
				DeviceID:           "abc",
				Locale:             "en-US",
				SoftwareVersion:    "1.0",
				StructureID:        "123",
				Name:               "test smoke alarm",
				NameLong:           "test smoke/co alarm",
				LastConnection:     "2019-01-02T14:27:53.729Z",
				IsOnline:           true,
				BatteryHealth:      "ok",
				COAlarmState:       "ok",
				SmokeAlarmState:    "ok",
				IsManualTestActive: true,
				LastManualTestTime: "2019-01-01T14:27:53.729Z",
				UIColorState:       "green",
				WhereID:            "qwerty",
				WhereName:          "Kitchen",
			},
		}

		returnData, err = json.Marshal(data)
		if err != nil {
			fmt.Println("ERR: ", err)
		}
	case "/devices/smoke_co_alarms/abc":
		data := SmokeCOAlarm{
			DeviceID:           "abc",
			Locale:             "en-US",
			SoftwareVersion:    "1.0",
			StructureID:        "123",
			Name:               "test smoke alarm",
			NameLong:           "test smoke/co alarm",
			LastConnection:     "2019-01-02T14:27:53.729Z",
			IsOnline:           true,
			BatteryHealth:      "ok",
			COAlarmState:       "ok",
			SmokeAlarmState:    "ok",
			IsManualTestActive: true,
			LastManualTestTime: "2019-01-01T14:27:53.729Z",
			UIColorState:       "green",
			WhereID:            "qwerty",
			WhereName:          "Kitchen",
		}

		returnData, err = json.Marshal(data)
		if err != nil {
			fmt.Println("ERR: ", err)
		}
	case "/devices/smoke_co_alarms/abc/locale":
		returnData = []byte("en-US")
	case "/devices/smoke_co_alarms/abc/software_version":
		returnData = []byte("1.0")
	case "/devices/smoke_co_alarms/abc/name":
		returnData = []byte("test smoke alarm")
	case "/devices/smoke_co_alarms/abc/last_connection":
		returnData = []byte("2019-01-02T14:27:53.729Z")
	case "/devices/smoke_co_alarms/abc/is_online":
		returnData = []byte("true")
	case "/devices/smoke_co_alarms/abc/battery_health":
		returnData = []byte("ok")
	case "/devices/smoke_co_alarms/abc/co_alarm_state":
		returnData = []byte("ok")
	case "/devices/smoke_co_alarms/abc/smoke_alarm_state":
		returnData = []byte("ok")
	case "/devices/cameras":
		data := cameraTestData{
			Abc: Camera{
				DeviceID:              "abc",
				SoftwareVersion:       "1.0",
				StructureID:           "abc123",
				WhereID:               "location",
				WhereName:             "location",
				Name:                  "test camera",
				NameLong:              "test camera",
				IsOnline:              true,
				IsStreaming:           true,
				IsAudioInputEnabled:   true,
				LastIsOnlineChange:    "2016-12-29T18:42:00.000Z",
				IsVideoHistoryEnabled: true,
				WebURL:                "https://home.nest.com/cameras/abc?auth=camera_token",
				AppURL:                "nestmobile://cameras/abc?auth=camera_token",
				IsPublicShareEnabled:  false,
			},
		}

		returnData, err = json.Marshal(data)
		if err != nil {
			fmt.Println("ERR: ", err)
		}
	case "/devices/cameras/abc":
		data := Camera{
			DeviceID:              "abc",
			SoftwareVersion:       "1.0",
			StructureID:           "abc123",
			WhereID:               "location",
			WhereName:             "location",
			Name:                  "test camera",
			NameLong:              "test camera",
			IsOnline:              true,
			IsStreaming:           true,
			IsAudioInputEnabled:   true,
			LastIsOnlineChange:    "2016-12-29T18:42:00.000Z",
			IsVideoHistoryEnabled: true,
			WebURL:                "https://home.nest.com/cameras/abc?auth=camera_token",
			AppURL:                "nestmobile://cameras/abc?auth=camera_token",
			IsPublicShareEnabled:  false,
		}

		returnData, err = json.Marshal(data)
		if err != nil {
			fmt.Println("ERR: ", err)
		}
	case "/devices/cameras/abc/software_version":
		returnData = []byte("1.0")
	case "/devices/cameras/abc/name":
		returnData = []byte("test camera")
	case "/devices/cameras/abc/is_online":
		returnData = []byte("true")
	case "/devices/cameras/abc/is_streaming":
		returnData = []byte("true")
	case "/devices/cameras/abc/is_audio_input_enabled":
		returnData = []byte("true")
	case "/devices/cameras/abc/is_video_history_enabled":
		returnData = []byte("true")
	case "/devices/cameras/abc/web_url":
		returnData = []byte("https://home.nest.com/cameras/abc?auth=camera_token")
	case "/devices/cameras/abc/app_url":
		returnData = []byte("nestmobile://cameras/abc?auth=camera_token")
	case "/devices/cameras/abc/last_event":
		returnData = []byte("event")
	case "/devices/structures":
		data := structureTestData{
			Abc: Structure{
				StructureID:         "abc123",
				Thermostats:         []string{"123"},
				SmokeCOAlarms:       []string{"456"},
				Cameras:             []string{"789"},
				Away:                "home",
				Name:                "test structure",
				CountryCode:         "US",
				PostalCode:          "12345",
				PeakPeriodStartTime: "2016-12-29T18:42:00.000Z",
				PeakPeriodEndTime:   "2016-12-39T18:42:00.000Z",
				TimeZone:            "America/Chicago",
				ETA:                 []string{"trip1"},
				ETABegin:            "2016-12-29T18:42:00.000Z",
				RHREnrollment:       true,
				WWNSecurityState:    "ok",
				COAlarmState:        "ok",
				SmokeAlarmState:     "ok",
			},
		}

		returnData, err = json.Marshal(data)
		if err != nil {
			fmt.Println("ERR: ", err)
		}
	case "/devices/structures/abc":
		data := Structure{
			StructureID:         "abc123",
			Thermostats:         []string{"123"},
			SmokeCOAlarms:       []string{"456"},
			Cameras:             []string{"789"},
			Away:                "home",
			Name:                "test structure",
			CountryCode:         "US",
			PostalCode:          "12345",
			PeakPeriodStartTime: "2016-12-29T18:42:00.000Z",
			PeakPeriodEndTime:   "2016-12-39T18:42:00.000Z",
			TimeZone:            "America/Chicago",
			ETA:                 []string{"trip1"},
			ETABegin:            "2016-12-29T18:42:00.000Z",
			RHREnrollment:       true,
			WWNSecurityState:    "ok",
			COAlarmState:        "ok",
			SmokeAlarmState:     "ok",
		}

		returnData, err = json.Marshal(data)
		if err != nil {
			fmt.Println("ERR: ", err)
		}
	case "/devices/structures/abc/thermostats":
		returnData = []byte("[\"123\"]")
	case "/devices/structures/abc/smoke_co_alarms":
		returnData = []byte("[\"456\"]")
	case "/devices/structures/abc/cameras":
		returnData = []byte("[\"789\"]")
	case "/devices/structures/abc/away":
		returnData = []byte("home")
	case "/devices/structures/abc/name":
		returnData = []byte("test structure")
	}

	return returnData
}
