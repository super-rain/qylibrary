package gcodex

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"text/template"

	"github.com/spf13/viper"
)

var (
	ViperCfg *viper.Viper
	IDL      *IDLConfig
)

func LoadIDLConfig(cfgPath, name string) error {
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
		return err
	}
	var idl IDLConfig
	err = ViperCfg.Unmarshal(&idl)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return err
	}

	IDL = &idl
	fmt.Println("load cofing success.")

	return nil
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

func createSrcFile(t *template.Template, idl *TableMetaData, fname, tmpTarget string) error {
	err := os.MkdirAll(tmpTarget, 0777)
	if err != nil && !os.IsExist(err) {
		log.Printf("make target file dir fail error:%v\n", err)
		return err
	}

	f, err := os.OpenFile(tmpTarget+fname, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Printf("create target file fail error:%v\n", err)
		return err
	}
	defer f.Close()

	ff := os.NewFile(f.Fd(), fname)
	err = t.Execute(ff, idl)
	if err != nil {
		log.Printf("tmpl exec fail,error:%v", err)
		return err
	}
	// 格式化
	runGoCMD("go", "fmt", tmpTarget+fname)
	// cmd := exec.Command("go", "fmt", fname)
	// cmd.Stdout = os.Stdout
	// cmd.Run()

	return nil
}

func runGoCMD(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}
