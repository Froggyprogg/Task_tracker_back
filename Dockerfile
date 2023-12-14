# sudo docker compose up -d  --build - исполнение докер файла, создание контейнеров
# sudo docker logs -f go - у меня не работает, но мало ли
# sudo docker exec -it postgres admin -U postgres - запуск
#По ощущениям работает на костылях надо посмортеть как можно доработать
FROM golang:1.21
#Глобальный путь до проекта
WORKDIR /home/FroggyProgg/GolandProjects/github.com/FroggyProgg/Task_tracker_back

COPY go.mod go.sum ./
RUN go mod download

COPY . .

#Глобальный путь до директории с main.go
WORKDIR /home/FroggyProgg/GolandProjects/github.com/FroggyProgg/Task_tracker_back/cmd/server
RUN go build -o main.go .

CMD ["./main.go"]