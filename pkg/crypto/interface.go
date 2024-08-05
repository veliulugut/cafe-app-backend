package crypto

type Crypto interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password string, hashedPassword string) bool
	GenerateOTP(length int) (string, error)
	GenerateToken(length int) (string, error)
}
