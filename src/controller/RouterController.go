package controller

import (
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	repositoryImpl "repository/impl"
	serviceImpl "service/impl"
)

func HandleRequest(db *sql.DB) {
	myRouter := mux.NewRouter().StrictSlash(true)

	userRepository := repositoryImpl.UserRepositoryImpl{DB: db}
	userService := serviceImpl.UserServiceImpl{UserRepository: userRepository}
	userHandel := UserController{UserService: userService}

	relationshipRepository := repositoryImpl.RelationshipRepositoryImpl{DB: db}
	relationshipService := serviceImpl.RelationshipServiceImpl{RelationshipRepository: relationshipRepository}
	friendService := serviceImpl.FriendServiceImpl{RelationshipService: relationshipService, UserService: userService}
	friendHandel := FriendController{FriendService: friendService}

	myRouter.HandleFunc("/users", userHandel.GetAllUsers).Methods("GET")
	myRouter.HandleFunc("/users/create-user", userHandel.CreateUser).Methods("POST")

	myRouter.HandleFunc("/friends/create-friend", friendHandel.CreateUser).Methods("POST")
	myRouter.HandleFunc("/friends/subscribe", friendHandel.CreateSubscribe).Methods("POST")
	myRouter.HandleFunc("/friends/block", friendHandel.CreateBlock).Methods("POST")
	myRouter.HandleFunc("/friends/get-friends-list", friendHandel.GetFriendsListByEmail).Methods("POST")
	myRouter.HandleFunc("/friends/get-common-friends-list", friendHandel.GetCommonFriends).Methods("POST")
	myRouter.HandleFunc("/friends/get-receivers-list", friendHandel.GetReceiversList).Methods("POST")

	log.Fatal(http.ListenAndServe(":8081", myRouter))
}