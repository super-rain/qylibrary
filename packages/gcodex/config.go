package gcodex

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type IDLConfig struct {
	DataBase *DBMetadata `json:"db_metadata"`
}

type DBMetadata struct {
	DataBase *DBConfig `json:"db"`
}

type DBConfig struct {
	Name   string           `json:"name"`
	Tables []*TableMetadata `json:"tables"`
}

type TableMetadata struct {
	Name    string   `json:"name"`
	Comment string   `json:"comment"`
	Columns []string `json:"columns"`
}

var ViperCfg *viper.Viper

func LoadIDLConfig(cfgPath, name string) {
	if name == "" {
		log.Fatalln("config name is empty.")
	}
	if cfgPath == "" {
		cfgPath = "."
		// return nil,Err_IDL_File_Not_Exists
	}

	fmt.Println("loading config begin...")

	ViperCfg = viper.New()
	ViperCfg.SetConfigName(name)
	// ViperCfg.SetConfigType("toml")
	ViperCfg.AddConfigPath(cfgPath)
	err := ViperCfg.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}
	allkeys := ViperCfg.AllKeys()
	fmt.Println(allkeys)
	// vv := ViperCfg.AllSettings()
	// fmt.Println(vv)
	db := ViperCfg.GetStringMap("db")
	dbname := db["name"]
	comment := db["name"]
	tables := ViperCfg.GetStringMap("db.tables")
	cols := ViperCfg.GetStringMapStringSlice("db.tables.columns")
	tname := tables["name"]
	tcomment := tables["name"]
	fmt.Printf("db:%v,name:%v,comment:%v,tables:%v,tcomment:%v,tablesS:%v,cols:%v", db, dbname, comment, tables, tname, tcomment, cols)

	var idl IDLConfig
	err = ViperCfg.Unmarshal(&idl)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	keys := ViperCfg.AllKeys()
	val := ViperCfg.AllSettings()
	fmt.Printf("keys:%s\n", keys)
	fmt.Printf("val:%s\n", val)
	fmt.Println("load cofing success.")

}
