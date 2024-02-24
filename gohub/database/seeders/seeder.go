package seeders

import "gohub/pkg/seed"

func Initializa() {
	seed.SetRunOrder([]string{
		"SeedUsersTable",
	})
}
