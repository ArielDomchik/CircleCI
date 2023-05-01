
# CircleCI
This project is on build..


Flask Application Deployment on EKS Cluster with CircleCI, Terraform and ArgoCD

This project is an example of deploying a Flask web application on an EKS cluster using CircleCI, Terraform and ArgoCD.


**Prerequisites**

  -  AWS account with permissions to create an EKS cluster
  -  Terraform installed on your local machine
  -  Terraform-aws-eks-blueprints cloned and present on your local machine
  -  GitLab account for container registry
  -  Discord server and webhook for notifications

**Application Architecture**

The application is a multi-container application that consists of a Flask web application, Gunicorn server and an Nginx service to reverse proxy to Gunicorn to retrieve the Flask web app. The application is tested using Selenium E2E tests.


**Deployment Pipeline**


The deployment pipeline is implemented using CircleCI and consists of the following stages:

  -  Build Docker Image: Building a Docker image of the application and pushing it to the GitLab container registry.

 -   Run Selenium Tests: Running Selenium E2E tests on the deployed application to ensure it is working as expected.

-    Deploy to EKS Cluster: Creating an EKS cluster using Terraform and deploying the application to the cluster. The Terraform configuration files are checked using Teratest and Checkov. Discord notifications are sent if the tests pass or fail.

 -   Deploy with ArgoCD: Deploying the application using ArgoCD, which manages the application by syncing changes from the repository.

-    Test Deployment: Running chaos tests using Netflix Chaos Monkey to check the system resilience. When a node goes down, the deployed Karpenter (with Terraform) controls the auto-scaling group and keeps the desired size of nodes even if some nodes go down spontaneously.

**Terraform Configuration**

The Terraform configuration files can be found in the terraform-eks-blueprints directory. They are used to create an EKS cluster and deploy the application to it.

**Outside of Deployment Pipeline**

    Deploy ArgoCD: Deploying ArgoCD in the EKS cluster to manage the deployment of the application.

**Conclusion**

This project demonstrates how to deploy a Flask web application on an EKS cluster using CircleCI, Terraform and ArgoCD. By using these tools, the deployment process can be automated, which makes it easier and faster to deploy applications.


 -  dummy object is for kubernetes cluster, keda addon is returning an empty response for metrics server, this will prevent it
