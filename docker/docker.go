package docker

import d "github.com/fsouza/go-dockerclient"

func NewDockerClient() (*d.Client, error) {
	return d.NewClientFromEnv()
}
