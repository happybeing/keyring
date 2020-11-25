package keyring

import "errors"

// PromptFunc is a function used to prompt the user for a password
type PromptFunc func(string) (string, error)

func terminalPrompt(prompt string) (string, error) {
	return "", errors.New("gitportal-stub: keyring terminalPrompt() not supported")
	// fmt.Printf("%s: ", prompt)
	// b, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	// if err != nil {
	// 	return "", err
	// }
	// fmt.Println()
	// return string(b), nil
}

func fixedStringPrompt(value string) PromptFunc {
	return func(_ string) (string, error) {
		return value, nil
	}
}
