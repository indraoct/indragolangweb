/**
 * Updated Date : 14 January 2017
 * Import os to get pwd path data
 */
package config
import ("os")

var PathView string

func GetPath(){
	pwd,_ := os.Getwd()
	PathView = pwd+"/apps/views"
}

