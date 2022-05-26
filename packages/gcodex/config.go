package gcodex

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var ViperCfg *viper.Viper

type Dtype struct {
	DB []*DD `json:"db"`
}
type DD struct {
	Name    string
	Comment string
	UserID  string `json:"user_id"`
	Tables  []interface{}
}

func LoadIDLConfig(cfgPath, name string) (*IDLConfig, error) {
	if name == "" {
		log.Fatal("metadata file  is not exists.")
	}
	if cfgPath == "" {
		cfgPath = "."
	}

	fmt.Println("loading config begin...")
	ViperCfg = viper.New()
	ViperCfg.SetConfigName(name)
	// ViperCfg.SetConfigType("toml")
	ViperCfg.AddConfigPath(cfgPath)
	err := ViperCfg.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
		return nil, err
	}
	var idl IDLConfig
	err = ViperCfg.Unmarshal(&idl)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	fmt.Println("load cofing success.")

	return &idl, nil
}

func InitIDLMetadata[T any](t *T) error {
	// if ViperCfg != nil {
	// 	return ErrIDLFileNotExists
	// }
	err := viper.Unmarshal(t)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return ErrInvalidIDLMetadata
	}

	return nil
}
