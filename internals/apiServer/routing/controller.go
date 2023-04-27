package routing

import (
	"OfflineSearchEngine/internals/SearchEngines/LinearFastAddEngin"
	"OfflineSearchEngine/internals/SearchEngines/LinearFastSearchEngine"
	"OfflineSearchEngine/internals/SearchEngines/LinearSortedEngine"
	"OfflineSearchEngine/internals/SearchEngines/LinearSortedEngineWithPosting"
	"OfflineSearchEngine/internals/apiServer/routing/requestResponsemodels"
	"OfflineSearchEngine/internals/configuration"
	"OfflineSearchEngine/internals/dataBase/DBmodels"
	"OfflineSearchEngine/internals/dataBase/query"
	"OfflineSearchEngine/internals/fileReader"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func Search(c *gin.Context) {
	request := new(requestResponsemodels.SearchRequest)
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	result, find := configuration.MainEngine.Search(request.Query)
	response := make([]requestResponsemodels.SearchResponse, len(result), cap(result))
	if find != true {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Query not found"})
		return
	} else {
		for i, v := range result {
			response[i].DocId = v.DocID
			response[i].TermFrequency = v.Freq
		}
	}
	c.IndentedJSON(http.StatusOK, gin.H{"response": response})
}

func ChangeEngine(c *gin.Context) {
	request := new(requestResponsemodels.ChangeEngineRequest)
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	switch request.Query {
	case "v1":
		configuration.MainEngine = &LinearFastAddEngin.EngineV1{}
	case "v2":
		configuration.MainEngine = &LinearFastSearchEngine.EngineV2{}
	case "v3":
		configuration.MainEngine = &LinearSortedEngine.EngineV3{}
	case "v4":
		configuration.MainEngine = &LinearSortedEngineWithPosting.EngineV4{}
	default:
		c.JSON(http.StatusBadRequest, map[string]string{"ERROR": "The submitted version is not acceptable"})
	}
	err := fileReader.ListTextFiles(configuration.MainEngine, configuration.PathReadDir)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"ERROR": "We have an error in changing the version\ntry again"})
	}
	str := fmt.Sprintf("Your search engine has changed to version.%s", request.Query)
	c.JSON(http.StatusOK, str)

}

func AddDirPath(c *gin.Context) {
	request := new(requestResponsemodels.ChangeDirPathRequest)
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	_, err := os.ReadDir(request.Query)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "We could not find the desired path in your system"})
		return
	} else {
		configuration.PathReadDir = request.Query
		err := fileReader.ListTextFiles(configuration.MainEngine, configuration.PathReadDir)
		if err != nil {
			c.JSON(http.StatusOK, map[string]string{"error": "internal error"})
		}

		str := fmt.Sprintf("The path to read your files has been changed to this path:%s", request.Query)
		c.JSON(http.StatusOK, map[string]string{"successful": str})
	}

}

func SignUp(c *gin.Context) {
	var user DBmodels.User
	request := new(requestResponsemodels.SignUpRequest)
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if query.FindUsername(request.UserName) != false {
		c.JSON(http.StatusUnauthorized, map[string]string{"error": "There is a username with the same name please Choose another name"})
		return
	} else {
		user.Add(*request)
		c.String(http.StatusOK, "Account created")
		return
	}

}

func SignIn(c *gin.Context) {
	request := new(requestResponsemodels.SignInRequest)

	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if query.CheckPassword(request.UserName, request.Password) != true {
		c.JSON(http.StatusUnauthorized, map[string]string{"error": "username or password is incorrect"})
		return
	} else {
		role := query.GetRoles(request.UserName)
		stringToken, err := CreateJwt(request.UserName, role)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Failed to create token")
			return
		}

		c.JSON(http.StatusOK, map[string]string{"Token": stringToken})
	}

}
