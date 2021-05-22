/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	kubekeyv1alpha1 "github.com/pixiake/kk-hosts-manager/apis/kubekey/v1alpha1"
	"github.com/pixiake/kk-hosts-manager/pkg/cmp"
)

func main() {
	demoCMP := &cmp.CMP{
		Name:     "demo",
		Endpoint: "http://ae00a910-e1b6-49e5-9d14-1c9bd6d44a6e.mock.pstmn.io",
		ListAPI: &cmp.API{
			Method: http.MethodGet,
			Path:   "/qcloud/server/list",
		},
		UpdateAPI: &cmp.API{
			Method: http.MethodPut,
			Path:   "/qcloud/server/action",
		},
	}

	router := gin.Default()

	router.GET("/api/v1alpha1/hosts", func(c *gin.Context) {
		// page := c.DefaultQuery("page", "1")
		// page_size := c.DefaultQuery("page_size", "10")
		// sort := c.DefaultQuery("sort", "-create_at")
		availablehostsList, err := demoCMP.List()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, gin.H{"items": availablehostsList.Items, "totalItems": len(availablehostsList.Items)})
	})

	router.POST("/api/v1alpha1/hosts", func(c *gin.Context) {
		var hostsAction kubekeyv1alpha1.HostsAction

		if err := c.ShouldBindJSON(&hostsAction); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		fmt.Printf("%+v", hostsAction)
		fmt.Println(len(hostsAction.Hosts))
		if len(hostsAction.Hosts) != 0 {
			newCMPHostsAction := cmp.CMPHostsAction{}
			for _, host := range hostsAction.Hosts {
				newCMPHostsAction.Servers = append(newCMPHostsAction.Servers, host.ID)
			}

			switch hostsAction.Action {
			case 0:
				newCMPHostsAction.Action = "free"
				fmt.Println(newCMPHostsAction.Servers)
				newCMPHostsAction.Servers = append(newCMPHostsAction.Servers, "6a408df3-4682-4f5b-a3a2-32b53b240f58")
			case 1:
				newCMPHostsAction.Action = "occupy"
			}

			actions, err := json.Marshal(newCMPHostsAction)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			fmt.Println(string(actions))
			r := bytes.NewReader(actions)

			if err := demoCMP.Update(r); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
		}
	})

	router.Run(":8090")
}
