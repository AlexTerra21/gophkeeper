package storage

type User struct {
	ID             int64
	Name           string `pg:",notnull,unique" json:"login"`
	Password       string `pg:"-" json:"password"`
	HashedPassword []byte `json:"-"`
	Salt           []byte `json:"-"`
}

type Secret struct {
	UserID     int64
	SecretData []byte `pg:",notnull"`
	SecretType int    `pg:",notnull"`
	SecretName string `pg:",notnull,unique"`
}
