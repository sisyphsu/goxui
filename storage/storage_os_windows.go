package storage

import (
	"errors"
	"github.com/boltdb/bolt"
	"os"
	"path/filepath"
)

// 初始化存储环境, 根据指定的厂商、应用名来创建Support、Cache目录
func Init(vendor string, appName string) (err error) {
	// find support dir
	var globalSettingFolder = os.Getenv("APPDATA")
	SupportDir = filepath.Join(globalSettingFolder, vendor, appName)
	if err = os.MkdirAll(SupportDir, 0755); err != nil {
		return errors.New("Can't create SupportDir: " + err.Error())
	}

	// find cache dir
	var cacheFolder = os.Getenv("LOCALAPPDATA")
	if cacheFolder != "" && cacheFolder != "%LOCALAPPDATA%" {
		CacheDir = filepath.Join(cacheFolder, vendor, appName) // win7
	} else {
		CacheDir = filepath.Join(SupportDir, "Temp") // xp
	}
	if err = os.MkdirAll(CacheDir, 0755); err != nil {
		return errors.New("Can't create CacheDir: " + err.Error())
	}

	// prepare storage
	dbFile := filepath.Join(SupportDir, appName+".db")
	if storage, err = bolt.Open(dbFile, 0600, nil); err != nil {
		return errors.New("Can't open db file: " + err.Error())
	}
	return
}