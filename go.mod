module github.com/OpenApis/m

go 1.14


replace (
	k8s.io/api => k8s.io/api v0.0.0-20191114100352-16d7abae0d2a
	k8s.io/apimachinery => k8s.io/apimachinery v0.16.5-beta.1
	k8s.io/client-go => k8s.io/client-go v0.0.0-20191016111102-bec269661e48
	github.com/OpenApis/m/kubernetes => /OpenApis/kubernetes
    github.com/OpenApis/m/dingtalk => /OpenApis/dingtalk
)
