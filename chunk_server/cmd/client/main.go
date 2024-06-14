package main

import "github.com/ishanmadhav/skyfs/client"

func main() {
	cli := client.NewClient()
	cli.CreateBucket("Buck")
	//cli.PutObject("obj5", "Buck", "google-chrome-stable_current_amd64.deb")
	cli.GetObject("obj5", "Buck", "google-chrome-stable_current_amd64.deb")
}
