
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
k9s --kubeconfig ~/.kube/k8s-bmpi-kubeconfig.yml

## Refference Links

- [Deploy to AKS from GitHub Actions | Thomas Stringer](https://trstringer.com/deploy-to-aks-from-github-actions/)

## HTTP

helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm repo update
helm install ingress-nginx ingress-nginx/ingress-nginx -f nginx-ingress.yaml --create-namespace --namespace ingress-nginx

kubectl apply -f echo-svc.yaml # create echo service for test

kubectl apply -f ingress-free4-chat.yaml # create ingress for echo service

```
Error from server (InternalError): error when creating "ingress-free4-chat.yaml": Internal error occurred: failed calling webhook "validate.nginx.ingress.kubernetes.io": Post "https://ingress-nginx-controller-admission.default.svc:443/networking/v1/ingresses?timeout=10s": context deadline exceeded
```

kubectl delete -A ValidatingWebhookConfiguration ingress-nginx-admission # for solving webhook validate error

kubectl port-forward svc/echo 8888:80 --address 0.0.0.0 # port forward for echo service

https://github.com/kubernetes/ingress-nginx/issues/2062 # cannot use NodePort 443 for ingress-nginx

## HTTPS

kubectl create namespace cert-manager
helm repo add jetstack https://charts.jetstack.io
helm repo update
helm install cert-manager jetstack/cert-manager --namespace cert-manager --version v1.2.0 --set installCRDs=true

kubectl apply -f production_issuer.yaml

## Debug

### Log in k8s node

kubectl get nodes
./nsenter-node.sh node-id-xxx
netstat -pan | grep :443