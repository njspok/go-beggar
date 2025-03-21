package main

import "fmt"

func assetFilePath(name string) string {
	return fmt.Sprintf("assets/%s", name)
}
