# Custom  Kubernetes Controller

A project to write a Kubernetes controller in GO
Team members:
Khushi S.
Srishti Dutta
Meghna Mandawra



## Overview

This project implements a Kubernetes CronJob controller in Go, designed to automate the creation, management, and scheduling of CronJobs within a Kubernetes cluster. CronJobs are scheduled tasks that run periodically based on a specified schedule, similar to cron jobs in Unix-like systems.

The controller watches for changes to custom resources defined as CronJobs and ensures that the actual state of CronJobs matches the desired state specified in these resources. It leverages Kubernetes client libraries to interact with the Kubernetes API server, facilitating seamless integration and management of scheduled jobs.

## Key Features:

- Automatically manages creation and deletion of CronJobs based on custom resource definitions.
- Supports defining CronJob schedules and other configuration parameters through custom resources.
- Implements a reconciliation loop to handle updates and ensure desired state convergence.
- Provides logging and basic error handling for operational visibility and reliability.
