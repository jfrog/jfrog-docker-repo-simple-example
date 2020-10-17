# Getting started with JFrog Docker Repos

This Dockerfile assumes the following:
* A local docker repository with the name "docker-dev-local" has been created in Artifactory
* A remote docker repository with the name "docker-hub-remote" has been created in Artifactory
* A virtual docker repository with the name "docker" has been created in Artifactory and includes the "docker-dev-local" and "docker-hub-remote" repositories with the "docker-dev-local" repository set as the Default Deployment Repository.

To learn more about JFrog Platform integration with Docker, visit [Docker Registry Documentation](https://www.jfrog.com/confluence/display/JFROG/Docker+Registry).

