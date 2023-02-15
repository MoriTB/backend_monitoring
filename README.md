# backend_monitoring
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

# Configuration
The config.json file contains the following parameters:

endpoints: a list of endpoints to monitor. Each endpoint should be a dictionary with the following keys:           
name: a descriptive name for the endpoint.          
url: the URL of the endpoint.           
headers (optional): a dictionary of headers to send with each request.           
interval (optional): the time interval between requests, in seconds. Default is 60 seconds.           
verify (optional): whether to verify SSL certificates. Default is true.              
timeout (optional): the timeout for requests, in seconds. Default is 10 seconds.               




