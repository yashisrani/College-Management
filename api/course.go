package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yashisrani/Go-Backend/model"
	"github.com/yashisrani/Go-Backend/utils"
)

func (api *Apiroutes) CourseRoutes(route *gin.Engine) {
	group := route.Group("Course")
	{
		group.POST("/create", api.AuthMiddleware(), api.CreateCourse)
		group.GET("/getCoursees", api.AuthMiddleware(), api.GetCourse)
		group.GET("/:id", api.AuthMiddleware(), api.GetCourseByID)
		group.GET("/filter", api.AuthMiddleware(), api.GetCourseByFilter)
		group.PUT("/update/:id", api.AuthMiddleware(), api.UpdateCourse)
		group.DELETE("/delete/:id", api.AuthMiddleware(), api.DeleteCourse)
	}
}

// Handler to create a Course
// @router /Course/create [post]
// @summary Create a new Course
// @tags Course
// @accept json
// @produce json
// @param Course body model.Course true "Course object"
// @success 201 {object} model.Course
// @Security ApiKeyAuth
func (api Apiroutes) CreateCourse(ctx *gin.Context) {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.CreateCourse, "Creating new Course", nil)
	api.Server.CreateCourse(ctx)
}

// Handler to get all Courses
// @router /Course/getCourses [get]
// @summary Get all Courses
// @tags Course
// @produce json
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.Course
// @Security ApiKeyAuth
func (api Apiroutes) GetCourse(ctx *gin.Context) {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.GetCourse, "Fetching Coursees", nil)
	api.Server.GetCourse(ctx)
}

// Handler to get a Course by ID
// @router /Course/{id} [get]
// @summary Get a Course by ID
// @tags Course
// @produce json
// @param id path string true "Course ID"
// @success 200 {object} model.Course
// @Security ApiKeyAuth
func (api Apiroutes) GetCourseByID(ctx *gin.Context) {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.GetCourseByID, "Fetching Course by ID", nil)
	api.Server.GetCourseByID(ctx)
}

// Handler to get Courses based on filters
// @router /Course/filter [get]
// @summary Get Courses based on given filters
// @tags Course
// @produce json
// @param course_name query string false "Course Name"
// @param school_id query string false "School ID"
// @param active query bool false "Active status"
// @param instructor query string false "Instructor"
// @param duration_weeks query int false "Duration in weeks"
// @param credits query int false "Course Credits"
// @param fee query number false "Course Fee"
// @param description query string false "Course Description"
// @param syllabus query string false "Syllabus"
// @param start_date query string false "Start Date (YYYY-MM-DD)"
// @param end_date query string false "End Date (YYYY-MM-DD)"
// @param created_by query string false "Created By"
// @param updated_by query string false "Updated By"
// @param course_id query string false "Course ID"
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.Course
// @Security ApiKeyAuth
func (api Apiroutes) GetCourseByFilter(ctx *gin.Context) {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.GetCourseByFilter, "Fetching Courses by filter", nil)
	api.Server.GetCourseByFilter(ctx)
}


// Handler to update a Course
// @router /Course/update/{id} [put]
// @summary Update a Course
// @tags Course
// @accept json
// @produce json
// @param id path string true "Course ID"
// @param Course body model.Course true "Course object"
// @success 200 {object} model.Course
// @Security ApiKeyAuth
func (api Apiroutes) UpdateCourse(ctx *gin.Context) {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.UpdateCourse, "Updating Course", nil)
	api.Server.UpdateCourse(ctx)
}

// Handler to delete a Course
// @router /Course/delete/{id} [delete]
// @summary Delete a Course
// @tags Course
// @param id path string true "Course ID"
// @Security ApiKeyAuth
func (api Apiroutes) DeleteCourse(ctx *gin.Context) {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.DeleteCourse, "Deleting Course", nil)
	api.Server.DeleteCourse(ctx)
}
