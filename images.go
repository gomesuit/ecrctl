package main

import (
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
	"os"
)

type ImagesCommand struct {
}

func (c *ImagesCommand) Synopsis() string {
	return "Get image list"
}

func (c *ImagesCommand) Help() string {
	return "Usage: images"
}

func (c *ImagesCommand) Run(args []string) int {
	var repository string
	f := flag.NewFlagSet("images", flag.ExitOnError)
	f.StringVar(&repository, "r", "", "repository name")
	f.Parse(args)

	if repository == "" {
		fmt.Println("repository name is require")
		os.Exit(1)
	}

	ecrCli := ecr.New(session.New(), aws.NewConfig().WithRegion("ap-northeast-1"))
	fmt.Println(getImages(ecrCli, repository))
	return 0
}

func getImages(ecrCli *ecr.ECR, repoName string) ([]*ecr.ImageIdentifier, error) {
	var (
		token    *string
		imageIDs = []*ecr.ImageIdentifier{}
	)
	for {
		resp, err := ecrCli.ListImages(&ecr.ListImagesInput{
			RepositoryName: aws.String(repoName),
			NextToken:      token,
		})
		if err != nil {
			return nil, fmt.Errorf("getting %v images: %v", repoName, err)
		}
		imageIDs = append(imageIDs, resp.ImageIds...)
		if token = resp.NextToken; token == nil {
			break
		}
	}
	return imageIDs, nil
}
