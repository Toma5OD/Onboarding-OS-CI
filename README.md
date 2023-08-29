# Onboarding OS-CI Repository

## Overview

This repository contains completed tasks and projects as part of the onboarding process for joining the OpenShift CI team. The tasks are based on a [structured onboarding document](https://docs.google.com/document/d/1WGH2cajZFkaR-dvPVIAZFUlgC5M0OhWLVTf-Zg9mI9w/edit#heading=h.37u1n9n9k12l) (Red Hat staff access only).

## Phases and Projects

### Phase 2 - Learn the Basics

- **[GNU Sort](./gnu-sort/)**: A Go program that reimplements GNU sort. Further details in the corresponding README file within the folder.
  
- **[HTTP Echo](./http-echo/)**: A Go HTTP server that serves as an echo server. Further details in the corresponding README file within the folder.

- **[Linux Shell Scripting Assignments](./Linux-Shell-Scripting-Assignments/)**: Various Bash scripts for different exercises. Further details in the corresponding README file within the folder.

### Phase 3 - Containers and Prow

- **[Docker](./docker/)**: Contains a Dockerfile and resources to containerize the HTTP Echo Server. Further details in the corresponding README file within the folder.

## How to Use

1. Fork this repository to your own GitHub account.
2. Clone your fork to your local machine: 
    ```bash
    git clone https://github.com/YOUR_USERNAME/Onboarding-OS-CI.git
    ```
3. Add the original repository as an upstream: 
    ```bash
    git remote add upstream https://github.com/Toma5OD/Onboarding-OS-CI.git
    ```
4. Fetch updates from upstream: 
    ```bash
    git fetch upstream
    ```
5. Check out your local `main` branch: 
    ```bash
    git checkout main
    ```
6. Merge updates from `upstream/main` into your local `main` branch: 
    ```bash
    git merge upstream/main
    ```
7. Navigate into the directory of the project you are interested in.
8. Follow the specific instructions in the corresponding README files within each folder.

### Submitting Pull Requests

- Create a new branch for the changes you want to make.
- Make your changes and commit them to your new branch.
- Push the changes to your fork on GitHub.
- Go to your fork on GitHub and create a new pull request for your changes.
- The pull request will be reviewed before it is merged into the `main` branch.

## License

This repository is licensed under the Apache License 2.0. The Apache License 2.0 is a permissive free software license written by the Apache Software Foundation. It allows users to use the software for any purpose, to distribute it, to modify it, and to distribute modified versions of the software under the terms of the license. For full details, please see the [LICENSE](./LICENSE) file.
