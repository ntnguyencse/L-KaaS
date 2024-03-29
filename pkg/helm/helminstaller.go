package helminstaller

import (
	"fmt"

	helmlib "github.com/ntnguyencse/helm-mod/cmd/helm"
)

func test() {
	kubePath := "/home/ubuntu/ntnguyen-helm/helm/test/kubeconfig/cluster-test"
	chartName2 := "chartname21"
	ciliumPath := "https://github.com/prometheus-community/helm-charts/releases/download/kube-prometheus-stack-46.7.0/kube-prometheus-stack-46.7.0.tgz"
	ciliumHelmArgs := []string{"install", chartName2, ciliumPath, "--kubeconfig", kubePath, "--debug", "--v", "5"}
	err := helmlib.ApplyHelmWrapper(kubePath, ciliumPath, true, true, ciliumHelmArgs, []string{})
	if err != nil {
		fmt.Println("error: 2", err)
	}
}
func Install(kubePath, chartName, chartPath, valueFilePath, namespace string) error {
	if len(valueFilePath) > 0 {
		installArgs := []string{"install", "-f", valueFilePath, chartName, chartPath, " --create-namespace", "--namespace", namespace, "--kubeconfig", kubePath, "--debug"}
		err := helmlib.ApplyHelmWrapper(kubePath, chartPath, true, false, installArgs, []string{})
		if err != nil {
			fmt.Println("Error when install application:", chartName, chartPath, kubePath, valueFilePath, err)
		}
		return err
	} else {
		installArgs := []string{"install", chartName, chartPath, "--create-namespace", "--namespace", namespace, "--kubeconfig", kubePath, "--debug"}
		err := helmlib.ApplyHelmWrapper(kubePath, chartPath, true, false, installArgs, []string{})
		if err != nil {
			fmt.Println("Error when install application:", chartName, chartPath, kubePath, err)
		}
		return err
	}

}
