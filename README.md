# Multiplayer-Game

To know which modes are being played the most at any moment of the day in their local area, which they provide as a three-digit area code

## Description

A backend service which will allow user to register/create and store a player in mongoDB and Also players who are currently playing the game are stored in cache with the players respective mode and areaCode. The areaCode code is a number. Different number means different area/location. The modes are Solo, DUO, TRIO and SQUAD. which is an enum where it should be given in the request data as 1 or 2 or 3 or 4. Then players name, areaCode and mode is used to begin the game and to end the game player name and areaCode is send to end the game.

## Getting Started

### Dependencies

* Golang application development, redis, mongoDB, Swagger-ui
* protobuf, Docker and Docker compose in high level

### Installing

* Install Docker in your system, pull redis and mongoDB image
* Change env values If you need to change the environment variables value in the docker-compose.yml file

### Executing program

* Then run "docker-compose build" in project's directory to build the application image
* Then finally run "docker-compose up"
* The application will be up and running in the port mentioned in docker-compose file
* Use /swagger/index.html get endpoint to access swagger-ui

### Design Choices Document

* I have desinged it in 3 layer where layers controller, service and data Layer.
* To separate the business logic, controller and databse operations.
* Which can be scaled-up later if required
* Load-balancer like ingress and auto-scaler in cloud environment can be used to manage the large workload



## Authors

Contributors names and contact info

Jayasuriyaa  
jsuriyaa2001@gmail.com
