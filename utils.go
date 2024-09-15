package main

func isValidName(name string) bool {
	return name != ""
}

func userExists(name string) bool {
	for _, u := range users {
		if u.name == name {
			return true
		}
	}

	return false

}
