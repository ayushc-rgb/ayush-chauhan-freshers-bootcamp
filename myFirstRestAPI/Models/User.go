package Models

import "firstAPI/Config"

// this is a very important file because it handles all the connection with database and have the actual logic

func GetAllUsers(user *[]User) error { // get all the user from the database using the gorm variable
	if err := Config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

func CreateUser(user *User) error { // to create a new user in the database
	if err := Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByID(user *User, id string) error { // this wil get the user with the given ID
	if err := Config.DB.Where("id=?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *User, id string) error { // to update the data of the user
	Config.DB.Save(user)
	return nil
}

func DeleteUser(user *User, id string) error { // to soft delete a user from the database
	if err := Config.DB.Where("id=?", id).Delete(user).Error; err != nil {
		return err
	}
	return nil
}
