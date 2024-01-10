package config

import (
	"log"
	"os"
)

var Logo = `
    ____             __        ______                              ___ ____ 
   / __ \_________  / /_____  / ____/___  ____ ___  ____     _   _<  // __ \
  / /_/ / ___/ __ \/ __/ __ \/ /   / __ \/ __ '__ \/ __ \   | | / / // / / /
 / ____/ /  / /_/ / /_/ /_/ / /___/ /_/ / / / / / / /_/ /   | |/ / // /_/ / 
/_/   /_/   \____/\__/\____/\____/\____/_/ /_/ /_/ .___/    |___/_(_)____/  
                                                /_/                         
`

var CompFilePath string
var CompFileName string
var DependencyProtos []string

var InfoLog = log.New(os.Stdout, "[X] INFO\t", log.LstdFlags)
var WarnLog = log.New(os.Stdout, "[X] WARN\t", log.LstdFlags)
var ErrorLog = log.New(os.Stdout, "[X] ERROR\t", log.LstdFlags)
