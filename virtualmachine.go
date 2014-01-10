package gopherstack

import (
	"net/url"
	"strings"
)

// Deploys a Virtual Machine and returns it's id
func (c CloudStackClient) DeployVirtualMachine(serviceofferingid string, templateid string, zoneid string, networkids []string, keypair string, displayname string, diskoffering string, projectid string) (string, string, error) {
	params := url.Values{}
	params.Set("serviceofferingid", serviceofferingid)
	params.Set("templateid", templateid)
	params.Set("zoneid", zoneid)
	params.Set("networkids", strings.Join(networkids, ","))
	params.Set("keypair", keypair)
	params.Set("displayname", displayname)
	if projectid != "" {
		params.Set("projectid", projectid)
	}
	if diskoffering != "" {
		params.Set("diskoffering", diskoffering)
	}
	response, err := NewRequest(c, "deployVirtualMachine", params)
	if err != nil {
		return "", "", err
	}
	vmid := response.(DeployVirtualMachineResponse).Deployvirtualmachineresponse.ID
	jobid := response.(DeployVirtualMachineResponse).Deployvirtualmachineresponse.Jobid
	return vmid, jobid, nil
}

// Stops a Virtual Machine
func (c CloudStackClient) StopVirtualMachine(id string) (string, error) {
	params := url.Values{}
	params.Set("id", id)
	response, err := NewRequest(c, "stopVirtualMachine", params)
	if err != nil {
		return "", err
	}
	jobid := response.(StopVirtualMachineResponse).Stopvirtualmachineresponse.Jobid
	return jobid, err
}

// Destroys a Virtual Machine
func (c CloudStackClient) DestroyVirtualMachine(id string) (string, error) {
	params := url.Values{}
	params.Set("id", id)
	response, err := NewRequest(c, "destroyVirtualMachine", params)
	if err != nil {
		return "", err
	}
	jobid := response.(DestroyVirtualMachineResponse).Destroyvirtualmachineresponse.Jobid
	return jobid, nil
}

// Returns CloudStack string representation of the Virtual Machine state
func (c CloudStackClient) VirtualMachineState(id string) (string, string, error) {
	params := url.Values{}
	params.Set("id", id)
	response, err := NewRequest(c, "listVirtualMachines", params)
	if err != nil {
		return "", "", err
	}

	count := response.(ListVirtualMachinesResponse).Listvirtualmachinesresponse.Count
	if count != 1 {
		// TODO: Return some like no virtual machines found.
		return "", "", err
	}

	state := response.(ListVirtualMachinesResponse).Listvirtualmachinesresponse.Virtualmachine[0].State
	ip := response.(ListVirtualMachinesResponse).Listvirtualmachinesresponse.Virtualmachine[0].Nic[0].Ipaddress

	return ip, state, err
}


type DeployVirtualMachineResponse struct {
	Deployvirtualmachineresponse struct {
		ID    string `json:"id"`
		Jobid string `json:"jobid"`
	} `json:"deployvirtualmachineresponse"`
}

type DestroyVirtualMachineResponse struct {
	Destroyvirtualmachineresponse struct {
		Jobid string `json:"jobid"`
	} `json:"destroyvirtualmachineresponse"`
}

type StopVirtualMachineResponse struct {
	Stopvirtualmachineresponse struct {
		Jobid string `json:"jobid"`
	} `json:"stopvirtualmachineresponse"`
}

type Nic struct {
	Gateway     string `json:"gateway"`
	ID          string `json:"id"`
	Ipaddress   string `json:"ipaddress"`
	Isdefault   bool   `json:"isdefault"`
	Macaddress  string `json:"macaddress"`
	Netmask     string `json:"netmask"`
	Networkid   string `json:"networkid"`
	Traffictype string `json:"traffictype"`
	Type        string `json:"type"`
}

type Virtualmachine struct {
	Account             string        `json:"account"`
	Cpunumber           float64       `json:"cpunumber"`
	Cpuspeed            float64       `json:"cpuspeed"`
	Created             string        `json:"created"`
	Displayname         string        `json:"displayname"`
	Domain              string        `json:"domain"`
	Domainid            string        `json:"domainid"`
	Guestosid           string        `json:"guestosid"`
	Haenable            bool          `json:"haenable"`
	Hypervisor          string        `json:"hypervisor"`
	ID                  string        `json:"id"`
	Keypair             string        `json:"keypair"`
	Memory              float64       `json:"memory"`
	Name                string        `json:"name"`
	Nic                 []Nic         `json:"nic"`
	Passwordenabled     bool          `json:"passwordenabled"`
	Rootdeviceid        float64       `json:"rootdeviceid"`
	Rootdevicetype      string        `json:"rootdevicetype"`
	Securitygroup       []interface{} `json:"securitygroup"`
	Serviceofferingid   string        `json:"serviceofferingid"`
	Serviceofferingname string        `json:"serviceofferingname"`
	State               string        `json:"state"`
	Tags                []interface{} `json:"tags"`
	Templatedisplaytext string        `json:"templatedisplaytext"`
	Templateid          string        `json:"templateid"`
	Templatename        string        `json:"templatename"`
	Zoneid              string        `json:"zoneid"`
	Zonename            string        `json:"zonename"`
}

type ListVirtualMachinesResponse struct {
	Listvirtualmachinesresponse struct {
		Count          float64          `json:"count"`
		Virtualmachine []Virtualmachine `json:"virtualmachine"`
	} `json:"listvirtualmachinesresponse"`
}