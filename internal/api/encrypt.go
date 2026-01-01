package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

const encryptionURL = "https://crypto.happ.su/api.php"

func EncryptString(plaintext string) (string, error) {
	reqBody, err := json.Marshal(Request{URL: plaintext})
	if err != nil {
		return "", err
	}

	resp, err := http.Post(encryptionURL, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var apiResponse Response
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return "", err
	}

	return apiResponse.EncryptedLink, nil
}
