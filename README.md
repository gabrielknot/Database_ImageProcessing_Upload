**Database Image Processing Upload**<h6>

 O objetivo desta aplicação é a manipulação de um banco de dados
 através de um frontend construido com javascript.
 

Para rodar esta aplicação são necessárias as seguintes **dependencias:**<h2>

	1 - Docker
	2 - Go (golang)
	3 - NPM (preferência com o yarn instalado de froma global)
**Siga estes passsos para rodar a aplicação**<h2>
		**No terminal digite:**<h1>
	
	 1 - git clone github.com/gabrielknot/Database_ImageProcessing_Upload
	 2 - cd DatabaseDatabase_ImageProcessing_Upload/backend
	 3 - docker run --rm -p 3001:5432 --name pg_test eg_postgresql
	 4 - go get github.com/gorilla/mux
	 5 - go get github.com/rs/cors
	 6 - go get github.com/lib/pq
	 7 - go run main.go
	 5 - cd ../fronted
	 6 - yarn start
**Depois que todos os processos concluirem a aplicação estará rodando em localhost:3003**



