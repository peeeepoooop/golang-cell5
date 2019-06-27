GCC=go
GCMD=run
GPATH=main.go

run:
	$(GCC) $(GCMD) $(GPATH)

install:
	make install_routes
	make install_db

install_routes:
	go get -u github.com/gorilla/mux

install_db: 
	go get -u github.com/go-xorm/xorm