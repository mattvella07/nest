package nest

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// Thermostat contains all the data for an individual Nest thermostat
type Thermostat struct {
	Humidity                  int     `json:"humidity"`
	Locale                    string  `json:"locale"`
	TemperatureScale          string  `json:"temperature_scale"`
	IsUsingEmergencyHeat      bool    `json:"is_using_emergency_heat"`
	HasFan                    bool    `json:"has_fan"`
	SoftwareVersion           string  `json:"software_version"`
	HasLeaf                   bool    `json:"has_leaf"`
	WhereID                   string  `json:"where_id"`
	DeviceID                  string  `json:"device_id"`
	Name                      string  `json:"name"`
	CanHeat                   bool    `json:"can_heat"`
	CanCool                   bool    `json:"can_cool"`
	TargetTemperatureC        float32 `json:"target_temperature_c"`
	TargetTemperatureF        int     `json:"target_temperature_f"`
	TargetTemperatureHighC    float32 `json:"target_temperature_high_c"`
	TargetTemperatureHighF    int     `json:"target_temperature_high_f"`
	TargetTemperatureLowC     float32 `json:"target_temperature_low_c"`
	TargetTemperatureLowF     int     `json:"target_temperature_low_f"`
	AmbientTemperatureC       float32 `json:"ambient_temperature_c"`
	AmbientTemperatureF       int     `json:"ambient_temperature_f"`
	AwayTemperatureHighC      float32 `json:"away_temperature_high_c"`
	AwayTemperatureHighF      int     `json:"away_temperature_high_f"`
	AwayTemperatureLowC       float32 `json:"away_temperature_low_c"`
	AwayTemperatureLowF       int     `json:"away_temperature_low_f"`
	EcoTemperatureHighC       float32 `json:"eco_temperature_high_c"`
	EcoTemperatureHighF       int     `json:"eco_temperature_high_f"`
	EcoTemperatureLowC        float32 `json:"eco_temperature_low_c"`
	EcoTemperatureLowF        int     `json:"eco_temperature_low_f"`
	IsLocked                  bool    `json:"is_locked"`
	LockedTempMinC            float32 `json:"locked_temp_min_c"`
	LockedTempMinF            int     `json:"locked_temp_min_f"`
	LockedTempMaxC            float32 `json:"locked_temp_max_c"`
	LockedTempMaxF            int     `json:"locked_temp_max_f"`
	SunlightCorrectionActive  bool    `json:"sunlight_correction_active"`
	SunlightCorrectionEnabled bool    `json:"sunlight_correction_enabled"`
	StructureID               string  `json:"structure_id"`
	FanTimerActive            bool    `json:"fan_timer_active"`
	FanTimerTimeout           string  `json:"fan_timer_timeout"`
	FanTimerDuration          int     `json:"fan_timer_duration"`
	PreviousHVACMode          string  `json:"previous_hvac_mode"`
	HVACMode                  string  `json:"hvac_mode"`
	TimeToTarget              string  `json:"time_to_target"`
	TimeToTargetTraining      string  `json:"time_to_target_training"`
	WhereName                 string  `json:"where_name"`
	Label                     string  `json:"label"`
	NameLong                  string  `json:"name_long"`
	IsOnline                  bool    `json:"is_online"`
	LastConnection            string  `json:"last_connection"`
	HVACState                 string  `json:"hvac_state"`
}

// GetThermostats returns all Nest thermostats along with all their data
func (n *Connection) GetThermostats() ([]Thermostat, error) {
	url := n.setURL("thermostats")

	data, err := n.execute(url, "GET", nil)
	if err != nil {
		return []Thermostat{}, err
	}

	// No thermostats found
	if len(data) == 0 {
		return []Thermostat{}, nil
	}

	// Create map to store JSON response
	fullResponse := make(map[string]interface{})

	err = json.Unmarshal(data, &fullResponse)
	if err != nil {
		return []Thermostat{}, err
	}

	allThermostats := []Thermostat{}

	for _, val := range fullResponse {
		thermostat := Thermostat{}

		t, err := json.Marshal(val)
		if err != nil {
			return []Thermostat{}, err
		}

		err = json.Unmarshal(t, &thermostat)
		if err != nil {
			return []Thermostat{}, err
		}

		allThermostats = append(allThermostats, thermostat)
	}

	return allThermostats, nil
}

