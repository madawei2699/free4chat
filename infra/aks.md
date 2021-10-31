
## Install

azure-cli
kubectl
helm
k9s

## Config

az login

az aks get-credentials --subscription 'i365-Pay-As-You-Go' --resource-group k8s-bmpi --name k8s-bmpi

cat ~/.kube/config

## Operate History

kubectl get namespaces

kubectl get nodes

kubectl describe node

kubectl get pods

kubectl create namespace free4chat

## k9s

mkdir ~/.k9s
k9s --kubeconfig ~/.kube/config