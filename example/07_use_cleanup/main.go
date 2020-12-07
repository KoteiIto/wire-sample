package main

import "github.com/KoteiIto/wire-sample/example/07_use_cleanup/di"

func main() {
	_, cleanup := di.InitializeAuthService()
	cleanup()
}
