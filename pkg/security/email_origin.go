package security

import "strings"

const gmailHost = "gmail.com"

// EmailOrigin is a function to get origin of gmail account.
// For example, email trisna.x2+github@gmail.com
// The origin is trisnax2@gmail.com
func EmailOrigin(email string) (string, error) {
	splitEmail := strings.Split(email, "@")
	emailHost := splitEmail[len(splitEmail)-1]
	splitEmail = strings.Split(splitEmail[0], "+")
	emailName := splitEmail[0]

	if emailHost == gmailHost {
		return strings.ReplaceAll(emailName, ".", "") + "@" + emailHost, nil
	}

	return email, nil
}
