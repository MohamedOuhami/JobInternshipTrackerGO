package controllers

import (
	"fmt"
	"net/http"

	"github.com/MohamedOuhami/JobInternshipTrackerGO/initializers"
	"github.com/MohamedOuhami/JobInternshipTrackerGO/models"
	"github.com/gin-gonic/gin"
)

// Get all the Jobs
var connectedUser models.User

func GetAllJobs(c *gin.Context) {

	var jobs []models.Job

	// Trying to get all of the jobs available in the database
	result := initializers.DB.Find(&jobs)

	if result.Error != nil {

		c.JSON(http.StatusBadRequest, gin.H{"message": "Error fetching the jobs"})

		return

	}
	c.JSON(http.StatusOK, gin.H{"Jobs": jobs})

}

// Get a Job by Id

func GetJobById(c *gin.Context) {

	owner, _ := c.Get("user")

	var job models.Job

	jobId := c.Param("id")

	initializers.DB.First(&job, jobId)

	if job.ID == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"message": "Error fetching the job of id " + jobId})

		return

	}

	if !checkOwner(int(owner.(models.User).ID), job) {

		c.JSON(http.StatusBadRequest, gin.H{"message": "This user is not authorized to see this opportunity"})

		return

	}
	c.JSON(http.StatusOK, gin.H{"Job": job})

}

// Get Jobs by OwnerId

func GetJobsByOwner(c *gin.Context) {
	// Get the user from the context
	owner, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User not authenticated"})
		return
	}

	// Type assert to models.User
	ownerUser, ok := owner.(models.User)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid user data"})
		return
	}

	// Fetch jobs by owner ID
	var jobs []models.Job
	ownerId := c.Param("ownerId")

	// Ensure the current user is the owner of the jobs they are requesting
	if ownerId != fmt.Sprintf("%d", ownerUser.ID) {
		c.JSON(http.StatusForbidden, gin.H{"message": "This user is not authorized to view these jobs"})
		return
	}

	result := initializers.DB.Where("owner_id = ?", ownerId).Find(&jobs)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error fetching the jobs of owner id " + ownerId})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Jobs": jobs})
}

// Create a new Job

func CreateJob(c *gin.Context) {

	var jobToCreate models.Job

	err := c.Bind(&jobToCreate)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"message": "Error binding the job"})

		return

	}

	var ownerId = 7

	jobToCreate.UserID = uint(ownerId)

	result := initializers.DB.Create(&jobToCreate)

	if result.Error != nil {

		c.JSON(http.StatusBadRequest, gin.H{"message": "Error creating the new Job " + result.Error.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"message": jobToCreate})

}

// Edit an existing Job

func EditJob(c *gin.Context) {

	// Get the user from the context
	owner, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User not authenticated"})
		return
	}

	// Type assert to models.User
	ownerUser, ok := owner.(models.User)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid user data"})
		return
	}

	// Fetching the wanted Job
	var job models.Job

	var editedJob models.Job

	jobId := c.Param("id")

	initializers.DB.First(&job, jobId)

	// Check If this job exists

	if job.ID == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"message": "Error fetching the job of id " + jobId})

		return
	}
	// Bind the data from the request

	// Ensure the user is the owner of the job before deleting
	if !checkOwner(int(ownerUser.ID), job) {
		c.JSON(http.StatusForbidden, gin.H{"message": "This user is not authorized to edit this job"})
		return
	}

	err := c.Bind(&editedJob)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"message": "Error binding the edited Job"})

		return
	}

	// Change the values of the existing job

	job.CompanyName = editedJob.CompanyName
	job.PostName = editedJob.PostName
	job.City = editedJob.City
	job.JobType = editedJob.JobType

	initializers.DB.Save(&job)

	// Return a response

	c.JSON(http.StatusOK, gin.H{"message": "Successfully edited the job"})

}

// Delete an Existing Job

func DeleteJob(c *gin.Context) {
	// Get the user from the context
	owner, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User not authenticated"})
		return
	}

	// Type assert to models.User
	ownerUser, ok := owner.(models.User)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid user data"})
		return
	}

	// Fetch the job to delete
	var job models.Job
	jobId := c.Param("id")
	initializers.DB.First(&job, jobId)

	// Check if job exists
	if job.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error fetching the job of id " + jobId})
		return
	}

	// Ensure the user is the owner of the job before deleting
	if !checkOwner(int(ownerUser.ID), job) {
		c.JSON(http.StatusForbidden, gin.H{"message": "This user is not authorized to delete this job"})
		return
	}

	// Delete the job
	err := initializers.DB.Delete(&job).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error deleting the job of id " + jobId})
		return
	}

	// Return the response
	c.JSON(http.StatusOK, gin.H{"message": "Successfully deleted job with id " + jobId})
}

// Mass Delete Jobs

func MassDeleteJobs(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "Mass deleting jobs"})

}

func checkOwner(ownerId int, job models.Job) bool {

	if job.UserID == uint(ownerId) {
		return true
	} else {
		return false
	}
}
