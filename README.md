# Electronic Survey
Electronic survey aims to help organizers to manage their poll and attendees.

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
  To stop the running backend service
  ```
  make stop
  ```
- Run the frontend service.
  ```
  make uirun
  ```
