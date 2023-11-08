package auth

type TokenStruct struct {
	Token string `json:"token"`
}

func TokenFormater(data string) TokenStruct {
	return TokenStruct{Token: data}
}
