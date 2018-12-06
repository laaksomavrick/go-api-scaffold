package auth

type Service struct {
	hmacSecret []byte
}

func NewService(hmacSecret []byte) *Service {
	return &Service{
		hmacSecret: hmacSecret,
	}
}
