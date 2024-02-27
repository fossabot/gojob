package main

import (
	"os"

	"github.com/WangYihang/gojob"
	"github.com/WangYihang/gojob/examples/complex-http-crawler/pkg/model"
	"github.com/WangYihang/gojob/pkg/util"
	"github.com/jessevdk/go-flags"
)

type Options struct {
	InputFilePath            string `short:"i" long:"input" description:"input file path" required:"true"`
	OutputFilePath           string `short:"o" long:"output" description:"output file path" required:"true"`
	MaxRetries               int    `short:"r" long:"max-retries" description:"max retries" default:"3"`
	MaxRuntimePerTaskSeconds int    `short:"t" long:"max-runtime-per-task-seconds" description:"max runtime per task seconds" default:"60"`
	NumWorkers               int    `short:"n" long:"num-workers" description:"number of workers" default:"32"`
	NumShards                int    `short:"s" long:"num-shards" description:"number of shards" default:"1"`
	Shard                    int    `short:"d" long:"shard" description:"shard" default:"0"`
}

var opts Options

func init() {
	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}
}

func main() {
	scheduler := gojob.NewScheduler().
		SetNumWorkers(opts.NumWorkers).
		SetMaxRetries(opts.MaxRetries).
		SetMaxRuntimePerTaskSeconds(opts.MaxRuntimePerTaskSeconds).
		SetNumShards(int64(opts.NumShards)).
		SetShard(int64(opts.Shard)).
		SetOutputFilePath(opts.OutputFilePath).
		Start()

	for line := range util.Cat(opts.InputFilePath) {
		scheduler.Submit(model.New(string(line)))
	}
	scheduler.Start()
	scheduler.Wait()
}