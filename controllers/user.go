package controllers

import (
	"go-server/models"
	"go-server/responses"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var validate = validator.New()

// @Tags user
// @Accept json
// @Produce json
// @Param user body models.User true "user"
// @Success 200 string ok
// @Router /api/v0/user [post]
func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := models.User{}
		result, err := user.Read(&gorm.DB{})
		if err != nil {
			c.JSON(http.StatusNotFound, responses.ApiResponse{
				Status:  http.StatusNotFound,
				Message: "error",
				Data: map[string]interface{}{
					"data": "Not found",
				},
			})
		}
		c.JSON(http.StatusOK, responses.ApiResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data: map[string]interface{}{
				"data": result,
			},
		})
	}
}

// @Tags user
// @Accept json
// @Produce json
// @Success 200 string ok
// @Router /api/v0/users [get]
func ReadUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := models.User{}
		result, err := user.Read(&gorm.DB{})
		if err != nil {
			c.JSON(http.StatusNotFound, responses.ApiResponse{
				Status:  http.StatusNotFound,
				Message: "error",
				Data: map[string]interface{}{
					"data": "Not found",
				},
			})
		}
		c.JSON(http.StatusOK, responses.ApiResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data: map[string]interface{}{
				"data": result,
			},
		})
	}
}

// // @Tags user
// // @Accept json
// // @Produce json
// // @Param userId path string true "userId"
// // @Success 200 string ok
// // @Router /api/v0/user/{userId} [get]
// func ReadUser() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 		userId := c.Param("userId")
// 		var user models.User
// 		defer cancel()

// 		objId, _ := primitive.ObjectIDFromHex(userId)

// 		err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
// 			return
// 		}

// 		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": user}})
// 	}
// }

// // @Tags user
// // @Accept json
// // @Produce json
// // @Success 200 string ok
// // @Router /api/v0/users [get]
// func ReadUsers() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 		var users []models.User
// 		defer cancel()

// 		results, err := userCollection.Find(ctx, bson.M{})

// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
// 			return
// 		}

// 		//reading from the db in an optimal way
// 		defer results.Close(ctx)
// 		for results.Next(ctx) {
// 			var singleUser models.User
// 			if err = results.Decode(&singleUser); err != nil {
// 				c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
// 			}

// 			users = append(users, singleUser)
// 		}

// 		c.JSON(http.StatusOK,
// 			responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": users}},
// 		)
// 	}
// }

// // @Tags user
// // @Accept json
// // @Produce json
// // @Param userId path string true "userId"
// // @Param JsonData body models.User false "name"
// // @Success 200 string ok
// // @Router /api/v0/user/{userId} [patch]
// func UpdateUser() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 		userId := c.Param("userId")
// 		var user models.User
// 		defer cancel()

// 		objId, _ := primitive.ObjectIDFromHex(userId)

// 		if err := c.BindJSON(&user); err != nil {
// 			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
// 			return
// 		}

// 		if validationErr := validate.Struct(&user); validationErr != nil {
// 			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
// 			return
// 		}

// 		update := bson.M{"name": user.Name, "location": user.Location, "title": user.Title}
// 		result, err := userCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})

// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
// 			return
// 		}

// 		var updatedUser models.User
// 		if result.MatchedCount == 1 {
// 			err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedUser)
// 			if err != nil {
// 				c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
// 				return
// 			}
// 		}

// 		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedUser}})
// 	}
// }

// // @Tags user
// // @Accept json
// // @Produce json
// // @Param userId path string true "userId"
// // @Success 200 string ok
// // @Router /api/v0/user/{userId} [delete]
// func DeleteUser() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 		userId := c.Param("userId")
// 		defer cancel()

// 		objId, _ := primitive.ObjectIDFromHex(userId)

// 		result, err := userCollection.DeleteOne(ctx, bson.M{"id": objId})

// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
// 			return
// 		}

// 		if result.DeletedCount < 1 {
// 			c.JSON(http.StatusNotFound,
// 				responses.UserResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "User with specified ID not found!"}},
// 			)
// 			return
// 		}

// 		c.JSON(http.StatusOK,
// 			responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "User successfully deleted!"}},
// 		)
// 	}
// }
