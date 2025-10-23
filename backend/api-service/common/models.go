package common

type Config struct {
	JwtSecretKey []byte
	PsqlHost     string
	PsqlPort     int
	PsqlUser     string
	PsqlPassword string
	PsqlDatabase string
}
