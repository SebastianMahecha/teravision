package config

import (	
	"os"	
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" 
	
)

var DB *gorm.DB
var err error

func InitDB(){		
	DB, err = gorm.Open("postgres", "host="+os.Getenv("DBHOST")+" user="+os.Getenv("USERNAME")+" dbname="+os.Getenv("DBNAME")+" password="+os.Getenv("USERPASS")+" sslmode=disable")
	if err != nil {
		panic("failed to connect database")		
		
	}	
}