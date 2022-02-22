package controller

import (
	"TodoList/JWT"
	"TodoList/helpers"

	"github.com/gin-gonic/gin"
)

func GetJWT(c *gin.Context) {
	type RequestData struct {
		PublicAddress    string `form:"public_address" json:"public_address" binding:"required"`
		Signature        string `form:"signature" json:"signature" binding:"required"`
	}  
	var json RequestData

	if err := c.ShouldBindJSON(&json); err == nil {
		// check signature is start with 0x
		if((len(json.Signature) < 2 || len(json.PublicAddress) < 2)){
			c.JSON(400, gin.H{
				"message": "Incorrect length of Signature or Public Address",
				"status": false,
			})
			return
		}

		if(json.Signature[:2] != "0x" && json.PublicAddress[:2] != "0x") {
			c.JSON(400, gin.H{
				"message": "Incorrect Format of Signature or Public Address",
				"status": false,
			})
			return
		}

		if helpers.VerifySig(json.PublicAddress, json.Signature, []byte("Login TodoList")) {
			jwtWrapper := JWT.JwtWrapper{
				SecretKey:       "IGhr2aIe6qnu8qBqXkC5X6DFbI4xxXQ7",
				Issuer:          "AuthService",
				ExpirationHours: 1000000000000000,
			}
			token, err := jwtWrapper.GenerateToken(json.PublicAddress)
			if err != nil {
				c.JSON(500, gin.H{
					"message": "Error generating token",
					"status": false,
				})
				return
			}
			c.JSON(200, gin.H{
				"message": "Login successfully",
				"status": true,
				"access_token": token,
			})
		} else {
			c.JSON(400, gin.H{
				"message": "Invalid signature",
				"status": false,
			})
		}
	} else if err.Error() == "EOF"{
		c.JSON(400, gin.H{
			"status": false,
			"message": "Invalid JSON",
		})
	} else {
		c.JSON(400, gin.H{
			"status": false,
			"message": err.Error(),
		})
	}	

}	