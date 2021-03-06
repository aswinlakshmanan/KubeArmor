# Deployment Guide

1. Deploy a custom resource definition

    Kubernetes cannot recognize the specification of security policies for KubeArmor unless you register the custom resource definition for KubeArmor's security policy. Thus, you first need to apply the custom resource definition for KubeArmor's security policy into Kubernetes.

    ```
    $ cd deployments/CRD
    (deployments/CRD) $ kubectl apply -f .
    ```

2. Deploy KubeArmor in your Kubernetes environment

    KubeArmor currently supports a bare-metal environment and Google Kubernetes Engine (GKE). It will support Amazon Elastic Kubernetes Service (EKS) and Azure Kubernetes Service (AKS) soon.

    - Deploy KubeArmor in a bare-metal environment (Docker)

        ```
        $ cd deployments/generic-docker
        (deployments/generic-docker) $ kubectl apply -f .
        ```

    - Deploy KubeArmor in a bare-metal environment (Containerd)

        ```
        $ cd deployments/generic-containerd
        (deployments/generic-containerd) $ kubectl apply -f .
        ```

    - Deploy KubeArmor in MicroK8s

        ```
        $ cd deployments/microk8s
        (deployments/microk8s) $ kubectl apply -f .
        ```

    - Deploy KubeArmor in GKE

        ```
        $ cd deployments/GKE
        (deployments/GKE) $ kubectl apply -f .
        ```

    - Deploy KubeArmor in EKS

        ```
        Coming soon
        ```

    - Deploy KubeAmor in AKS

        ```
        Coming soon
        ```
