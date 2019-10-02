package transloadit

import (
	"context"
	"fmt"

	transloadit "gopkg.in/transloadit/go-sdk.v1"
)

// ResizeImage would resize your imgpath with your verdict setting.
func ResizeImage(imgpath string, verdict string) (url string, err error) {
	// Create client
	options := transloadit.DefaultConfig
	options.AuthKey = "b211384720824c36ae7e844f4d2cb0c1"
	options.AuthSecret = "2c43615948c97542356e030685a451a95b3e09d0"

	client := transloadit.NewClient(options)

	assembly := transloadit.NewAssembly()

	assembly.AddFile("image", imgpath)
	assembly.AddStep("resize", verdictMap[verdict])

	info, err := client.StartAssembly(context.Background(), assembly)
	if err != nil {
		return "", err
	}

	info, err = client.WaitForAssembly(context.Background(), info)
	if err != nil {
		return "", err
	}

	fmt.Printf("You can view the result at %s\n", info.Results["resize"][0].SSLURL)
	return info.Results["resize"][0].SSLURL, nil
}
