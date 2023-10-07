package config

import "os"

var nameSpaceFilePath = "/var/run/secrets/kubernetes.io/serviceaccount/namespace"

func IsK8s() bool {
	_, err := os.ReadFile(nameSpaceFilePath)
	return err == nil
}

func K8sNameSpaces() string {
	nsRead, _ := os.ReadFile(nameSpaceFilePath)
	return string(nsRead)
}
