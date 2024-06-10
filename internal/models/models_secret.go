package models

import (
	"bytes"
	"encoding/gob"
)

// Типы хранимой информации
const (
	SecretTypePassword SecretType = iota + 1
	SecretTypeCard
	SecretTypeText
)

type SecretType int

// Пароль
type PasswordSecret struct {
	Login    string
	Password string
}

// Кодирование пароля
func (p *PasswordSecret) ToBinary() ([]byte, error) {
	var buff bytes.Buffer

	enc := gob.NewEncoder(&buff)
	err := enc.Encode(p)

	return buff.Bytes(), err
}

// Информация о банковской карте
type CardSecret struct {
	Number     string
	HolderName string
	CCV        string
	Date       string
}

// Кодирование информации о банковской карте
func (c *CardSecret) ToBinary() ([]byte, error) {
	var buff bytes.Buffer

	enc := gob.NewEncoder(&buff)
	err := enc.Encode(c)

	return buff.Bytes(), err
}

// Любая текстовая информация
type TextSecret struct {
	Text string
}

// Кодирование любой текстовой информации
func (t *TextSecret) ToBinary() ([]byte, error) {
	var buff bytes.Buffer

	enc := gob.NewEncoder(&buff)
	err := enc.Encode(t)

	return buff.Bytes(), err
}
