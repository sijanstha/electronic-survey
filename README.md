# Electronic Survey
Electronic survey aims to help organizers to manage their poll and attendees.

# Tools used
- Go (v 1.19.4)
- Mysql
- React
- Docker

## Installation
If you wish to run the project in docker then, you must ensure that docker is setup into your local machine and is up and running.

- Check if docker is installed into your system by issuing command: 
  ```
  docker -v
  ```
- Clone the project
  ```
  git clone https://github.com/sijanstha/electronic-survey.git
  git fetch origin
  ```
- Run the backend service. This runs the go backend services and exposes it on port 4000
  ```
  make start
  ```
  If you are running the service for the first time, before starting you need to create docker network using:
  ```
  make create-network
  ```
  To stop the running backend service
  ```
  make stop
  ```
- Run the frontend service. This spins up UI backed by react on port 3000. 
  ```
  make uirun
  ```
- Navigate to http://localhost:3000

## Default users
After service gets fully loaded, there are some default users that will get ingested to the system.

|Username   |  Password   |   Role   |
|------------|------------- |---------------|
|test@evs.com|password|ROLE_ORGANIZER|
|bob@evs.com|password|ROLE_ORGANIZER|
|john.doe@evs.com|password|ROLE_ORGANIZER|
