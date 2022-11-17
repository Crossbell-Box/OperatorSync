package utils

import (
	"fmt"
	"strings"
)

func SplitFediverseUsernameInstance(usernameAtInstance string) (string, string, error) {
	var (
		username string
		instance string
	)

	splits := strings.Split(usernameAtInstance, "@")

	if len(splits) != 2 {
		return "", "", fmt.Errorf("invalid format")
	}

	username = splits[0]
	instance = splits[1]

	return username, instance, nil
}
