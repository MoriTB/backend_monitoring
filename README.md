# backend_monitoring
A HTTP endpoint monitor service written in go with RESTful API.
Backend Monitoring is a Python script for monitoring web services and APIs. It uses the requests library to check if the endpoints are responding correctly and logs any errors or anomalies.

Features
Monitor endpoints and log any errors or anomalies.
Set custom headers for requests.
Set the time interval between requests.
Enable or disable SSL verification.
Set the timeout for requests.
Getting Started
# Prerequisites
Python 3.6 or higher.
You can install it via pip: pip install requests
PostgreSql for dealing with the database.
# Installation
Clone the repository: git clone https://github.com/MoriTB/backend_monitoring.git                     
Navigate to the directory: cd backend_monitoring                  
Install the required dependencies: pip install -r requirements.txt           
Copy the config.example.json file and rename it to config.json: cp config.example.json config.json              
Edit the config.json file to specify the endpoints to monitor and set any optional parameters.            

# Usage
you can use the frontend that has been designed.                         
or send GET/POST request via POSTMAN and work with postman to check the Rest API.                                 
The script will run indefinitely and log any errors or anomalies to the console.

# Database
This project uses PostgreSQL as its database management system. PostgreSQL is an open source relational database that is known for its stability and scalability, making it a popular choice for many web applications.

The database is hosted on a separate server and is accessed by the application using a connection string, which is stored as an environment variable. The database schema and tables are created using SQL migration scripts that are run when the application is deployed.

In the config directory of the project, you will find a database.yml file that contains the configuration settings for the database connection. The database.yml file specifies the host, port, username, password, and database name that the application uses to connect to the database.
# RESTful API
The server exposes a RESTful API to interact with the monitored system. The API is used by the monitoring client to report the status of the system and by the monitoring dashboard to display the status of the monitored systems.

/health
This endpoint returns the health status of the server.

URL

/health

Method:

GET

Success Response:

Code: 200 <br />
Content: { "status": "ok" }
/api/v1/systems
This endpoint is used to create and retrieve monitored systems.

URL

/api/v1/systems

Method:

POST | GET

URL Params

None

Data Params

POST request:

json
Copy code
{
  "name": "system-name",
  "url": "http://system-url.com"
}
name (required): The name of the monitored system.

url (required): The URL of the system to be monitored.

Success Response:

Code: 201 <br />
Content: { "id": 1, "name": "system-name", "url": "http://system-url.com", "status": "unknown", "last_checked": null }
OR

Code: 200 <br />
Content: [ { "id": 1, "name": "system-name", "url": "http://system-url.com", "status": "unknown", "last_checked": null } ]
Error Response:

Code: 400 BAD REQUEST <br />
Content: { "error": "name is required" }
OR

Code: 400 BAD REQUEST <br />
Content: { "error": "url is required" }
/api/v1/systems/{id}
This endpoint is used to retrieve and update a monitored system.

URL

/api/v1/systems/{id}

Method:

GET | PUT

URL Params

id (required): The ID of the system to be retrieved or updated.

Data Params

PUT request:

json
Copy code
{
  "name": "new-name",
  "url": "http://new-url.com"
}
name (optional): The new name of the monitored system.

url (optional): The new URL of the system to be monitored.

Success Response:

Code: 200 <br />
Content: { "id": 1, "name": "new-name", "url": "http://new-url.com", "status": "unknown", "last_checked": null }
Error Response:

Code: 404 NOT FOUND <br />
Content: { "error": "system not found" }
/api/v1/systems/{id}/status
This endpoint is used to retrieve the status of a monitored system.

URL

/api/v1/systems/{id}/status

Method:

GET

URL Params

id (required): The ID of the system to be checked.

Success Response:

Code: 200 <br />
Content: { "status": "up" }
OR

Code: 200 <br />
Content: { "status": "down" }
# Configuration
The config.json file contains the following parameters:

endpoints: a list of endpoints to monitor. Each endpoint should be a dictionary with the following keys:           
name: a descriptive name for the endpoint.          
url: the URL of the endpoint.           
headers (optional): a dictionary of headers to send with each request.           
interval (optional): the time interval between requests, in seconds. Default is 60 seconds.           
verify (optional): whether to verify SSL certificates. Default is true.              
timeout (optional): the timeout for requests, in seconds. Default is 10 seconds.               




