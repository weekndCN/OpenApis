module github.com/OpenApis/m

go 1.14

replace (
	github.com/OpenApis/m/dingtalk => /OpenApis/dingtalk
	github.com/OpenApis/m/loki => /OpenApis/loki
	github.com/OpenApis/m/mailer => /OpenApis/mailer
	github.com/OpenApis/m/kubernetes => /OpenApis/kubernetes
	k8s.io/api => k8s.io/api v0.0.0-20191114100352-16d7abae0d2a
	k8s.io/apimachinery => k8s.io/apimachinery v0.16.5-beta.1
	k8s.io/client-go => k8s.io/client-go v0.0.0-20191016111102-bec269661e48
)

require github.com/gorilla/websocket v1.4.2 // indirect
