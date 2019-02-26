package controller

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "min-services/models"
    "min-services/pkg/e"
    "min-services/pkg/util"
)

// Example for binding JSON ({"user": "manu", "password": "123"})
func LoginJson(c *gin.Context) {
    var auth models.Login
    if err := c.ShouldBindJSON(&auth); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"code": e.INVALID_PARAMS, "Msg": e.GetMsg(e.INVALID_PARAMS), "error": err.Error()})
        return
    }
    
    if auth.User != "manu" || auth.Password != "123" {
        c.JSON(http.StatusUnauthorized, gin.H{"code": e.ERROR_AUTH_CHECK_TOKEN_FAIL, "Msg": e.GetMsg(e.ERROR_AUTH_CHECK_TOKEN_FAIL)})
        return
    } 
    
    token, err := util.GenerateToken(auth.User, auth.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"code": e.ERROR_AUTH_TOKEN, "Msg": e.GetMsg(e.ERROR_AUTH_TOKEN), "error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"code": e.SUCCESS, "data": token})
}

// Example for binding XML (
//  <?xml version="1.0" encoding="UTF-8"?>
//  <root>
//      <user>user</user>
//      <password>123</password>
//  </root>)
func LoginXML(c *gin.Context) {
    var xml models.Login
    if err := c.ShouldBindXML(&xml); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    if xml.User != "manu" || xml.Password != "123" {
        c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
        return
    } 
    
    c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}


// Example for binding a HTML form (user=manu&password=123)
func LoginForm(c *gin.Context) {
    var form models.Login
    // This will infer what binder to use depending on the content-type header.
    if err := c.ShouldBind(&form); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    if form.User != "manu" || form.Password != "123" {
        c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
        return
    } 
    
    c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}


