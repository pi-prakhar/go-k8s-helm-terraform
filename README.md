# Go-K8s-Helm-Terraform

This project allows you to deploy two services (`server1` and `server2`) on a Google Kubernetes Engine (GKE) cluster with an NGINX Ingress Controller using Terraform and Helm. The setup process involves setting up Google Cloud, creating the GKE cluster using Terraform, installing the NGINX Ingress Controller, and deploying your workload using Helm.

## Prerequisites

Before you start, ensure you have the following installed on your local machine:

- **Google Cloud SDK**: For interacting with Google Cloud and GKE.
- **Terraform**: To create and manage the GKE cluster.
- **Helm**: To manage Kubernetes packages (charts).
- **kubectl**: To interact with your Kubernetes cluster.
- **git**: To clone this repository.

## Setup Google Cloud

Follow the steps below to set up Google Cloud on your local machine:

1. **Install Google Cloud SDK**:
   - Follow the [official guide](https://cloud.google.com/sdk/docs/install) to install the Google Cloud SDK.

2. **Authenticate with Google Cloud**:
   ```bash
   gcloud auth login
   ```

3. **Set your Google Cloud project**:
   Set your Google Cloud project ID (replace `<PROJECT_ID>` with your actual project ID):
   ```bash
   gcloud config set project <PROJECT_ID>
   ```

4. **Set the default region and zone**:
   For example, set the region to `us-central1` and zone to `us-central1-a`:
   ```bash
   gcloud config set compute/region us-central1
   gcloud config set compute/zone us-central1-a
   ```

## Creating the GKE Cluster

Once the prerequisites are set up, navigate to the `/deployments/terraform` directory:

```bash
cd deployments/terraform
```

### Terraform Commands

1. **Initialize Terraform**:
   This will download the necessary plugins and providers.
   ```bash
   terraform init
   ```

2. **Plan the Infrastructure**:
   Terraform will show you what changes will be made. Review the plan.
   ```bash
   terraform plan
   ```

3. **Apply the Plan**:
   Apply the changes to create the GKE cluster.
   ```bash
   terraform apply
   ```

   After applying, Terraform will output the necessary information about your GKE cluster.

### Verifying Cluster Health

Once the cluster is created, you can check its health by running the following command:

```bash
kubectl get nodes
```

This will show the nodes in your GKE cluster. Ensure all nodes are in a `Ready` state.

## Installing the NGINX Ingress Controller

To deploy the NGINX Ingress Controller, follow these steps:

1. **Add the Ingress NGINX Helm Chart**:
   ```bash
   helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
   helm repo update
   ```

2. **Install the NGINX Ingress Controller**:
   ```bash
   kubectl create namespace ingress-nginx
   helm install ingress-nginx ingress-nginx/ingress-nginx -n ingress-nginx
   ```

   This will deploy the NGINX Ingress Controller in the `ingress-nginx` namespace.

## Deploying the Workload

Now, let's deploy the two services (`server1` and `server2`) to the GKE cluster using Helm.

1. **Navigate to the Helm directory**:
   ```bash
   cd ../helm
   ```

2. **Install the Helm Chart**:
   Deploy your workload (server1 and server2) using the following Helm command:
   ```bash
   helm install go-test . -n go-test --create-namespace
   ```

3. **Verify the Deployment**:
   You can check the status of your services with:
   ```bash
   kubectl get all -n go-test
   ```

   This will list all the resources deployed in the `go-test` namespace, including the pods, services, and deployments.

## Checking External IP of Ingress

To access the services externally through the NGINX Ingress, you need to find the external IP of the Ingress controller:

1. **Check the External IP**:
   ```bash
   kubectl get svc -n ingress-nginx
   ```

   Look for the `EXTERNAL-IP` of the NGINX Ingress service.

2. **Access the Services**:
   Once you have the external IP (e.g., `X.X.X.X`), you can access the services:

   - `server1`: `https://<EXTERNAL-IP>/server1`
   - `server2`: `https://<EXTERNAL-IP>/server2`

   Replace `<EXTERNAL-IP>` with the actual IP you obtained from the previous step.

## Clean Up

To clean up resources after you are done, you can delete the Helm release and the Terraform-managed infrastructure.

1. **Delete the Helm Release**:
   ```bash
   helm uninstall go-test -n go-test
   ```

2. **Delete the Terraform-managed Infrastructure**:
   Navigate to the `deployments/terraform` directory and run:
   ```bash
   terraform destroy
   ```

   This will destroy the GKE cluster and all associated resources.

## Troubleshooting

- **GKE Cluster Not Found**: Ensure you have properly authenticated with Google Cloud and set the correct project, region, and zone.
- **Ingress Controller Not Working**: Double-check the NGINX Ingress installation and make sure the service has an external IP.
- **Helm Errors**: Make sure the Helm repository is added and updated before trying to install charts.

## Conclusion

This project provides an automated way to deploy `server1` and `server2` on a GKE Kubernetes cluster with an NGINX Ingress Controller. By following the steps outlined, you can easily manage the infrastructure and services using Terraform and Helm.

For further questions or contributions, feel free to open an issue or submit a pull request.
