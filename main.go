package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"k8s.io/apiserver/pkg/apis/audit"
	"k8s.io/klog"
	"log"
	"net/http"
	"os"
)

func main() {
	logger := log.New(os.Stderr, "", 0)
	logger.Println("[WARNING] DON'T USE THE EMBED CERTS FROM THIS EXAMPLE IN PRODUCTION ENVIRONMENT, GENERATE YOUR OWN!")

	r := gin.Default()
	r.GET("/sink", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.POST("/sink", func(c *gin.Context) {
		bodyBytes, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			log.Fatal(err)
		}
		defer c.Request.Body.Close()
		var eventList audit.EventList
		err = json.Unmarshal(bodyBytes, &eventList)
		if err != nil {
			log.Fatal(err)
		}
		for _, event := range eventList.Items {
			klog.Infof("this event is %+v\n", event)
			// asyncProducer(string(jsonBytes))
		}
		//bodyString := string(bodyBytes)
		//log.Printf("kubernetes sink msg: %v", bodyString)
		c.String(http.StatusOK, "pong")
	})

	// Listen and Server in https://127.0.0.1:8080
	r.RunTLS(":443", "./webhook-server-tls.crt", "./webhook-server-tls.key")
}
