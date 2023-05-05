package main

import (
	"fmt"
	"github.com/mburbidg/goplug/ops"
	"log"
	"os"
	"path/filepath"
	"plugin"
)

func main() {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatalf("Error: getting working directory: err=%s\n", err)
	}
	opList, err := loadPlugins(filepath.Dir(exePath) + "/plugins")
	if err != nil {
		log.Fatalf("Error: loading plugins: err=%s\n", err)
	}
	for name, op := range opList {
		log.Printf("%s=%s\n", name, op())
	}
}

func loadPlugins(pluginsDir string) (ops.OpList, error) {
	log.Printf("Info: loading plugins: dir=%s", pluginsDir)
	entries, err := os.ReadDir(pluginsDir)
	if err != nil {
		return nil, fmt.Errorf("reading plugin directory: err=%s\n", err)
	}
	opList := ops.OpList{}
	for _, entry := range entries {
		err := loadPlugin(opList, pluginsDir, entry)
		if err != nil {
			log.Printf("Warning: loading pluging: name=%s, err=%s\n", entry.Name(), err)
		}
	}
	return opList, nil
}

func loadPlugin(opList ops.OpList, dir string, entry os.DirEntry) error {
	log.Printf("Info: loading plugin: name=%s\n", entry.Name())
	p, err := plugin.Open(dir + "/" + entry.Name())
	if err != nil {
		return err
	}
	list, err := p.Lookup("OperationList")
	if err != nil {
		return fmt.Errorf("loading plugin: %w", err)
	}
	for name, op := range *list.(*ops.OpList) {
		opList[name] = op
	}
	return nil
}
