package commands

import (
	"errors"
	"fmt"

	"github.com/manifoldco/promptui"
)

// Вывод приглашения для ввода строки с текстом
func getStringData(label string) string {
	prompt := promptui.Prompt{
		Label:       label,
		HideEntered: true,
	}
	result, err := prompt.Run()
	if err != nil {
		fmt.Println(err)
	}
	return result
}

// Вывод приглашения для ввода пароля с маскированием.
func getPasswordData(label string) string {
	validate := func(input string) error {
		if len(input) < 1 {
			return errors.New("password must have more than 0 characters")
		}
		return nil
	}
	prompt := promptui.Prompt{
		Label:       label,
		HideEntered: true,
		Validate:    validate,
		Mask:        '*',
	}
	result, err := prompt.Run()
	if err != nil {
		fmt.Println(err)
	}
	return result
}
