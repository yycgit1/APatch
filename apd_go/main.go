package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func printbanner() {
	banner := `
    _    ____       _       _     
   / \  |  _ \ __ _| |_ ___| |__  
  / _ \ | |_) / _` + "`" + ` | __/ __| '_ \ 
 / ___ \|  __/ (_| | || (__| | | |
/_/   \_\_|   \__,_|\__\___|_| |_|
   `
	fmt.Println(banner)
}
func main() {

	if len(os.Args) < 3 {
		fmt.Println("Usage: apd module install <module-name>")
		return
	}

	if os.Args[1] == "module" {
		if os.Args[2] == "test" {
			//install_modules(os.Args[3])
			test := os.Args[3]
			fmt.Printf("test function: %s\n", test)

			return
		}
		if os.Args[2] == "install" {
			//install_modules(os.Args[3])
			modulepath := os.Args[3]
			fmt.Printf("Installing module: %s\n", modulepath)
			installModule(modulepath)
			return
		}
		if os.Args[2] == "list" {
			//install_modules(os.Args[3])

			modules, err := listModules()
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}
			// 将模块转换为JSON格式
			jsonOutput, err := json.MarshalIndent(modules, "", "  ")
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}
			fmt.Println(string(jsonOutput))
			return
		}
		if os.Args[2] == "enable" {
			if err := enableModule(os.Args[3], true); err != nil {
				fmt.Printf("Error: %v\n", err)
			}
			return
		}
		if os.Args[2] == "disable" {
			if err := enableModule(os.Args[3], false); err != nil {
				fmt.Printf("Error: %v\n", err)
			}
			return
		}
		if os.Args[2] == "disable_all_modules" {
			if err := disableAllModulesUpdate(); err != nil {
				fmt.Printf("Error: %v\n", err)
			}
			return
		}

		fmt.Println("Usage: apd module install <module-name>")
		return
	}
	if os.Args[1] == "post-fs-data" { //Trigger `post-fs-data` event
		on_postdata_fs(os.Args[2])
	}
	if os.Args[1] == "services" { //Trigger `services` event
		on_services(os.Args[2])
	}
	if os.Args[1] == "boot-completed" { //Trigger `boot-completed` event
		on_boot_completed(os.Args[2])
	}
	if os.Args[1] == "getprop" {
		value, err := getprop(os.Args[2])
		//value, err := getprop("vendor.post_boot.parsed")
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Printf("%s: %s\n", os.Args[2], value)

	}

}
