package services

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/devlipe/encoder/application/repositories"
	"github.com/devlipe/encoder/domain"
	"io"
	"log"
	"os"
	"os/exec"
)

type VideoService struct {
	Video           *domain.Video
	VideoRepository repositories.VideoRepository
}

func NewVideoService() VideoService {
	return VideoService{}
}

func (v *VideoService) Download(bucketName string) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)

	if err != nil {
		return err
	}

	bkt := client.Bucket(bucketName)
	obj := bkt.Object(v.Video.FilePath)

	r, err := obj.NewReader(ctx)
	if err != nil {
		return err
	}

	defer func(r *storage.Reader) {
		err := r.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(r)

	body, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	f, err := os.Create(fmt.Sprintf("%s/%s.mp4", os.Getenv("LOCAL_STORAGE_PATH"), v.Video.ID))
	if err != nil {
		return err
	}

	_, err = f.Write(body)
	if err != nil {
		return err
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(f)

	log.Printf("Video %v has been stored", v.Video.ID)

	return nil
}

func (v *VideoService) Fragment() error {

	err := os.Mkdir(os.Getenv("LOCAL_STORAGE_PATH")+"/"+v.Video.ID, os.ModePerm)

	if err != nil {
		return err
	}

	source := fmt.Sprintf("%s/%s.mp4", os.Getenv("LOCAL_STORAGE_PATH"), v.Video.ID)
	target := fmt.Sprintf("%s/%s.frag", os.Getenv("LOCAL_STORAGE_PATH"), v.Video.ID)

	cmd := exec.Command("./mp4fragment", source, target)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	printOutput(output)

	return nil
}

func printOutput(out []byte) {
	if len(out) > 0 {
		log.Printf("======> Output: %s\n", out)
	}
}
