# Translator
This project is a study for the Cloud Computing 2 module. 
It offers a simple translator with a web interface (using HTMX) and an API endpoint for translation.

The application uses a PostgreSQL database as a cache to store the query hash, the target language, and the corresponding translation.
The Google Cloud Platform API is utilized for the translation process.
You can find more information about it at [https://github.com/googleapis/google-cloud-go/tree/main/translate](https://github.com/googleapis/google-cloud-go/tree/main/translate)


## Install
1. Clone the repository
```bash
git clone https://github.com/MetaEMK/DHBW-GCP-Translate
```

2. Define all necessary parameters for the application in `./ansible/templates/config.yaml.j2`

3. Create the ssh keys for the translator server and the monitoring server
```bash
ssh-keygen
# default names are: translator / monitoring
# You can adjust them in the terraform script if you want to use different names or the same key for both servers
```

4. Adjust the `./terraform/secrets.tfvars` file with the necessary information

5. Run the following commands to deploy the application

```bash
cd terraform
terraform init
terraform apply --var-file=secrets.tfvars
```

6. Fill in the correct ip address in the `./ansible/node_inventory.ini` file
7. Run the following command to deploy the application
```bash
# You may need to adjust the private key path
ansible-playbook -i node_inventory.ini
ansible-playbook -i node_inventory.ini
```
