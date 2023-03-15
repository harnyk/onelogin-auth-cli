package utils

import (
	"fmt"
	"log"
	"strconv"

	"github.com/manifoldco/promptui"
)

func PromptForInt(label string) (*int, error) {

	validate := func(input string) error {
		//ToDo validate
		return nil
	}

	prompt := promptui.Prompt{
		Label:    label,
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		log.Printf("Prompt failed %v\n", err)
		return nil, err
	}
	numResult, err := strconv.Atoi(result)
	if err != nil {
		return nil, err
	}

	return &numResult, nil
}

func PromptForString(label string) (string, error) {

	validate := func(input string) error {
		if input == "" {
			return fmt.Errorf("Error: Empty input")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    label,
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		log.Printf("Prompt failed %v\n", err)
		return "", err
	}

	return result, nil
}

func PromptForSecretString(label string) (string, error) {

	validate := func(input string) error {
		if input == "" {
			return fmt.Errorf("Error: Empty input")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    label,
		Validate: validate,
		Mask:     rune('*'),
	}

	result, err := prompt.Run()

	if err != nil {
		log.Printf("Prompt failed %v\n", err)
		return "", err
	}

	return result, nil
}

func PromptSelect(label string, options []string, skipSingleChoice bool) (string, error) {

	if skipSingleChoice && len(options) == 1 {
		return options[0], nil
	}

	prompt := promptui.Select{
		Label: label,
		Items: options,
	}

	_, result, err := prompt.Run()

	if err != nil {
		log.Printf("Prompt failed %v\n", err)
		return "", err
	}

	return result, nil
}
