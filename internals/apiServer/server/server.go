package server

import (
	"OfflineSearchEngine/internals/apiServer/CreateEngine"
	"OfflineSearchEngine/internals/apiServer/server/Authentication"
	"OfflineSearchEngine/internals/apiServer/server/requestResponsemodels"
	"OfflineSearchEngine/internals/dataBase/DBmodels"
	"OfflineSearchEngine/internals/dataBase/query"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	Engine  *CreateEngine.ISearchEngine
	Gin     *gin.Engine
	Context *gin.Context
}

func CreateServer(engine *CreateEngine.ISearchEngine) *Server {
	server := &Server{
		Engine: engine,
		Gin:    gin.Default(),
	}
	server.SetRouts()
	return server
}

func (s *Server) Run(address string) error {
	return s.Gin.Run(address)
}

func (s *Server) SetRouts() {

	s.Gin.GET("signup", SignUp)
	s.Gin.GET("signin", SignIn)
	signInGroup := s.Gin.Group("signin")
	signInGroup.POST("search", Authentication.AuthMiddleware(false), s.Search)
	//signInGroup.POST("newdir", Authentication.AuthMiddleware(false), s.AddDirPath)
	//signInGroup.POST("change-engine", Authentication.AuthMiddleware(true), routing.ChangeEngine)
}

func (s *Server) Search(c *gin.Context) {
	request := new(requestResponsemodels.SearchRequest)
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	result, find := s.Engine.Engine.Search(request.Query)
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
		stringToken, err := Authentication.CreateJwt(request.UserName, role)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Failed to create token")
			return
		}

		c.JSON(http.StatusOK, map[string]string{"Token": stringToken})
	}

}

//func (s *Server) ChangeEngine(c *gin.Context) {
//	request := new(requestResponsemodels.ChangeEngineRequest)
//	if err := c.ShouldBindJSON(request); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
//		return
//	}
//	switch request.Query {
//	case "v1":
//		s.Engine.Engine = &LinearFastAddEngin.EngineV1{}
//	case "v2":
//		s.Engine.Engine = &LinearFastSearchEngine.EngineV2{}
//	case "v3":
//		s.Engine.Engine = &LinearSortedEngine.EngineV3{}
//	case "v4":
//		s.Engine.Engine = &LinearSortedEngineWithPosting.EngineV4{}
//	default:
//		s.Context.JSON(http.StatusBadRequest, map[string]string{"ERROR": "The submitted version is not acceptable"})
//	}
//	err := fileReader.ListTextFiles(s.Engine.Engine,s.Engine.PathReadDir)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, map[string]string{"ERROR": "We have an error in changing the version\ntry again"})
//	}
//	str := fmt.Sprintf("Your search engine has changed to version.%s", request.Query)
//	c.JSON(http.StatusOK, str)
//
//}

//func (s *Server) AddDirPath(c *gin.Context) {
//	request := new(requestResponsemodels.ChangeDirPathRequest)
//	if err := c.ShouldBindJSON(request); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
//		return
//	}
//	_, err := os.ReadDir(request.Query)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "We could not find the desired path in your system"})
//		return
//	} else {
//		s.Engine.PathReadDir = request.Query
//		err := fileReader.ListTextFiles(s.Engine.Engine, s.Engine.PathReadDir)
//		if err != nil {
//			c.JSON(http.StatusOK, map[string]string{"error": "internal error"})
//		}
//
//		str := fmt.Sprintf("The path to read your files has been changed to this path:%s", request.Query)
//		c.JSON(http.StatusOK, map[string]string{"successful": str})
//	}
//
//}
