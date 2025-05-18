package auth

func extractSession() string { //private only to this packages

	return "Loggedin"
}

func GetSession() string {
	return extractSession()
}
