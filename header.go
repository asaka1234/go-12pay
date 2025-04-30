package go_12pay

func getHeaders() map[string]string {
	return map[string]string{
		"Content-Type": "application/jason",
	}
}

// notice:  input params is fixed
func getAuthHeaders(auth string, channel string, device string, partnerCode string) map[string]string {
	return map[string]string{
		"Content-Type":  "application/jason",
		"channel":       channel, //"WEB",
		"partner_code":  partnerCode,
		"authorization": auth,
		"device":        device, //WEB  option
	}
}
