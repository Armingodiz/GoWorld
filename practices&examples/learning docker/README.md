# Gook

Rest microservice for saving items informations and signup&Login user .

## Services  

* [Users-API](https://github.com/Armingodiz/Gook-Users-API)
* [Oauth-API](https://github.com/Armingodiz/Gook-oauth-API)
* [Items-API](https://github.com/Armingodiz/Gook-Items-API)


## Used Databases 

* redis
* mysql


## Practiced design patterns

*  MVC in Users-API & Items-API
*  Domain driven development in Oauth-API



## Dependencies

name     | repo
------------- | -------------
  gin-gonic   | https://github.com/gin-gonic/gin
  mysql driver| github.com/go-sql-driver/mysql
  crypto/md5  | https://golang.org/pkg/crypto/md5/
  redis       | https://github.com/go-redis/redis
  gorilla-mux | https://github.com/gorilla/mux
  


## Installation 

First make sure you have installed all dependencies ,
Go to each service repository and install and run it .


## USER API EndPoints 

	* POST == > localhost:1111/users (create or signup the user with given informations as json)
	* GET ==> localhost:1111/users/:user_id (return user with given id)
	* PUT == > localhost:1111/users/:user_id  (update)
	* PATCH ==> localhost:1111/users/:user_id (partial update)
	* DELETE ==> localhost:1111/users/:user_id (Delete user with given id)
	* GET == > localhost:1111/internal/users/search (Search for users)
	* POST ==> localhost:1111/users/login (oauth API use this endpoint for Login)
  
## Oauth API EndPoints 
  
  	* GET ==> localhost:2222/oauth/access_token/:access_token_id (Get Access token by id)
  	* POST ==> localhost:2222/oauth/access_token (Create Access token)
  
## Oauth API EndPoints 
    
 	 * POST ==> 127.0.0.1:3333/items  (Create item)
  	 * GET ==> 127.0.0.1:3333/items/{id} (Get item by id)
  	 * POST ==> 127.0.0.1:3333/items/search (Search for users)

# Overal view 

![view](view.png)

