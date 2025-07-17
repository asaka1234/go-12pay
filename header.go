package go_12pay

func getHeaders() map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
	}
}

// notice:  input params is fixed
func getAuthHeaders(partnerCode string, auth string, channel string, device string) map[string]string {
	return map[string]string{
		"Content-Type":  "application/json",
		"channel":       channel, //"WEB",
		"partner_code":  partnerCode,
		"Authorization": auth,
		"device":        device, //WEB  option
	}
}

func getPayoutAuthHeaders(partnerCode string, auth string, channel string) map[string]string {
	return map[string]string{
		"Content-Type":  "application/json",
		"Channel":       channel,     //"WEB",
		"Partnercode":   partnerCode, //名字不一样
		"Authorization": auth,
		//"device":        device, //WEB  option
	}
}
