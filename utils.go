package main

func isValidName(name string) bool {
	return name != ""
}

func userExists(name string) (bool, error) {
	users, err := userRepository.GetAllUsers()

	if err != nil {
		return false, err
	}

	for _, u := range users {
		if u.Name == name {
			return true, nil
		}
	}

	return false, nil

}
