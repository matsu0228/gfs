package main

import (
	"context"
	"flag"
	"log"
	"os"
)

var crtFile, firestorePath string

func init() {
	flag.StringVar(&crtFile, "certFile", "secrets/spolive-dev.json", "set secrets json for firestore")
	flag.StringVar(&firestorePath, "p", "sports/rugby_topleague/games/4ayoQUGbmNAC1l9lITjB", "set firetstore path(containts collection's path and documentID")
}

func errorCheck(err error) {
	if err != nil {
		errorExit(err)
	}
}

func errorExit(err error) {
	log.Fatalf("[ERROR] %v", err)
}

func main() {

	flag.Parse()
	ctx := context.Background() //TODO: 外から設定?
	repo, err := NewFirebase(ctx, crtFile)
	errorCheck(err)
	log.Printf("[DEBUG] repo %#v", repo)

	outStream := os.Stdout
	err = repo.ToStruct(ctx, firestorePath, outStream)
	errorCheck(err)
}
