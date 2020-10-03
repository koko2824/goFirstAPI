package db

import (
	"context"
	"database/sql"
	"github.com/volatiletech/sqlboiler/boil"
	_ "github.com/volatiletech/sqlboiler/drivers/sqlboiler-mysql/driver"
	"test_members2_rest/config"
	"test_members2_rest/models"
)

//Init
func Init()  {
	db, err := sql.Open("mysql", config.DBUser())
	if err != nil {
		panic(err)
	}
	boil.SetDB(db)
}

func GetAll() models.UserSlice {
	db, err := sql.Open("mysql", config.DBUser())
	if err != nil {
		panic(err)
	}

	users, err := models.Users().All(context.Background(), db)
	if err != nil {
		panic(err)
	}

	return users
}

func GetOne(id int) models.User {
	db, err := sql.Open("mysql", config.DBUser())
	if err != nil {
		panic(err)
	}

	user, err := models.FindUser(context.Background(), db, id)
	if err != nil {
		panic(err)
	}

	return *user
}

func Insert(name string, role string)  {
	db, err := sql.Open("mysql", config.DBUser())
	if err != nil {
		panic(err)
	}

	user := models.User{
		Name: name,
		Role: role,
	}

	user.Insert(context.Background(), db, boil.Infer())
}

func Update(id int, name string, role string)  {
	
	db, err := sql.Open("mysql", config.DBUser())
	if err != nil {
		panic(err)
	}

	user := new(models.User)
	user.ID = id
	user.Name = name
	user.Role = role

	_, err = user.Update(context.Background(), db, boil.Infer())
	if err != nil {
		panic(err)
	}
}

func Delete(id int)  {
	
	db, err := sql.Open("mysql", config.DBUser())
	if err != nil {
		panic(err)
	}

	user := models.User{ID: id}
	_, err = user.Delete(context.Background(), db)
	if err != nil{
		panic(err)
	}
}