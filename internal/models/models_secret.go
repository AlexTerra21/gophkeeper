package models

import (
	"bytes"
	"encoding/gob"
)

const (
	SecretTypePassword SecretType = iota + 1
	SecretTypeCard
	SecretTypeText
)

type SecretType int

type Secret interface {
	ToBinary() ([]byte, error)
}

type PasswordSecret struct {
	Login    string
	Password string
}

func (p *PasswordSecret) ToBinary() ([]byte, error) {
	var buff bytes.Buffer

	enc := gob.NewEncoder(&buff)
	err := enc.Encode(p)

	return buff.Bytes(), err
}

type CardSecret struct {
	Number     string
	HolderName string
	CCV        string
	Date       string
}

func (c *CardSecret) ToBinary() ([]byte, error) {
	var buff bytes.Buffer

	enc := gob.NewEncoder(&buff)
	err := enc.Encode(c)

	return buff.Bytes(), err
}

type TextSecret struct {
	Text string
}

func (t *TextSecret) ToBinary() ([]byte, error) {
	var buff bytes.Buffer

	enc := gob.NewEncoder(&buff)
	err := enc.Encode(t)

	return buff.Bytes(), err
}
