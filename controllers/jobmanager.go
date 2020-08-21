package controllers

import (
	"gin-admin/utils"
	"github.com/bamzi/jobrunner"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type jobTest struct {
	jobID int
}

func (j jobTest) Run() {
	now := time.Now().Format("2006-01-02 15:04:05")
	//utils.Logger.Printf("time: %s, jobID: %d, Every 2 sec execute job --->\n", now, j.jobID)
	utils.Logger.WithFields(logrus.Fields{
		"time": now,
		"jobID": j.jobID,
	}).Info("Every 2 sec execute job -->")
}

func AddJob(c *gin.Context) {
	utils.Logger.Info("add job")
	err := jobrunner.Schedule("@every 2s", jobTest{jobID: rand.Int()})
	if err != nil {

		c.JSON(http.StatusNotImplemented, err.Error())
	} else {
		c.JSON(200, jobrunner.StatusJson())
	}

}

func ListJobs(c *gin.Context) {
	c.JSON(200, jobrunner.StatusJson())
}

func RemoveJob(c *gin.Context) {
	jobID := c.Params.ByName("jobID")
	if jobID != "" {
		removeID, _ := strconv.Atoi(jobID)
		id := cron.EntryID(removeID)

		jobrunner.Remove(id)
		c.JSON(200, gin.H{"remove job ok": jobID})
	}
}
