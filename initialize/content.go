package initialize

import (
	"gin-basics/global"
	"github.com/docker/docker/client"
	"os"
	"runtime"
)

func InitContent(t string) {
	var err error
	if t == "local" {
		url := "unix:///var/run/docker.sock"
		if runtime.GOOS == "windows" {
			url = "npipe:////./pipe/docker_engine"
		}

		global.GB_CIL, err = client.NewClientWithOpts(
			client.FromEnv,
			client.WithAPIVersionNegotiation(),
			client.WithHost(url),
		)

		if err != nil {
			global.GB_LOG.Error(err.Error())
			os.Exit(-1)
		}
	} else {
		url := "unix:///var/run/docker.sock"
		if runtime.GOOS == "windows" {
			url = "npipe:////./pipe/docker_engine"
		}

		global.GB_CIL, err = client.NewClientWithOpts(
			client.FromEnv,
			client.WithAPIVersionNegotiation(),
			client.WithHost(url),
		)

		if err != nil {
			global.GB_LOG.Error(err.Error())
			os.Exit(-1)
		}
	}

	//_ = http.Client{
	//	Transport: &http.Transport{
	//		DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
	//			if runtime.GOOS == "darwin" || runtime.GOOS == "linux" {
	//				return net.Dial("unix", "/var/run/docker.sock")
	//			} else {
	//				return net.Dial("unix", "/var/run/docker.sock")
	//			}
	//		},
	//	},
	//}
}
