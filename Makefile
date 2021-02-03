
gen:
	mockgen -source=pkg/common.go -destination=src/executor/common_mock_test.go -package=executor
