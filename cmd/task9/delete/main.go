package main

import (
	"DatabaseCourse/internal/task9"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-co-op/gocron/v2"
)

const ExperimantCount = 60
const ExperimantSeconds = time.Second * 5

const ChangeSeconds = time.Second * 10

const ResFile = "delete.log"

type GetTimeUnit struct {
	DoneAt   time.Time
	Duration time.Duration
}

var pgTimes = 0
var redisTimes = 0

func pgPipeline(ch chan<- *GetTimeUnit, st *task9.Task9Storage) {
	if pgTimes >= ExperimantCount {
		return
	}
	stt := time.Now()
	_, err := st.GetTopWealthyClientsPostgres()
	if err != nil {
		log.Println("Error getting clients from PostgreSQL:", err)
		return
	}
	ent := time.Now()
	ch <- &GetTimeUnit{DoneAt: stt, Duration: ent.Sub(stt)}
	pgTimes++
	if pgTimes >= ExperimantCount {
		close(ch)
	}
}

func redisPipeline(ch chan<- *GetTimeUnit, st *task9.Task9Storage) {
	if redisTimes >= ExperimantCount {
		return
	}
	stt := time.Now()
	_, err := st.GetTopWealthyClients()
	if err != nil {
		log.Println("Error getting clients from Redis:", err)
		return
	}
	ent := time.Now()
	ch <- &GetTimeUnit{DoneAt: stt, Duration: ent.Sub(stt)}
	redisTimes++
	if redisTimes >= ExperimantCount {
		close(ch)
	}
}

func DeletePipeline(st *task9.Task9Storage) {
	err := st.DeleteRandomAccount()
	if err != nil {
		log.Println("Error deleting account:", err)
		return
	}
}

func main() {
	pg, err := task9.NewPgStorage(context.TODO())
	if err != nil {
		panic(err)
	}
	redis, err := task9.NewRedisStorage()
	if err != nil {
		panic(err)
	}

	storage := task9.NewTask9Storage(pg, redis)
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		panic(err)
	}

	pgch := make(chan *GetTimeUnit)
	_, err = scheduler.NewJob(gocron.DurationJob(ExperimantSeconds), gocron.NewTask(pgPipeline, pgch, storage))
	if err != nil {
		panic(err)
	}
	redisch := make(chan *GetTimeUnit)
	_, err = scheduler.NewJob(gocron.DurationJob(ExperimantSeconds), gocron.NewTask(redisPipeline, redisch, storage))
	if err != nil {
		panic(err)
	}

	_, err = scheduler.NewJob(gocron.DurationJob(ChangeSeconds), gocron.NewTask(DeletePipeline, storage))
	if err != nil {
		panic(err)
	}

	scheduler.Start()
	defer scheduler.Shutdown()
	pgStat := make([]GetTimeUnit, 0, ExperimantCount)
	for un := range pgch {
		pgStat = append(pgStat, *un)
	}
	redisStat := make([]GetTimeUnit, 0, ExperimantCount)
	for un := range redisch {
		redisStat = append(redisStat, *un)
	}

	out, err := os.Create(ResFile)
	if err != nil {
		panic(err)
	}
	for i := 0; i < ExperimantCount; i++ {
		out.WriteString(fmt.Sprintf("%d %f %f\n", i*int(ExperimantSeconds.Seconds()), pgStat[i].Duration.Seconds(), redisStat[i].Duration.Seconds()))
	}

}
