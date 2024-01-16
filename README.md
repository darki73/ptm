# Proxmox Templates Maker
This application allows you to automate templates creation for Proxmox VE.  
It provides a simple way to create templates for Proxmox VE by using a simple YAML file or command line arguments.  

# Table of Contents
- [Proxmox Templates Maker](#proxmox-templates-maker)
- [Installation](#installation)
- [Configuration](#configuration)
    * [Downloader Configuration](#downloader-configuration)
    * [Base Image Configuration](#base-image-configuration)
    * [Extra Packages Configuration](#extra-packages-configuration)
    * [Repositories Configuration](#repositories-configuration)
    * [Qemu Configuration](#qemu-configuration)
    * [Cloud-Init Configuration](#cloud-init-configuration)
    * [Unattended Upgrades Configuration](#unattended-upgrades-configuration)
    * [Minimal Configuration](#minimal-configuration)
- [Usage](#usage)
    * [Commands](#commands)
    * [Customize](#customize)
    * [Make](#make)
        + [Prompt Flow](#prompt-flow)
        + [Configuration Flow](#configuration-flow)
        + [Flags Flow](#flags-flow)

# Installation
You can install the application by downloading the latest version from [Releases](https://github.com/darki73/ptm/releases) page.

# Configuration
By default, application expects configuration file to be located at `/etc/ptm/config.yaml`.  
You can override:
1. Configuration file location by using `--configuration-path` argument. (defaults to `/etc/ptm`)
2. Configuration file name by using `--configuration-name` argument. (defaults to `config`)
3. Configuration file extension by using `--configuration-extension` argument. (defaults to `yaml`)

Below, you would be able to find examples of configuration file sections.  
**All examples are using YAML format.**  
***Most of them are presented with the default values that application would use if no configuration is provided.***

## Downloader Configuration
Downloader configuration is located under `downloader` key.  
Downloader is responsible for downloading and saving images from remote location.  

```yaml
downloader:
  save_to: /etc/ptm/images
```

**Keys:**  
- `save_to` - path to directory where downloaded images should be saved. (defaults to `/etc/ptm/images`)

## Base Image Configuration
Base image configuration is located under `base_image` key.  
It is responsible for providing information on what type of image you want to download.  

```yaml
base_image:
  distribution: ubuntu
  release: jammy
  minimal: true
  architecture: amd64
  format: img
```

**Keys:**
- `distribution` - distribution of the image. (defaults to `ubuntu`)
- `release` - release of the image. (defaults to `jammy`)
- `minimal` - whether image should you want to download the "minimal" version of the image. (defaults to `true`)
- `architecture` - architecture of the image. (defaults to `amd64`)
- `format` - format of the image. (defaults to `img`)

## Extra Packages Configuration
Extra packages configuration is located under `extra_packages` key.  
It is responsible for providing information on what extra packages should be installed on the image.  
It is a list of strings, which are package names.  


```yaml
extra_packages: []
```

## Repositories Configuration
Repositories configuration is located under `repositories` key.  
It is responsible for providing information on what repositories should be added to the image.  

```yaml
repositories:
  - name: docker
    gpg: https://download.docker.com/linux/ubuntu/gpg
    url: https://download.docker.com/linux/ubuntu
    release: jammy
    component: stable
    key_name: docker-archive-keyring
```

**Keys:**
- `name` - name of the repository. (used to generate lists file name in `/etc/apt/sources.list.d/`)
- `gpg` - URL to GPG key file.
- `url` - URL to repository.
- `release` - release of the repository.
- `component` - component of the repository.
- `key_name` - name of the key file in `/usr/share/keyrings/`. (extension is added automatically)

## Qemu Configuration
Qemu configuration is located under `qemu` key.  
It is responsible for providing information on what qemu options should be used when creating the template.  
Most of the time, you would use this configuration to skip some of the configuration steps to speed up the process.  
If any of the options are not provided, application would prompt you to provide them during the process.  
If all of the options are provided, application would skip the configuration step and use provided values.

```yaml
qemu:
  identifier: 9000
  name: ubuntu-cloudinit
  image: /etc/ptm/images/ubuntu-21.10-minimal-amd64.img
  resources:
    cores: 2
    memory: 2G
    cpu_type: host
  storage:
    name: local-lvm
    resize: 4G
  network:
    driver: virtio
    bridge: vmbr0
```

**Keys:**
- `identifier` - identifier of the template.
- `name` - name of the template.
- `image` - path to the image file.
- `resources` - resources configuration.
  - `cores` - number of cores.
  - `memory` - amount of memory.
  - `cpu_type` - type of CPU.
- `storage` - storage configuration.
  - `name` - name of the storage. (refers to the storage name in Proxmox VE, for example, `local-lvm`)
  - `resize` - amount of disk space to allocate for the template. (used to resize the image file)
- `network` - network configuration.
  - `driver` - driver to use for the network interface.
  - `bridge` - bridge to use for the network interface.

## Cloud-Init Configuration
Cloud-Init configuration is located under `cloud_init` key.  
It is responsible for providing information on what cloud-init configuration should be used when creating the template.  
Most of the time, you would use this configuration to skip some of the configuration steps to speed up the process.  
If any of the options are not provided, application would prompt you to provide them during the process.  
If all of the options are provided, application would skip the configuration step and use provided values.  
All of them are optional, but if `enabled` is set to `true`, application will "autoconfigure" IPv4 and IPv6 addresses. (DHCP)

```yaml
cloud_init:
  enabled: true
  username: administrator
  password: 12345678
  ssh_authorized_keys:
    - /root/.ssh/administrator.pub
    - ssh-rsa AAAAB3NzaC1yc2EAAA
  network:
    ipv4:
      auto_configure: true
      ip: 127.0.0.10/24
      gateway: 127.0.0.1
    ipv6:
      auto_configure: true
      ip: ::1/128
      gateway: ::1
```

**Keys:**
- `enabled` - whether cloud-init should be enabled.
- `username` - username of the user that should be created.
- `password` - password of the user that should be created.
- `ssh_authorized_keys` - list of paths to SSH public keys that should be added to the user.
  - You can provide keys as a string.
  - You can provide keys as a path to a file.
- `network` - network configuration.
  - `ipv4` - IPv4 configuration.
    - `auto_configure` - whether IPv4 should be autoconfigured. (manual settings for `ip` and `gateway` will be ignored if set to `true`)
    - `ip` - IPv4 address.
    - `gateway` - IPv4 gateway.
  - `ipv6` - IPv6 configuration.
    - `auto_configure` - whether IPv6 should be autoconfigured. (manual settings for `ip` and `gateway` will be ignored if set to `true`)
    - `ip` - IPv6 address.
    - `gateway` - IPv6 gateway.

## Unattended Upgrades Configuration
Unattended Upgrades configuration is located under `unattended_upgrades` key.  
It is responsible for providing information on what unattended upgrades configuration should be used when customizing the image.  
If not present, application will not try to setup unattended upgrades.  

```yaml
unattended_upgrades:
  enabled: true
  whitelist:
    - "${distro_id}:${distro_codename}"
    - "${distro_id}:${distro_codename}-security"
    - "${distro_id}ESMApps:${distro_codename}-apps-security"
    - "${distro_id}ESM:${distro_codename}-infra-security"
    - "${distro_id}:${distro_codename}-updates"
    - "Docker:${distro_codename}"
    - "InfluxDB:stable"
  blacklist: []
  dev_release: "auto"
  fix_interrupted: true
  minimal_steps: true
  install_on_shutdown: false
  remove_unused_kernel: true
  remove_unused_dependencies: true
  remove_unused_auto_depend: true
  automatic_reboot: true
  automatic_reboot_with_users: true
  automatic_reboot_time: "04:00"
```

**Keys:**
- `enabled` - whether unattended upgrades should be enabled.
- `whitelist` - list of origins that should be whitelisted.
- `blacklist` - list of packages that should be blacklisted.
- `dev_release` - whether unattended upgrades should use development release.
- `fix_interrupted` - whether unattended upgrades should fix interrupted upgrades.
- `minimal_steps` - whether unattended upgrades should use minimal steps.
- `install_on_shutdown` - whether unattended upgrades should install updates on shutdown.
- `remove_unused_kernel` - whether unattended upgrades should remove unused kernel.
- `remove_unused_dependencies` - whether unattended upgrades should remove unused dependencies.
- `remove_unused_auto_depend` - whether unattended upgrades should remove unused auto dependencies.
- `automatic_reboot` - whether system should automatically reboot.
- `automatic_reboot_with_users` - whether system should automatically reboot with users.
- `automatic_reboot_time` - time when system should automatically reboot.

## Minimal Configuration
Although application does not require you to have any configuration, it is still required to have configuration file created.

# Usage
This section describes how to use the application.

## Commands
The following commands are available:
- `version` - displays application version.
- `help` - displays help message.
- `customize` - allows user to customize the image.
- `make` - allows user to create the template.

## Customize
This command allows you to customize the image.  
It will ask you for all the required information and then it will download (if missing) and customize the image.  
You can use `--help` argument to display help message for this command.  

## Make
This command allows you to create the template.  
It will ask you for all the required information and then it will create the template.
You can use `--help` argument to display help message for this command.  

### Prompt Flow
In prompt flow mode (assuming you have not provided flags or configuration options to make configuration valid), application will ask you for all the required information.  
It will guide you through the process and ask you for all the required information.  
Upon completion, it will create the template.  

### Configuration Flow
In configuration flow mode (assuming you have provided complete configuration - this includes `qemu` and `cloud_init` configuration), application will skip the prompt flow and use provided configuration.  
It will create the template using provided configuration.

### Flags Flow
In flags flow mode (assuming you have provided flags), application will skip the prompt flow and use provided flags.  
The caveat of this flow is that you must provide almost all of the flags.  
It will create the template using provided flags.  
It is ideal for automation purposes.

**Here is the list of flags that you must provide:**
- `--identifier` - Identifier of template ***(required)***
- `--name` - Name of the template ***(required)***
- `--cores` - Number of cpu cores ***(required)***
- `--cpu-type` - Desired cpu type (host / kvm64 / etc) ***(required)***
- `--memory` - Amount of memory (example: 1024 / 1024M / 1G) ***(required)***
- `--storage` - Disk storage (local-lvm / local / etc) ***(required)***
- `--image` - Path to the image (/etc/ptm/images/image.qcow2) ***(required)***
- `--mage-new-size` - Size to which the image should be resized (example: 4G) *(optional)*
- `--network-driver` - Network driver (virtio / e1000 / etc) ***(required)***
- `--network-bridge` - Network bridge (vmbr0 / vmbr1 / etc) ***(required)***
- `--ci-username` - Username for cloud-init *(optional)*
- `--ci-password` - Password for cloud-init *(optional)*
- `--ci-ssh-keys` - Comma-separated list of SSH keys for cloud-init *(optional)*
- `--ci-ipv4-auto` - Automatically configure IPv4 for cloud-init
- `--ci-ipv6-auto` - Automatically configure IPv6 for cloud-init
- `--ci-ipv4-address` - Manually set IPv4 address for cloud-init (example: 10.10.10.10/24) (not required when `--ci-ipv4-auto` flag is used)
- `--ci-ipv6-address` - Manually set IPv6 address for cloud-init (example: 2001:db8::1/64) (not required when `--ci-ipv4-auto` flag is used)
- `--ci-ipv4-gateway` - Manually set IPv4 gateway for cloud-init (example: 10.10.10.1) (not required when `--ci-ipv6-auto` flag is used)
- `--ci-ipv6-gateway` - Manually set IPv6 gateway for cloud-init (example: 2001:db8::1) (not required when `--ci-ipv6-auto` flag is used)