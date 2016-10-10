package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
)

type ReposCommand struct {
}

func (c *ReposCommand) Synopsis() string {
	return "Get repository name list"
}

func (c *ReposCommand) Help() string {
	return "Usage: todo repos"
}

func (c *ReposCommand) Run(args []string) int {
	ecrCli := ecr.New(session.New(), aws.NewConfig().WithRegion("ap-northeast-1"))
	fmt.Println(getAllRepoNames(ecrCli))
	return 0
}

func getAllRepoNames(ecrCli *ecr.ECR) ([]string, error) {
	resp, err := ecrCli.DescribeRepositories(&ecr.DescribeRepositoriesInput{})
	if err != nil {
		return []string{}, fmt.Errorf("getting ecr repos: %v", err)
	}

	repos := make([]string, 0, len(resp.Repositories))
	for _, repo := range resp.Repositories {
		repos = append(repos, *repo.RepositoryName)
	}
	return repos, nil
}
