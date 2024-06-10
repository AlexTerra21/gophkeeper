package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

// Криптограф шифрует и расшифровывает данные
type Cryptographer interface {
	Encrypt(src []byte) ([]byte, error)
	Decrypt(src []byte) ([]byte, error)
}

// GCMAESCryptographer - это криптограф, основанный на AES с режимом Galois/Counter.
//
// AES с режимом Galois/Counter (AES-GCM) обеспечивает как аутентифицированное шифрование (конфиденциальность, так и аутентификацию).
// и возможность проверять целостность и аутентификацию дополнительных
// аутентифицированных данных (AAD), которые отправляются в открытом виде.
type GCMAESCryptographer struct {
	Random Generator // генератор случайных чисел. Используется для генерации одноразового
	Key    []byte    // ключ, используемый для шифрования и расшифровки данных
}

// Encrypt шифрует открытый текст с помощью AES-GCM.
func (c *GCMAESCryptographer) Encrypt(plaintext []byte) ([]byte, error) {
	aesblock, err := aes.NewCipher(c.Key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		return nil, err
	}

	nonce, err := c.Random.GenerateRandomBytes(aesgcm.NonceSize())
	if err != nil {
		return nil, err
	}

	return aesgcm.Seal(nonce, nonce, plaintext, nil), nil
}

// Decrypt расшифровывает зашифрованный текст с помощью AES-GCM.
func (c *GCMAESCryptographer) Decrypt(ciphertext []byte) ([]byte, error) {
	aesblock, err := aes.NewCipher(c.Key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		return nil, err
	}

	nonceSize := aesgcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext is too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	return aesgcm.Open(nil, nonce, ciphertext, nil)
}
