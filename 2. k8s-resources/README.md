# Shop Backend Kubernetes Deployment

This project defines a Kubernetes Deployment and CronJob to deploy and manage a shop backend API. The API runs on port 80 and uses the container image `strm/helloworld-http:latest`. The deployment is configured to meet the requirements provided.

---

## Features

1. **Deployment**:
   - Configured for **three replicas** to ensure high availability.
   - Uses the container image `strm/helloworld-http:latest`.
   - Exposes port **80** for incoming traffic.
   - Configured with **liveness and readiness probes**:
     - **Liveness Probe**: Ensures the container is running by checking TCP requests to port 80.
     - **Readiness Probe**: Ensures the application is ready to serve traffic by sending HTTP requests to the root endpoint (`/`) on port 80.

2. **CronJob**:
   - Outputs the message **"Hello SRE"** every 30 minutes.
   - Uses the container image `busybox` with a simple shell command.

---

## File Structure

- **`values.yaml`**: Contains configurable parameters for the Helm chart, such as replica count, image details, and liveness/readiness probe settings.
- **Helm Templates**:
  - `deployment`: Defines the deployment resource with the container specifications and probes.
  - `cronjob`: Defines the CronJob resource with the schedule and command.

---

## Configuration

You can modify the following parameters in the `values.yaml` file:

### **Deployment Parameters**
| Parameter               | Default Value                  | Description                              |
|-------------------------|--------------------------------|------------------------------------------|
| `deployment.replicaCount` | `3`                          | Number of pod replicas.                  |
| `deployment.port`        | `80`                         | Port exposed by the container.           |
| `deployment.image.repository` | `strm/helloworld-http`  | Container image repository.              |
| `deployment.image.tag`   | `latest`                     | Container image tag.                     |
| `deployment.image.imagePullPolicy` | `IfNotPresent`   | Image pull policy.                       |
| `deployment.livenessProbe.initialDelaySeconds` | `10` | Initial delay for liveness probe.        |
| `deployment.livenessProbe.periodSeconds` | `10`       | Liveness probe check frequency.          |
| `deployment.readinessProbe.initialDelaySeconds` | `10` | Initial delay for readiness probe.       |
| `deployment.readinessProbe.periodSeconds` | `1`       | Readiness probe check frequency.         |
| `deployment.readinessProbe.failureThreshold` | `2`    | Maximum failures for readiness probe.    |
| `deployment.readinessProbe.path` | `/`                | Path to check readiness.                 |

### **CronJob Parameters**
| Parameter               | Default Value                  | Description                              |
|-------------------------|--------------------------------|------------------------------------------|
| `cronjob.name`           | `sre-job`                    | Name of the CronJob.                     |
| `cronjob.image`          | `busybox`                    | Container image for the CronJob.         |
| `cronjob.schedule`       | `*/30 * * * *`               | Cron schedule (every 30 minutes).        |
| `cronjob.command`        | `["/bin/sh", "-c", "echo Hello SRE"]` | Command executed by the CronJob. |
| `cronjob.restartPolicy`  | `OnFailure`                  | Restart policy for the CronJob.          |

---

## Usage

### Prerequisites:
- Kubernetes cluster (any cloud provider or local setup like Minikube).
- Helm CLI installed and configured.

### Deployment:
```bash
helm install shop-backend .
```
Verify the deployment:

```bash
kubectl get deployments
kubectl get pods
```

Verify the cronJob:
```bash
kubectl get cronjobs
```

Check the logs of the cron pod:

```bash
kubectl logs $cron_pod
```
