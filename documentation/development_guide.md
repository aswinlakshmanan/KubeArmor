# Development Guide

1. Bare-metal Environment

  - Requirements

    Here is the list of requirements for a bare-metal environment.

    ```
    Kubernetes - v1.19 (or newer)
    Docker - 18.03 (or newer) / Containerd - 1.3.7 (or newer)
    Linux Kernel - v4.15 (or newer)
    LSM - AppArmor (+ SELinux and KRSI)
    ```

    First of all, KubeArmor is designed for Kubernetes, which means that Kubernetes should be ready in your environment. If Kubernetes is not prepared yet, please refer to [Kubernetes installation guide](./k8s_installation_guide.md). KubeArmor also requires Docker or Containerd since it internally uses their APIs. If you have other container platforms (e.g., Podman), please make an issue in this repository. While we are going to adopt other container platforms in KubeArmor, we may be able to adjust the priorities of our planned tasks on demand. KubeArmor requires LSMs to operate properly; thus, please make sure that your environment supports LSMs (at least, AppArmor).

    <font color=red>Note that we may not be able to run KubeArmor if you are using MiniKube because MiniKube does not support AppArmor. In addition, KubeArmor does not work if you use Docker Desktops on Windows and MacOS. It is because such environments do not have a full Linux file system.</font>

  - Environmetal Setup

    In order to install all dependencies, please run the following command.

    ```
    $ cd contributions/bare-metal
    (bare-metal) $ ./setup.sh
    ```

    [setup.sh](../contributions/bare-metal/setup.sh) will automatically install BCC (latest), Go (v1.15.2), AppArmor (with AppArmor-utils), and Auditd.

    Now, you are ready to develop the code for KubeArmor. Please enjoy the journey with KubeArmor.

  - (Optional) MicroK8s Setup

    In order to install MicroK8s, please run the following command.

    ```
    $ cd contributions/bare-metal/microk8s
    (microk8s) $ ./install_microk8s.sh
    ```

2. Vagrant Environment

  - Requirements

    Here is the list of requirements for a Vagrant environment

    ```
    Vagrant - v2.2.9
    VirtualBox - v6.0
    ```

    If you do not have Vagrant and VirtualBox in your environment, you can easily install them by running the following command.

    ```
    cd contributions/vagrant
    (vagrant) $ ./setup.sh
    ```

  - VM Setup using Vagrant

    Now, it is time to create a VM for development.

    ```
    cd contributions/vagrant
    (vagrant) $ ./create.sh
    ```

    You can directly use the vagrant command to create a VM.

    ```
    (vagrant) $ vagrant up
    ```

    In this case, please make sure that you have ssh keys in '~/.ssh'. If you do not have ssh keys yet, please run the following command before running the above command (i.e., 'vagrant up').

    ```
    (vagrant) $ ssh-keygen -> [Enter] -> [Enter] -> [Enter]
    ```

    If you want to remove the created VM, please run the following command.

    ```
    cd contributions/vagrant
    (vagrant) $ ./remove.sh
    ```

    You are ready to develop the code for KubeArmor. Please enjoy the journey with KubeArmor.

    ```
    cd contributions/vagrant
    (vagrant) $ vagrant ssh
    ```

# Code Directories

Here, we briefly give you the overview of KubeArmor's directories.

  - Source code for KubeArmor (/KubeArmor)

    ```
    KubeArmor/
      core        - The main body (start point) of KubeArmor
      audit       - Audit Logger (getting audit logs and sending them to somewhere)
      discovery   - Automated security policy discovery (under development)
      enforcer    - Runtime policy enforcer (enforcing security policies into LSMs)
      feeder      - gRPC-based feeder (sending audit/system logs to a log server)
      monitor     - eBPF-based container monitor (mapping process IDs to container IDs)
      BPF         - eBPF code for container monitor
      log         - Message logger (stdout) for KubeArmor
      common      - Libraries internally used
      types       - Type definitions
    ```

  - Source code for KubeArmor's log server

    ```
    LogServer     - gRPC-based log server
    protobuf      - Protocol buffer
    ```

  - Source code for KubeArmor's custom resource defintion (CRD)

    ```
    pkg/
      k8s         - CRD Code generated by Kube-Builder
    ```

  - Configurations for AppArmor and Auditd

    ```
    KubeArmor/
      AppArmor    - Script files to create and delete AppArmor profiles
      Auditd      - Configurations for Auditd
    ```

  - Scripts for GKE

    ```
    KubeArmor/
      GKE         - scripts to set up the enforcer in a container-optimized OS (COS)
    ```

  - Files for testing

    ```
    KubeArmor/
      build       - Scripts to create a container image for KubeArmor
    examples/
      multiubuntu - Example microservice for testing
    ```
