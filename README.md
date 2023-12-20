<a href="https://fcexperience.fullcycle.com.br/">
   <img src="https://events-fullcycle.s3.amazonaws.com/events-fullcycle/media/images/962edb195c0448df860fbea9304a7f24.png" alt="FCLX" title="Full Cycle Learning Experience" align="right" height="60" />
</a>

# Full Cycle Learning Experience

Para acessar todas as aulas e lives, acesse: https://fcexperience.fullcycle.com.br/

### Aula 01

Docker e Containers e Microsserviço do ChatGPT

### Aula 02

gRPC vs REST: Comunicação entre microsserviços do ChatGPT na prática

### Aula 03

Backend for Frontend com Next.js

### Aula 04

Frontend do ChatGPT com Next.js e React

### Aula 05

KeyCloak: Integrando servidor de identidade

<img src="./arquitetura do projeto.png" alt="FCLX" title="Full Cycle Learning Experience" />

#MYSQL
https://www.baeldung.com/docker-cant-connect-local-mysql
docker pull mysql:latest

docker run --name mysql -p 3306:3306 -v ~/mysql-data:/var/lib/mysql -e MYSQL*ROOT_PASSWORD=root -d mysql:latest
mysql -h 172.17.0.2 -P 3306 --protocol=tcp -u root -p
CREATE USER 'fols'@'%' IDENTIFIED BY 'fols';
GRANT ALL PRIVILEGES ON *.\_ TO 'fols'@'%';
create database chat_test;

docker compose exec chatservice bash
go run cmd/chatservice/main.go
