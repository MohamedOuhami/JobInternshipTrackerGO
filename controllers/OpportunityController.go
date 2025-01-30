package controllers

import (
	"fmt"
	"net/http"

	"github.com/MohamedOuhami/JobInternshipTrackerGO/initializers"
	"github.com/MohamedOuhami/JobInternshipTrackerGO/models"
	"github.com/gin-gonic/gin"
)

// Get all the Opportunities

func GetAllOpportunities(c *gin.Context) {

	var opportunities []models.Opportunity

	// Trying to get all of the opportunities available in the database
	result := initializers.DB.Find(&opportunities)

	if result.Error != nil {

		c.JSON(http.StatusBadRequest, gin.H{"message": "Error fetching the opportunities"})

		return

	}
	c.JSON(http.StatusOK, gin.H{"Opportunities": opportunities})

}

// Get a opportunity by Id

func GetopportunityById(c *gin.Context) {
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

	var opportunity models.Opportunity

	opportunityId := c.Param("id")

	initializers.DB.First(&opportunity, opportunityId)

	if opportunity.ID == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"message": "Error fetching the opportunity of id " + opportunityId})

		return

	}

	// Ensure the user is the owner of the job before deleting
	if !checkOwnerOpportunity(int(ownerUser.ID), opportunity) {
		c.JSON(http.StatusForbidden, gin.H{"message": "This user is not authorized to get these opportunities"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"opportunity": opportunity})

}

// Get Opportunities by OwnerId

func GetOpportunitiesByOwner(c *gin.Context) {
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

	var opportunities []models.Opportunity

	ownerId := c.Param("ownerId")

	fmt.Println(ownerId)

	result := initializers.DB.Where("owner_id = ?", ownerId).Find(&opportunities)

	if result.Error != nil {

		c.JSON(http.StatusBadRequest, gin.H{"message": "Error fetching the opportunity of owner id " + ownerId})

		return

	}

	// Ensure the current user is the owner of the jobs they are requesting
	if ownerId != fmt.Sprintf("%d", ownerUser.ID) {
		c.JSON(http.StatusForbidden, gin.H{"message": "This user is not authorized to view these jobs"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Opportunities": opportunities})
}

// Create a new opportunity

func Createopportunity(c *gin.Context) {

	var opportunityToCreate models.Opportunity

	err := c.Bind(&opportunityToCreate)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"message": "Error binding the opportunity"})

		return

	}

	var ownerId = 7

	opportunityToCreate.UserID = uint(ownerId)

	result := initializers.DB.Create(&opportunityToCreate)

	if result.Error != nil {

		c.JSON(http.StatusBadRequest, gin.H{"message": "Error creating the new opportunity " + result.Error.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"message": opportunityToCreate})

}

// Edit an existing opportunity

func Editopportunity(c *gin.Context) {
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

	// Fetching the wanted opportunity
	var opportunity models.Opportunity

	var editedopportunity models.Opportunity

	opportunityId := c.Param("id")

	initializers.DB.First(&opportunity, opportunityId)

	// Check If this opportunity exists

	if opportunity.ID == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"message": "Error fetching the opportunity of id " + opportunityId})

		return
	}

	// Ensure the user is the owner of the job before deleting
	if !checkOwnerOpportunity(int(ownerUser.ID), opportunity) {
		c.JSON(http.StatusForbidden, gin.H{"message": "This user is not authorized to edit this opportunity"})
		return
	}

	// Bind the data from the request

	err := c.Bind(&editedopportunity)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"message": "Error binding the edited opportunity"})

		return
	}

	// Change the values of the existing opportunity

	opportunity.CompanyName = editedopportunity.CompanyName
	opportunity.PostName = editedopportunity.PostName
	opportunity.City = editedopportunity.City
	opportunity.JobType = editedopportunity.JobType

	initializers.DB.Save(&opportunity)

	// Return a response

	c.JSON(http.StatusOK, gin.H{"message": "Successfully edited the opportunity"})

}

// Delete an Existing opportunity

func Deleteopportunity(c *gin.Context) {
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

	// Fetching the wanted opportunity
	var opportunity models.Opportunity

	opportunityId := c.Param("id")

	initializers.DB.First(&opportunity, opportunityId)

	// Check If this opportunity exists

	if opportunity.ID == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"message": "Error fetching the opportunity of id " + opportunityId})

		return
	}

	// Ensure the user is the owner of the job before deleting
	if !checkOwnerOpportunity(int(ownerUser.ID), opportunity) {
		c.JSON(http.StatusForbidden, gin.H{"message": "This user is not authorized to delete this opportunity"})
		return
	}

	err := initializers.DB.Delete(&opportunity)

	if err.Error != nil {

		c.JSON(http.StatusBadRequest, gin.H{"message": "Error deleting the opportunity of id " + opportunityId})

		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleting opportunity of id " + c.Param("id")})

}

// Mass Delete Opportunities

func MassDeleteOpportunities(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "Mass deleting opportunities"})

}

// Turn an Opportunity to Job
func TurnOpportunityToJob(c *gin.Context) {

	var opportunity models.Opportunity
	var job models.Job

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

	opportunityId := c.Param("id")

	// This process should be a transaction. as If one process fail, roll back the database to the original state

	tx := initializers.DB.Begin()

	initializers.DB.First(&opportunity, opportunityId)

	if err := tx.First(&opportunity, opportunityId).Error; err != nil {
		tx.Rollback() // Rollback the transaction in case of an error
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error fetching the opportunity of id " + opportunityId})
		return
	}

	if opportunity.ID == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not find the opportunity of id " + opportunityId})

		return

	}

	// Ensure the user is the owner of the job before deleting
	if !checkOwnerOpportunity(int(ownerUser.ID), opportunity) {
		tx.Rollback() // Rollback the transaction in case of an error
		c.JSON(http.StatusForbidden, gin.H{"message": "This user is not authorized to turn this opportunity to a job"})
		return
	}

	job.City = opportunity.City
	job.JobType = opportunity.JobType
	job.CompanyName = opportunity.CompanyName
	job.PostName = opportunity.PostName
	job.UserID = opportunity.UserID

	if err := tx.Create(&job).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"message": "Error Turning the Opportunity of id " + opportunityId})

		return

	}

	if err := tx.Delete(&opportunity).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"message": "Error Turning the Opportunity of id " + opportunityId})

		return

	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"NewJob": job})
}

func checkOwnerOpportunity(ownerId int, job models.Opportunity) bool {

	if job.UserID == uint(ownerId) {
		return true
	} else {
		return false
	}
}
