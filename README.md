Aim : Creating an API where an user can sign-up the required fields - id, name, mobile number, address, username, password.

Task:-
Firstly storing the data in MySql database.
I have database configuration, model, routes, controller and main file.

Process:-
config : I have added the database configuration in config.go file under config directory.

model : The model for the user consisting of the following fields - id, name, mobile number, address, username, password, is created in model.go under model directory.

routes : I have created a group to adjust my routing related to users and have to take care of the group while consuming the api. This is created in routes.go file under routes directory.

model, controllers : http requests coming from front end is handled in controller. I have created different functions that handles specific requests routed to controller by the router. 
I have made user.go file inside model directory to interact with the database. Response to the user is according to the data that is received from my database. 
If no error is encountered, I have supplied response as StatusOK and if we get an error, error status is supplied. 
I have created profile_handler.go file inside controllers directory.

main : The starter function of my project is the main.go file. I connect mysql and setup router form here. This is created in the root of the project.
