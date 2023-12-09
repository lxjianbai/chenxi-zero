GOCTL_HOME=deploy/goctl/1.5.6
model:
	goctl model mysql ddl -dir="pkg/model/cloudmodel" -cache=false -src=data/ddl/cloud/*.sql -strict=true --style=goZero -home=${GOCTL_HOME}
	goctl model mysql ddl -dir="pkg/model/gamemodel" -cache=false -src=data/ddl/game/*.sql -strict=true --style=goZero -home=${GOCTL_HOME}
	goimports -w pkg/model/cloudmodel/*.go
	goimports -w pkg/model/gamemodel/*.go

proto:
	protoc pkg/cpb/*.proto --proto_path=pkg/cpb --go_out=pkg 

api:
	goctl api go -dir="app/api" -api="app/api/desc/api.api" -home=${GOCTL_HOME}

cms:
	goctl api go -dir="app/cms" -api="app/cms/desc/cms.api" -home=${GOCTL_HOME}

usercenter:
	goctl rpc protoc "app/usercenter/usercenter.proto"  --go_out="app/usercenter/" --go-grpc_out="app/usercenter/" --zrpc_out="app/usercenter/"