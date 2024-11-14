package model

func ValidateCard(id string) (bool, string) {
	if id == "" {
		return false, "missing id in request"
	}

	return true, ""
}
