# CYDERES-challenge

**February 20, 2022**
Using Golang, I'm creating a simple API that is deployed as a serverless function on a cloud platform.


I decided to go with Serverless Framework from https://www.serverless.com/framework/ after looking through their documentation. Easy setup with AWS Lambda and start a serverless 
Go API. 
Started with: 'sls create --template aws-go-mod --path myService' which created a template directory with files that utilize Go Modules 

This is my first project utilizing Go, Serverless, Terraform. The only service I was familiar with was AWS Lambda. Coming from Frontend Engineering, I believe I could have done things better:

        Organization: Organizing the repository and file directories to have a more clean and readable project
        Testing: I have only utilized automated testing once with NodeJS, couple with learning Go, Go API's, serverless, and Terraform I was unable to achieve testing. 
        Experience: It is difficult to approach a problem with no experience in any of the tools required to solve it. So researching every step within the timeframe was                               challenging, yet rewarding to learn as much as I did.

Overall I'm happy with my implementation and am looking forward to building on top of the foundation this project has laid out for me.

**Starting with the environment**

        NodeJS v16.14.0

        npm v8.5.1   

        serverless Framework Core: 3.2.1  Plugin: 6.1.0   SDK: 4.3.1

        Terraform 1.1.6
        
 **Installation**

        chmod +x ./deploy.sh

        chmod +x ./destroy.sh

**Usage**
       
        To deploy run sh deploy.sh

        To destory run sh destroy.sh

**Testing**
        
        I've been using Postman to POST valid IPv4, IPv6, and URL's to my endpoint to get information back. 

        Some of the tests I used:
                8.8.8.8
                8.8.4.4
                139.130.4.5
                24.131.210.219
                105.249.217.234


                http://www.testingmcafeesites.com/testcat_al.html
                http://www.testingmcafeesites.com/index.html
                http://www.Wikia.com
                http://www.Example.com
