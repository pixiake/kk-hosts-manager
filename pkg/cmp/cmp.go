package cmp

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	kubekeyv1alpha1 "github.com/pixiake/kk-hosts-manager/apis/kubekey/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type HostsManager interface {
	List() (kubekeyv1alpha1.AvailableHostList, error)
	Update(hostsAction io.Reader) error
}

type CMP struct {
	Name      string `json:"name,omitempty"`
	Endpoint  string `json:"apiAddress,omitempty"`
	ListAPI   *API   `json:"list,omitempty"`
	UpdateAPI *API   `json:"update,omitempty"`
}

type API struct {
	Method string `json:"method,omitempty"`
	Path   string `json:"path,omitempty"`
}

type CMPHosts struct {
	Status     bool    `json:"status,omitempty"`
	Message    string  `json:"message,omitempty"`
	Data       Servers `json:"data,omitempty"`
	Code       int     `json:"code,omitempty"`
	Attributes string  `json:"attributes,omitempty"`
}

type Servers struct {
	Servers []Host `json:"servers,omitempty"`
}

type Host struct {
	ID                 string `json:"id,omitempty"`
	ResourceID         string `json:"resourceId,omitempty"`
	ServerName         string `json:"serverName,omitempty"`
	Alias              string `json:"alias,omitempty"`
	AzoneID            string `json:"azoneId,omitempty"`
	ZONE               string `json:"zone,omitempty"`
	LabelName          string `json:"labelName,omitempty"`
	ResouceType        string `json:"resouceType,omitempty"`
	VirtType           string `json:"virtType,omitempty"`
	Compute            string `json:"compute,omitempty"`
	Description        string `json:"description,omitempty"`
	CPU                int    `json:"cpu,omitempty"`
	Memory             int    `json:"memory,omitempty"`
	Storage            int    `json:"storage,omitempty"`
	OSUserName         string `json:"osUserName,omitempty"`
	Password           string `json:"password,omitempty"`
	PasswordType       string `json:"passwordType,omitempty"`
	StorageType        string `json:"storageType,omitempty"`
	Network            string `json:"network,omitempty"`
	IPAddr             string `json:"ipAddr,omitempty"`
	NetName            string `json:"netName,omitempty"`
	NetworkID          string `json:"networkId,omitempty"`
	OSName             string `json:"osName,omitempty"`
	OSType             string `json:"osType,omitempty"`
	ImageID            string `json:"imageId,omitempty"`
	ImageName          string `json:"imageName,omitempty"`
	MinCPU             string `json:"minCpu,omitempty"`
	MinRAM             string `json:"minRam,omitempty"`
	MinDisk            string `json:"minDisks,omitempty"`
	FlavorID           string `json:"flavorId,omitempty"`
	FlavorInfo         string `json:"flavorInfo,omitempty"`
	ResTenancy         string `json:"resTenancy,omitempty"`
	Status             string `json:"status,omitempty"`
	TenantID           string `json:"tenantId,omitempty"`
	UserID             string `json:"userId,omitempty"`
	UserName           string `json:"userName,omitempty"`
	CreateDate         string `json:"createDate,omitempty"`
	EndDate            string `json:"endDate,omitempty"`
	Host               string `json:"host,omitempty"`
	ServerNum          string `json:"serverNum,omitempty"`
	ProjectName        string `json:"projectName,omitempty"`
	DiskStorage        string `json:"diskStorage,omitempty"`
	SoftwareSystemName string `json:"softwareSystemName,omitempty"`
	SoftwareSystemIDs  string `json:"softwareSystemId,omitempty"`
	Cvk                string `json:"cvk,omitempty"`
	VmxVersion         string `json:"vmxVersion,omitempty"`
	ProcessStatus      string `json:"processStatus,omitempty"`
	ExistTag           bool   `json:"existTag,omitempty"`
}

type CMPHostsAction struct {
	Servers []string `json:"servers,omitempty"`
	Action  string   `json:"action,omitempty"`
}

func (c CMP) List() (kubekeyv1alpha1.AvailableHostList, error) {
	availableHostsList := kubekeyv1alpha1.AvailableHostList{}
	req, err := http.NewRequest(http.MethodGet, c.Endpoint+c.ListAPI.Path, nil)
	if err != nil {
		return availableHostsList, err
	}

	req.Header.Add("Content-Type", "application/json")

	r, err := http.DefaultClient.Do(req)
	if err != nil {
		return availableHostsList, err
	}

	defer r.Body.Close()

	rbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return availableHostsList, err
	}

	hostsList := CMPHosts{}

	err = json.Unmarshal(rbody, &hostsList)
	if err != nil {
		return availableHostsList, err
	}

	for _, host := range hostsList.Data.Servers {
		var (
			user   string = "root"
			osName string = "linux"
			port   int    = 22
		)
		if host.OSUserName != "" {
			user = host.OSUserName
		}
		if host.ImageName != "" {
			osName = host.ImageName
		}
		availableHostsList.Items = append(availableHostsList.Items, kubekeyv1alpha1.AvailableHost{
			ObjectMeta: metav1.ObjectMeta{
				Name: host.ServerName,
			},
			Spec: kubekeyv1alpha1.AvailableHostSpec{
				ID:              host.ID,
				Zone:            host.ZONE,
				Address:         host.IPAddr,
				InternalAddress: host.IPAddr,
				User:            user,
				Port:            port,
				Password:        host.Password,
				OSName:          osName,
				CPU:             host.CPU,
				Memory:          host.Memory,
				Storage:         host.Storage,
				ARCH:            "amd64",
			},
		})
	}

	return availableHostsList, nil
}

func (c CMP) Update(hostsAction io.Reader) error {

	req, err := http.NewRequest(c.UpdateAPI.Method, c.Endpoint+c.UpdateAPI.Path, hostsAction)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	fmt.Printf("%+v", res)
	return nil
}
