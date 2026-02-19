package sistem

import (
	"os"
	"os/user"
)

func DeteksiUser() (string, string, error) {
	u, err := user.Current()
	if err != nil {
		return "user", "host", err
	}

	h, err := os.Hostname()
	if err != nil {
		return u.Username, "host", err
	}

	return u.Username, h, nil
}
