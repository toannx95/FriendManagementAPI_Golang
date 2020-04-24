package router

import (
	"database/sql"
	"friend/controller"
	"friend/repository"
	"friend/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func initUserController(db *sql.DB) controller.UserController {
	userRepository := repository.UserRepository{DB: db}
	userService := service.UserService{IUserRepository: userRepository}
	return controller.UserController{UserService: userService}
}

func initFriendController(db *sql.DB) controller.FriendController {
	relationshipRepository := repository.RelationshipRepository{DB: db}
	relationshipService := service.RelationshipService{IRelationshipRepository: relationshipRepository}

	userRepository := repository.UserRepository{DB: db}
	userService := service.UserService{IUserRepository: userRepository}

	friendService := service.FriendService{IRelationshipService: relationshipService, IUserService: userService}
	return controller.FriendController{FriendService: friendService}
}

func HandleRequest(db *sql.DB) {
	myRouter := mux.NewRouter().StrictSlash(true)

	userHandel := initUserController(db)
	friendHandel := initFriendController(db)

	myRouter.HandleFunc("/users", userHandel.GetAllUsers).Methods("GET")
	myRouter.HandleFunc("/users/create-user", userHandel.CreateUser).Methods("POST")

	myRouter.HandleFunc("/friends/create-friend", friendHandel.CreateFriend).Methods("POST")
	myRouter.HandleFunc("/friends/subscribe", friendHandel.CreateSubscribe).Methods("POST")
	myRouter.HandleFunc("/friends/block", friendHandel.CreateBlock).Methods("POST")
	myRouter.HandleFunc("/friends/get-friends-list", friendHandel.GetFriendsListByEmail).Methods("POST")
	myRouter.HandleFunc("/friends/get-common-friends-list", friendHandel.GetCommonFriends).Methods("POST")
	myRouter.HandleFunc("/friends/get-receivers-list", friendHandel.GetReceiversList).Methods("POST")

	log.Fatal(http.ListenAndServe(":8081", myRouter))
}