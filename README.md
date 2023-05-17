# README
- go gin default project
- project default URL: http://127.0.0.1:8080/

- Integrate its basic functions
```text
    1. config reload
    2. metrics
    3. BasicAuth
    4. build scripts
    5. docker / kubernetes deploy
```

- default URL path
```text
/check
    health status check
/-/reload
    reload config file
/metrics
/
    default page
/api/v1/
    BasicAuth page
```


### build/deploy
- build package
```
# 执行 go build, 并制作 images
bash cmd/linux_build.sh v0.2
```

- deploy
```
kubectl create namespace go
kubectl -n go create configmap go-default-service-configmap --from-file=configs/config.yaml
kubectl apply -f build/go_default_service-deploy.yaml
```

- deploy 本地部署
```sh
    version=v0.8
    TIME=$(echo -n $(date +%Y-%m-%d-%H-%M))
    server_manager_ip="8.8.8.8"
    project_name="default-project"
    scp -P 49999 ${project_name}-linux-amd64-$version.tar.gz root@${server_manager_ip}:
    ssh -p 49999 root@${server_manager_ip} "
        rm -rf /data/service/${project_name}_data/${TIME}
        mkdir -p /data/service/${project_name}_data/${TIME}
        tar -zxf ${project_name}-linux-amd64-$version.tar.gz -C /data/service/${project_name}_data/${TIME}
        rm -f /data/service/${project_name}
        ln -s /data/service/${project_name}_data/${TIME} /data/service/${project_name}
        ln -s /data/service/${project_name}_data/data/ /data/service/${project_name}/data
        chmod +x /data/service/${project_name}/build/${project_name}
        systemctl restart ${project_name}
    "
```

### demo
- initial page

![初始页面演示](./doc/img/init_demo.png)