// GetThermostat returns the specified Nest thermostat by device id along with its data
func (n *Connection) GetThermostat(deviceID string) (Thermostat, error) {
	// Error checking
	if strings.Trim(deviceID, " ") == "" {
		return Thermostat{}, errors.New("Device ID must not be empty")
	}

	url := n.setURL(fmt.Sprintf("thermostats/%s", deviceID))

	data, err := n.execute(url, "GET", nil)
	if err != nil {
		return Thermostat{}, err
	}

	// Thermostat not found
	if len(data) == 0 {
		return Thermostat{}, fmt.Errorf("Thermostat %s not found", deviceID)
	}

	thermostat := Thermostat{}

	err = json.Unmarshal(data, &thermostat)
	if err != nil {
		return Thermostat{}, err
	}

	return thermostat, nil
}

// GetThermostatLocale returns the locale of the specified thermostat
func (n *Connection) GetThermostatLocale(deviceID string) (string, error) {
	return n.getValue("thermostats", deviceID, "locale")
}

// GetThermostatSoftwareVersion returns the software version of the specified thermostat
func (n *Connection) GetThermostatSoftwareVersion(deviceID string) (string, error) {
	return n.getValue("thermostats", deviceID, "software_version")
}

// GetThermostatName returns the name of the specified thermostat
func (n *Connection) GetThermostatName(deviceID string) (string, error) {
	return n.getValue("thermostats", deviceID, "name")
}

// GetThermostatLastConnection gets the last connection time of the specified thermostat
func (n *Connection) GetThermostatLastConnection(deviceID string) (string, error) {
	return n.getValue("thermostats", deviceID, "last_connection")
}

// IsThermostatOnline returns true if the specified thermostat is online, false if it isn't
func (n *Connection) IsThermostatOnline(deviceID string) (string, error) {
	return n.getValue("thermostats", deviceID, "is_online")
}

// GetTemperatureScale returns the temperature scale of the specified thermostat
func (n *Connection) GetTemperatureScale(deviceID string) (string, error) {
	return n.getValue("thermostats", deviceID, "temperature_scale")
}

// GetTargetTemperature returns the target temperature of the specified thermostat
func (n *Connection) GetTargetTemperature(deviceID string) (string, error) {
	scale, err := n.GetTemperatureScale(deviceID)
	if err != nil {
		return "", err
	}

	scale = strings.ToLower(strings.Replace(scale, "\"", "", -1))

	return n.getValue("thermostats", deviceID, fmt.Sprintf("target_temperature_%s", scale))
}

// GetTargetHighLowTemperature returns the target high and low temperatures of the specified thermostat
func (n *Connection) GetTargetHighLowTemperature(deviceID string) (string, string, error) {
	scale, err := n.GetTemperatureScale(deviceID)
	if err != nil {
		return "", "", err
	}

	scale = strings.ToLower(strings.Replace(scale, "\"", "", -1))

	high, err := n.getValue("thermostats", deviceID, fmt.Sprintf("target_temperature_high_%s", scale))
	if err != nil {
		return "", "", err
	}

	low, err := n.getValue("thermostats", deviceID, fmt.Sprintf("target_temperature_low_%s", scale))
	if err != nil {
		return "", "", err
	}

	return high, low, nil
}

// GetHVACMode returns the hvac mode of the specified thermostat
func (n *Connection) GetHVACMode(deviceID string) (string, error) {
	return n.getValue("thermostats", deviceID, "hvac_mode")
}

// GetThermostatLabel returns the lebl of the specified thermostat
func (n *Connection) GetThermostatLabel(deviceID string) (string, error) {
	return n.getValue("thermostats", deviceID, "label")
}

// SetTemperatureScale sets the temperature scale of the specified thermostat
func (n *Connection) SetTemperatureScale(deviceID, scale string) error {
	// Error checking
	scale = strings.Trim(scale, " ")
	validVals := []string{"F", "C"}

	if scale == "" {
		return errors.New("Scale must not be empty")
	}

	valid := false
	for _, v := range validVals {
		if scale == v {
			valid = true
		}
	}

	if !valid {
		return fmt.Errorf("Temperature Scale must be one of the following: %s", validVals)
	}

	vals := make(map[string]interface{})
	vals["temperature_scale"] = scale

	return n.setValue("thermostats", deviceID, vals)
}

// SetTargetTemperatureF changes the target temperature (F) of the specified thermostat
func (n *Connection) SetTargetTemperatureF(deviceID string, temp int) error {
	// Error checking
	if temp < 50 || temp > 90 {
		return errors.New("Target Temperature must be in the range of 50 - 90")
	}

	scale, err := n.GetTemperatureScale(deviceID)
	if err != nil {
		return err
	}

	if scale != "F" {
		return errors.New("Temperature Scale must be set to F")
	}

	vals := make(map[string]interface{})
	vals["target_temperature_f"] = temp

	return n.setValue("thermostats", deviceID, vals)
}

