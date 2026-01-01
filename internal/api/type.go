package api

type Request struct {
	URL string `json:"url"`
}

type Response struct {
	EncryptedLink string `json:"encrypted_link"`
}
