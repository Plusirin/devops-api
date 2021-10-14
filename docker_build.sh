swag init && \
docker login -u admin -p Harbor12345 harbor.xwsoftlan.com && \
docker build --rm -t harbor.xwsoftlan.com/devops/devops-api:latest .  && \
docker push harbor.xwsoftlan.com/devops/devops-api:latest && \
kubectl apply -f deployment.yaml
