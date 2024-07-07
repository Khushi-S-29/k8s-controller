# Custom  Kubernetes Controller

A project to write a Kubernetes controller in GO
Team members:
Khushi S.
Srishti Dutta
Meghna Mandawra



## Overview
This project implements a custom Kubernetes controller designed to manage CronJobs within a Kubernetes cluster. The controller ensures that CronJobs are created, updated, or deleted based on defined specifications, providing automated management of scheduled tasks.

## Features

- Automated Management:   Watches for changes to CronJob resources and automatically reconciles them with the desired state.
- Flexible Scheduling:   Supports flexible scheduling configurations using cron expressions.
- Error Handling:   Implements robust error handling and reconciliation mechanisms to maintain system reliability.
- Scalability: Designed to handle multiple CronJobs and scale with cluster demands.
