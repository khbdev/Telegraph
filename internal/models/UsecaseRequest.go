package models

type CreateDataInput struct {
	Title     string `json:"title" binding:"required"`
	YourName  string `json:"your_name" binding:"required"`
	YourStory string `json:"your_story" binding:"required"`
}