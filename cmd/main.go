package main

import (
	"crypt/pkg/env"
	"flag"
	"fmt"
	"os"
)

func main() {
	path := flag.String("o", ".env", "Path to .env file")
	enc := flag.Bool("e", false, "Encode the .env file")
	dec := flag.Bool("d", false, "Decode the .env file")
	flag.Parse()

	envs := env.ReadEnv(*path)

	if !*enc && !*dec {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *enc && *dec {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *enc {
		all := env.EncodeEnv(envs)
		env.SaveEnv(all, *path)
		fmt.Println("Envs encoded in: " + *path)
	}

	if *dec {
		all := env.DecodeEnv(envs)
		env.SaveEnv(all, *path+".dec")
		fmt.Println("Envs decoded in: " + *path + ".dec")
	}
}
