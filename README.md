# README
- go gin default project
- project default URL: http://127.0.0.1:8080/


### build/deploy
- build package
```
# 执行 go build, 并制作 images
bash cmd/linux_build.sh
```

- deploy
```
kubectl create namespace go
kubectl -n go create configmap go-default-service-configmap --from-file=configs/config.yaml
kubectl apply -f build/go_default_service-deploy.yaml
```

### 演示
- 初始页面

![初始页面演示](https://github.com/weiqiang333/go-web-init-template/doc/img/init_demo.png)