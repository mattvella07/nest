package nest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

type errorResponse struct {
	Error    string `json:"error"`
	Type     string `json:"type"`
	Message  string `json:"message"`
	Instance string `json:"instance"`
}

func (n *Connection) setURL(endpoint string) string {
	url := fmt.Sprintf("%s/%s", BaseURL, endpoint)

	// Base URL to use for tests
	if n.testURL != "" {
		url = fmt.Sprintf("%s/%s", n.testURL, endpoint)
	}

	return url
}

func (n *Connection) execute(url, method string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return []byte{}, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", n.AccessToken))

	// Need to create a custom client because defualt http client
	// doesn't forward headers when a redirect 3xx is received
	client := &http.Client{
		CheckRedirect: func(redirRequest *http.Request, via []*http.Request) error {
			redirRequest.Header = req.Header

			if len(via) >= 10 {
				return errors.New("stopped after 10 redirects")
			}
			return nil
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	// Check for errors
	if resp.StatusCode != 200 {
		errMsg := errorResponse{}

		err = json.Unmarshal(data, &errMsg)
		if err != nil {
			return []byte{}, fmt.Errorf("Error: %s", string(data))
		}

		return []byte{}, fmt.Errorf("Error: %s", errMsg.Message)
	}

	return data, nil
}

func (n *Connection) getValue(deviceType, deviceID, field string) (string, error) {
	// Error checking
	if strings.Trim(deviceID, " ") == "" {
		if deviceType == "structures" {
			return "", errors.New("Structure ID must not be empty")
		}

		return "", errors.New("Device ID must not be empty")
	}

	url := n.setURL(fmt.Sprintf("%s/%s/%s", deviceType, deviceID, field))

	data, err := n.execute(url, "GET", nil)
	if err != nil {
		return "", err
	}

	// Thermostat field not found
	if len(data) == 0 {
		return "", fmt.Errorf("%s %s not found", n.toTitleCase(deviceType), field)
	}

	return string(data), nil
}

func (n *Connection) setValue(deviceType, deviceID string, vals map[string]interface{}) error {
	// Error checking
	if strings.Trim(deviceID, " ") == "" {
		return errors.New("Device ID must not be empty")
	}

	url := n.setURL(fmt.Sprintf("%s/%s", deviceType, deviceID))

	body := strings.NewReader(n.formatMap(vals))

	data, err := n.execute(url, "PUT", body)
	if err != nil {
		return err
	}

	// Thermostat field not found
	if len(data) == 0 {
		return fmt.Errorf("Thermostat %s not found", deviceID)
	}

	return nil
}

// formatMap formats a map of string keys and interface{} values as a JSON string
func (n *Connection) formatMap(vals map[string]interface{}) string {
	str := "{"

	for key, val := range vals {
		str += fmt.Sprintf("\"%s\":", key)

		dataType := reflect.ValueOf(val)
		switch dataType.Kind() {
		case reflect.String:
			str += fmt.Sprintf("\"%s\",", val)
		case reflect.Int, reflect.Float32:
			str += fmt.Sprintf("%d,", val)
		default:
			str += fmt.Sprintf("\"%v\",", val)
		}
	}

	// Remove last comma and add closing curly brace
	str = str[:len(str)-1]
	str += "}"

	return str
}

func (n *Connection) toTitleCase(str string) string {
	returnStr := ""

	// Replace _ with space and capitalize first letter of each word
	for _, s := range strings.Split(str, "_") {
		returnStr += strings.ToUpper(s[0:1]) + s[1:] + " "
	}

	// Replace trailing s and space
	returnStr = strings.Trim(returnStr, "s ")

	return returnStr
}
