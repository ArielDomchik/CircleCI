- GitOps-dev

This repository contains the development version of the application that is deployed using ArgoCD. When a new release is available,
the values.yaml file in the private repository is replaced, and the chart in the GitOps-prod repository is updated using the workflow we created.


- Usage

To use this repository, you can clone it and start developing the application. Once you have made changes to the application, you can create a pull request and merge it to the main branch.
This will trigger the GitHub Actions workflow, which will update the values.yaml file in the private repository and update the chart in the GitOps-prod repository.
