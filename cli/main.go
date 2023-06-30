package main

import (
	"envdost/cmd"
	_ "envdost/cmd/delete"
)

func main() {
	cmd.Execute()




	// for reference
	// key, err := registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlSet\Control\Session Manager\Environment`, registry.ALL_ACCESS)
	// if err != nil {
	// 	fmt.Println("Failed to open registry key:", err)
	// 	return
	// }
	// defer key.Close()

	// // // Set the environment variable
	// // err = key.SetStringValue("MY_VARIABLE", "my_value")
	// // if err != nil {
	// // 	fmt.Println("Failed to set environment variable:", err)
	// // 	return
	// // }

	// out, _,err := key.GetStringValue("MY_VARIABLE")
	
	// fmt.Println(out)


	 
}