// SetTargetTemperatureC changes the target temperature (C) of the specified thermostat
func (n *Connection) SetTargetTemperatureC(deviceID string, temp float32) error {
	// Error checking
	if temp < 9 || temp > 32 {
		return errors.New("Target Temperature must be in the range of 9 - 32")
	}

	scale, err := n.GetTemperatureScale(deviceID)
	if err != nil {
		return err
	}

	if scale != "C" {
		return errors.New("Temperature Scale must be set to C")
	}

	vals := make(map[string]interface{})
	vals["target_temperature_c"] = temp

	return n.setValue("thermostats", deviceID, vals)
}

// SetTargetHighLowTemperatureF changes the target high and low temperatures (F) of
// the specified thermostat
func (n *Connection) SetTargetHighLowTemperatureF(deviceID string, high, low int) error {
	// Error checking
	if high < 50 || high > 90 {
		return errors.New("Target High Temperature must be in the range of 50 - 90")
	}

	if low < 50 || low > 90 {
		return errors.New("Target Low Temperature must be in the range of 50 - 90")
	}

	scale, err := n.GetTemperatureScale(deviceID)
	if err != nil {
		return err
	}

	if scale != "F" {
		return errors.New("Temperature Scale must be set to F")
	}

	vals := make(map[string]interface{})
	vals["target_temperature_high_f"] = high
	vals["target_temperature_low_f"] = low

	return n.setValue("thermostats", deviceID, vals)
}

// SetTargetHighLowTemperatureC changes the target high and low temperatures (C) of
// the specified thermostat
func (n *Connection) SetTargetHighLowTemperatureC(deviceID string, high, low float32) error {
	// Error checking
	if high < 9 || high > 32 {
		return errors.New("Target High Temperature must be in the range of 9 - 32")
	}

	if low < 9 || low > 32 {
		return errors.New("Target Low Temperature must be in the range of 9 - 32")
	}

	scale, err := n.GetTemperatureScale(deviceID)
	if err != nil {
		return err
	}

	if scale != "C" {
		return errors.New("Temperature Scale must be set to C")
	}

	vals := make(map[string]interface{})
	vals["target_temperature_high_c"] = high
	vals["target_temperature_low_c"] = low

	return n.setValue("thermostats", deviceID, vals)
}

// SetHVACMode changes the HVAC mode of the specified thermostat
func (n *Connection) SetHVACMode(deviceID, mode string) error {
	// Error checking
	mode = strings.Trim(mode, " ")
	validVals := []string{"heat", "cool", "heat-cool", "eco", "off"}

	if mode == "" {
		return errors.New("HVAC Mode must not be empty")
	}

	valid := false
	for _, v := range validVals {
		if mode == v {
			valid = true
		}
	}

	if !valid {
		return fmt.Errorf("HVAC Mode must be one of the following: %s", validVals)
	}

	vals := make(map[string]interface{})
	vals["hvac_mode"] = mode

	return n.setValue("thermostats", deviceID, vals)
}

// SetThermostatLabel - label
func (n *Connection) SetThermostatLabel(deviceID, label string) error {
	// Error checking
	if strings.Trim(label, " ") == "" {
		return errors.New("Label must not be empty")
	}

	vals := make(map[string]interface{})
	vals["label"] = label

	return n.setValue("thermostats", deviceID, vals)
}

// TurnOnFanTimer turns on the fan timer and sets the duration of the specified thermostat
func (n *Connection) TurnOnFanTimer(deviceID string, duration int) error {
	// Error checking
	validVals := []int{15, 30, 45, 60, 120, 240, 480, 720}

	valid := false
	for _, v := range validVals {
		if duration == v {
			valid = true
		}
	}

	if !valid {
		return fmt.Errorf("Fan Timer Duration must be one of the following: %d", validVals)
	}

	vals := make(map[string]interface{})
	vals["fan_timer_active"] = true
	vals["fan_timer_duration"] = duration

	return n.setValue("thermostats", deviceID, vals)
}

// TurnOffFanTimer turns off the fant timer of the specified thermostat
func (n *Connection) TurnOffFanTimer(deviceID string) error {
	vals := make(map[string]interface{})
	vals["fan_timer_active"] = false

	return n.setValue("thermostats", deviceID, vals)
}
