# How to use kubectl & k9s to connect to a kubernetes cluster on AWS EKS

- <https://kubernetes.io/docs/reference/kubectl/>

```bash
# Install kubectl
# see documentation

# commands
aws eks update-kubeconfig --name <cluster-name> --region <region>
kubectl config use-context <cluster-name>
kubectl cluster-info
kubectl get nodes
kubectl get pods --all-namespaces
kubectl apply -f <file-name>
```

- <https://k9scli.io/>
