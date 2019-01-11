package nest

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// SmokeCOAlarm contains all the data for an individual Nest smoke/co alarm
type SmokeCOAlarm struct {
	DeviceID           string `json:"device_id"`
	Locale             string `json:"locale"`
	SoftwareVersion    string `json:"software_version"`
	StructureID        string `json:"structure_id"`
	Name               string `json:"name"`
	NameLong           string `json:"name_long"`
	LastConnection     string `json:"last_connection"`
	IsOnline           bool   `json:"is_online"`
	BatteryHealth      string `json:"battery_health"`
	COAlarmState       string `json:"co_alarm_state"`
	SmokeAlarmState    string `json:"smoke_alarm_state"`
	IsManualTestActive bool   `json:"is_manual_test_active"`
	LastManualTestTime string `json:"last_manual_test_time"`
	UIColorState       string `json:"ui_color_state"`
	WhereID            string `json:"where_id"`
	WhereName          string `json:"where_name"`
}

// GetSmokeCOAlarms returns all Nest smoke/co alarms along with all their data
func (n *Connection) GetSmokeCOAlarms() ([]SmokeCOAlarm, error) {
	url := n.setURL("smoke_co_alarms")

	data, err := n.execute(url, "GET", nil)
	if err != nil {
		return []SmokeCOAlarm{}, err
	}

	// No smoke/co alarms found
	if len(data) == 0 {
		return []SmokeCOAlarm{}, nil
	}

	// Create map to store JSON response
	fullResponse := make(map[string]interface{})

	err = json.Unmarshal(data, &fullResponse)
	if err != nil {
		return []SmokeCOAlarm{}, err
	}

	allSmokeCOAlarms := []SmokeCOAlarm{}

	for _, val := range fullResponse {
		smokeCOAlarm := SmokeCOAlarm{}

		s, err := json.Marshal(val)
		if err != nil {
			return []SmokeCOAlarm{}, err
		}

		err = json.Unmarshal(s, &smokeCOAlarm)
		if err != nil {
			return []SmokeCOAlarm{}, err
		}

		allSmokeCOAlarms = append(allSmokeCOAlarms, smokeCOAlarm)
	}

	return allSmokeCOAlarms, nil
}

// GetSmokeCOAlarm returns the specified Nest smoke/co alarm by device id along with its data
func (n *Connection) GetSmokeCOAlarm(deviceID string) (SmokeCOAlarm, error) {
	// Error checking
	if strings.Trim(deviceID, " ") == "" {
		return SmokeCOAlarm{}, errors.New("Device ID must not be empty")
	}

	url := n.setURL(fmt.Sprintf("smoke_co_alarms/%s", deviceID))

	data, err := n.execute(url, "GET", nil)
	if err != nil {
		return SmokeCOAlarm{}, err
	}

	// Smoke/CO alarm not found
	if len(data) == 0 {
		return SmokeCOAlarm{}, fmt.Errorf("Smoke/CO Alarm %s not found", deviceID)
	}

	smokeCOAlarm := SmokeCOAlarm{}

	err = json.Unmarshal(data, &smokeCOAlarm)
	if err != nil {
		return SmokeCOAlarm{}, err
	}

	return smokeCOAlarm, nil
}

// GetSmokeCOAlarmLocale returns the locale of the specified smoke/co alarm
func (n *Connection) GetSmokeCOAlarmLocale(deviceID string) (string, error) {
	return n.getValue("smoke_co_alarms", deviceID, "locale")
}

// GetSmokeCOAlarmSoftwareVersion returns the software version of the specified smoke/co alarm
func (n *Connection) GetSmokeCOAlarmSoftwareVersion(deviceID string) (string, error) {
	return n.getValue("smoke_co_alarms", deviceID, "software_version")
}

// GetSmokeCOAlarmName returns the name of the specified smoke/co alarm
func (n *Connection) GetSmokeCOAlarmName(deviceID string) (string, error) {
	return n.getValue("smoke_co_alarms", deviceID, "name")
}

// GetSmokeCOAlarmLastConnection gets the last connection time of the specified smoke/co alarm
func (n *Connection) GetSmokeCOAlarmLastConnection(deviceID string) (string, error) {
	return n.getValue("smoke_co_alarms", deviceID, "last_connection")
}

// IsSmokeCOAlarmOnline returns true if the specified smoke/co alarm is online, false if it isn't
func (n *Connection) IsSmokeCOAlarmOnline(deviceID string) (string, error) {
	return n.getValue("smoke_co_alarms", deviceID, "is_online")
}

// GetSmokeCOAlarmBatteryHealth gets the battery health of the specified smoke/co alarm
func (n *Connection) GetSmokeCOAlarmBatteryHealth(deviceID string) (string, error) {
	return n.getValue("smoke_co_alarms", deviceID, "battery_health")
}

// GetCOAlarmState gets the Carbon Monoxide (CO) alarm status
func (n *Connection) GetCOAlarmState(deviceID string) (string, error) {
	return n.getValue("smoke_co_alarms", deviceID, "co_alarm_state")
}

// GetSmokeAlarmState gets the Smoke alarm status
func (n *Connection) GetSmokeAlarmState(deviceID string) (string, error) {
	return n.getValue("smoke_co_alarms", deviceID, "smoke_alarm_state")
}
