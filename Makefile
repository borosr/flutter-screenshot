
gen:
	mockgen -source=pkg/common.go -destination=src/executor/common_mock_test.go -package=executor
	mockgen -source=pkg/exec/cmd.go -destination=pkg/ios/ios_mock_test.go -package=ios
	mockgen -source=pkg/exec/cmd.go -destination=src/executor/cmd_mock_test.go -package=executor
