## Team Geauxinue:
Siddhant Keshkar, Biljith Thadichi, Aakarsh Fadnis, Abhas Prasad

## MarkMyLink Overview
This app will be used to organize bookmarks as a service. A user will be able to log in and view all their bookmarks in a dashboard. The dashboard will display all the statistics about the bookmark links usage, access time, summary, and, recommendations.

## Features

This application will automatically categorize the bookmarks into different bins like Social, Education, Work, Entertainment.  In the dashboard, the user will be able to view the summary of each bookmark. The user will be able to filter bookmarks with different categories.
Depending on the user's bookmark usage the application will recommend other similar websites to view. The user will also be able to set reminders to read specific bookmarks and the application will remind them at that time. The application will also recommend a bookmark as a suggestion to view when asked for. The user will be able to sort their bookmarks on the dashboard as well. The dashboard will also provide an option to search from different bookmarks as they are added.

## Technologies used

1. Heroku app deploy
2. Golang web app
3. RabbitMQ
4. Google Kubernetes Engine

## Deployment details
We have set up continuous deployment on Google App Engine as described below :
1. We used the Google cloud build github application to link our repo with Google cloud build. 
2. Google cloud build deployed the application to app engine using instructions in cloudbuild-noloop.yaml file. 
3. Whenever a push into git repo is made, it triggers a build in Google cloud engine and deploys the system accordingly. 

### ProcessLinkWorker Service
https://github.com/Casecarsid/ProcessLinkWorker


Points of Focus
<br>
| Area | Product/Service used |
| ----------- | --------- |
| Web Application used | Native GoLang  |
| Data Collection Batch Worker | NA |
| Data Analyzer | LinkPreview and Klazify |
| Unittests | Go tests (available in this project wiki) |
| Data Persistence | MongoDB Atlas (available in this project wiki) |
| API Endpoints | LinkPreview, Klazify API and REST API for the web service we have deployed on Heroku |
| Product Environment | Heroku Deployment Service |
| Integration Testing | NA |
| Mock Objects | Mocked Bookmark objects for categorizing and describing bookmarks |
| Continuous Integration | Heroku CI |
| Production Monitoring | Heroku |
| Acceptance Tests | Demo video /Acceptance test provided in presentation |
| Event Collaboratiob Messaging | RabbitMQ (details available in this wiki) |
| Continuous Delivery | Heroku CD |
| Participation | details available in wiki |

