package helpers

import "fmt"

func GenerateProtocolIDRequest(namespace string, version string) string {
	return fmt.Sprintf("/%s/request/%s", namespace, version)

}

func GenerateProtocolIDResponse(namespace string, version string) string {
	return fmt.Sprintf("/%s/response/%s", namespace, version)
}
