## Install

Generate your sshkey for the translator server
```bash
ssh-keygen
```

Generate your sshkey for the monitoring server. You can use also the same.
```bash
ssh-keygen
```

Change the correct paths in the specific terraform file.
The default SSH_KEY_PATH points to the project route path but can be adjusted in the `secrets.tfvars` file
