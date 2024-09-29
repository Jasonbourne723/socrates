package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Jasonbourne723/socrates/app/common/request"
	"github.com/Jasonbourne723/socrates/app/common/response"
	"github.com/Jasonbourne723/socrates/app/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type AuthApi struct {
}

func (a *AuthApi) Register(c *gin.Context) {
	var form request.Register
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.UserService.Register(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		tokenData, err, _ := services.JwtService.CreateToken(services.AppGuardName, user)
		if err != nil {
			response.BusinessFail(c, err.Error())
			return
		}
		response.Success(c, tokenData)
	}
}

func (a *AuthApi) Login(c *gin.Context) {
	var form request.Login
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.UserService.Login(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		tokenData, err, _ := services.JwtService.CreateToken(services.AppGuardName, user)
		if err != nil {
			response.BusinessFail(c, err.Error())
			return
		}
		response.Success(c, tokenData)
	}
}

func (a *AuthApi) Info(c *gin.Context) {
	err, user := services.UserService.GetUserInfo(c.Keys["id"].(string))
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, user)
}

func (a *AuthApi) Logout(c *gin.Context) {
	err := services.JwtService.JoinBlackList(c.Keys["token"].(*jwt.Token))
	if err != nil {
		response.BusinessFail(c, "登出失败")
		return
	}
	response.Success[any](c, nil)
}

func (a *AuthApi) GithubLogin(c *gin.Context) {
	githubLogin := request.GithubLogin{}
	if err := c.ShouldBindJSON(&githubLogin); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(githubLogin, err))
		return
	}
	ssoApplicationService := services.NewSsoApplicationService()
	githubInfo, err := ssoApplicationService.GetByName("github")
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}

	url := fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s", githubInfo.AppKey, githubInfo.AppSecret, githubLogin.Code)

	token, err := getGitHubAccessToken(url)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	githubUser, err := getGitHubUserInfo(token)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	user, err := services.UserService.RegisterByGitHub(githubUser)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	tokenData, err, _ := services.JwtService.CreateToken(services.AppGuardName, user)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, tokenData)
}

func getGitHubAccessToken(uri string) (token string, err error) {

	client := &http.Client{}

	req, err := http.NewRequest("POST", uri, bytes.NewBuffer([]byte{}))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Accept", "application/json")

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	githubToken := response.GitHubToken{}
	err = json.Unmarshal(body, &githubToken)
	if err != nil {
		return
	}
	token = githubToken.AccessToken

	return
}

func getGitHubUserInfo(token string) (user response.GitHubUser, err error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.github.com/user", bytes.NewBuffer([]byte{}))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "bearer "+token)

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	// 解析 JSON 响应
	user = response.GitHubUser{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		return
	}
	return
}
