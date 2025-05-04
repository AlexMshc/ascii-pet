package gen

//go:generate swagger generate server --spec ../../swagger.json --target . --name Pet --principal interface{} --exclude-main
