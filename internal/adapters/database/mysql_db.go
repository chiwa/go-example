package database

import "fmt"

type MysqlDatabase struct{}

func NewMysqlDatabaseAdapter() *MysqlDatabase {
	return &MysqlDatabase{}
}

func (db *MysqlDatabase) SaveInfo(message string) error {
	fmt.Println("Save to Mysql : ", message)
	return nil
}
