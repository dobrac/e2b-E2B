# Orchestration Services
Backend services for handling VM sessions, environments' pipeline, and the API for them.

## Development
Architecture - https://www.figma.com/file/pr02o1okRpScOmNpAmgvCL/Architecture

### GCE Terraform troubleshooting
If you get stuck with a non-working Nomad on all VMs and therefore you are not able to change the infrastructure because the Nomad jobs that are supposedly running there cannot be reached use the:
```
sudo nomad agent -dev -bind 0.0.0.0 -log-level INFO
```

to start temporary Nomad on any server and tear it all down with `terraform apply -destroy`.

If the Nomad is running but it has no leader, delete Load balancer and Instance group then destroy the infrastructure.

If you have problems with variables when redeploying from scratch (missing api key variables in hcl/nomad jobs, etc) comment out the api in the top `main.tf` file and deploy it after you deploy all the other infrastructure.

### Subtrees
#### firecracker-task-driver
The subtree commands you need for controling this repo are:
```bash
git subtree add --prefix cluster/disk-image/firecracker-task-driver https://github.com/devbookhq/firecracker-task-driver.git master
```

```bash
git subtree pull --prefix cluster/disk-image/firecracker-task-driver https://github.com/devbookhq/firecracker-task-driver.git master
```

```bash
git subtree push --prefix cluster/disk-image/firecracker-task-driver https://github.com/devbookhq/firecracker-task-driver.git master
```

#### shared
FC task driver is a subtree made from https://github.com/devbookhq/shared repository.

The subtree commands you need for controling this repo are:
```bash
git subtree add --prefix shared https://github.com/devbookhq/shared.git master
```

```bash
git subtree pull --prefix shared https://github.com/devbookhq/shared.git master
```

```bash
git subtree push --prefix shared https://github.com/devbookhq/shared.git master
```
